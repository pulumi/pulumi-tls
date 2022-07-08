// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.LocallySignedCertArgs;
import com.pulumi.tls.Utilities;
import com.pulumi.tls.inputs.LocallySignedCertState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="tls:index/locallySignedCert:LocallySignedCert")
public class LocallySignedCert extends com.pulumi.resources.CustomResource {
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
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421)
     * format.
     * 
     */
    @Export(name="caCertPem", type=String.class, parameters={})
    private Output<String> caCertPem;

    /**
     * @return Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421)
     * format.
     * 
     */
    public Output<String> caCertPem() {
        return this.caCertPem;
    }
    /**
     * Name of the algorithm used when generating the private key provided in `ca_private_key_pem`. **NOTE**: this is
     * deprecated and ignored, as the key algorithm is now inferred from the key.
     * 
     * @deprecated
     * This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
     * 
     */
    @Deprecated /* This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`. */
    @Export(name="caKeyAlgorithm", type=String.class, parameters={})
    private Output<String> caKeyAlgorithm;

    /**
     * @return Name of the algorithm used when generating the private key provided in `ca_private_key_pem`. **NOTE**: this is
     * deprecated and ignored, as the key algorithm is now inferred from the key.
     * 
     */
    public Output<String> caKeyAlgorithm() {
        return this.caKeyAlgorithm;
    }
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC
     * 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Export(name="caPrivateKeyPem", type=String.class, parameters={})
    private Output<String> caPrivateKeyPem;

    /**
     * @return Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC
     * 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> caPrivateKeyPem() {
        return this.caPrivateKeyPem;
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
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Export(name="certRequestPem", type=String.class, parameters={})
    private Output<String> certRequestPem;

    /**
     * @return Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> certRequestPem() {
        return this.certRequestPem;
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
    public LocallySignedCert(String name) {
        this(name, LocallySignedCertArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocallySignedCert(String name, LocallySignedCertArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocallySignedCert(String name, LocallySignedCertArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/locallySignedCert:LocallySignedCert", name, args == null ? LocallySignedCertArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocallySignedCert(String name, Output<String> id, @Nullable LocallySignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/locallySignedCert:LocallySignedCert", name, state, makeResourceOptions(options, id));
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
    public static LocallySignedCert get(String name, Output<String> id, @Nullable LocallySignedCertState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocallySignedCert(name, id, state, options);
    }
}