package main

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		_, err := tls.NewPrivateKey(ctx, "private-key", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
		})

		if err != nil {
			return err
		}
		return nil
	})
}
