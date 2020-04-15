// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    public partial class CertRequest : Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate request data in PEM format.
        /// </summary>
        [Output("certRequestPem")]
        public Output<string> CertRequestPem { get; private set; } = null!;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        [Output("dnsNames")]
        public Output<ImmutableArray<string>> DnsNames { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// The name of the algorithm for the key provided
        /// in `private_key_pem`.
        /// </summary>
        [Output("keyAlgorithm")]
        public Output<string> KeyAlgorithm { get; private set; } = null!;

        /// <summary>
        /// PEM-encoded private key that the certificate will belong to
        /// </summary>
        [Output("privateKeyPem")]
        public Output<string> PrivateKeyPem { get; private set; } = null!;

        /// <summary>
        /// The subject for which a certificate is being requested. This is
        /// a nested configuration block whose structure is described below.
        /// </summary>
        [Output("subjects")]
        public Output<ImmutableArray<Outputs.CertRequestSubject>> Subjects { get; private set; } = null!;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
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

    public sealed class CertRequestArgs : Pulumi.ResourceArgs
    {
        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The name of the algorithm for the key provided
        /// in `private_key_pem`.
        /// </summary>
        [Input("keyAlgorithm", required: true)]
        public Input<string> KeyAlgorithm { get; set; } = null!;

        /// <summary>
        /// PEM-encoded private key that the certificate will belong to
        /// </summary>
        [Input("privateKeyPem", required: true)]
        public Input<string> PrivateKeyPem { get; set; } = null!;

        [Input("subjects", required: true)]
        private InputList<Inputs.CertRequestSubjectArgs>? _subjects;

        /// <summary>
        /// The subject for which a certificate is being requested. This is
        /// a nested configuration block whose structure is described below.
        /// </summary>
        public InputList<Inputs.CertRequestSubjectArgs> Subjects
        {
            get => _subjects ?? (_subjects = new InputList<Inputs.CertRequestSubjectArgs>());
            set => _subjects = value;
        }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public CertRequestArgs()
        {
        }
    }

    public sealed class CertRequestState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate request data in PEM format.
        /// </summary>
        [Input("certRequestPem")]
        public Input<string>? CertRequestPem { get; set; }

        [Input("dnsNames")]
        private InputList<string>? _dnsNames;

        /// <summary>
        /// List of DNS names for which a certificate is being requested.
        /// </summary>
        public InputList<string> DnsNames
        {
            get => _dnsNames ?? (_dnsNames = new InputList<string>());
            set => _dnsNames = value;
        }

        [Input("ipAddresses")]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested.
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        /// <summary>
        /// The name of the algorithm for the key provided
        /// in `private_key_pem`.
        /// </summary>
        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        /// <summary>
        /// PEM-encoded private key that the certificate will belong to
        /// </summary>
        [Input("privateKeyPem")]
        public Input<string>? PrivateKeyPem { get; set; }

        [Input("subjects")]
        private InputList<Inputs.CertRequestSubjectGetArgs>? _subjects;

        /// <summary>
        /// The subject for which a certificate is being requested. This is
        /// a nested configuration block whose structure is described below.
        /// </summary>
        public InputList<Inputs.CertRequestSubjectGetArgs> Subjects
        {
            get => _subjects ?? (_subjects = new InputList<Inputs.CertRequestSubjectGetArgs>());
            set => _subjects = value;
        }

        [Input("uris")]
        private InputList<string>? _uris;

        /// <summary>
        /// List of URIs for which a certificate is being requested.
        /// </summary>
        public InputList<string> Uris
        {
            get => _uris ?? (_uris = new InputList<string>());
            set => _uris = value;
        }

        public CertRequestState()
        {
        }
    }
}
