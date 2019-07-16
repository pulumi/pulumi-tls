// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-tls/blob/master/website/docs/r/cert_request.html.markdown.
type CertRequest struct {
	s *pulumi.ResourceState
}

// NewCertRequest registers a new resource with the given unique name, arguments, and options.
func NewCertRequest(ctx *pulumi.Context,
	name string, args *CertRequestArgs, opts ...pulumi.ResourceOpt) (*CertRequest, error) {
	if args == nil || args.KeyAlgorithm == nil {
		return nil, errors.New("missing required argument 'KeyAlgorithm'")
	}
	if args == nil || args.PrivateKeyPem == nil {
		return nil, errors.New("missing required argument 'PrivateKeyPem'")
	}
	if args == nil || args.Subjects == nil {
		return nil, errors.New("missing required argument 'Subjects'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dnsNames"] = nil
		inputs["ipAddresses"] = nil
		inputs["keyAlgorithm"] = nil
		inputs["privateKeyPem"] = nil
		inputs["subjects"] = nil
	} else {
		inputs["dnsNames"] = args.DnsNames
		inputs["ipAddresses"] = args.IpAddresses
		inputs["keyAlgorithm"] = args.KeyAlgorithm
		inputs["privateKeyPem"] = args.PrivateKeyPem
		inputs["subjects"] = args.Subjects
	}
	inputs["certRequestPem"] = nil
	s, err := ctx.RegisterResource("tls:index/certRequest:CertRequest", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CertRequest{s: s}, nil
}

// GetCertRequest gets an existing CertRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertRequest(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CertRequestState, opts ...pulumi.ResourceOpt) (*CertRequest, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["certRequestPem"] = state.CertRequestPem
		inputs["dnsNames"] = state.DnsNames
		inputs["ipAddresses"] = state.IpAddresses
		inputs["keyAlgorithm"] = state.KeyAlgorithm
		inputs["privateKeyPem"] = state.PrivateKeyPem
		inputs["subjects"] = state.Subjects
	}
	s, err := ctx.ReadResource("tls:index/certRequest:CertRequest", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &CertRequest{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *CertRequest) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *CertRequest) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The certificate request data in PEM format.
func (r *CertRequest) CertRequestPem() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["certRequestPem"])
}

// List of DNS names for which a certificate is being requested.
func (r *CertRequest) DnsNames() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsNames"])
}

// List of IP addresses for which a certificate is being requested.
func (r *CertRequest) IpAddresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipAddresses"])
}

// The name of the algorithm for the key provided
// in `private_key_pem`.
func (r *CertRequest) KeyAlgorithm() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyAlgorithm"])
}

// PEM-encoded private key that the certificate will belong to
func (r *CertRequest) PrivateKeyPem() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["privateKeyPem"])
}

// The subject for which a certificate is being requested. This is
// a nested configuration block whose structure is described below.
func (r *CertRequest) Subjects() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["subjects"])
}

// Input properties used for looking up and filtering CertRequest resources.
type CertRequestState struct {
	// The certificate request data in PEM format.
	CertRequestPem interface{}
	// List of DNS names for which a certificate is being requested.
	DnsNames interface{}
	// List of IP addresses for which a certificate is being requested.
	IpAddresses interface{}
	// The name of the algorithm for the key provided
	// in `private_key_pem`.
	KeyAlgorithm interface{}
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem interface{}
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects interface{}
}

// The set of arguments for constructing a CertRequest resource.
type CertRequestArgs struct {
	// List of DNS names for which a certificate is being requested.
	DnsNames interface{}
	// List of IP addresses for which a certificate is being requested.
	IpAddresses interface{}
	// The name of the algorithm for the key provided
	// in `private_key_pem`.
	KeyAlgorithm interface{}
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem interface{}
	// The subject for which a certificate is being requested. This is
	// a nested configuration block whose structure is described below.
	Subjects interface{}
}
