// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type CertRequestSubject struct {
	// Distinguished name: `CN`
	CommonName *string `pulumi:"commonName"`
	// Distinguished name: `C`
	Country *string `pulumi:"country"`
	// Distinguished name: `L`
	Locality *string `pulumi:"locality"`
	// Distinguished name: `O`
	Organization *string `pulumi:"organization"`
	// Distinguished name: `OU`
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	// Distinguished name: `PC`
	PostalCode *string `pulumi:"postalCode"`
	// Distinguished name: `ST`
	Province *string `pulumi:"province"`
	// Distinguished name: `SERIALNUMBER`
	SerialNumber *string `pulumi:"serialNumber"`
	// Distinguished name: `STREET`
	StreetAddresses []string `pulumi:"streetAddresses"`
}

// CertRequestSubjectInput is an input type that accepts CertRequestSubjectArgs and CertRequestSubjectOutput values.
// You can construct a concrete instance of `CertRequestSubjectInput` via:
//
//	CertRequestSubjectArgs{...}
type CertRequestSubjectInput interface {
	pulumi.Input

	ToCertRequestSubjectOutput() CertRequestSubjectOutput
	ToCertRequestSubjectOutputWithContext(context.Context) CertRequestSubjectOutput
}

type CertRequestSubjectArgs struct {
	// Distinguished name: `CN`
	CommonName pulumi.StringPtrInput `pulumi:"commonName"`
	// Distinguished name: `C`
	Country pulumi.StringPtrInput `pulumi:"country"`
	// Distinguished name: `L`
	Locality pulumi.StringPtrInput `pulumi:"locality"`
	// Distinguished name: `O`
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// Distinguished name: `OU`
	OrganizationalUnit pulumi.StringPtrInput `pulumi:"organizationalUnit"`
	// Distinguished name: `PC`
	PostalCode pulumi.StringPtrInput `pulumi:"postalCode"`
	// Distinguished name: `ST`
	Province pulumi.StringPtrInput `pulumi:"province"`
	// Distinguished name: `SERIALNUMBER`
	SerialNumber pulumi.StringPtrInput `pulumi:"serialNumber"`
	// Distinguished name: `STREET`
	StreetAddresses pulumi.StringArrayInput `pulumi:"streetAddresses"`
}

func (CertRequestSubjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertRequestSubject)(nil)).Elem()
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectOutput() CertRequestSubjectOutput {
	return i.ToCertRequestSubjectOutputWithContext(context.Background())
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectOutputWithContext(ctx context.Context) CertRequestSubjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectOutput)
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return i.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (i CertRequestSubjectArgs) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectOutput).ToCertRequestSubjectPtrOutputWithContext(ctx)
}

// CertRequestSubjectPtrInput is an input type that accepts CertRequestSubjectArgs, CertRequestSubjectPtr and CertRequestSubjectPtrOutput values.
// You can construct a concrete instance of `CertRequestSubjectPtrInput` via:
//
//	        CertRequestSubjectArgs{...}
//
//	or:
//
//	        nil
type CertRequestSubjectPtrInput interface {
	pulumi.Input

	ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput
	ToCertRequestSubjectPtrOutputWithContext(context.Context) CertRequestSubjectPtrOutput
}

type certRequestSubjectPtrType CertRequestSubjectArgs

func CertRequestSubjectPtr(v *CertRequestSubjectArgs) CertRequestSubjectPtrInput {
	return (*certRequestSubjectPtrType)(v)
}

func (*certRequestSubjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequestSubject)(nil)).Elem()
}

func (i *certRequestSubjectPtrType) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return i.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (i *certRequestSubjectPtrType) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertRequestSubjectPtrOutput)
}

type CertRequestSubjectOutput struct{ *pulumi.OutputState }

func (CertRequestSubjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertRequestSubject)(nil)).Elem()
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectOutput() CertRequestSubjectOutput {
	return o
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectOutputWithContext(ctx context.Context) CertRequestSubjectOutput {
	return o
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return o.ToCertRequestSubjectPtrOutputWithContext(context.Background())
}

func (o CertRequestSubjectOutput) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertRequestSubject) *CertRequestSubject {
		return &v
	}).(CertRequestSubjectPtrOutput)
}

