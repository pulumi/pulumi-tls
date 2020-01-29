// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package tls

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-tls/blob/master/website/docs/r/cert_request.html.markdown.
type CertRequest struct {
	pulumi.CustomResourceState

	// The certificate request data in PEM format.
	CertRequestPem pulumi.StringOutput `pulumi:"certRequestPem"`
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayOutput `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects CertRequestSubjectArrayOutput `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
}

// NewCertRequest registers a new resource with the given unique name, arguments, and options.
func NewCertRequest(ctx *pulumi.Context,
	name string, args *CertRequestArgs, opts ...pulumi.ResourceOption) (*CertRequest, error) {
	if args == nil || args.KeyAlgorithm == nil {
		return nil, errors.New("missing required argument 'KeyAlgorithm'")
	}
	if args == nil || args.PrivateKeyPem == nil {
		return nil, errors.New("missing required argument 'PrivateKeyPem'")
	}
	if args == nil || args.Subjects == nil {
		return nil, errors.New("missing required argument 'Subjects'")
	}
	if args == nil {
		args = &CertRequestArgs{}
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
	// The certificate request data in PEM format.
	CertRequestPem *string `pulumi:"certRequestPem"`
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects []CertRequestSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
}

type CertRequestState struct {
	// The certificate request data in PEM format.
	CertRequestPem pulumi.StringPtrInput
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayInput
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringPtrInput
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem pulumi.StringPtrInput
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects CertRequestSubjectArrayInput
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput
}

func (CertRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*certRequestState)(nil)).Elem()
}

type certRequestArgs struct {
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses []string `pulumi:"ipAddresses"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm string `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem string `pulumi:"privateKeyPem"`
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects []CertRequestSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
}

// The set of arguments for constructing a CertRequest resource.
type CertRequestArgs struct {
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayInput
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringInput
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem pulumi.StringInput
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects CertRequestSubjectArrayInput
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput
}

func (CertRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certRequestArgs)(nil)).Elem()
}

