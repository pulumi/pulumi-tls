// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateKey struct {
	pulumi.CustomResourceState

	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve pulumi.StringOutput `pulumi:"ecdsaCurve"`
	// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
	PrivateKeyOpenssh pulumi.StringOutput `pulumi:"privateKeyOpenssh"`
	// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
	PrivateKeyPemPkcs8 pulumi.StringOutput `pulumi:"privateKeyPemPkcs8"`
	// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintMd5 pulumi.StringOutput `pulumi:"publicKeyFingerprintMd5"`
	// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintSha256 pulumi.StringOutput `pulumi:"publicKeyFingerprintSha256"`
	// The public key data in ["Authorized
	// Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not
	// populated for `ECDSA` with curve `P224`, as it is [not supported](../../docs#limitations). **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyOpenssh pulumi.StringOutput `pulumi:"publicKeyOpenssh"`
	// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyPem pulumi.StringOutput `pulumi:"publicKeyPem"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntOutput `pulumi:"rsaBits"`
}

// NewPrivateKey registers a new resource with the given unique name, arguments, and options.
func NewPrivateKey(ctx *pulumi.Context,
	name string, args *PrivateKeyArgs, opts ...pulumi.ResourceOption) (*PrivateKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKeyOpenssh",
		"privateKeyPem",
		"privateKeyPemPkcs8",
	})
	opts = append(opts, secrets)
	var resource PrivateKey
	err := ctx.RegisterResource("tls:index/privateKey:PrivateKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateKey gets an existing PrivateKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateKeyState, opts ...pulumi.ResourceOption) (*PrivateKey, error) {
	var resource PrivateKey
	err := ctx.ReadResource("tls:index/privateKey:PrivateKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateKey resources.
type privateKeyState struct {
	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm *string `pulumi:"algorithm"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve *string `pulumi:"ecdsaCurve"`
	// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
	PrivateKeyOpenssh *string `pulumi:"privateKeyOpenssh"`
	// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
	PrivateKeyPemPkcs8 *string `pulumi:"privateKeyPemPkcs8"`
	// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintMd5 *string `pulumi:"publicKeyFingerprintMd5"`
	// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintSha256 *string `pulumi:"publicKeyFingerprintSha256"`
	// The public key data in ["Authorized
	// Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not
	// populated for `ECDSA` with curve `P224`, as it is [not supported](../../docs#limitations). **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyOpenssh *string `pulumi:"publicKeyOpenssh"`
	// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyPem *string `pulumi:"publicKeyPem"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits *int `pulumi:"rsaBits"`
}

type PrivateKeyState struct {
	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm pulumi.StringPtrInput
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve pulumi.StringPtrInput
	// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
	PrivateKeyOpenssh pulumi.StringPtrInput
	// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
	PrivateKeyPem pulumi.StringPtrInput
	// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
	PrivateKeyPemPkcs8 pulumi.StringPtrInput
	// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintMd5 pulumi.StringPtrInput
	// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
	PublicKeyFingerprintSha256 pulumi.StringPtrInput
	// The public key data in ["Authorized
	// Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not
	// populated for `ECDSA` with curve `P224`, as it is [not supported](../../docs#limitations). **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyOpenssh pulumi.StringPtrInput
	// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
	// [underlying](https://pkg.go.dev/encoding/pem#Encode)
	// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
	// the end of the PEM. In case this disrupts your use case, we recommend using
	// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
	PublicKeyPem pulumi.StringPtrInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntPtrInput
}

func (PrivateKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateKeyState)(nil)).Elem()
}

type privateKeyArgs struct {
	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm string `pulumi:"algorithm"`
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve *string `pulumi:"ecdsaCurve"`
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits *int `pulumi:"rsaBits"`
}

// The set of arguments for constructing a PrivateKey resource.
type PrivateKeyArgs struct {
	// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
	Algorithm pulumi.StringInput
	// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
	EcdsaCurve pulumi.StringPtrInput
	// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
	RsaBits pulumi.IntPtrInput
}

func (PrivateKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateKeyArgs)(nil)).Elem()
}

type PrivateKeyInput interface {
	pulumi.Input

	ToPrivateKeyOutput() PrivateKeyOutput
	ToPrivateKeyOutputWithContext(ctx context.Context) PrivateKeyOutput
}

func (*PrivateKey) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateKey)(nil)).Elem()
}

func (i *PrivateKey) ToPrivateKeyOutput() PrivateKeyOutput {
	return i.ToPrivateKeyOutputWithContext(context.Background())
}

func (i *PrivateKey) ToPrivateKeyOutputWithContext(ctx context.Context) PrivateKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateKeyOutput)
}

// PrivateKeyArrayInput is an input type that accepts PrivateKeyArray and PrivateKeyArrayOutput values.
// You can construct a concrete instance of `PrivateKeyArrayInput` via:
//
//	PrivateKeyArray{ PrivateKeyArgs{...} }
type PrivateKeyArrayInput interface {
	pulumi.Input

	ToPrivateKeyArrayOutput() PrivateKeyArrayOutput
	ToPrivateKeyArrayOutputWithContext(context.Context) PrivateKeyArrayOutput
}

type PrivateKeyArray []PrivateKeyInput

