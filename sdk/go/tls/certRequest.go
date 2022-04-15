// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tls.NewCertRequest(ctx, "example", &tls.CertRequestArgs{
// 			PrivateKeyPem: readFileOrPanic("private_key.pem"),
// 			Subjects: CertRequestSubjectArray{
// 				&CertRequestSubjectArgs{
// 					CommonName:   pulumi.String("example.com"),
// 					Organization: pulumi.String("ACME Examples, Inc"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CertRequest struct {
	pulumi.CustomResourceState

	// The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringOutput `pulumi:"certRequestPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames pulumi.StringArrayOutput `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
	// and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
	// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
	// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
	// based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subjects CertRequestSubjectArrayOutput `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
}

// NewCertRequest registers a new resource with the given unique name, arguments, and options.
func NewCertRequest(ctx *pulumi.Context,
	name string, args *CertRequestArgs, opts ...pulumi.ResourceOption) (*CertRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKeyPem == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKeyPem'")
	}
	if args.Subjects == nil {
		return nil, errors.New("invalid value for required argument 'Subjects'")
	}
	var resource CertRequest
	err := ctx.RegisterResource("tls:index/certRequest:CertRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertRequest gets an existing CertRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertRequestState, opts ...pulumi.ResourceOption) (*CertRequest, error) {
	var resource CertRequest
	err := ctx.ReadResource("tls:index/certRequest:CertRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertRequest resources.
type certRequestState struct {
	// The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem *string `pulumi:"certRequestPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames []string `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
	// and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
	// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
	// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
	// based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subjects []CertRequestSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
}

type CertRequestState struct {
	// The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	CertRequestPem pulumi.StringPtrInput
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames pulumi.StringArrayInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
	// and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
	KeyAlgorithm pulumi.StringPtrInput
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
	// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
	// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	PrivateKeyPem pulumi.StringPtrInput
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
	// based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subjects CertRequestSubjectArrayInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
}

func (CertRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certRequestState)(nil)).Elem()
}

type certRequestArgs struct {
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames []string `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
	// and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
	// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
	// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	PrivateKeyPem string `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
	// based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subjects []CertRequestSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
}

// The set of arguments for constructing a CertRequest resource.
type CertRequestArgs struct {
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames pulumi.StringArrayInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
	// and ignored, as the key algorithm is now inferred from the key.
	//
	// Deprecated: This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
	KeyAlgorithm pulumi.StringPtrInput
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
	// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
	// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
	PrivateKeyPem pulumi.StringInput
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
	// based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subjects CertRequestSubjectArrayInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
}

func (CertRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certRequestArgs)(nil)).Elem()
}

type CertRequestInput interface {
	pulumi.Input

	ToCertRequestOutput() CertRequestOutput
	ToCertRequestOutputWithContext(ctx context.Context) CertRequestOutput
}

func (*CertRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequest)(nil)).Elem()
}

func (i *CertRequest) ToCertRequestOutput() CertRequestOutput {
	return i.ToCertRequestOutputWithContext(context.Background())
}

func (i *CertRequest) ToCertRequestOutputWithContext(ctx context.Context) CertRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestOutput)
}

// CertRequestArrayInput is an input type that accepts CertRequestArray and CertRequestArrayOutput values.
// You can construct a concrete instance of `CertRequestArrayInput` via:
//
//          CertRequestArray{ CertRequestArgs{...} }
type CertRequestArrayInput interface {
	pulumi.Input

	ToCertRequestArrayOutput() CertRequestArrayOutput
	ToCertRequestArrayOutputWithContext(context.Context) CertRequestArrayOutput
}

type CertRequestArray []CertRequestInput

func (CertRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertRequest)(nil)).Elem()
}

func (i CertRequestArray) ToCertRequestArrayOutput() CertRequestArrayOutput {
	return i.ToCertRequestArrayOutputWithContext(context.Background())
}

func (i CertRequestArray) ToCertRequestArrayOutputWithContext(ctx context.Context) CertRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestArrayOutput)
}

// CertRequestMapInput is an input type that accepts CertRequestMap and CertRequestMapOutput values.
// You can construct a concrete instance of `CertRequestMapInput` via:
//
//          CertRequestMap{ "key": CertRequestArgs{...} }
type CertRequestMapInput interface {
	pulumi.Input

	ToCertRequestMapOutput() CertRequestMapOutput
	ToCertRequestMapOutputWithContext(context.Context) CertRequestMapOutput
}

type CertRequestMap map[string]CertRequestInput

func (CertRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertRequest)(nil)).Elem()
}

func (i CertRequestMap) ToCertRequestMapOutput() CertRequestMapOutput {
	return i.ToCertRequestMapOutputWithContext(context.Background())
}

func (i CertRequestMap) ToCertRequestMapOutputWithContext(ctx context.Context) CertRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestMapOutput)
}

type CertRequestOutput struct{ *pulumi.OutputState }

func (CertRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequest)(nil)).Elem()
}

func (o CertRequestOutput) ToCertRequestOutput() CertRequestOutput {
	return o
}

func (o CertRequestOutput) ToCertRequestOutputWithContext(ctx context.Context) CertRequestOutput {
	return o
}

type CertRequestArrayOutput struct{ *pulumi.OutputState }

func (CertRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertRequest)(nil)).Elem()
}

func (o CertRequestArrayOutput) ToCertRequestArrayOutput() CertRequestArrayOutput {
	return o
}

func (o CertRequestArrayOutput) ToCertRequestArrayOutputWithContext(ctx context.Context) CertRequestArrayOutput {
	return o
}

func (o CertRequestArrayOutput) Index(i pulumi.IntInput) CertRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertRequest {
		return vs[0].([]*CertRequest)[vs[1].(int)]
	}).(CertRequestOutput)
}

type CertRequestMapOutput struct{ *pulumi.OutputState }

func (CertRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertRequest)(nil)).Elem()
}

func (o CertRequestMapOutput) ToCertRequestMapOutput() CertRequestMapOutput {
	return o
}

func (o CertRequestMapOutput) ToCertRequestMapOutputWithContext(ctx context.Context) CertRequestMapOutput {
	return o
}

func (o CertRequestMapOutput) MapIndex(k pulumi.StringInput) CertRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertRequest {
		return vs[0].(map[string]*CertRequest)[vs[1].(string)]
	}).(CertRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestInput)(nil)).Elem(), &CertRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestArrayInput)(nil)).Elem(), CertRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestMapInput)(nil)).Elem(), CertRequestMap{})
	pulumi.RegisterOutputType(CertRequestOutput{})
	pulumi.RegisterOutputType(CertRequestArrayOutput{})
	pulumi.RegisterOutputType(CertRequestMapOutput{})
}
