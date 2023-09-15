// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPublicKeyResult {
    /**
     * @return The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA` and `ED25519`.
     * 
     */
    private String algorithm;
    /**
     * @return Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
     * 
     */
    private String id;
    /**
     * @return The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `private_key_pem`.
     * 
     */
    private @Nullable String privateKeyOpenssh;
    /**
     * @return The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `private_key_openssh`.
     * 
     */
    private @Nullable String privateKeyPem;
    /**
     * @return The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
     * 
     */
    private String publicKeyFingerprintMd5;
    /**
     * @return The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
     * 
     */
    private String publicKeyFingerprintSha256;
    /**
     * @return The public key, in  OpenSSH PEM (RFC 4716).
     * 
     */
    private String publicKeyOpenssh;
    /**
     * @return The public key, in PEM (RFC 1421).
     * 
     */
    private String publicKeyPem;

    private GetPublicKeyResult() {}
    /**
     * @return The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA` and `ED25519`.
     * 
     */
    public String algorithm() {
        return this.algorithm;
    }
    /**
     * @return Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `private_key_pem`.
     * 
     */
    public Optional<String> privateKeyOpenssh() {
        return Optional.ofNullable(this.privateKeyOpenssh);
    }
    /**
     * @return The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. Currently-supported algorithms for keys are `RSA`, `ECDSA` and `ED25519`. This is *mutually exclusive* with `private_key_openssh`.
     * 
     */
    public Optional<String> privateKeyPem() {
        return Optional.ofNullable(this.privateKeyPem);
    }
    /**
     * @return The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
     * 
     */
    public String publicKeyFingerprintMd5() {
        return this.publicKeyFingerprintMd5;
    }
    /**
     * @return The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
     * 
     */
    public String publicKeyFingerprintSha256() {
        return this.publicKeyFingerprintSha256;
    }
    /**
     * @return The public key, in  OpenSSH PEM (RFC 4716).
     * 
     */
    public String publicKeyOpenssh() {
        return this.publicKeyOpenssh;
    }
    /**
     * @return The public key, in PEM (RFC 1421).
     * 
     */
    public String publicKeyPem() {
        return this.publicKeyPem;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPublicKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String algorithm;
        private String id;
        private @Nullable String privateKeyOpenssh;
        private @Nullable String privateKeyPem;
        private String publicKeyFingerprintMd5;
        private String publicKeyFingerprintSha256;
        private String publicKeyOpenssh;
        private String publicKeyPem;
        public Builder() {}
        public Builder(GetPublicKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.id = defaults.id;
    	      this.privateKeyOpenssh = defaults.privateKeyOpenssh;
    	      this.privateKeyPem = defaults.privateKeyPem;
    	      this.publicKeyFingerprintMd5 = defaults.publicKeyFingerprintMd5;
    	      this.publicKeyFingerprintSha256 = defaults.publicKeyFingerprintSha256;
    	      this.publicKeyOpenssh = defaults.publicKeyOpenssh;
    	      this.publicKeyPem = defaults.publicKeyPem;
        }

        @CustomType.Setter
        public Builder algorithm(String algorithm) {
            this.algorithm = Objects.requireNonNull(algorithm);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder privateKeyOpenssh(@Nullable String privateKeyOpenssh) {
            this.privateKeyOpenssh = privateKeyOpenssh;
            return this;
        }
        @CustomType.Setter
        public Builder privateKeyPem(@Nullable String privateKeyPem) {
            this.privateKeyPem = privateKeyPem;
            return this;
        }
        @CustomType.Setter
        public Builder publicKeyFingerprintMd5(String publicKeyFingerprintMd5) {
            this.publicKeyFingerprintMd5 = Objects.requireNonNull(publicKeyFingerprintMd5);
            return this;
        }
        @CustomType.Setter
        public Builder publicKeyFingerprintSha256(String publicKeyFingerprintSha256) {
            this.publicKeyFingerprintSha256 = Objects.requireNonNull(publicKeyFingerprintSha256);
            return this;
        }
        @CustomType.Setter
        public Builder publicKeyOpenssh(String publicKeyOpenssh) {
            this.publicKeyOpenssh = Objects.requireNonNull(publicKeyOpenssh);
            return this;
        }
        @CustomType.Setter
        public Builder publicKeyPem(String publicKeyPem) {
            this.publicKeyPem = Objects.requireNonNull(publicKeyPem);
            return this;
        }
        public GetPublicKeyResult build() {
            final var o = new GetPublicKeyResult();
            o.algorithm = algorithm;
            o.id = id;
            o.privateKeyOpenssh = privateKeyOpenssh;
            o.privateKeyPem = privateKeyPem;
            o.publicKeyFingerprintMd5 = publicKeyFingerprintMd5;
            o.publicKeyFingerprintSha256 = publicKeyFingerprintSha256;
            o.publicKeyOpenssh = publicKeyOpenssh;
            o.publicKeyPem = publicKeyPem;
            return o;
        }
    }
}
