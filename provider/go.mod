module github.com/pulumi/pulumi-tls/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.1.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.1-0.20200424001829-090f390d7b1a
	github.com/terraform-providers/terraform-provider-tls v0.0.0-20190925211901-afd9e9546f57
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