// Distinguished name: `CN`
func (o CertRequestSubjectOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

// Distinguished name: `C`
func (o CertRequestSubjectOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Country }).(pulumi.StringPtrOutput)
}

// Distinguished name: `L`
func (o CertRequestSubjectOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Locality }).(pulumi.StringPtrOutput)
}

// Distinguished name: `O`
func (o CertRequestSubjectOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// Distinguished name: `OU`
func (o CertRequestSubjectOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

// Distinguished name: `PC`
func (o CertRequestSubjectOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

// Distinguished name: `ST`
func (o CertRequestSubjectOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.Province }).(pulumi.StringPtrOutput)
}

// Distinguished name: `SERIALNUMBER`
func (o CertRequestSubjectOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertRequestSubject) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// Distinguished name: `STREET`
func (o CertRequestSubjectOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertRequestSubject) []string { return v.StreetAddresses }).(pulumi.StringArrayOutput)
}

type CertRequestSubjectPtrOutput struct{ *pulumi.OutputState }

func (CertRequestSubjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertRequestSubject)(nil)).Elem()
}

func (o CertRequestSubjectPtrOutput) ToCertRequestSubjectPtrOutput() CertRequestSubjectPtrOutput {
	return o
}

func (o CertRequestSubjectPtrOutput) ToCertRequestSubjectPtrOutputWithContext(ctx context.Context) CertRequestSubjectPtrOutput {
	return o
}

func (o CertRequestSubjectPtrOutput) Elem() CertRequestSubjectOutput {
	return o.ApplyT(func(v *CertRequestSubject) CertRequestSubject {
		if v != nil {
			return *v
		}
		var ret CertRequestSubject
		return ret
	}).(CertRequestSubjectOutput)
}

// Distinguished name: `CN`
func (o CertRequestSubjectPtrOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.CommonName
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `C`
func (o CertRequestSubjectPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Country
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `L`
func (o CertRequestSubjectPtrOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Locality
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `O`
func (o CertRequestSubjectPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `OU`
func (o CertRequestSubjectPtrOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnit
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `PC`
func (o CertRequestSubjectPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `ST`
func (o CertRequestSubjectPtrOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.Province
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `SERIALNUMBER`
func (o CertRequestSubjectPtrOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertRequestSubject) *string {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `STREET`
func (o CertRequestSubjectPtrOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertRequestSubject) []string {
		if v == nil {
			return nil
		}
		return v.StreetAddresses
	}).(pulumi.StringArrayOutput)
}

type ProviderProxy struct {
	// When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
	FromEnv *bool `pulumi:"fromEnv"`
	// Password used for Basic authentication against the Proxy.
	Password *string `pulumi:"password"`
	// URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
	Url *string `pulumi:"url"`
	// Username (or Token) used for Basic authentication against the Proxy.
	Username *string `pulumi:"username"`
}

// ProviderProxyInput is an input type that accepts ProviderProxyArgs and ProviderProxyOutput values.
// You can construct a concrete instance of `ProviderProxyInput` via:
//
//	ProviderProxyArgs{...}
type ProviderProxyInput interface {
	pulumi.Input

	ToProviderProxyOutput() ProviderProxyOutput
	ToProviderProxyOutputWithContext(context.Context) ProviderProxyOutput
}

type ProviderProxyArgs struct {
	// When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
	FromEnv pulumi.BoolPtrInput `pulumi:"fromEnv"`
	// Password used for Basic authentication against the Proxy.
	Password pulumi.StringPtrInput `pulumi:"password"`
	// URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
	Url pulumi.StringPtrInput `pulumi:"url"`
	// Username (or Token) used for Basic authentication against the Proxy.
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (ProviderProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderProxy)(nil)).Elem()
}

func (i ProviderProxyArgs) ToProviderProxyOutput() ProviderProxyOutput {
	return i.ToProviderProxyOutputWithContext(context.Background())
}

func (i ProviderProxyArgs) ToProviderProxyOutputWithContext(ctx context.Context) ProviderProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyOutput)
}

func (i ProviderProxyArgs) ToProviderProxyPtrOutput() ProviderProxyPtrOutput {
	return i.ToProviderProxyPtrOutputWithContext(context.Background())
}

func (i ProviderProxyArgs) ToProviderProxyPtrOutputWithContext(ctx context.Context) ProviderProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyOutput).ToProviderProxyPtrOutputWithContext(ctx)
}

// ProviderProxyPtrInput is an input type that accepts ProviderProxyArgs, ProviderProxyPtr and ProviderProxyPtrOutput values.
// You can construct a concrete instance of `ProviderProxyPtrInput` via:
//
//	        ProviderProxyArgs{...}
//
//	or:
//
//	        nil
type ProviderProxyPtrInput interface {
	pulumi.Input

	ToProviderProxyPtrOutput() ProviderProxyPtrOutput
	ToProviderProxyPtrOutputWithContext(context.Context) ProviderProxyPtrOutput
}

type providerProxyPtrType ProviderProxyArgs

func ProviderProxyPtr(v *ProviderProxyArgs) ProviderProxyPtrInput {
	return (*providerProxyPtrType)(v)
}

func (*providerProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderProxy)(nil)).Elem()
}

