module github.com/pulumi/pulumi-tls/provider/v4

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.23.0
	github.com/pulumi/pulumi/sdk/v3 v3.32.1
	github.com/terraform-providers/terraform-provider-tls/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.15.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220523144019-a9dc436975cc
	github.com/pulumi/pulumi/pkg/v3 => github.com/pulumi/pulumi/pkg/v3 v3.32.1
	github.com/terraform-providers/terraform-provider-tls/shim => ./shim
)
