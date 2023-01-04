// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using Pulumi;
    /// using Tls = Pulumi.Tls;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Tls.CertRequest("example", new()
    ///     {
    ///         PrivateKeyPem = File.ReadAllText("private_key.pem"),
    ///         Subject = new Tls.Inputs.CertRequestSubjectArgs
    ///         {
    ///             CommonName = "example.com",
    ///             Organization = "ACME Examples, Inc",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TlsResourceType("tls:index/certRequest:CertRequest")]
    public partial class CertRequest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
        /// [underlying](https://pkg.go.dev/encoding/pem#Encode)
        /// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
        /// the end of the PEM. In case this disrupts your use case, we recommend using
        /// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
        /// </summary>
        [Output("certRequestPem")]
        public Output<string> CertRequestPem { get; private set; } = null!;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("dnsNames")]
        public Output<ImmutableArray<string>> DnsNames { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
        /// </summary>
        [Output("keyAlgorithm")]
        public Output<string> KeyAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
        /// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
        /// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
        /// </summary>
        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Output("subject")]
        public Output<Outputs.CertRequestSubject?> Subject { get; private set; } = null!;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("uris")]
        public Output<ImmutableArray<string>> Uris { get; private set; } = null!;


        /// <summary>
        /// Create a CertRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertRequest(string name, CertRequestArgs args, CustomResourceOptions? options = null)
            : base("tls:index/certRequest:CertRequest", name, args ?? new CertRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertRequest(string name, Input<string> id, CertRequestState? state = null, CustomResourceOptions? options = null)
            : base("tls:index/certRequest:CertRequest", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "privateKeyPem",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertRequest Get(string name, Input<string> id, CertRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new CertRequest(name, id, state, options);
        }
    }

    public sealed class CertRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
        /// </summary>
        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        [Input("privateKeyPem", required: true)]
        private Input<string>? _privateKeyPem;

        /// <summary>
        /// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
        /// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
        /// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
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

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Input("subject")]
        public Input<Inputs.CertRequestSubjectArgs>? Subject { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public CertRequestArgs()
        {
        }
        public static new CertRequestArgs Empty => new CertRequestArgs();
    }

    public sealed class CertRequestState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
        /// [underlying](https://pkg.go.dev/encoding/pem#Encode)
        /// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
        /// the end of the PEM. In case this disrupts your use case, we recommend using
        /// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
        /// </summary>
        [Input("certRequestPem")]
        public Input<string>? CertRequestPem { get; set; }

        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`. **NOTE**: this is deprecated and ignored, as the key algorithm is now inferred from the key.
        /// </summary>
        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        [Input("privateKeyPem")]
        private Input<string>? _privateKeyPem;

        /// <summary>
        /// Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
        /// to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
        /// interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
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

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Input("subject")]
        public Input<Inputs.CertRequestSubjectGetArgs>? Subject { get; set; }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public CertRequestState()
        {
        }
        public static new CertRequestState Empty => new CertRequestState();
    }
}
