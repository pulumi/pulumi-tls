module github.com/pulumi/pulumi-tls

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/pkg v1.13.1
	github.com/pulumi/pulumi/sdk v1.13.1
	github.com/terraform-providers/terraform-provider-tls v0.0.0-20190925211901-afd9e9546f57
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
