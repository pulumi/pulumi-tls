module github.com/pulumi/pulumi-tls

go 1.12

require (
	github.com/hashicorp/terraform v0.12.0-rc1.0.20190509225429-28b2383eacae
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.26-0.20190723034217-ed5b8437d126
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190724233644-fdafc00b3f6a
	github.com/stretchr/testify v1.3.1-0.20190709195754-d84e815d441d
	github.com/terraform-providers/terraform-provider-tls v0.0.0-20190430200932-a63e85603781
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)