func (i *providerProxyPtrType) ToProviderProxyPtrOutput() ProviderProxyPtrOutput {
	return i.ToProviderProxyPtrOutputWithContext(context.Background())
}

func (i *providerProxyPtrType) ToProviderProxyPtrOutputWithContext(ctx context.Context) ProviderProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderProxyPtrOutput)
}

type ProviderProxyOutput struct{ *pulumi.OutputState }

func (ProviderProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderProxy)(nil)).Elem()
}

func (o ProviderProxyOutput) ToProviderProxyOutput() ProviderProxyOutput {
	return o
}

func (o ProviderProxyOutput) ToProviderProxyOutputWithContext(ctx context.Context) ProviderProxyOutput {
	return o
}

func (o ProviderProxyOutput) ToProviderProxyPtrOutput() ProviderProxyPtrOutput {
	return o.ToProviderProxyPtrOutputWithContext(context.Background())
}

func (o ProviderProxyOutput) ToProviderProxyPtrOutputWithContext(ctx context.Context) ProviderProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProviderProxy) *ProviderProxy {
		return &v
	}).(ProviderProxyPtrOutput)
}

// When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
func (o ProviderProxyOutput) FromEnv() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ProviderProxy) *bool { return v.FromEnv }).(pulumi.BoolPtrOutput)
}

// Password used for Basic authentication against the Proxy.
func (o ProviderProxyOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderProxy) *string { return v.Password }).(pulumi.StringPtrOutput)
}

// URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
func (o ProviderProxyOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderProxy) *string { return v.Url }).(pulumi.StringPtrOutput)
}

// Username (or Token) used for Basic authentication against the Proxy.
func (o ProviderProxyOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderProxy) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ProviderProxyPtrOutput struct{ *pulumi.OutputState }

func (ProviderProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderProxy)(nil)).Elem()
}

func (o ProviderProxyPtrOutput) ToProviderProxyPtrOutput() ProviderProxyPtrOutput {
	return o
}

func (o ProviderProxyPtrOutput) ToProviderProxyPtrOutputWithContext(ctx context.Context) ProviderProxyPtrOutput {
	return o
}

func (o ProviderProxyPtrOutput) Elem() ProviderProxyOutput {
	return o.ApplyT(func(v *ProviderProxy) ProviderProxy {
		if v != nil {
			return *v
		}
		var ret ProviderProxy
		return ret
	}).(ProviderProxyOutput)
}

// When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
func (o ProviderProxyPtrOutput) FromEnv() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) *bool {
		if v == nil {
			return nil
		}
		return v.FromEnv
	}).(pulumi.BoolPtrOutput)
}

