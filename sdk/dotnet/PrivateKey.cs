// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-tls/blob/master/website/docs/r/private_key.html.markdown.
    /// </summary>
    public partial class PrivateKey : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the algorithm to use for
        /// the key. Currently-supported values are "RSA" and "ECDSA".
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is "ECDSA", the name of the elliptic
        /// curve to use. May be any one of "P224", "P256", "P384" or "P521", with "P224" as the
        /// default.
        /// </summary>
        [Output("ecdsaCurve")]
        public Output<string?> EcdsaCurve { get; private set; } = null!;

        /// <summary>
        /// The private key data in PEM format.
        /// </summary>
        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// The md5 hash of the public key data in
        /// OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the
        /// selected private key format is compatible, as per the rules for
        /// `public_key_openssh`.
        /// </summary>
        [Output("publicKeyFingerprintMd5")]
        public Output<string> PublicKeyFingerprintMd5 { get; private set; } = null!;

        /// <summary>
        /// The public key data in OpenSSH `authorized_keys`
        /// format, if the selected private key format is compatible. All RSA keys
        /// are supported, and ECDSA keys with curves "P256", "P384" and "P521"
        /// are supported. This attribute is empty if an incompatible ECDSA curve
        /// is selected.
        /// </summary>
        [Output("publicKeyOpenssh")]
        public Output<string> PublicKeyOpenssh { get; private set; } = null!;

        /// <summary>
        /// The public key data in PEM format.
        /// </summary>
        [Output("publicKeyPem")]
        public Output<string> PublicKeyPem { get; private set; } = null!;

        /// <summary>
        /// When `algorithm` is "RSA", the size of the generated
        /// RSA key in bits. Defaults to 2048.
        /// </summary>
        [Output("rsaBits")]
        public Output<int?> RsaBits { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateKey(string name, PrivateKeyArgs args, CustomResourceOptions? options = null)
            : base("tls:index/privateKey:PrivateKey", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class PrivateKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the algorithm to use for
        /// the key. Currently-supported values are "RSA" and "ECDSA".
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// When `algorithm` is "ECDSA", the name of the elliptic
        /// curve to use. May be any one of "P224", "P256", "P384" or "P521", with "P224" as the
        /// default.
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<string>? EcdsaCurve { get; set; }

        /// <summary>
        /// When `algorithm` is "RSA", the size of the generated
        /// RSA key in bits. Defaults to 2048.
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        public PrivateKeyArgs()
        {
        }
    }

    public sealed class PrivateKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the algorithm to use for
        /// the key. Currently-supported values are "RSA" and "ECDSA".
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// When `algorithm` is "ECDSA", the name of the elliptic
        /// curve to use. May be any one of "P224", "P256", "P384" or "P521", with "P224" as the
        /// default.
        /// </summary>
        [Input("ecdsaCurve")]
        public Input<string>? EcdsaCurve { get; set; }

        /// <summary>
        /// The private key data in PEM format.
        /// </summary>
        [Input("privateKeyPem")]
        public Input<string>? PrivateKeyPem { get; set; }

        /// <summary>
        /// The md5 hash of the public key data in
        /// OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the
        /// selected private key format is compatible, as per the rules for
        /// `public_key_openssh`.
        /// </summary>
        [Input("publicKeyFingerprintMd5")]
        public Input<string>? PublicKeyFingerprintMd5 { get; set; }

        /// <summary>
        /// The public key data in OpenSSH `authorized_keys`
        /// format, if the selected private key format is compatible. All RSA keys
        /// are supported, and ECDSA keys with curves "P256", "P384" and "P521"
        /// are supported. This attribute is empty if an incompatible ECDSA curve
        /// is selected.
        /// </summary>
        [Input("publicKeyOpenssh")]
        public Input<string>? PublicKeyOpenssh { get; set; }

        /// <summary>
        /// The public key data in PEM format.
        /// </summary>
        [Input("publicKeyPem")]
        public Input<string>? PublicKeyPem { get; set; }

        /// <summary>
        /// When `algorithm` is "RSA", the size of the generated
        /// RSA key in bits. Defaults to 2048.
        /// </summary>
        [Input("rsaBits")]
        public Input<int>? RsaBits { get; set; }

        public PrivateKeyState()
        {
        }
    }
}
