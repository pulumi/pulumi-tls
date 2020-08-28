// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information, such as SHA1 fingerprint or serial number, about the TLS certificates that
// protect an HTTPS website. Note that the certificate chain isn't verified.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/eks"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi-tls/sdk/v2/go/tls"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleCluster, err := eks.NewCluster(ctx, "exampleCluster", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewOpenIdConnectProvider(ctx, "exampleOpenIdConnectProvider", &iam.OpenIdConnectProviderArgs{
// 			ClientIdLists: pulumi.StringArray{
// 				pulumi.String("sts.amazonaws.com"),
// 			},
// 			ThumbprintLists: pulumi.StringArray{
// 				exampleCertificate.ApplyT(func(exampleCertificate tls.GetCertificateResult) (string, error) {
// 					return exampleCertificate.Certificates[0].Sha1Fingerprint, nil
// 				}).(pulumi.StringOutput),
// 			},
// 			Url: pulumi.String(exampleCluster.Identities.ApplyT(func(identities []eks.ClusterIdentity) (string, error) {
// 				return identities[0].Oidcs[0].Issuer, nil
// 			}).(pulumi.StringOutput)),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCertificate(ctx *pulumi.Context, args *GetCertificateArgs, opts ...pulumi.InvokeOption) (*GetCertificateResult, error) {
	var rv GetCertificateResult
	err := ctx.Invoke("tls:index/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type GetCertificateArgs struct {
	// The URL of the website to get the certificates from.
	Url string `pulumi:"url"`
	// Whether to verify the certificate chain while parsing it or not
	VerifyChain *bool `pulumi:"verifyChain"`
}

// A collection of values returned by getCertificate.
type GetCertificateResult struct {
	// The certificates protecting the site, with the root of the chain first.
	// * `certificates.#.not_after` - The time until which the certificate is invalid, as an
	//   [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	// * `certificates.#.not_before` - The time after which the certificate is valid, as an
	//   [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	// * `certificates.#.is_ca` - `true` if this certificate is a ca certificate.
	// * `certificates.#.issuer` - Who verified and signed the certificate, roughly following
	//   [RFC2253](https://tools.ietf.org/html/rfc2253).
	// * `certificates.#.public_key_algorithm` - The algorithm used to create the certificate.
	// * `certificates.#.serial_number` - Number that uniquely identifies the certificate with the CA's system. The `format`
	//   function can be used to convert this base 10 number into other bases, such as hex.
	// * `certificates.#.sha1_fingerprint` - The SHA1 fingerprint of the public key of the certificate.
	// * `certificates.#.signature_algorithm` - The algorithm used to sign the certificate.
	// * `certificates.#.subject` - The entity the certificate belongs to, roughly following
	//   [RFC2253](https://tools.ietf.org/html/rfc2253).
	// * `certificates.#.version` - The version the certificate is in.
	Certificates []GetCertificateCertificate `pulumi:"certificates"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	Url         string `pulumi:"url"`
	VerifyChain *bool  `pulumi:"verifyChain"`
}