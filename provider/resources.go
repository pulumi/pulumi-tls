// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tls

import (
	"bytes"
	"fmt"
	"path/filepath"
	"unicode"

	_ "embed" // used to store bridge-metadata.json in the compiled binary

	"github.com/hashicorp/terraform-provider-tls/shim"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-tls/provider/v5/pkg/version"
)

// all of the tls token components used below.
const (
	tlsPkg = "tls"
	tlsMod = "index"
)

// tlsMember manufactures a type token for the TLS package and the given module and type.
func tlsMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(tlsPkg + ":" + mod + ":" + mem)
}

// tlsType manufactures a type token for the TLS package and the given module and type.
func tlsType(mod string, typ string) tokens.Type {
	return tokens.Type(tlsMember(mod, typ))
}

// tlsDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the data
// source's first character.
func tlsDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return tlsMember(mod+"/"+fn, res)
}

// tlsResource manufactures a standard resource token given a module and resource name.
// It automatically uses the TLS package and names the file by simply lower casing the resource's
// first character.
func tlsResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return tlsType(mod+"/"+fn, res)
}

//go:embed cmd/pulumi-resource-tls/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		P:            pf.ShimProvider(shim.NewProvider()),
		Name:         "tls",
		Description:  "A Pulumi package to create TLS resources in Pulumi programs.",
		Keywords:     []string{"pulumi", "tls"},
		License:      "Apache-2.0",
		Homepage:     "https://pulumi.io",
		Repository:   "https://github.com/pulumi/pulumi-tls",
		Version:      version.Version,
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		GitHubOrg:    "hashicorp",
		DocRules:     &tfbridge.DocRuleInfo{EditRules: docEditRules},
		Resources: map[string]*tfbridge.ResourceInfo{
			"tls_cert_request":        {Tok: tlsResource(tlsMod, "CertRequest")},
			"tls_locally_signed_cert": {Tok: tlsResource(tlsMod, "LocallySignedCert")},
			"tls_private_key":         {Tok: tlsResource(tlsMod, "PrivateKey")},

			"tls_self_signed_cert": {
				Tok:                 tlsResource(tlsMod, "SelfSignedCert"),
				PreStateUpgradeHook: selfSignedCertPreStateUpgradeHook,
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"tls_public_key":  {Tok: tlsDataSource(tlsMod, "getPublicKey")},
			"tls_certificate": {Tok: tlsDataSource(tlsMod, "getCertificate")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				RespectSchemaVersion: true,
			}
			i.PyProject.Enabled = true
			return i
		})(),

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", tlsPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				tlsPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"tls": "Tls",
			},
		},
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
	}

	return info
}

func selfSignedCertPreStateUpgradeHook(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	// When upstream starts versioning the schemas, do nothing. Only attempt to fixup state for the transition to
	// Plugin Framework.
	if args.ResourceSchemaVersion != 0 || args.PriorStateSchemaVersion != 0 {
		return args.PriorStateSchemaVersion, args.PriorState, nil
	}

	s := args.PriorState

	// Computed+Optional booleans defaulting to false. Prior to PF, when unset they are not present in Pulumi state,
	// causing non-empty diffs (nil vs false). Fixup to false to get empty diffs.
	for _, f := range []string{
		"setSubjectKeyId",
		"setAuthorityKeyId",
		"isCaCertificate",
	} {
		k := resource.PropertyKey(f)
		if _, ok := s[k]; !ok {
			s[k] = resource.NewBoolProperty(false)
		}
	}

	if subject, ok := s["subject"]; ok && subject.IsObject() {
		su := subject.ObjectValue()

		// All these are optional strings, but manifest as empty strings in Pulumi state. This causes spurious
		// diffs. Delete empty strings if found.
		for _, country := range []string{
			"commonName",
			"country",
			"locality",
			"organization",
			"organizationalUnit",
			"postalCode",
			"province",
			"serialNumber",
		} {
			country := resource.PropertyKey(country)
			if x, ok := su[country]; ok && x.IsString() && x.StringValue() == "" {
				delete(su, country)
			}
		}

		if sa, ok := su["streetAddresses"]; ok && sa.IsArray() && len(sa.ArrayValue()) == 0 {
			delete(su, "streetAddresses")
		}
	}

	// allowedUses is a required property. If the state does not have it, replace with an empty array.
	if _, ok := s["allowedUses"]; !ok {
		s["allowedUses"] = resource.NewArrayProperty([]resource.PropertyValue{})
	}

	return 0, s, nil
}

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		skipSecretsSection,
		noFileFunction,
	)
}

// Removes a "Secrets and Terraform state" section that does not apply to Pulumi
var skipSecretsSection = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		return tfgen.SkipSectionByHeaderContent(content, func(headerText string) bool {
			return headerText == "Secrets and Terraform state"
		})
	},
}

// Removes an unnecessary reference to the TF file interpolation function
var noFileFunction = tfbridge.DocsEdit{
	Path: "self_signed_cert.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		inputByte := []byte("This can be read from a separate file using the `file` interpolation function.")
		if bytes.Contains(content, inputByte) {
			content = bytes.ReplaceAll(content, inputByte, nil)
		} else {
			// Hard error to ensure we keep this content up to date
			return nil, fmt.Errorf("could not find text in upstream self_signed_cert.md, "+
				"please verify upstream self_signed_cert.md contains:\n\n %s", string(inputByte),
			)
		}
		return content, nil
	},
}
