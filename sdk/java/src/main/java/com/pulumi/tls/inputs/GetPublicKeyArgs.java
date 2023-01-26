// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tls.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPublicKeyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPublicKeyArgs Empty = new GetPublicKeyArgs();

    /**
     * The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    @Import(name="privateKeyOpenssh")
    private @Nullable Output<String> privateKeyOpenssh;

    /**
     * @return The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    public Optional<Output<String>> privateKeyOpenssh() {
        return Optional.ofNullable(this.privateKeyOpenssh);
    }

    /**
     * The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    @Import(name="privateKeyPem")
    private @Nullable Output<String> privateKeyPem;

    /**
     * @return The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
     * 
     */
    public Optional<Output<String>> privateKeyPem() {
        return Optional.ofNullable(this.privateKeyPem);
    }

    private GetPublicKeyArgs() {}

    private GetPublicKeyArgs(GetPublicKeyArgs $) {
        this.privateKeyOpenssh = $.privateKeyOpenssh;
        this.privateKeyPem = $.privateKeyPem;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPublicKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPublicKeyArgs $;

        public Builder() {
            $ = new GetPublicKeyArgs();
        }

        public Builder(GetPublicKeyArgs defaults) {
            $ = new GetPublicKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param privateKeyOpenssh The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyOpenssh(@Nullable Output<String> privateKeyOpenssh) {
            $.privateKeyOpenssh = privateKeyOpenssh;
            return this;
        }

        /**
         * @param privateKeyOpenssh The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyOpenssh(String privateKeyOpenssh) {
            return privateKeyOpenssh(Output.of(privateKeyOpenssh));
        }

        /**
         * @param privateKeyPem The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(@Nullable Output<String> privateKeyPem) {
            $.privateKeyPem = privateKeyPem;
            return this;
        }

        /**
         * @param privateKeyPem The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
         * 
         * @return builder
         * 
         */
        public Builder privateKeyPem(String privateKeyPem) {
            return privateKeyPem(Output.of(privateKeyPem));
        }

        public GetPublicKeyArgs build() {
            return $;
        }
    }

}
