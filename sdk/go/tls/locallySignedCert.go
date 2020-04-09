// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package tls

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LocallySignedCert struct {
	pulumi.CustomResourceState

	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayOutput `pulumi:"allowedUses"`
	// PEM-encoded certificate data for the CA.
	CaCertPem pulumi.StringOutput `pulumi:"caCertPem"`
	// The name of the algorithm for the key provided
	// in `caPrivateKeyPem`.
	CaKeyAlgorithm pulumi.StringOutput `pulumi:"caKeyAlgorithm"`
	// PEM-encoded private key data for the CA.
	// This can be read from a separate file using the ``file`` interpolation
	// function.
	CaPrivateKeyPem pulumi.StringOutput `pulumi:"caPrivateKeyPem"`
	// The certificate data in PEM format.
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// PEM-encoded request certificate data.
	CertRequestPem pulumi.StringOutput `pulumi:"certRequestPem"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrOutput `pulumi:"earlyRenewalHours"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrOutput `pulumi:"isCaCertificate"`
	ReadyForRenewal pulumi.BoolOutput    `pulumi:"readyForRenewal"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrOutput `pulumi:"setSubjectKeyId"`
	// The time until which the certificate is invalid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringOutput `pulumi:"validityEndTime"`
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringOutput `pulumi:"validityStartTime"`
}

// NewLocallySignedCert registers a new resource with the given unique name, arguments, and options.
func NewLocallySignedCert(ctx *pulumi.Context,
	name string, args *LocallySignedCertArgs, opts ...pulumi.ResourceOption) (*LocallySignedCert, error) {
	if args == nil || args.AllowedUses == nil {
		return nil, errors.New("missing required argument 'AllowedUses'")
	}
	if args == nil || args.CaCertPem == nil {
		return nil, errors.New("missing required argument 'CaCertPem'")
	}
	if args == nil || args.CaKeyAlgorithm == nil {
		return nil, errors.New("missing required argument 'CaKeyAlgorithm'")
	}
	if args == nil || args.CaPrivateKeyPem == nil {
		return nil, errors.New("missing required argument 'CaPrivateKeyPem'")
	}
	if args == nil || args.CertRequestPem == nil {
		return nil, errors.New("missing required argument 'CertRequestPem'")
	}
	if args == nil || args.ValidityPeriodHours == nil {
		return nil, errors.New("missing required argument 'ValidityPeriodHours'")
	}
	if args == nil {
		args = &LocallySignedCertArgs{}
	}
	var resource LocallySignedCert
	err := ctx.RegisterResource("tls:index/locallySignedCert:LocallySignedCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocallySignedCert gets an existing LocallySignedCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocallySignedCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocallySignedCertState, opts ...pulumi.ResourceOption) (*LocallySignedCert, error) {
	var resource LocallySignedCert
	err := ctx.ReadResource("tls:index/locallySignedCert:LocallySignedCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocallySignedCert resources.
type locallySignedCertState struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses []string `pulumi:"allowedUses"`
	// PEM-encoded certificate data for the CA.
	CaCertPem *string `pulumi:"caCertPem"`
	// The name of the algorithm for the key provided
	// in `caPrivateKeyPem`.
	CaKeyAlgorithm *string `pulumi:"caKeyAlgorithm"`
	// PEM-encoded private key data for the CA.
	// This can be read from a separate file using the ``file`` interpolation
	// function.
	CaPrivateKeyPem *string `pulumi:"caPrivateKeyPem"`
	// The certificate data in PEM format.
	CertPem *string `pulumi:"certPem"`
	// PEM-encoded request certificate data.
	CertRequestPem *string `pulumi:"certRequestPem"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	ReadyForRenewal *bool `pulumi:"readyForRenewal"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The time until which the certificate is invalid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime *string `pulumi:"validityEndTime"`
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours *int `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime *string `pulumi:"validityStartTime"`
}

type LocallySignedCertState struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayInput
	// PEM-encoded certificate data for the CA.
	CaCertPem pulumi.StringPtrInput
	// The name of the algorithm for the key provided
	// in `caPrivateKeyPem`.
	CaKeyAlgorithm pulumi.StringPtrInput
	// PEM-encoded private key data for the CA.
	// This can be read from a separate file using the ``file`` interpolation
	// function.
	CaPrivateKeyPem pulumi.StringPtrInput
	// The certificate data in PEM format.
	CertPem pulumi.StringPtrInput
	// PEM-encoded request certificate data.
	CertRequestPem pulumi.StringPtrInput
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrInput
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrInput
	ReadyForRenewal pulumi.BoolPtrInput
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrInput
	// The time until which the certificate is invalid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringPtrInput
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours pulumi.IntPtrInput
	// The time after which the certificate is valid, as an
	// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringPtrInput
}

func (LocallySignedCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*locallySignedCertState)(nil)).Elem()
}

type locallySignedCertArgs struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses []string `pulumi:"allowedUses"`
	// PEM-encoded certificate data for the CA.
	CaCertPem string `pulumi:"caCertPem"`
	// The name of the algorithm for the key provided
	// in `caPrivateKeyPem`.
	CaKeyAlgorithm string `pulumi:"caKeyAlgorithm"`
	// PEM-encoded private key data for the CA.
	// This can be read from a separate file using the ``file`` interpolation
	// function.
	CaPrivateKeyPem string `pulumi:"caPrivateKeyPem"`
	// PEM-encoded request certificate data.
	CertRequestPem string `pulumi:"certRequestPem"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a LocallySignedCert resource.
type LocallySignedCertArgs struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayInput
	// PEM-encoded certificate data for the CA.
	CaCertPem pulumi.StringInput
	// The name of the algorithm for the key provided
	// in `caPrivateKeyPem`.
	CaKeyAlgorithm pulumi.StringInput
	// PEM-encoded private key data for the CA.
	// This can be read from a separate file using the ``file`` interpolation
	// function.
	CaPrivateKeyPem pulumi.StringInput
	// PEM-encoded request certificate data.
	CertRequestPem pulumi.StringInput
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrInput
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrInput
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrInput
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours pulumi.IntInput
}

func (LocallySignedCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locallySignedCertArgs)(nil)).Elem()
}
