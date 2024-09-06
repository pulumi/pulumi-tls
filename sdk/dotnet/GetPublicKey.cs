// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    public static class GetPublicKey
    {
        /// <summary>
        /// Get a public key from a PEM-encoded private key.
        /// 
        /// Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Std = Pulumi.Std;
        /// using Tls = Pulumi.Tls;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ed25519_example = new Tls.PrivateKey("ed25519-example", new()
        ///     {
        ///         Algorithm = "ED25519",
        ///     });
        /// 
        ///     // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
        ///     var privateKeyPem_example = Tls.GetPublicKey.Invoke(new()
        ///     {
        ///         PrivateKeyPem = ed25519_example.PrivateKeyPem,
        ///     });
        /// 
        ///     // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
        ///     var privateKeyOpenssh_example = Tls.GetPublicKey.Invoke(new()
        ///     {
        ///         PrivateKeyOpenssh = Std.File.Invoke(new()
        ///         {
        ///             Input = "~/.ssh/id_rsa_rfc4716",
        ///         }).Result,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPublicKeyResult> InvokeAsync(GetPublicKeyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPublicKeyResult>("tls:index/getPublicKey:getPublicKey", args ?? new GetPublicKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Get a public key from a PEM-encoded private key.
        /// 
        /// Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Std = Pulumi.Std;
        /// using Tls = Pulumi.Tls;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var ed25519_example = new Tls.PrivateKey("ed25519-example", new()
        ///     {
        ///         Algorithm = "ED25519",
        ///     });
        /// 
        ///     // Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
        ///     var privateKeyPem_example = Tls.GetPublicKey.Invoke(new()
        ///     {
        ///         PrivateKeyPem = ed25519_example.PrivateKeyPem,
        ///     });
        /// 
        ///     // Public key loaded from filesystem, using the Open SSH (RFC 4716) format
        ///     var privateKeyOpenssh_example = Tls.GetPublicKey.Invoke(new()
        ///     {
        ///         PrivateKeyOpenssh = Std.File.Invoke(new()
        ///         {
        ///             Input = "~/.ssh/id_rsa_rfc4716",
        ///         }).Result,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPublicKeyResult> Invoke(GetPublicKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPublicKeyResult>("tls:index/getPublicKey:getPublicKey", args ?? new GetPublicKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPublicKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("privateKeyOpenssh")]
        private string? _privateKeyOpenssh;

        /// <summary>
        /// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        public string? PrivateKeyOpenssh
        {
            get => _privateKeyOpenssh;
            set => _privateKeyOpenssh = value;
        }

        [Input("privateKeyPem")]
        private string? _privateKeyPem;

        /// <summary>
        /// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        public string? PrivateKeyPem
        {
            get => _privateKeyPem;
            set => _privateKeyPem = value;
        }

        public GetPublicKeyArgs()
        {
        }
        public static new GetPublicKeyArgs Empty => new GetPublicKeyArgs();
    }

    public sealed class GetPublicKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("privateKeyOpenssh")]
        private Input<string>? _privateKeyOpenssh;

        /// <summary>
        /// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
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
        /// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
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

        public GetPublicKeyInvokeArgs()
        {
        }
        public static new GetPublicKeyInvokeArgs Empty => new GetPublicKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPublicKeyResult
    {
        /// <summary>
        /// The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        public readonly string? PrivateKeyOpenssh;
        /// <summary>
        /// The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        /// </summary>
        public readonly string? PrivateKeyPem;
        /// <summary>
        /// The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
        /// </summary>
        public readonly string PublicKeyFingerprintMd5;
        /// <summary>
        /// The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
        /// </summary>
        public readonly string PublicKeyFingerprintSha256;
        /// <summary>
        /// The public key, in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format. This is also known as ['Authorized Keys'](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        public readonly string PublicKeyOpenssh;
        /// <summary>
        /// The public key, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        /// </summary>
        public readonly string PublicKeyPem;

        [OutputConstructor]
        private GetPublicKeyResult(
            string algorithm,

            string id,

            string? privateKeyOpenssh,

            string? privateKeyPem,

            string publicKeyFingerprintMd5,

            string publicKeyFingerprintSha256,

            string publicKeyOpenssh,

            string publicKeyPem)
        {
            Algorithm = algorithm;
            Id = id;
            PrivateKeyOpenssh = privateKeyOpenssh;
            PrivateKeyPem = privateKeyPem;
            PublicKeyFingerprintMd5 = publicKeyFingerprintMd5;
            PublicKeyFingerprintSha256 = publicKeyFingerprintSha256;
            PublicKeyOpenssh = publicKeyOpenssh;
            PublicKeyPem = publicKeyPem;
        }
    }
}
