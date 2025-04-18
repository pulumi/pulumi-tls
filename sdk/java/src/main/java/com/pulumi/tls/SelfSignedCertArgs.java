// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.tls.inputs.SelfSignedCertSubjectArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SelfSignedCertArgs extends com.pulumi.resources.ResourceArgs {

    public static final SelfSignedCertArgs Empty = new SelfSignedCertArgs();

    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    @Import(name="allowedUses", required=true)
    private Output<List<String>> allowedUses;

    /**
     * @return List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
     * 
     */
    public Output<List<String>> allowedUses() {
        return this.allowedUses;
    }

    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="dnsNames")
    private @Nullable Output<List<String>> dnsNames;

    /**
     * @return List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> dnsNames() {
        return Optional.ofNullable(this.dnsNames);
    }

    @Import(name="earlyRenewalHours")
    private @Nullable Output<Integer> earlyRenewalHours;

    public Optional<Output<Integer>> earlyRenewalHours() {
        return Optional.ofNullable(this.earlyRenewalHours);
    }

    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="ipAddresses")
    private @Nullable Output<List<String>> ipAddresses;

    /**
     * @return List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> ipAddresses() {
        return Optional.ofNullable(this.ipAddresses);
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
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
     * 
     */
    @Import(name="privateKeyPem", required=true)
    private Output<String> privateKeyPem;

    /**
     * @return Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }

    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Import(name="setAuthorityKeyId")
    private @Nullable Output<Boolean> setAuthorityKeyId;

    /**
     * @return Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> setAuthorityKeyId() {
        return Optional.ofNullable(this.setAuthorityKeyId);
    }

    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    @Import(name="setSubjectKeyId")
    private @Nullable Output<Boolean> setSubjectKeyId;

    /**
     * @return Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     * 
     */
    public Optional<Output<Boolean>> setSubjectKeyId() {
        return Optional.ofNullable(this.setSubjectKeyId);
    }

    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    @Import(name="subject")
    private @Nullable Output<SelfSignedCertSubjectArgs> subject;

    /**
     * @return The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     * 
     */
    public Optional<Output<SelfSignedCertSubjectArgs>> subject() {
        return Optional.ofNullable(this.subject);
    }

    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    @Import(name="uris")
    private @Nullable Output<List<String>> uris;

    /**
     * @return List of URIs for which a certificate is being requested (i.e. certificate subjects).
     * 
     */
    public Optional<Output<List<String>>> uris() {
        return Optional.ofNullable(this.uris);
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

    private SelfSignedCertArgs() {}

    private SelfSignedCertArgs(SelfSignedCertArgs $) {
        this.allowedUses = $.allowedUses;
        this.dnsNames = $.dnsNames;
        this.earlyRenewalHours = $.earlyRenewalHours;
        this.ipAddresses = $.ipAddresses;
        this.isCaCertificate = $.isCaCertificate;
        this.privateKeyPem = $.privateKeyPem;
        this.setAuthorityKeyId = $.setAuthorityKeyId;
        this.setSubjectKeyId = $.setSubjectKeyId;
        this.subject = $.subject;
        this.uris = $.uris;
        this.validityPeriodHours = $.validityPeriodHours;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SelfSignedCertArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SelfSignedCertArgs $;

        public Builder() {
            $ = new SelfSignedCertArgs();
        }

        public Builder(SelfSignedCertArgs defaults) {
            $ = new SelfSignedCertArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(Output<List<String>> allowedUses) {
            $.allowedUses = allowedUses;
            return this;
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(List<String> allowedUses) {
            return allowedUses(Output.of(allowedUses));
        }

        /**
         * @param allowedUses List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
         * 
         * @return builder
         * 
         */
        public Builder allowedUses(String... allowedUses) {
            return allowedUses(List.of(allowedUses));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(@Nullable Output<List<String>> dnsNames) {
            $.dnsNames = dnsNames;
            return this;
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(List<String> dnsNames) {
            return dnsNames(Output.of(dnsNames));
        }

        /**
         * @param dnsNames List of DNS names for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder dnsNames(String... dnsNames) {
            return dnsNames(List.of(dnsNames));
        }

        public Builder earlyRenewalHours(@Nullable Output<Integer> earlyRenewalHours) {
            $.earlyRenewalHours = earlyRenewalHours;
            return this;
        }

        public Builder earlyRenewalHours(Integer earlyRenewalHours) {
            return earlyRenewalHours(Output.of(earlyRenewalHours));
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(@Nullable Output<List<String>> ipAddresses) {
            $.ipAddresses = ipAddresses;
            return this;
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(List<String> ipAddresses) {
            return ipAddresses(Output.of(ipAddresses));
        }

        /**
         * @param ipAddresses List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
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
         * @param privateKeyPem Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(Output<String> privateKeyPem) {
            $.privateKeyPem = privateKeyPem;
            return this;
        }

        /**
         * @param privateKeyPem Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong to.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(String privateKeyPem) {
            return privateKeyPem(Output.of(privateKeyPem));
        }

        /**
         * @param setAuthorityKeyId Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setAuthorityKeyId(@Nullable Output<Boolean> setAuthorityKeyId) {
            $.setAuthorityKeyId = setAuthorityKeyId;
            return this;
        }

        /**
         * @param setAuthorityKeyId Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setAuthorityKeyId(Boolean setAuthorityKeyId) {
            return setAuthorityKeyId(Output.of(setAuthorityKeyId));
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(@Nullable Output<Boolean> setSubjectKeyId) {
            $.setSubjectKeyId = setSubjectKeyId;
            return this;
        }

        /**
         * @param setSubjectKeyId Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
         * 
         * @return builder
         * 
         */
        public Builder setSubjectKeyId(Boolean setSubjectKeyId) {
            return setSubjectKeyId(Output.of(setSubjectKeyId));
        }

        /**
         * @param subject The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
         * 
         * @return builder
         * 
         */
        public Builder subject(@Nullable Output<SelfSignedCertSubjectArgs> subject) {
            $.subject = subject;
            return this;
        }

        /**
         * @param subject The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
         * 
         * @return builder
         * 
         */
        public Builder subject(SelfSignedCertSubjectArgs subject) {
            return subject(Output.of(subject));
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(@Nullable Output<List<String>> uris) {
            $.uris = uris;
            return this;
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(List<String> uris) {
            return uris(Output.of(uris));
        }

        /**
         * @param uris List of URIs for which a certificate is being requested (i.e. certificate subjects).
         * 
         * @return builder
         * 
         */
        public Builder uris(String... uris) {
            return uris(List.of(uris));
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

        public SelfSignedCertArgs build() {
            if ($.allowedUses == null) {
                throw new MissingRequiredPropertyException("SelfSignedCertArgs", "allowedUses");
            }
            if ($.privateKeyPem == null) {
                throw new MissingRequiredPropertyException("SelfSignedCertArgs", "privateKeyPem");
            }
            if ($.validityPeriodHours == null) {
                throw new MissingRequiredPropertyException("SelfSignedCertArgs", "validityPeriodHours");
            }
            return $;
        }
    }

}
