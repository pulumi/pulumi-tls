module github.com/pulumi/pulumi-tls/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0-beta.1
	github.com/pulumi/pulumi/sdk/v2 v2.0.0-beta.3
	github.com/terraform-providers/terraform-provider-tls v0.0.0-20190925211901-afd9e9546f57
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
