// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCertificateCertificate {
    /**
     * @return Certificate data in PEM (RFC 1421).
     * 
     */
    private String certPem;
    /**
     * @return `true` if the certificate is of a CA (Certificate Authority).
     * 
     */
    private Boolean isCa;
    /**
     * @return Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
     * 
     */
    private String issuer;
    /**
     * @return The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    private String notAfter;
    /**
     * @return The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    private String notBefore;
    /**
     * @return The key algorithm used to create the certificate.
     * 
     */
    private String publicKeyAlgorithm;
    /**
     * @return Number that uniquely identifies the certificate with the CA&#39;s system.
     * The `format` function can be used to convert this *base 10* number into other bases, such as hex.
     * 
     */
    private String serialNumber;
    /**
     * @return The SHA1 fingerprint of the public key of the certificate.
     * 
     */
    private String sha1Fingerprint;
    /**
     * @return The algorithm used to sign the certificate.
     * 
     */
    private String signatureAlgorithm;
    /**
     * @return The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
     * 
     */
    private String subject;
    /**
     * @return The version the certificate is in.
     * 
     */
    private Integer version;

    private GetCertificateCertificate() {}
    /**
     * @return Certificate data in PEM (RFC 1421).
     * 
     */
    public String certPem() {
        return this.certPem;
    }
    /**
     * @return `true` if the certificate is of a CA (Certificate Authority).
     * 
     */
    public Boolean isCa() {
        return this.isCa;
    }
    /**
     * @return Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
     * 
     */
    public String issuer() {
        return this.issuer;
    }
    /**
     * @return The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    public String notAfter() {
        return this.notAfter;
    }
    /**
     * @return The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     * 
     */
    public String notBefore() {
        return this.notBefore;
    }
    /**
     * @return The key algorithm used to create the certificate.
     * 
     */
    public String publicKeyAlgorithm() {
        return this.publicKeyAlgorithm;
    }
    /**
     * @return Number that uniquely identifies the certificate with the CA&#39;s system.
     * The `format` function can be used to convert this *base 10* number into other bases, such as hex.
     * 
     */
    public String serialNumber() {
        return this.serialNumber;
    }
    /**
     * @return The SHA1 fingerprint of the public key of the certificate.
     * 
     */
    public String sha1Fingerprint() {
        return this.sha1Fingerprint;
    }
    /**
     * @return The algorithm used to sign the certificate.
     * 
     */
    public String signatureAlgorithm() {
        return this.signatureAlgorithm;
    }
    /**
     * @return The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
     * 
     */
    public String subject() {
        return this.subject;
    }
    /**
     * @return The version the certificate is in.
     * 
     */
    public Integer version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCertificateCertificate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String certPem;
        private Boolean isCa;
        private String issuer;
        private String notAfter;
        private String notBefore;
        private String publicKeyAlgorithm;
        private String serialNumber;
        private String sha1Fingerprint;
        private String signatureAlgorithm;
        private String subject;
        private Integer version;
        public Builder() {}
        public Builder(GetCertificateCertificate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certPem = defaults.certPem;
    	      this.isCa = defaults.isCa;
    	      this.issuer = defaults.issuer;
    	      this.notAfter = defaults.notAfter;
    	      this.notBefore = defaults.notBefore;
    	      this.publicKeyAlgorithm = defaults.publicKeyAlgorithm;
    	      this.serialNumber = defaults.serialNumber;
    	      this.sha1Fingerprint = defaults.sha1Fingerprint;
    	      this.signatureAlgorithm = defaults.signatureAlgorithm;
    	      this.subject = defaults.subject;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder certPem(String certPem) {
            this.certPem = Objects.requireNonNull(certPem);
            return this;
        }
        @CustomType.Setter
        public Builder isCa(Boolean isCa) {
            this.isCa = Objects.requireNonNull(isCa);
            return this;
        }
        @CustomType.Setter
        public Builder issuer(String issuer) {
            this.issuer = Objects.requireNonNull(issuer);
            return this;
        }
        @CustomType.Setter
        public Builder notAfter(String notAfter) {
            this.notAfter = Objects.requireNonNull(notAfter);
            return this;
        }
        @CustomType.Setter
        public Builder notBefore(String notBefore) {
            this.notBefore = Objects.requireNonNull(notBefore);
            return this;
        }
        @CustomType.Setter
        public Builder publicKeyAlgorithm(String publicKeyAlgorithm) {
            this.publicKeyAlgorithm = Objects.requireNonNull(publicKeyAlgorithm);
            return this;
        }
        @CustomType.Setter
        public Builder serialNumber(String serialNumber) {
            this.serialNumber = Objects.requireNonNull(serialNumber);
            return this;
        }
        @CustomType.Setter
        public Builder sha1Fingerprint(String sha1Fingerprint) {
            this.sha1Fingerprint = Objects.requireNonNull(sha1Fingerprint);
            return this;
        }
        @CustomType.Setter
        public Builder signatureAlgorithm(String signatureAlgorithm) {
            this.signatureAlgorithm = Objects.requireNonNull(signatureAlgorithm);
            return this;
        }
        @CustomType.Setter
        public Builder subject(String subject) {
            this.subject = Objects.requireNonNull(subject);
            return this;
        }
        @CustomType.Setter
        public Builder version(Integer version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetCertificateCertificate build() {
            final var _resultValue = new GetCertificateCertificate();
            _resultValue.certPem = certPem;
            _resultValue.isCa = isCa;
            _resultValue.issuer = issuer;
            _resultValue.notAfter = notAfter;
            _resultValue.notBefore = notBefore;
            _resultValue.publicKeyAlgorithm = publicKeyAlgorithm;
            _resultValue.serialNumber = serialNumber;
            _resultValue.sha1Fingerprint = sha1Fingerprint;
            _resultValue.signatureAlgorithm = signatureAlgorithm;
            _resultValue.subject = subject;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