// Password used for Basic authentication against the Proxy.
func (o ProviderProxyPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

// URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
func (o ProviderProxyPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

// Username (or Token) used for Basic authentication against the Proxy.
func (o ProviderProxyPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderProxy) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type SelfSignedCertSubject struct {
	// Distinguished name: `CN`
	CommonName *string `pulumi:"commonName"`
	// Distinguished name: `C`
	Country *string `pulumi:"country"`
	// Distinguished name: `L`
	Locality *string `pulumi:"locality"`
	// Distinguished name: `O`
	Organization *string `pulumi:"organization"`
	// Distinguished name: `OU`
	OrganizationalUnit *string `pulumi:"organizationalUnit"`
	// Distinguished name: `PC`
	PostalCode *string `pulumi:"postalCode"`
	// Distinguished name: `ST`
	Province *string `pulumi:"province"`
	// Distinguished name: `SERIALNUMBER`
	SerialNumber *string `pulumi:"serialNumber"`
	// Distinguished name: `STREET`
	StreetAddresses []string `pulumi:"streetAddresses"`
}

// SelfSignedCertSubjectInput is an input type that accepts SelfSignedCertSubjectArgs and SelfSignedCertSubjectOutput values.
// You can construct a concrete instance of `SelfSignedCertSubjectInput` via:
//
//	SelfSignedCertSubjectArgs{...}
type SelfSignedCertSubjectInput interface {
	pulumi.Input

	ToSelfSignedCertSubjectOutput() SelfSignedCertSubjectOutput
	ToSelfSignedCertSubjectOutputWithContext(context.Context) SelfSignedCertSubjectOutput
}

type SelfSignedCertSubjectArgs struct {
	// Distinguished name: `CN`
	CommonName pulumi.StringPtrInput `pulumi:"commonName"`
	// Distinguished name: `C`
	Country pulumi.StringPtrInput `pulumi:"country"`
	// Distinguished name: `L`
	Locality pulumi.StringPtrInput `pulumi:"locality"`
	// Distinguished name: `O`
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// Distinguished name: `OU`
	OrganizationalUnit pulumi.StringPtrInput `pulumi:"organizationalUnit"`
	// Distinguished name: `PC`
	PostalCode pulumi.StringPtrInput `pulumi:"postalCode"`
	// Distinguished name: `ST`
	Province pulumi.StringPtrInput `pulumi:"province"`
	// Distinguished name: `SERIALNUMBER`
	SerialNumber pulumi.StringPtrInput `pulumi:"serialNumber"`
	// Distinguished name: `STREET`
	StreetAddresses pulumi.StringArrayInput `pulumi:"streetAddresses"`
}

func (SelfSignedCertSubjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSignedCertSubject)(nil)).Elem()
}

func (i SelfSignedCertSubjectArgs) ToSelfSignedCertSubjectOutput() SelfSignedCertSubjectOutput {
	return i.ToSelfSignedCertSubjectOutputWithContext(context.Background())
}

func (i SelfSignedCertSubjectArgs) ToSelfSignedCertSubjectOutputWithContext(ctx context.Context) SelfSignedCertSubjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertSubjectOutput)
}

func (i SelfSignedCertSubjectArgs) ToSelfSignedCertSubjectPtrOutput() SelfSignedCertSubjectPtrOutput {
	return i.ToSelfSignedCertSubjectPtrOutputWithContext(context.Background())
}

func (i SelfSignedCertSubjectArgs) ToSelfSignedCertSubjectPtrOutputWithContext(ctx context.Context) SelfSignedCertSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertSubjectOutput).ToSelfSignedCertSubjectPtrOutputWithContext(ctx)
}

// SelfSignedCertSubjectPtrInput is an input type that accepts SelfSignedCertSubjectArgs, SelfSignedCertSubjectPtr and SelfSignedCertSubjectPtrOutput values.
// You can construct a concrete instance of `SelfSignedCertSubjectPtrInput` via:
//
//	        SelfSignedCertSubjectArgs{...}
//
//	or:
//
//	        nil
type SelfSignedCertSubjectPtrInput interface {
	pulumi.Input

	ToSelfSignedCertSubjectPtrOutput() SelfSignedCertSubjectPtrOutput
	ToSelfSignedCertSubjectPtrOutputWithContext(context.Context) SelfSignedCertSubjectPtrOutput
}

type selfSignedCertSubjectPtrType SelfSignedCertSubjectArgs

func SelfSignedCertSubjectPtr(v *SelfSignedCertSubjectArgs) SelfSignedCertSubjectPtrInput {
	return (*selfSignedCertSubjectPtrType)(v)
}

func (*selfSignedCertSubjectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSignedCertSubject)(nil)).Elem()
}

func (i *selfSignedCertSubjectPtrType) ToSelfSignedCertSubjectPtrOutput() SelfSignedCertSubjectPtrOutput {
	return i.ToSelfSignedCertSubjectPtrOutputWithContext(context.Background())
}

