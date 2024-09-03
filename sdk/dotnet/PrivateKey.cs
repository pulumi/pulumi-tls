// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    [TlsResourceType("tls:index/privateKey:PrivateKey")]
    public partial class PrivateKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        /// </summary>
        [Output("ecdsaCurve")]
        public Output<string> EcdsaCurve { get; private set; } = null!;

        /// <summary>
        /// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        /// </summary>
        [Output("privateKeyOpenssh")]
        public Output<string> PrivateKeyOpenssh { get; private set; } = null!;

        /// <summary>
        /// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        /// </summary>
        [Output("privateKeyPemPkcs8")]
        public Output<string> PrivateKeyPemPkcs8 { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Output("publicKeyFingerprintMd5")]
        public Output<string> PublicKeyFingerprintMd5 { get; private set; } = null!;

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Output("publicKeyFingerprintSha256")]
        public Output<string> PublicKeyFingerprintSha256 { get; private set; } = null!;

        /// <summary>
        /// The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        [Output("publicKeyOpenssh")]
        public Output<string> PublicKeyOpenssh { get; private set; } = null!;

        /// <summary>
        /// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        [Output("publicKeyPem")]
        public Output<string> PublicKeyPem { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        /// </summary>
        [Output("rsaBits")]
        public Output<int> RsaBits { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateKey(string name, PrivateKeyArgs args, CustomResourceOptions? options = null)
            : base("tls:index/privateKey:PrivateKey", name, args ?? new PrivateKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateKey(string name, Input<string> id, PrivateKeyState? state = null, CustomResourceOptions? options = null)
            : base("tls:index/privateKey:PrivateKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "privateKeyOpenssh",
                    "privateKeyPem",
                    "privateKeyPemPkcs8",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PrivateKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateKey Get(string name, Input<string> id, PrivateKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateKey(name, id, state, options);
        }
    }

    public sealed class PrivateKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<string>? EcdsaCurve { get; set; }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        public PrivateKeyArgs()
        {
        }
        public static new PrivateKeyArgs Empty => new PrivateKeyArgs();
    }

    public sealed class PrivateKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<string>? EcdsaCurve { get; set; }

        [Input("privateKeyOpenssh")]
        private Input<string>? _privateKeyOpenssh;

        /// <summary>
        /// Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        /// </summary>
        public Input<string>? PrivateKeyOpenssh
        {
            get => _privateKeyOpenssh;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyOpenssh = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyPem")]
        private Input<string>? _privateKeyPem;

        /// <summary>
        /// Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        /// </summary>
        public Input<string>? PrivateKeyPem
        {
            get => _privateKeyPem;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyPem = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateKeyPemPkcs8")]
        private Input<string>? _privateKeyPemPkcs8;

        /// <summary>
        /// Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        /// </summary>
        public Input<string>? PrivateKeyPemPkcs8
        {
            get => _privateKeyPemPkcs8;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKeyPemPkcs8 = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Input("publicKeyFingerprintMd5")]
        public Input<string>? PublicKeyFingerprintMd5 { get; set; }

        /// <summary>
        /// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        /// </summary>
        [Input("publicKeyFingerprintSha256")]
        public Input<string>? PublicKeyFingerprintSha256 { get; set; }

        /// <summary>
        /// The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        [Input("publicKeyOpenssh")]
        public Input<string>? PublicKeyOpenssh { get; set; }

        /// <summary>
        /// Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        /// <summary>
        /// When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        public PrivateKeyState()
        {
        }
        public static new PrivateKeyState Empty => new PrivateKeyState();
    }
}
