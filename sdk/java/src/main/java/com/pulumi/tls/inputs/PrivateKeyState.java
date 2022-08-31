// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateKeyState extends com.pulumi.resources.ResourceArgs {

    public static final PrivateKeyState Empty = new PrivateKeyState();

    /**
     * Name of the algorithm to use when generating the private key. Currently-supported values are `RSA`, `ECDSA` and `ED25519`.
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key. Currently-supported values are `RSA`, `ECDSA` and `ED25519`.
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are `P224`, `P256`, `P384` or `P521` (default: `P224`).
     * 
     */
    @Import(name="ecdsaCurve")
    private @Nullable Output<String> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are `P224`, `P256`, `P384` or `P521` (default: `P224`).
     * 
     */
    public Optional<Output<String>> ecdsaCurve() {
        return Optional.ofNullable(this.ecdsaCurve);
    }

    /**
     * Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    @Import(name="privateKeyOpenssh")
    private @Nullable Output<String> privateKeyOpenssh;

    /**
     * @return Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    public Optional<Output<String>> privateKeyOpenssh() {
        return Optional.ofNullable(this.privateKeyOpenssh);
    }

    /**
     * Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Import(name="privateKeyPem")
    private @Nullable Output<String> privateKeyPem;

    /**
     * @return Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Optional<Output<String>> privateKeyPem() {
        return Optional.ofNullable(this.privateKeyPem);
    }

    /**
     * The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Import(name="publicKeyFingerprintMd5")
    private @Nullable Output<String> publicKeyFingerprintMd5;

    /**
     * @return The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Optional<Output<String>> publicKeyFingerprintMd5() {
        return Optional.ofNullable(this.publicKeyFingerprintMd5);
    }

    /**
     * The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Import(name="publicKeyFingerprintSha256")
    private @Nullable Output<String> publicKeyFingerprintSha256;

    /**
     * @return The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Optional<Output<String>> publicKeyFingerprintSha256() {
        return Optional.ofNullable(this.publicKeyFingerprintSha256);
    }

    /**
     * The public key data in [&#34;Authorized
     * Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is
     * populated only if the configured private key is supported: this includes all `RSA` and `ED25519` keys, as well as
     * `ECDSA` keys with curves `P256`, `P384` and `P521`. `ECDSA` with curve `P224` [is not
     * supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    @Import(name="publicKeyOpenssh")
    private @Nullable Output<String> publicKeyOpenssh;

    /**
     * @return The public key data in [&#34;Authorized
     * Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is
     * populated only if the configured private key is supported: this includes all `RSA` and `ED25519` keys, as well as
     * `ECDSA` keys with curves `P256`, `P384` and `P521`. `ECDSA` with curve `P224` [is not
     * supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    public Optional<Output<String>> publicKeyOpenssh() {
        return Optional.ofNullable(this.publicKeyOpenssh);
    }

    /**
     * Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    @Import(name="publicKeyPem")
    private @Nullable Output<String> publicKeyPem;

    /**
     * @return Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     * 
     */
    public Optional<Output<String>> publicKeyPem() {
        return Optional.ofNullable(this.publicKeyPem);
    }

    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    @Import(name="rsaBits")
    private @Nullable Output<Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    public Optional<Output<Integer>> rsaBits() {
        return Optional.ofNullable(this.rsaBits);
    }

    private PrivateKeyState() {}

    private PrivateKeyState(PrivateKeyState $) {
        this.algorithm = $.algorithm;
        this.ecdsaCurve = $.ecdsaCurve;
        this.privateKeyOpenssh = $.privateKeyOpenssh;
        this.privateKeyPem = $.privateKeyPem;
        this.publicKeyFingerprintMd5 = $.publicKeyFingerprintMd5;
        this.publicKeyFingerprintSha256 = $.publicKeyFingerprintSha256;
        this.publicKeyOpenssh = $.publicKeyOpenssh;
        this.publicKeyPem = $.publicKeyPem;
        this.rsaBits = $.rsaBits;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateKeyState $;

        public Builder() {
            $ = new PrivateKeyState();
        }

        public Builder(PrivateKeyState defaults) {
            $ = new PrivateKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key. Currently-supported values are `RSA`, `ECDSA` and `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm Name of the algorithm to use when generating the private key. Currently-supported values are `RSA`, `ECDSA` and `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are `P224`, `P256`, `P384` or `P521` (default: `P224`).
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(@Nullable Output<String> ecdsaCurve) {
            $.ecdsaCurve = ecdsaCurve;
            return this;
        }

        /**
         * @param ecdsaCurve When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are `P224`, `P256`, `P384` or `P521` (default: `P224`).
         * 
         * @return builder
         * 
         */
        public Builder ecdsaCurve(String ecdsaCurve) {
            return ecdsaCurve(Output.of(ecdsaCurve));
        }

        /**
         * @param privateKeyOpenssh Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyOpenssh(@Nullable Output<String> privateKeyOpenssh) {
            $.privateKeyOpenssh = privateKeyOpenssh;
            return this;
        }

        /**
         * @param privateKeyOpenssh Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyOpenssh(String privateKeyOpenssh) {
            return privateKeyOpenssh(Output.of(privateKeyOpenssh));
        }

        /**
         * @param privateKeyPem Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(@Nullable Output<String> privateKeyPem) {
            $.privateKeyPem = privateKeyPem;
            return this;
        }

        /**
         * @param privateKeyPem Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(String privateKeyPem) {
            return privateKeyPem(Output.of(privateKeyPem));
        }

        /**
         * @param publicKeyFingerprintMd5 The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
         * 
         * @return builder
         * 
         */
        public Builder publicKeyFingerprintMd5(@Nullable Output<String> publicKeyFingerprintMd5) {
            $.publicKeyFingerprintMd5 = publicKeyFingerprintMd5;
            return this;
        }

        /**
         * @param publicKeyFingerprintMd5 The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
         * 
         * @return builder
         * 
         */
        public Builder publicKeyFingerprintMd5(String publicKeyFingerprintMd5) {
            return publicKeyFingerprintMd5(Output.of(publicKeyFingerprintMd5));
        }

        /**
         * @param publicKeyFingerprintSha256 The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
         * 
         * @return builder
         * 
         */
        public Builder publicKeyFingerprintSha256(@Nullable Output<String> publicKeyFingerprintSha256) {
            $.publicKeyFingerprintSha256 = publicKeyFingerprintSha256;
            return this;
        }

        /**
         * @param publicKeyFingerprintSha256 The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
         * 
         * @return builder
         * 
         */
        public Builder publicKeyFingerprintSha256(String publicKeyFingerprintSha256) {
            return publicKeyFingerprintSha256(Output.of(publicKeyFingerprintSha256));
        }

        /**
         * @param publicKeyOpenssh The public key data in [&#34;Authorized
         * Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is
         * populated only if the configured private key is supported: this includes all `RSA` and `ED25519` keys, as well as
         * `ECDSA` keys with curves `P256`, `P384` and `P521`. `ECDSA` with curve `P224` [is not
         * supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode)
         * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
         * the end of the PEM. In case this disrupts your use case, we recommend using
         * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
         * 
         * @return builder
         * 
         */
        public Builder publicKeyOpenssh(@Nullable Output<String> publicKeyOpenssh) {
            $.publicKeyOpenssh = publicKeyOpenssh;
            return this;
        }

        /**
         * @param publicKeyOpenssh The public key data in [&#34;Authorized
         * Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is
         * populated only if the configured private key is supported: this includes all `RSA` and `ED25519` keys, as well as
         * `ECDSA` keys with curves `P256`, `P384` and `P521`. `ECDSA` with curve `P224` [is not
         * supported](../../docs#limitations). **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode)
         * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
         * the end of the PEM. In case this disrupts your use case, we recommend using
         * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
         * 
         * @return builder
         * 
         */
        public Builder publicKeyOpenssh(String publicKeyOpenssh) {
            return publicKeyOpenssh(Output.of(publicKeyOpenssh));
        }

        /**
         * @param publicKeyPem Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
         * [underlying](https://pkg.go.dev/encoding/pem#Encode)
         * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
         * the end of the PEM. In case this disrupts your use case, we recommend using
         * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
         * 
         * @return builder
         * 
         */
        public Builder publicKeyPem(@Nullable Output<String> publicKeyPem) {
            $.publicKeyPem = publicKeyPem;
            return this;
        }

        /**
         * @param publicKeyPem Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
         * [underlying](https://pkg.go.dev/encoding/pem#Encode)
         * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
         * the end of the PEM. In case this disrupts your use case, we recommend using
         * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
         * 
         * @return builder
         * 
         */
        public Builder publicKeyPem(String publicKeyPem) {
            return publicKeyPem(Output.of(publicKeyPem));
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(@Nullable Output<Integer> rsaBits) {
            $.rsaBits = rsaBits;
            return this;
        }

        /**
         * @param rsaBits When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
         * 
         * @return builder
         * 
         */
        public Builder rsaBits(Integer rsaBits) {
            return rsaBits(Output.of(rsaBits));
        }

        public PrivateKeyState build() {
            return $;
        }
    }

}