func (i *selfSignedCertSubjectPtrType) ToSelfSignedCertSubjectPtrOutputWithContext(ctx context.Context) SelfSignedCertSubjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSignedCertSubjectPtrOutput)
}

type SelfSignedCertSubjectOutput struct{ *pulumi.OutputState }

func (SelfSignedCertSubjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSignedCertSubject)(nil)).Elem()
}

func (o SelfSignedCertSubjectOutput) ToSelfSignedCertSubjectOutput() SelfSignedCertSubjectOutput {
	return o
}

func (o SelfSignedCertSubjectOutput) ToSelfSignedCertSubjectOutputWithContext(ctx context.Context) SelfSignedCertSubjectOutput {
	return o
}

func (o SelfSignedCertSubjectOutput) ToSelfSignedCertSubjectPtrOutput() SelfSignedCertSubjectPtrOutput {
	return o.ToSelfSignedCertSubjectPtrOutputWithContext(context.Background())
}

func (o SelfSignedCertSubjectOutput) ToSelfSignedCertSubjectPtrOutputWithContext(ctx context.Context) SelfSignedCertSubjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SelfSignedCertSubject) *SelfSignedCertSubject {
		return &v
	}).(SelfSignedCertSubjectPtrOutput)
}

// Distinguished name: `CN`
func (o SelfSignedCertSubjectOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

// Distinguished name: `C`
func (o SelfSignedCertSubjectOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.Country }).(pulumi.StringPtrOutput)
}

// Distinguished name: `L`
func (o SelfSignedCertSubjectOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.Locality }).(pulumi.StringPtrOutput)
}

// Distinguished name: `O`
func (o SelfSignedCertSubjectOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// Distinguished name: `OU`
func (o SelfSignedCertSubjectOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.OrganizationalUnit }).(pulumi.StringPtrOutput)
}

// Distinguished name: `PC`
func (o SelfSignedCertSubjectOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.PostalCode }).(pulumi.StringPtrOutput)
}

// Distinguished name: `ST`
func (o SelfSignedCertSubjectOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.Province }).(pulumi.StringPtrOutput)
}

// Distinguished name: `SERIALNUMBER`
func (o SelfSignedCertSubjectOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// Distinguished name: `STREET`
func (o SelfSignedCertSubjectOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelfSignedCertSubject) []string { return v.StreetAddresses }).(pulumi.StringArrayOutput)
}

type SelfSignedCertSubjectPtrOutput struct{ *pulumi.OutputState }

func (SelfSignedCertSubjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSignedCertSubject)(nil)).Elem()
}

func (o SelfSignedCertSubjectPtrOutput) ToSelfSignedCertSubjectPtrOutput() SelfSignedCertSubjectPtrOutput {
	return o
}

func (o SelfSignedCertSubjectPtrOutput) ToSelfSignedCertSubjectPtrOutputWithContext(ctx context.Context) SelfSignedCertSubjectPtrOutput {
	return o
}

func (o SelfSignedCertSubjectPtrOutput) Elem() SelfSignedCertSubjectOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) SelfSignedCertSubject {
		if v != nil {
			return *v
		}
		var ret SelfSignedCertSubject
		return ret
	}).(SelfSignedCertSubjectOutput)
}

