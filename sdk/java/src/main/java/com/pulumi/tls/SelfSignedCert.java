// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.SelfSignedCertArgs;
import com.pulumi.tls.Utilities;
import com.pulumi.tls.inputs.SelfSignedCertState;
import com.pulumi.tls.outputs.SelfSignedCertSubject;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="tls:index/selfSignedCert:SelfSignedCert")
public class SelfSignedCert extends com.pulumi.resources.CustomResource {
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC
     * 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key
     * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key
     * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`,
     * `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`,
     * `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`,
     * `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`,
     * `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    @Export(name="allowedUses", type=List.class, parameters={String.class})
    private Output<List<String>> allowedUses;

    /**
     * @return List of key usages allowed for the issued certificate. Values are defined in [RFC
     * 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key
     * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key
     * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`,
     * `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`,
     * `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`,
     * `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`,
     * `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    public Output<List<String>> allowedUses() {
        return this.allowedUses;
    }
    /**
     * Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    @Export(name="certPem", type=String.class, parameters={})
    private Output<String> certPem;

    /**
     * @return Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    public Output<String> certPem() {
        return this.certPem;
    }
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="dnsNames", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> dnsNames() {
        return Codegen.optional(this.dnsNames);
    }
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    @Export(name="earlyRenewalHours", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> earlyRenewalHours;

    /**
     * @return The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    public Output<Optional<Integer>> earlyRenewalHours() {
        return Codegen.optional(this.earlyRenewalHours);
    }
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="ipAddresses", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> ipAddresses() {
        return Codegen.optional(this.ipAddresses);
    }
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    @Export(name="isCaCertificate", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isCaCertificate;

    /**
     * @return Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    public Output<Optional<Boolean>> isCaCertificate() {
        return Codegen.optional(this.isCaCertificate);
    }
    /**
     * Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
     * and ignored, as the key algorithm is now inferred from the key.
     * 
     * @deprecated
     * This is now ignored, as the key algorithm is inferred from the `private_key_pem`.
     * 
     */
    @Deprecated /* This is now ignored, as the key algorithm is inferred from the `private_key_pem`. */
    @Export(name="keyAlgorithm", type=String.class, parameters={})
    private Output<String> keyAlgorithm;

    /**
     * @return Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated
     * and ignored, as the key algorithm is now inferred from the key.
     * 
     */
    public Output<String> keyAlgorithm() {
        return this.keyAlgorithm;
    }
    /**
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     * 
     */
    @Export(name="privateKeyPem", type=String.class, parameters={})
    private Output<String> privateKeyPem;

    /**
     * @return Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    /**
     * Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within
     * the `early_renewal_hours`)?
     * 
     */
    @Export(name="readyForRenewal", type=Boolean.class, parameters={})
    private Output<Boolean> readyForRenewal;

    /**
     * @return Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within
     * the `early_renewal_hours`)?
     * 
     */
    public Output<Boolean> readyForRenewal() {
        return this.readyForRenewal;
    }
    /**
     * Should the generated certificate include an [authority key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the
     * same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default:
     * `false`).
     * 
     */
    @Export(name="setAuthorityKeyId", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> setAuthorityKeyId;

    /**
     * @return Should the generated certificate include an [authority key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the
     * same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default:
     * `false`).
     * 
     */
    public Output<Optional<Boolean>> setAuthorityKeyId() {
        return Codegen.optional(this.setAuthorityKeyId);
    }
    /**
     * Should the generated certificate include a [subject key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Export(name="setSubjectKeyId", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a [subject key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Output<Optional<Boolean>> setSubjectKeyId() {
        return Codegen.optional(this.setSubjectKeyId);
    }
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
     * based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    @Export(name="subject", type=SelfSignedCertSubject.class, parameters={})
    private Output</* @Nullable */ SelfSignedCertSubject> subject;

    /**
     * @return The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is
     * based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    public Output<Optional<SelfSignedCertSubject>> subject() {
        return Codegen.optional(this.subject);
    }
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Export(name="uris", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Output<Optional<List<String>>> uris() {
        return Codegen.optional(this.uris);
    }
    /**
     * The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339)
     * timestamp.
     * 
     */
    @Export(name="validityEndTime", type=String.class, parameters={})
    private Output<String> validityEndTime;

    /**
     * @return The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339)
     * timestamp.
     * 
     */
    public Output<String> validityEndTime() {
        return this.validityEndTime;
    }
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    @Export(name="validityPeriodHours", type=Integer.class, parameters={})
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }
    /**
     * The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    @Export(name="validityStartTime", type=String.class, parameters={})
    private Output<String> validityStartTime;

    /**
     * @return The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    public Output<String> validityStartTime() {
        return this.validityStartTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SelfSignedCert(String name) {
        this(name, SelfSignedCertArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SelfSignedCert(String name, SelfSignedCertArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SelfSignedCert(String name, SelfSignedCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/selfSignedCert:SelfSignedCert", name, args == null ? SelfSignedCertArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SelfSignedCert(String name, Output<String> id, @Nullable SelfSignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/selfSignedCert:SelfSignedCert", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static SelfSignedCert get(String name, Output<String> id, @Nullable SelfSignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SelfSignedCert(name, id, state, options);
    }
}
