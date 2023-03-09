package main

import (
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		caKey, err := tls.NewPrivateKey(ctx, "ca-key", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA").ToStringOutput(),
			RsaBits:   pulumi.Int(2048).ToIntPtrOutput(),
		})

		if err != nil {
			return err
		}

		caRoot, err := tls.NewSelfSignedCert(ctx, "self-signed", &tls.SelfSignedCertArgs{
			PrivateKeyPem: caKey.PrivateKeyPem,

			ValidityPeriodHours: pulumi.Int(87600).ToIntOutput(),
			IsCaCertificate:     pulumi.Bool(true).ToBoolPtrOutput(),
			AllowedUses:         pulumi.ToStringArray([]string{"cert_signing"}).ToStringArrayOutput(),

			Subject: tls.SelfSignedCertSubjectArgs{
				CommonName:   pulumi.String("CustomizeDiff repro").ToStringPtrOutput(),
				Organization: pulumi.String("Pulumi").ToStringPtrOutput(),
			},
		})

		if err != nil {
			return err
		}

		certKey, err := tls.NewPrivateKey(ctx, "cert-key", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA").ToStringOutput(),
			RsaBits:   pulumi.Int(2048).ToIntPtrOutput(),
		})

		if err != nil {
			return err
		}

		localCertReq, err := tls.NewCertRequest(ctx, "local", &tls.CertRequestArgs{
			PrivateKeyPem: certKey.PrivateKeyPem,

			Subject: tls.CertRequestSubjectArgs{
				CommonName:   pulumi.String("sometest.pulumi.com").ToStringPtrOutput(),
				Organization: pulumi.String("Pulumi").ToStringPtrOutput(),
			},
		})
		if err != nil {
			return err
		}

		localCert, err := tls.NewLocallySignedCert(ctx, "local", &tls.LocallySignedCertArgs{
			CertRequestPem: localCertReq.CertRequestPem,

			CaPrivateKeyPem: caKey.PrivateKeyPem,
			CaCertPem:       caRoot.CertPem,

			ValidityPeriodHours: pulumi.Int(2).ToIntOutput(),
			EarlyRenewalHours:   pulumi.Int(3).ToIntOutput(),
			AllowedUses:         pulumi.ToStringArray([]string{"server_auth"}),
		})

		if err != nil {
			return err
		}

		ctx.Export("cert-id", localCert.ID())

		return nil
	})
}
