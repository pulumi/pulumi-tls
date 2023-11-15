// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a public key from a PEM-encoded private key.
//
// Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.NewPrivateKey(ctx, "ed25519-example", &tls.PrivateKeyArgs{
//				Algorithm: pulumi.String("ED25519"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = tls.GetPublicKeyOutput(ctx, tls.GetPublicKeyOutputArgs{
//				PrivateKeyPem: ed25519_example.PrivateKeyPem,
//			}, nil)
//			_, err = tls.GetPublicKey(ctx, &tls.GetPublicKeyArgs{
//				PrivateKeyOpenssh: pulumi.StringRef(readFileOrPanic("~/.ssh/id_rsa_rfc4716")),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPublicKey(ctx *pulumi.Context, args *GetPublicKeyArgs, opts ...pulumi.InvokeOption) (*GetPublicKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPublicKeyResult
	err := ctx.Invoke("tls:index/getPublicKey:getPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicKey.
type GetPublicKeyArgs struct {
	// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyPem`.
	PrivateKeyOpenssh *string `pulumi:"privateKeyOpenssh"`
	// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyOpenssh`.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
}

// A collection of values returned by getPublicKey.
type GetPublicKeyResult struct {
	// The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA` and `ED25519`.
	Algorithm string `pulumi:"algorithm"`
	// Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
	Id string `pulumi:"id"`
	// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyPem`.
	PrivateKeyOpenssh *string `pulumi:"privateKeyOpenssh"`
	// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyOpenssh`.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `publicKeyOpenssh` and ECDSA P224 limitations.
	PublicKeyFingerprintMd5 string `pulumi:"publicKeyFingerprintMd5"`
	// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `publicKeyOpenssh` and ECDSA P224 limitations.
	PublicKeyFingerprintSha256 string `pulumi:"publicKeyFingerprintSha256"`
	// The public key, in  OpenSSH PEM (RFC 4716).
	PublicKeyOpenssh string `pulumi:"publicKeyOpenssh"`
	// The public key, in PEM (RFC 1421).
	PublicKeyPem string `pulumi:"publicKeyPem"`
}

func GetPublicKeyOutput(ctx *pulumi.Context, args GetPublicKeyOutputArgs, opts ...pulumi.InvokeOption) GetPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicKeyResult, error) {
			args := v.(GetPublicKeyArgs)
			r, err := GetPublicKey(ctx, &args, opts...)
			var s GetPublicKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPublicKeyResultOutput)
}

// A collection of arguments for invoking getPublicKey.
type GetPublicKeyOutputArgs struct {
	// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyPem`.
	PrivateKeyOpenssh pulumi.StringPtrInput `pulumi:"privateKeyOpenssh"`
	// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyOpenssh`.
	PrivateKeyPem pulumi.StringPtrInput `pulumi:"privateKeyPem"`
}

func (GetPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicKeyArgs)(nil)).Elem()
}

// A collection of values returned by getPublicKey.
type GetPublicKeyResultOutput struct{ *pulumi.OutputState }

func (GetPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicKeyResult)(nil)).Elem()
}

func (o GetPublicKeyResultOutput) ToGetPublicKeyResultOutput() GetPublicKeyResultOutput {
	return o
}

func (o GetPublicKeyResultOutput) ToGetPublicKeyResultOutputWithContext(ctx context.Context) GetPublicKeyResultOutput {
	return o
}

// The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA` and `ED25519`.
func (o GetPublicKeyResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

// Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
func (o GetPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyPem`.
func (o GetPublicKeyResultOutput) PrivateKeyOpenssh() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPublicKeyResult) *string { return v.PrivateKeyOpenssh }).(pulumi.StringPtrOutput)
}

// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `privateKeyOpenssh`.
func (o GetPublicKeyResultOutput) PrivateKeyPem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPublicKeyResult) *string { return v.PrivateKeyPem }).(pulumi.StringPtrOutput)
}

// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `publicKeyOpenssh` and ECDSA P224 limitations.
func (o GetPublicKeyResultOutput) PublicKeyFingerprintMd5() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.PublicKeyFingerprintMd5 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `publicKeyOpenssh` and ECDSA P224 limitations.
func (o GetPublicKeyResultOutput) PublicKeyFingerprintSha256() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.PublicKeyFingerprintSha256 }).(pulumi.StringOutput)
}

// The public key, in  OpenSSH PEM (RFC 4716).
func (o GetPublicKeyResultOutput) PublicKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.PublicKeyOpenssh }).(pulumi.StringOutput)
}

// The public key, in PEM (RFC 1421).
func (o GetPublicKeyResultOutput) PublicKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicKeyResult) string { return v.PublicKeyPem }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicKeyResultOutput{})
}
