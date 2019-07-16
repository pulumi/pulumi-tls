module github.com/pulumi/pulumi-tls

go 1.12

require (
	github.com/hashicorp/terraform v0.12.0-rc1.0.20190509225429-28b2383eacae
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.23-0.20190715212628-02ffff88409f
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190716112909-08d502e9b427
	github.com/stretchr/testify v1.3.1-0.20190709195754-d84e815d441d
	github.com/terraform-providers/terraform-provider-tls v0.0.0-20190430200932-a63e85603781
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)