// Distinguished name: `CN`
func (o SelfSignedCertSubjectPtrOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.CommonName
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `C`
func (o SelfSignedCertSubjectPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.Country
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `L`
func (o SelfSignedCertSubjectPtrOutput) Locality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.Locality
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `O`
func (o SelfSignedCertSubjectPtrOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.Organization
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `OU`
func (o SelfSignedCertSubjectPtrOutput) OrganizationalUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.OrganizationalUnit
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `PC`
func (o SelfSignedCertSubjectPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.PostalCode
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `ST`
func (o SelfSignedCertSubjectPtrOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.Province
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `SERIALNUMBER`
func (o SelfSignedCertSubjectPtrOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) *string {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.StringPtrOutput)
}

// Distinguished name: `STREET`
func (o SelfSignedCertSubjectPtrOutput) StreetAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SelfSignedCertSubject) []string {
		if v == nil {
			return nil
		}
		return v.StreetAddresses
	}).(pulumi.StringArrayOutput)
}

type GetCertificateCertificate struct {
	// Certificate data in PEM (RFC 1421).
	CertPem string `pulumi:"certPem"`
	// `true` if the certificate is of a CA (Certificate Authority).
	IsCa bool `pulumi:"isCa"`
	// Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
	Issuer string `pulumi:"issuer"`
	// The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	NotAfter string `pulumi:"notAfter"`
	// The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	NotBefore string `pulumi:"notBefore"`
	// The key algorithm used to create the certificate.
	PublicKeyAlgorithm string `pulumi:"publicKeyAlgorithm"`
	// Number that uniquely identifies the certificate with the CA's system.
	// The `format` function can be used to convert this *base 10* number into other bases, such as hex.
	SerialNumber string `pulumi:"serialNumber"`
	// The SHA1 fingerprint of the public key of the certificate.
	Sha1Fingerprint string `pulumi:"sha1Fingerprint"`
	// The algorithm used to sign the certificate.
	SignatureAlgorithm string `pulumi:"signatureAlgorithm"`
	// The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
	Subject string `pulumi:"subject"`
	// The version the certificate is in.
	Version int `pulumi:"version"`
}

// GetCertificateCertificateInput is an input type that accepts GetCertificateCertificateArgs and GetCertificateCertificateOutput values.
// You can construct a concrete instance of `GetCertificateCertificateInput` via:
//
//	GetCertificateCertificateArgs{...}
type GetCertificateCertificateInput interface {
	pulumi.Input

	ToGetCertificateCertificateOutput() GetCertificateCertificateOutput
	ToGetCertificateCertificateOutputWithContext(context.Context) GetCertificateCertificateOutput
}

type GetCertificateCertificateArgs struct {
	// Certificate data in PEM (RFC 1421).
	CertPem pulumi.StringInput `pulumi:"certPem"`
	// `true` if the certificate is of a CA (Certificate Authority).
	IsCa pulumi.BoolInput `pulumi:"isCa"`
	// Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
	Issuer pulumi.StringInput `pulumi:"issuer"`
	// The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	NotAfter pulumi.StringInput `pulumi:"notAfter"`
	// The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
	NotBefore pulumi.StringInput `pulumi:"notBefore"`
	// The key algorithm used to create the certificate.
	PublicKeyAlgorithm pulumi.StringInput `pulumi:"publicKeyAlgorithm"`
	// Number that uniquely identifies the certificate with the CA's system.
	// The `format` function can be used to convert this *base 10* number into other bases, such as hex.
	SerialNumber pulumi.StringInput `pulumi:"serialNumber"`
	// The SHA1 fingerprint of the public key of the certificate.
	Sha1Fingerprint pulumi.StringInput `pulumi:"sha1Fingerprint"`
	// The algorithm used to sign the certificate.
	SignatureAlgorithm pulumi.StringInput `pulumi:"signatureAlgorithm"`
	// The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
	Subject pulumi.StringInput `pulumi:"subject"`
	// The version the certificate is in.
	Version pulumi.IntInput `pulumi:"version"`
}

func (GetCertificateCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificateCertificate)(nil)).Elem()
}

func (i GetCertificateCertificateArgs) ToGetCertificateCertificateOutput() GetCertificateCertificateOutput {
	return i.ToGetCertificateCertificateOutputWithContext(context.Background())
}

func (i GetCertificateCertificateArgs) ToGetCertificateCertificateOutputWithContext(ctx context.Context) GetCertificateCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCertificateCertificateOutput)
}

// GetCertificateCertificateArrayInput is an input type that accepts GetCertificateCertificateArray and GetCertificateCertificateArrayOutput values.
// You can construct a concrete instance of `GetCertificateCertificateArrayInput` via:
//
//	GetCertificateCertificateArray{ GetCertificateCertificateArgs{...} }
type GetCertificateCertificateArrayInput interface {
	pulumi.Input

	ToGetCertificateCertificateArrayOutput() GetCertificateCertificateArrayOutput
	ToGetCertificateCertificateArrayOutputWithContext(context.Context) GetCertificateCertificateArrayOutput
}

type GetCertificateCertificateArray []GetCertificateCertificateInput

func (GetCertificateCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCertificateCertificate)(nil)).Elem()
}

func (i GetCertificateCertificateArray) ToGetCertificateCertificateArrayOutput() GetCertificateCertificateArrayOutput {
	return i.ToGetCertificateCertificateArrayOutputWithContext(context.Background())
}

