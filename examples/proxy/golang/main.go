// Copyright 2016-2023, Pulumi Corporation.
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

package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		prov, err := tls.NewProvider(ctx, "tlsprovider", &tls.ProviderArgs{
			Proxy: tls.ProviderProxyArgs{
				Url:     pulumi.String("http://localhost:8080"),
				FromEnv: pulumi.Bool(false),
			},
		})

		if err != nil {
			return err
		}

		url := "https://pulumi.com"
		verifyChain := false
		res, err := tls.GetCertificate(ctx, &tls.GetCertificateArgs{
			Url:         &url,
			VerifyChain: &verifyChain,
		}, pulumi.Provider(prov))
		if err != nil {
			return err
		}

		ctx.Export("issuer", pulumi.String(res.Certificates[0].Issuer))
		return nil
	})
}
