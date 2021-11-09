module github.com/pulumi/pulumi-tls/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
	github.com/terraform-providers/terraform-provider-tls/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
	github.com/terraform-providers/terraform-provider-tls/shim => ./shim
)
