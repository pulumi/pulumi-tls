// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LocallySignedCertArgs extends com.pulumi.resources.ResourceArgs {

    public static final LocallySignedCertArgs Empty = new LocallySignedCertArgs();

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
    @Import(name="allowedUses", required=true)
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
    @Import(name="caCertPem", required=true)
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
    @Import(name="caKeyAlgorithm")
    private @Nullable Output<String> caKeyAlgorithm;

    /**
     * @return Name of the algorithm used when generating the private key provided in `ca_private_key_pem`. **NOTE**: this is
     * deprecated and ignored, as the key algorithm is now inferred from the key.
     * 
     * @deprecated
     * This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
     * 
     */
    @Deprecated /* This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`. */
    public Optional<Output<String>> caKeyAlgorithm() {
        return Optional.ofNullable(this.caKeyAlgorithm);
    }

    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC
     * 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Import(name="caPrivateKeyPem", required=true)
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
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Import(name="certRequestPem", required=true)
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
    @Import(name="earlyRenewalHours")
    private @Nullable Output<Integer> earlyRenewalHours;

    /**
     * @return The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     * 
     */
    public Optional<Output<Integer>> earlyRenewalHours() {
        return Optional.ofNullable(this.earlyRenewalHours);
    }

    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    @Import(name="isCaCertificate")
    private @Nullable Output<Boolean> isCaCertificate;

    /**
     * @return Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> isCaCertificate() {
        return Optional.ofNullable(this.isCaCertificate);
    }

    /**
     * Should the generated certificate include a [subject key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Import(name="setSubjectKeyId")
    private @Nullable Output<Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a [subject key
     * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> setSubjectKeyId() {
        return Optional.ofNullable(this.setSubjectKeyId);
    }

    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    @Import(name="validityPeriodHours", required=true)
    private Output<Integer> validityPeriodHours;

    /**
     * @return Number of hours, after initial issuing, that the certificate will remain valid for.
     * 
     */
    public Output<Integer> validityPeriodHours() {
        return this.validityPeriodHours;
    }

    private LocallySignedCertArgs() {}

    private LocallySignedCertArgs(LocallySignedCertArgs $) {
        this.allowedUses = $.allowedUses;
        this.caCertPem = $.caCertPem;
        this.caKeyAlgorithm = $.caKeyAlgorithm;
        this.caPrivateKeyPem = $.caPrivateKeyPem;
        this.certRequestPem = $.certRequestPem;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.isCaCertificate = $.isCaCertificate;
        this.setSubjectKeyId = $.setSubjectKeyId;
        this.validityPeriodHours = $.validityPeriodHours;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LocallySignedCertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LocallySignedCertArgs $;

        public Builder() {
            $ = new LocallySignedCertArgs();
        }

        public Builder(LocallySignedCertArgs defaults) {
            $ = new LocallySignedCertArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC
         * 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`,
         * `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`,
         * `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`,
         * `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`,
         * `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(Output<List<String>> allowedUses) {
            $.allowedUses = allowedUses;
            return this;
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC
         * 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`,
         * `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`,
         * `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`,
         * `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`,
         * `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(List<String> allowedUses) {
            return allowedUses(Output.of(allowedUses));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC
         * 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key
         * Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`,
         * `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`,
         * `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`,
         * `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`,
         * `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(String... allowedUses) {
            return allowedUses(List.of(allowedUses));
        }

        /**
         * @param caCertPem Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421)
         * format.
         * 
         * @return builder
         * 
         */
        public Builder caCertPem(Output<String> caCertPem) {
            $.caCertPem = caCertPem;
            return this;
        }

        /**
         * @param caCertPem Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421)
         * format.
         * 
         * @return builder
         * 
         */
        public Builder caCertPem(String caCertPem) {
            return caCertPem(Output.of(caCertPem));
        }

        /**
         * @param caKeyAlgorithm Name of the algorithm used when generating the private key provided in `ca_private_key_pem`. **NOTE**: this is
         * deprecated and ignored, as the key algorithm is now inferred from the key.
         * 
         * @return builder
         * 
         * @deprecated
         * This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
         * 
         */
        @Deprecated /* This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`. */
        public Builder caKeyAlgorithm(@Nullable Output<String> caKeyAlgorithm) {
            $.caKeyAlgorithm = caKeyAlgorithm;
            return this;
        }

        /**
         * @param caKeyAlgorithm Name of the algorithm used when generating the private key provided in `ca_private_key_pem`. **NOTE**: this is
         * deprecated and ignored, as the key algorithm is now inferred from the key.
         * 
         * @return builder
         * 
         * @deprecated
         * This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`.
         * 
         */
        @Deprecated /* This is now ignored, as the key algorithm is inferred from the `ca_private_key_pem`. */
        public Builder caKeyAlgorithm(String caKeyAlgorithm) {
            return caKeyAlgorithm(Output.of(caKeyAlgorithm));
        }

        /**
         * @param caPrivateKeyPem Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC
         * 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caPrivateKeyPem(Output<String> caPrivateKeyPem) {
            $.caPrivateKeyPem = caPrivateKeyPem;
            return this;
        }

        /**
         * @param caPrivateKeyPem Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC
         * 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder caPrivateKeyPem(String caPrivateKeyPem) {
            return caPrivateKeyPem(Output.of(caPrivateKeyPem));
        }

        /**
         * @param certRequestPem Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder certRequestPem(Output<String> certRequestPem) {
            $.certRequestPem = certRequestPem;
            return this;
        }

        /**
         * @param certRequestPem Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder certRequestPem(String certRequestPem) {
            return certRequestPem(Output.of(certRequestPem));
        }

        /**
         * @param earlyRenewalHours The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
         * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
         * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
         * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
         * early renewal period. (default: `0`)
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(@Nullable Output<Integer> earlyRenewalHours) {
            $.earlyRenewalHours = earlyRenewalHours;
            return this;
        }

        /**
         * @param earlyRenewalHours The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
         * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
         * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
         * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
         * early renewal period. (default: `0`)
         * 
         * @return builder
         * 
         */
        public Builder earlyRenewalHours(Integer earlyRenewalHours) {
            return earlyRenewalHours(Output.of(earlyRenewalHours));
        }

        /**
         * @param isCaCertificate Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder isCaCertificate(@Nullable Output<Boolean> isCaCertificate) {
            $.isCaCertificate = isCaCertificate;
            return this;
        }

        /**
         * @param isCaCertificate Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder isCaCertificate(Boolean isCaCertificate) {
            return isCaCertificate(Output.of(isCaCertificate));
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key
         * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(@Nullable Output<Boolean> setSubjectKeyId) {
            $.setSubjectKeyId = setSubjectKeyId;
            return this;
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key
         * identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(Boolean setSubjectKeyId) {
            return setSubjectKeyId(Output.of(setSubjectKeyId));
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid for.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Output<Integer> validityPeriodHours) {
            $.validityPeriodHours = validityPeriodHours;
            return this;
        }

        /**
         * @param validityPeriodHours Number of hours, after initial issuing, that the certificate will remain valid for.
         * 
         * @return builder
         * 
         */
        public Builder validityPeriodHours(Integer validityPeriodHours) {
            return validityPeriodHours(Output.of(validityPeriodHours));
        }

        public LocallySignedCertArgs build() {
            $.allowedUses = Objects.requireNonNull($.allowedUses, "expected parameter 'allowedUses' to be non-null");
            $.caCertPem = Objects.requireNonNull($.caCertPem, "expected parameter 'caCertPem' to be non-null");
            $.caPrivateKeyPem = Objects.requireNonNull($.caPrivateKeyPem, "expected parameter 'caPrivateKeyPem' to be non-null");
            $.certRequestPem = Objects.requireNonNull($.certRequestPem, "expected parameter 'certRequestPem' to be non-null");
            $.validityPeriodHours = Objects.requireNonNull($.validityPeriodHours, "expected parameter 'validityPeriodHours' to be non-null");
            return $;
        }
    }

}
