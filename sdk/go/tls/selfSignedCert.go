// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SelfSignedCert struct {
	pulumi.CustomResourceState

	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayOutput `pulumi:"allowedUses"`
	// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
	CertPem pulumi.StringOutput `pulumi:"certPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayOutput `pulumi:"dnsNames"`
	EarlyRenewalHours pulumi.IntOutput         `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayOutput `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolOutput `pulumi:"isCaCertificate"`
	// Name of the algorithm used when generating the private key provided in `privateKeyPem`.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal pulumi.BoolOutput `pulumi:"readyForRenewal"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolOutput `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolOutput `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject SelfSignedCertSubjectPtrOutput `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayOutput `pulumi:"uris"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringOutput `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntOutput `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringOutput `pulumi:"validityStartTime"`
}

// NewSelfSignedCert registers a new resource with the given unique name, arguments, and options.
func NewSelfSignedCert(ctx *pulumi.Context,
	name string, args *SelfSignedCertArgs, opts ...pulumi.ResourceOption) (*SelfSignedCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedUses == nil {
		return nil, errors.New("invalid value for required argument 'AllowedUses'")
	}
	if args.PrivateKeyPem == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKeyPem'")
	}
	if args.ValidityPeriodHours == nil {
		return nil, errors.New("invalid value for required argument 'ValidityPeriodHours'")
	}
	if args.PrivateKeyPem != nil {
		args.PrivateKeyPem = pulumi.ToSecret(args.PrivateKeyPem).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKeyPem",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses []string `pulumi:"allowedUses"`
	// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
	CertPem *string `pulumi:"certPem"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          []string `pulumi:"dnsNames"`
	EarlyRenewalHours *int     `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// Name of the algorithm used when generating the private key provided in `privateKeyPem`.
	KeyAlgorithm *string `pulumi:"keyAlgorithm"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal *bool `pulumi:"readyForRenewal"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject *SelfSignedCertSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime *string `pulumi:"validityEndTime"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours *int `pulumi:"validityPeriodHours"`
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime *string `pulumi:"validityStartTime"`
}

type SelfSignedCertState struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayInput
	// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
	CertPem pulumi.StringPtrInput
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayInput
	EarlyRenewalHours pulumi.IntPtrInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrInput
	// Name of the algorithm used when generating the private key provided in `privateKeyPem`.
	KeyAlgorithm pulumi.StringPtrInput
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
	PrivateKeyPem pulumi.StringPtrInput
	// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
	ReadyForRenewal pulumi.BoolPtrInput
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject SelfSignedCertSubjectPtrInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
	// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityEndTime pulumi.StringPtrInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntPtrInput
	// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	ValidityStartTime pulumi.StringPtrInput
}

func (SelfSignedCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSignedCertState)(nil)).Elem()
}

type selfSignedCertArgs struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses []string `pulumi:"allowedUses"`
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          []string `pulumi:"dnsNames"`
	EarlyRenewalHours *int     `pulumi:"earlyRenewalHours"`
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses []string `pulumi:"ipAddresses"`
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate *bool `pulumi:"isCaCertificate"`
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
	PrivateKeyPem string `pulumi:"privateKeyPem"`
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId *bool `pulumi:"setAuthorityKeyId"`
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId *bool `pulumi:"setSubjectKeyId"`
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject *SelfSignedCertSubject `pulumi:"subject"`
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris []string `pulumi:"uris"`
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours int `pulumi:"validityPeriodHours"`
}

// The set of arguments for constructing a SelfSignedCert resource.
type SelfSignedCertArgs struct {
	// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
	AllowedUses pulumi.StringArrayInput
	// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
	DnsNames          pulumi.StringArrayInput
	EarlyRenewalHours pulumi.IntPtrInput
	// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
	IpAddresses pulumi.StringArrayInput
	// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
	IsCaCertificate pulumi.BoolPtrInput
	// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
	PrivateKeyPem pulumi.StringInput
	// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetAuthorityKeyId pulumi.BoolPtrInput
	// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
	SetSubjectKeyId pulumi.BoolPtrInput
	// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
	Subject SelfSignedCertSubjectPtrInput
	// List of URIs for which a certificate is being requested (i.e. certificate subjects).
	Uris pulumi.StringArrayInput
	// Number of hours, after initial issuing, that the certificate will remain valid for.
	ValidityPeriodHours pulumi.IntInput
}

func (SelfSignedCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSignedCertArgs)(nil)).Elem()
}

type SelfSignedCertInput interface {
	pulumi.Input

	ToSelfSignedCertOutput() SelfSignedCertOutput
	ToSelfSignedCertOutputWithContext(ctx context.Context) SelfSignedCertOutput
}

func (*SelfSignedCert) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSignedCert)(nil)).Elem()
}

func (i *SelfSignedCert) ToSelfSignedCertOutput() SelfSignedCertOutput {
	return i.ToSelfSignedCertOutputWithContext(context.Background())
}

func (i *SelfSignedCert) ToSelfSignedCertOutputWithContext(ctx context.Context) SelfSignedCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertOutput)
}

// SelfSignedCertArrayInput is an input type that accepts SelfSignedCertArray and SelfSignedCertArrayOutput values.
// You can construct a concrete instance of `SelfSignedCertArrayInput` via:
//
//	SelfSignedCertArray{ SelfSignedCertArgs{...} }
type SelfSignedCertArrayInput interface {
	pulumi.Input

	ToSelfSignedCertArrayOutput() SelfSignedCertArrayOutput
	ToSelfSignedCertArrayOutputWithContext(context.Context) SelfSignedCertArrayOutput
}

type SelfSignedCertArray []SelfSignedCertInput

func (SelfSignedCertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSignedCert)(nil)).Elem()
}

func (i SelfSignedCertArray) ToSelfSignedCertArrayOutput() SelfSignedCertArrayOutput {
	return i.ToSelfSignedCertArrayOutputWithContext(context.Background())
}

func (i SelfSignedCertArray) ToSelfSignedCertArrayOutputWithContext(ctx context.Context) SelfSignedCertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertArrayOutput)
}

// SelfSignedCertMapInput is an input type that accepts SelfSignedCertMap and SelfSignedCertMapOutput values.
// You can construct a concrete instance of `SelfSignedCertMapInput` via:
//
//	SelfSignedCertMap{ "key": SelfSignedCertArgs{...} }
type SelfSignedCertMapInput interface {
	pulumi.Input

	ToSelfSignedCertMapOutput() SelfSignedCertMapOutput
	ToSelfSignedCertMapOutputWithContext(context.Context) SelfSignedCertMapOutput
}

type SelfSignedCertMap map[string]SelfSignedCertInput

func (SelfSignedCertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSignedCert)(nil)).Elem()
}

func (i SelfSignedCertMap) ToSelfSignedCertMapOutput() SelfSignedCertMapOutput {
	return i.ToSelfSignedCertMapOutputWithContext(context.Background())
}

func (i SelfSignedCertMap) ToSelfSignedCertMapOutputWithContext(ctx context.Context) SelfSignedCertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertMapOutput)
}

type SelfSignedCertOutput struct{ *pulumi.OutputState }

func (SelfSignedCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSignedCert)(nil)).Elem()
}

func (o SelfSignedCertOutput) ToSelfSignedCertOutput() SelfSignedCertOutput {
	return o
}

func (o SelfSignedCertOutput) ToSelfSignedCertOutputWithContext(ctx context.Context) SelfSignedCertOutput {
	return o
}

// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
func (o SelfSignedCertOutput) AllowedUses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringArrayOutput { return v.AllowedUses }).(pulumi.StringArrayOutput)
}

// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
func (o SelfSignedCertOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringOutput { return v.CertPem }).(pulumi.StringOutput)
}

// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
func (o SelfSignedCertOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringArrayOutput { return v.DnsNames }).(pulumi.StringArrayOutput)
}

func (o SelfSignedCertOutput) EarlyRenewalHours() pulumi.IntOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.IntOutput { return v.EarlyRenewalHours }).(pulumi.IntOutput)
}

// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
func (o SelfSignedCertOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
func (o SelfSignedCertOutput) IsCaCertificate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.BoolOutput { return v.IsCaCertificate }).(pulumi.BoolOutput)
}

// Name of the algorithm used when generating the private key provided in `privateKeyPem`.
func (o SelfSignedCertOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
func (o SelfSignedCertOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

// Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
func (o SelfSignedCertOutput) ReadyForRenewal() pulumi.BoolOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.BoolOutput { return v.ReadyForRenewal }).(pulumi.BoolOutput)
}

// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o SelfSignedCertOutput) SetAuthorityKeyId() pulumi.BoolOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.BoolOutput { return v.SetAuthorityKeyId }).(pulumi.BoolOutput)
}

// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
func (o SelfSignedCertOutput) SetSubjectKeyId() pulumi.BoolOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.BoolOutput { return v.SetSubjectKeyId }).(pulumi.BoolOutput)
}

// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
func (o SelfSignedCertOutput) Subject() SelfSignedCertSubjectPtrOutput {
	return o.ApplyT(func(v *SelfSignedCert) SelfSignedCertSubjectPtrOutput { return v.Subject }).(SelfSignedCertSubjectPtrOutput)
}

// List of URIs for which a certificate is being requested (i.e. certificate subjects).
func (o SelfSignedCertOutput) Uris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringArrayOutput { return v.Uris }).(pulumi.StringArrayOutput)
}

// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o SelfSignedCertOutput) ValidityEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringOutput { return v.ValidityEndTime }).(pulumi.StringOutput)
}

// Number of hours, after initial issuing, that the certificate will remain valid for.
func (o SelfSignedCertOutput) ValidityPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.IntOutput { return v.ValidityPeriodHours }).(pulumi.IntOutput)
}

// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o SelfSignedCertOutput) ValidityStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SelfSignedCert) pulumi.StringOutput { return v.ValidityStartTime }).(pulumi.StringOutput)
}

type SelfSignedCertArrayOutput struct{ *pulumi.OutputState }

func (SelfSignedCertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSignedCert)(nil)).Elem()
}

func (o SelfSignedCertArrayOutput) ToSelfSignedCertArrayOutput() SelfSignedCertArrayOutput {
	return o
}

func (o SelfSignedCertArrayOutput) ToSelfSignedCertArrayOutputWithContext(ctx context.Context) SelfSignedCertArrayOutput {
	return o
}

func (o SelfSignedCertArrayOutput) Index(i pulumi.IntInput) SelfSignedCertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SelfSignedCert {
		return vs[0].([]*SelfSignedCert)[vs[1].(int)]
	}).(SelfSignedCertOutput)
}

type SelfSignedCertMapOutput struct{ *pulumi.OutputState }

func (SelfSignedCertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSignedCert)(nil)).Elem()
}

func (o SelfSignedCertMapOutput) ToSelfSignedCertMapOutput() SelfSignedCertMapOutput {
	return o
}

func (o SelfSignedCertMapOutput) ToSelfSignedCertMapOutputWithContext(ctx context.Context) SelfSignedCertMapOutput {
	return o
}

func (o SelfSignedCertMapOutput) MapIndex(k pulumi.StringInput) SelfSignedCertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SelfSignedCert {
		return vs[0].(map[string]*SelfSignedCert)[vs[1].(string)]
	}).(SelfSignedCertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSignedCertInput)(nil)).Elem(), &SelfSignedCert{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSignedCertArrayInput)(nil)).Elem(), SelfSignedCertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSignedCertMapInput)(nil)).Elem(), SelfSignedCertMap{})
	pulumi.RegisterOutputType(SelfSignedCertOutput{})
	pulumi.RegisterOutputType(SelfSignedCertArrayOutput{})
	pulumi.RegisterOutputType(SelfSignedCertMapOutput{})
}
