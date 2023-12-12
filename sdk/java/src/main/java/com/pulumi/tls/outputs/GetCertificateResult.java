// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.tls.outputs.GetCertificateCertificate;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCertificateResult {
    /**
     * @return The certificates protecting the site, with the root of the chain first.
     * 
     */
    private List<GetCertificateCertificate> certificates;
    /**
     * @return The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. Cannot be used with `url`.
     * 
     */
    private @Nullable String content;
    /**
     * @return Unique identifier of this data source: hashing of the certificates in the chain.
     * 
     */
    private String id;
    /**
     * @return The URL of the website to get the certificates from. Cannot be used with `content`.
     * 
     */
    private @Nullable String url;
    /**
     * @return Whether to verify the certificate chain while parsing it or not (default: `true`). Cannot be used with `content`.
     * 
     */
    private @Nullable Boolean verifyChain;

    private GetCertificateResult() {}
    /**
     * @return The certificates protecting the site, with the root of the chain first.
     * 
     */
    public List<GetCertificateCertificate> certificates() {
        return this.certificates;
    }
    /**
     * @return The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. Cannot be used with `url`.
     * 
     */
    public Optional<String> content() {
        return Optional.ofNullable(this.content);
    }
    /**
     * @return Unique identifier of this data source: hashing of the certificates in the chain.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The URL of the website to get the certificates from. Cannot be used with `content`.
     * 
     */
    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }
    /**
     * @return Whether to verify the certificate chain while parsing it or not (default: `true`). Cannot be used with `content`.
     * 
     */
    public Optional<Boolean> verifyChain() {
        return Optional.ofNullable(this.verifyChain);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCertificateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetCertificateCertificate> certificates;
        private @Nullable String content;
        private String id;
        private @Nullable String url;
        private @Nullable Boolean verifyChain;
        public Builder() {}
        public Builder(GetCertificateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certificates = defaults.certificates;
    	      this.content = defaults.content;
    	      this.id = defaults.id;
    	      this.url = defaults.url;
    	      this.verifyChain = defaults.verifyChain;
        }

        @CustomType.Setter
        public Builder certificates(List<GetCertificateCertificate> certificates) {
            this.certificates = Objects.requireNonNull(certificates);
            return this;
        }
        public Builder certificates(GetCertificateCertificate... certificates) {
            return certificates(List.of(certificates));
        }
        @CustomType.Setter
        public Builder content(@Nullable String content) {
            this.content = content;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder url(@Nullable String url) {
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder verifyChain(@Nullable Boolean verifyChain) {
            this.verifyChain = verifyChain;
            return this;
        }
        public GetCertificateResult build() {
            final var _resultValue = new GetCertificateResult();
            _resultValue.certificates = certificates;
            _resultValue.content = content;
            _resultValue.id = id;
            _resultValue.url = url;
            _resultValue.verifyChain = verifyChain;
            return _resultValue;
        }
    }
}