func (PrivateKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateKey)(nil)).Elem()
}

func (i PrivateKeyArray) ToPrivateKeyArrayOutput() PrivateKeyArrayOutput {
	return i.ToPrivateKeyArrayOutputWithContext(context.Background())
}

func (i PrivateKeyArray) ToPrivateKeyArrayOutputWithContext(ctx context.Context) PrivateKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateKeyArrayOutput)
}

// PrivateKeyMapInput is an input type that accepts PrivateKeyMap and PrivateKeyMapOutput values.
// You can construct a concrete instance of `PrivateKeyMapInput` via:
//
//	PrivateKeyMap{ "key": PrivateKeyArgs{...} }
type PrivateKeyMapInput interface {
	pulumi.Input

	ToPrivateKeyMapOutput() PrivateKeyMapOutput
	ToPrivateKeyMapOutputWithContext(context.Context) PrivateKeyMapOutput
}

type PrivateKeyMap map[string]PrivateKeyInput

func (PrivateKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateKey)(nil)).Elem()
}

func (i PrivateKeyMap) ToPrivateKeyMapOutput() PrivateKeyMapOutput {
	return i.ToPrivateKeyMapOutputWithContext(context.Background())
}

func (i PrivateKeyMap) ToPrivateKeyMapOutputWithContext(ctx context.Context) PrivateKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateKeyMapOutput)
}

type PrivateKeyOutput struct{ *pulumi.OutputState }

func (PrivateKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateKey)(nil)).Elem()
}

func (o PrivateKeyOutput) ToPrivateKeyOutput() PrivateKeyOutput {
	return o
}

func (o PrivateKeyOutput) ToPrivateKeyOutputWithContext(ctx context.Context) PrivateKeyOutput {
	return o
}

// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
func (o PrivateKeyOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
func (o PrivateKeyOutput) EcdsaCurve() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.EcdsaCurve }).(pulumi.StringOutput)
}

// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
func (o PrivateKeyOutput) PrivateKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PrivateKeyOpenssh }).(pulumi.StringOutput)
}

// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
func (o PrivateKeyOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
func (o PrivateKeyOutput) PrivateKeyPemPkcs8() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PrivateKeyPemPkcs8 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
func (o PrivateKeyOutput) PublicKeyFingerprintMd5() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PublicKeyFingerprintMd5 }).(pulumi.StringOutput)
}

// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `publicKeyOpenssh` and the ECDSA P224 limitations.
func (o PrivateKeyOutput) PublicKeyFingerprintSha256() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PublicKeyFingerprintSha256 }).(pulumi.StringOutput)
}

// The public key data in ["Authorized
// Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not
// populated for `ECDSA` with curve `P224`, as it is [not supported](../../docs#limitations). **NOTE**: the
// [underlying](https://pkg.go.dev/encoding/pem#Encode)
// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
// the end of the PEM. In case this disrupts your use case, we recommend using
// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
func (o PrivateKeyOutput) PublicKeyOpenssh() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PublicKeyOpenssh }).(pulumi.StringOutput)
}

// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
// [underlying](https://pkg.go.dev/encoding/pem#Encode)
// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
// the end of the PEM. In case this disrupts your use case, we recommend using
// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
func (o PrivateKeyOutput) PublicKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.StringOutput { return v.PublicKeyPem }).(pulumi.StringOutput)
}

// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
func (o PrivateKeyOutput) RsaBits() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateKey) pulumi.IntOutput { return v.RsaBits }).(pulumi.IntOutput)
}

type PrivateKeyArrayOutput struct{ *pulumi.OutputState }

func (PrivateKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateKey)(nil)).Elem()
}

func (o PrivateKeyArrayOutput) ToPrivateKeyArrayOutput() PrivateKeyArrayOutput {
	return o
}

func (o PrivateKeyArrayOutput) ToPrivateKeyArrayOutputWithContext(ctx context.Context) PrivateKeyArrayOutput {
	return o
}

func (o PrivateKeyArrayOutput) Index(i pulumi.IntInput) PrivateKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateKey {
		return vs[0].([]*PrivateKey)[vs[1].(int)]
	}).(PrivateKeyOutput)
}

type PrivateKeyMapOutput struct{ *pulumi.OutputState }

func (PrivateKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateKey)(nil)).Elem()
}

func (o PrivateKeyMapOutput) ToPrivateKeyMapOutput() PrivateKeyMapOutput {
	return o
}

func (o PrivateKeyMapOutput) ToPrivateKeyMapOutputWithContext(ctx context.Context) PrivateKeyMapOutput {
	return o
}

func (o PrivateKeyMapOutput) MapIndex(k pulumi.StringInput) PrivateKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateKey {
		return vs[0].(map[string]*PrivateKey)[vs[1].(string)]
	}).(PrivateKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateKeyInput)(nil)).Elem(), &PrivateKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateKeyArrayInput)(nil)).Elem(), PrivateKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateKeyMapInput)(nil)).Elem(), PrivateKeyMap{})
	pulumi.RegisterOutputType(PrivateKeyOutput{})
	pulumi.RegisterOutputType(PrivateKeyArrayOutput{})
	pulumi.RegisterOutputType(PrivateKeyMapOutput{})
}
