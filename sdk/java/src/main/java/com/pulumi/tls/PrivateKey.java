// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tls.PrivateKeyArgs;
import com.pulumi.tls.Utilities;
import com.pulumi.tls.inputs.PrivateKeyState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

@ResourceType(type="tls:index/privateKey:PrivateKey")
public class PrivateKey extends com.pulumi.resources.CustomResource {
    /**
     * Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    @Export(name="algorithm", refs={String.class}, tree="[0]")
    private Output<String> algorithm;

    /**
     * @return Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    public Output<String> algorithm() {
        return this.algorithm;
    }
    /**
     * When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    @Export(name="ecdsaCurve", refs={String.class}, tree="[0]")
    private Output<String> ecdsaCurve;

    /**
     * @return When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
     * 
     */
    public Output<String> ecdsaCurve() {
        return this.ecdsaCurve;
    }
    /**
     * Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    @Export(name="privateKeyOpenssh", refs={String.class}, tree="[0]")
    private Output<String> privateKeyOpenssh;

    /**
     * @return Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
     * 
     */
    public Output<String> privateKeyOpenssh() {
        return this.privateKeyOpenssh;
    }
    /**
     * Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    @Export(name="privateKeyPem", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPem;

    /**
     * @return Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     * 
     */
    public Output<String> privateKeyPem() {
        return this.privateKeyPem;
    }
    /**
     * Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
     * 
     */
    @Export(name="privateKeyPemPkcs8", refs={String.class}, tree="[0]")
    private Output<String> privateKeyPemPkcs8;

    /**
     * @return Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
     * 
     */
    public Output<String> privateKeyPemPkcs8() {
        return this.privateKeyPemPkcs8;
    }
    /**
     * The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Export(name="publicKeyFingerprintMd5", refs={String.class}, tree="[0]")
    private Output<String> publicKeyFingerprintMd5;

    /**
     * @return The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Output<String> publicKeyFingerprintMd5() {
        return this.publicKeyFingerprintMd5;
    }
    /**
     * The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    @Export(name="publicKeyFingerprintSha256", refs={String.class}, tree="[0]")
    private Output<String> publicKeyFingerprintSha256;

    /**
     * @return The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
     * 
     */
    public Output<String> publicKeyFingerprintSha256() {
        return this.publicKeyFingerprintSha256;
    }
    /**
     * The public key data in [&#34;Authorized Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     * 
     */
    @Export(name="publicKeyOpenssh", refs={String.class}, tree="[0]")
    private Output<String> publicKeyOpenssh;

    /**
     * @return The public key data in [&#34;Authorized Keys&#34;](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     * 
     */
    public Output<String> publicKeyOpenssh() {
        return this.publicKeyOpenssh;
    }
    /**
     * Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     * 
     */
    @Export(name="publicKeyPem", refs={String.class}, tree="[0]")
    private Output<String> publicKeyPem;

    /**
     * @return Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     * 
     */
    public Output<String> publicKeyPem() {
        return this.publicKeyPem;
    }
    /**
     * When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    @Export(name="rsaBits", refs={Integer.class}, tree="[0]")
    private Output<Integer> rsaBits;

    /**
     * @return When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
     * 
     */
    public Output<Integer> rsaBits() {
        return this.rsaBits;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateKey(java.lang.String name) {
        this(name, PrivateKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateKey(java.lang.String name, PrivateKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateKey(java.lang.String name, PrivateKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/privateKey:PrivateKey", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PrivateKey(java.lang.String name, Output<java.lang.String> id, @Nullable PrivateKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tls:index/privateKey:PrivateKey", name, state, makeResourceOptions(options, id), false);
    }

    private static PrivateKeyArgs makeArgs(PrivateKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PrivateKeyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKeyOpenssh",
                "privateKeyPem",
                "privateKeyPemPkcs8"
            ))
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
    public static PrivateKey get(java.lang.String name, Output<java.lang.String> id, @Nullable PrivateKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateKey(name, id, state, options);
    }
}
