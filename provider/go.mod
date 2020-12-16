module github.com/pulumi/pulumi-tls/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.16.0
	github.com/pulumi/pulumi/sdk/v2 v2.15.1-0.20201202214525-260620430c4c
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/terraform-providers/terraform-provider-tls v1.2.1-0.20200724134453-f1b91f494149
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
