module github.com/pulumi/pulumi-tls/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.8.0
	github.com/pulumi/pulumi/sdk/v2 v2.10.0
	github.com/terraform-providers/terraform-provider-tls v1.2.1-0.20200724134453-f1b91f494149
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