func (i GetCertificateCertificateArray) ToGetCertificateCertificateArrayOutputWithContext(ctx context.Context) GetCertificateCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetCertificateCertificateArrayOutput)
}

type GetCertificateCertificateOutput struct{ *pulumi.OutputState }

func (GetCertificateCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificateCertificate)(nil)).Elem()
}

func (o GetCertificateCertificateOutput) ToGetCertificateCertificateOutput() GetCertificateCertificateOutput {
	return o
}

func (o GetCertificateCertificateOutput) ToGetCertificateCertificateOutputWithContext(ctx context.Context) GetCertificateCertificateOutput {
	return o
}

// Certificate data in PEM (RFC 1421).
func (o GetCertificateCertificateOutput) CertPem() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.CertPem }).(pulumi.StringOutput)
}

// `true` if the certificate is of a CA (Certificate Authority).
func (o GetCertificateCertificateOutput) IsCa() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCertificateCertificate) bool { return v.IsCa }).(pulumi.BoolOutput)
}

// Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
func (o GetCertificateCertificateOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.Issuer }).(pulumi.StringOutput)
}

// The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o GetCertificateCertificateOutput) NotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.NotAfter }).(pulumi.StringOutput)
}

// The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
func (o GetCertificateCertificateOutput) NotBefore() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.NotBefore }).(pulumi.StringOutput)
}

// The key algorithm used to create the certificate.
func (o GetCertificateCertificateOutput) PublicKeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.PublicKeyAlgorithm }).(pulumi.StringOutput)
}

// Number that uniquely identifies the certificate with the CA's system.
// The `format` function can be used to convert this *base 10* number into other bases, such as hex.
func (o GetCertificateCertificateOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// The SHA1 fingerprint of the public key of the certificate.
func (o GetCertificateCertificateOutput) Sha1Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.Sha1Fingerprint }).(pulumi.StringOutput)
}

// The algorithm used to sign the certificate.
func (o GetCertificateCertificateOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

// The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
func (o GetCertificateCertificateOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificateCertificate) string { return v.Subject }).(pulumi.StringOutput)
}

// The version the certificate is in.
func (o GetCertificateCertificateOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v GetCertificateCertificate) int { return v.Version }).(pulumi.IntOutput)
}

type GetCertificateCertificateArrayOutput struct{ *pulumi.OutputState }

func (GetCertificateCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetCertificateCertificate)(nil)).Elem()
}

func (o GetCertificateCertificateArrayOutput) ToGetCertificateCertificateArrayOutput() GetCertificateCertificateArrayOutput {
	return o
}

func (o GetCertificateCertificateArrayOutput) ToGetCertificateCertificateArrayOutputWithContext(ctx context.Context) GetCertificateCertificateArrayOutput {
	return o
}

func (o GetCertificateCertificateArrayOutput) Index(i pulumi.IntInput) GetCertificateCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetCertificateCertificate {
		return vs[0].([]GetCertificateCertificate)[vs[1].(int)]
	}).(GetCertificateCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestSubjectInput)(nil)).Elem(), CertRequestSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertRequestSubjectPtrInput)(nil)).Elem(), CertRequestSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderProxyInput)(nil)).Elem(), ProviderProxyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderProxyPtrInput)(nil)).Elem(), ProviderProxyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSignedCertSubjectInput)(nil)).Elem(), SelfSignedCertSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSignedCertSubjectPtrInput)(nil)).Elem(), SelfSignedCertSubjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCertificateCertificateInput)(nil)).Elem(), GetCertificateCertificateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetCertificateCertificateArrayInput)(nil)).Elem(), GetCertificateCertificateArray{})
	pulumi.RegisterOutputType(CertRequestSubjectOutput{})
	pulumi.RegisterOutputType(CertRequestSubjectPtrOutput{})
	pulumi.RegisterOutputType(ProviderProxyOutput{})
	pulumi.RegisterOutputType(ProviderProxyPtrOutput{})
	pulumi.RegisterOutputType(SelfSignedCertSubjectOutput{})
	pulumi.RegisterOutputType(SelfSignedCertSubjectPtrOutput{})
	pulumi.RegisterOutputType(GetCertificateCertificateOutput{})
	pulumi.RegisterOutputType(GetCertificateCertificateArrayOutput{})
}
