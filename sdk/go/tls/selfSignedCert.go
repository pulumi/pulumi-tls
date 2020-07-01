// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SelfSignedCert struct {
	pulumi.CustomResourceState

	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayOutput `pulumi:"allowedUses"`
	// The certificate data in PEM format.
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayOutput `pulumi:"dnsNames"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrOutput `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrOutput `pulumi:"isCaCertificate"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem   pulumi.StringOutput `pulumi:"privateKeyPem"`
	ReadyForRenewal pulumi.BoolOutput   `pulumi:"readyForRenewal"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrOutput `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested.
	// This is a nested configuration block whose structure matches the
	// corresponding block for `CertRequest`.
	Subjects SelfSignedCertSubjectArrayOutput `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
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

// NewSelfSignedCert registers a new resource with the given unique name, arguments, and options.
func NewSelfSignedCert(ctx *pulumi.Context,
	name string, args *SelfSignedCertArgs, opts ...pulumi.ResourceOption) (*SelfSignedCert, error) {
	if args == nil || args.AllowedUses == nil {
		return nil, errors.New("missing required argument 'AllowedUses'")
	}
	if args == nil || args.KeyAlgorithm == nil {
		return nil, errors.New("missing required argument 'KeyAlgorithm'")
	}
	if args == nil || args.PrivateKeyPem == nil {
		return nil, errors.New("missing required argument 'PrivateKeyPem'")
	}
	if args == nil || args.Subjects == nil {
		return nil, errors.New("missing required argument 'Subjects'")
	}
	if args == nil || args.ValidityPeriodHours == nil {
		return nil, errors.New("missing required argument 'ValidityPeriodHours'")
	}
	if args == nil {
		args = &SelfSignedCertArgs{}
	}
	var resource SelfSignedCert
	err := ctx.RegisterResource("tls:index/selfSignedCert:SelfSignedCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSignedCert gets an existing SelfSignedCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSignedCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSignedCertState, opts ...pulumi.ResourceOption) (*SelfSignedCert, error) {
	var resource SelfSignedCert
	err := ctx.ReadResource("tls:index/selfSignedCert:SelfSignedCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSignedCert resources.
type selfSignedCertState struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses []string `pulumi:"allowedUses"`
	// The certificate data in PEM format.
	CertPem *string `pulumi:"certPem"`
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem   *string `pulumi:"privateKeyPem"`
	ReadyForRenewal *bool   `pulumi:"readyForRenewal"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested.
	// This is a nested configuration block whose structure matches the
	// corresponding block for `CertRequest`.
	Subjects []SelfSignedCertSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
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

type SelfSignedCertState struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayInput
	// The certificate data in PEM format.
	CertPem pulumi.StringPtrInput
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrInput
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayInput
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrInput
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringPtrInput
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem   pulumi.StringPtrInput
	ReadyForRenewal pulumi.BoolPtrInput
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrInput
	// The subject for which a certificate is being requested.
	// This is a nested configuration block whose structure matches the
	// corresponding block for `CertRequest`.
	Subjects SelfSignedCertSubjectArrayInput
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput
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

func (SelfSignedCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSignedCertState)(nil)).Elem()
}

type selfSignedCertArgs struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses []string `pulumi:"allowedUses"`
	// List of DNS names for which a certificate is being requested.
	DnsNames []string `pulumi:"dnsNames"`
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours *int `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm string `pulumi:"keyAlgorithm"`
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem string `pulumi:"privateKeyPem"`
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested.
	// This is a nested configuration block whose structure matches the
	// corresponding block for `CertRequest`.
	Subjects []SelfSignedCertSubject `pulumi:"subjects"`
	// List of URIs for which a certificate is being requested.
	Uris []string `pulumi:"uris"`
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a SelfSignedCert resource.
type SelfSignedCertArgs struct {
	// List of keywords each describing a use that is permitted
	// for the issued certificate. The valid keywords are listed below.
	AllowedUses pulumi.StringArrayInput
	// List of DNS names for which a certificate is being requested.
	DnsNames pulumi.StringArrayInput
	// Number of hours before the certificates expiry when a new certificate will be generated
	EarlyRenewalHours pulumi.IntPtrInput
	// List of IP addresses for which a certificate is being requested.
	IpAddresses pulumi.StringArrayInput
	// Boolean controlling whether the CA flag will be set in the
	// generated certificate. Defaults to `false`, meaning that the certificate does not represent
	// a certificate authority.
	IsCaCertificate pulumi.BoolPtrInput
	// The name of the algorithm for the key provided
	// in `privateKeyPem`.
	KeyAlgorithm pulumi.StringInput
	// PEM-encoded private key that the certificate will belong to
	PrivateKeyPem pulumi.StringInput
	// If `true`, the certificate will include
	// the subject key identifier. Defaults to `false`, in which case the subject
	// key identifier is not set at all.
	SetSubjectKeyId pulumi.BoolPtrInput
	// The subject for which a certificate is being requested.
	// This is a nested configuration block whose structure matches the
	// corresponding block for `CertRequest`.
	Subjects SelfSignedCertSubjectArrayInput
	// List of URIs for which a certificate is being requested.
	Uris pulumi.StringArrayInput
	// The number of hours after initial issuing that the
	// certificate will become invalid.
	ValidityPeriodHours pulumi.IntInput
}

func (SelfSignedCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSignedCertArgs)(nil)).Elem()
}
