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
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-provider-tls/shim"
	tfpfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-tls/provider/v5/pkg/version"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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

// Provider returns additional overlaid schema and metadata associated with the tls package.
func Provider() tfpfbridge.ProviderInfo {
	info := tfbridge.ProviderInfo{
		Name:        "tls",
		Description: "A Pulumi package to create TLS resources in Pulumi programs.",
		Keywords:    []string{"pulumi", "tls"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-tls",
		Version:     version.Version,
		GitHubOrg:   "hashicorp",
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
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", tlsPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				tlsPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"tls": "Tls",
			},
		},
	}

	return tfpfbridge.ProviderInfo{
		ProviderInfo: info,
		NewProvider:  shim.NewProvider,
	}
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
