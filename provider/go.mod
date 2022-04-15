module github.com/pulumi/pulumi-tls/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.1-0.20220415111455-5c2d9875a867
	github.com/pulumi/pulumi/sdk/v3 v3.28.0
	github.com/terraform-providers/terraform-provider-tls/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.15.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
	github.com/terraform-providers/terraform-provider-tls/shim => ./shim
)
