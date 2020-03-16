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
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pulumi/pulumi-terraform-bridge/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-tls/tls"
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
func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		P:           tls.Provider().(*schema.Provider),
		Name:        "tls",
		Description: "A Pulumi package to create TLS resources in Pulumi programs.",
		Keywords:    []string{"pulumi", "tls"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-tls",
		Resources: map[string]*tfbridge.ResourceInfo{
			"tls_cert_request":        {Tok: tlsResource(tlsMod, "CertRequest")},
			"tls_locally_signed_cert": {Tok: tlsResource(tlsMod, "LocallySignedCert")},
			"tls_private_key":         {Tok: tlsResource(tlsMod, "PrivateKey")},
			"tls_self_signed_cert":    {Tok: tlsResource(tlsMod, "SelfSignedCert")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"tls_public_key": {Tok: tlsDataSource(tlsMod, "getPublicKey")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.12.1-preview",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				"tls": "Tls",
			},
		},
	}
}
