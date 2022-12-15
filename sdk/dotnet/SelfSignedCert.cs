// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    [TlsResourceType("tls:index/selfSignedCert:SelfSignedCert")]
    public partial class SelfSignedCert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        /// </summary>
        [Output("allowedUses")]
        public Output<ImmutableArray<string>> AllowedUses { get; private set; } = null!;

        /// <summary>
        /// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
        /// [underlying](https://pkg.go.dev/encoding/pem#Encode)
        /// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
        /// the end of the PEM. In case this disrupts your use case, we recommend using
        /// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
        /// </summary>
        [Output("certPem")]
        public Output<string> CertPem { get; private set; } = null!;

        /// <summary>
        /// List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("dnsNames")]
        public Output<ImmutableArray<string>> DnsNames { get; private set; } = null!;

        /// <summary>
        /// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
        /// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
        /// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
        /// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
        /// early renewal period. (default: `0`)
        /// </summary>
        [Output("earlyRenewalHours")]
        public Output<int> EarlyRenewalHours { get; private set; } = null!;

        /// <summary>
        /// List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("ipAddresses")]
        public Output<ImmutableArray<string>> IpAddresses { get; private set; } = null!;

        /// <summary>
        /// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        /// </summary>
        [Output("isCaCertificate")]
        public Output<bool> IsCaCertificate { get; private set; } = null!;

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`.
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
        /// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        /// </summary>
        [Output("readyForRenewal")]
        public Output<bool> ReadyForRenewal { get; private set; } = null!;

        /// <summary>
        /// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Output("setAuthorityKeyId")]
        public Output<bool> SetAuthorityKeyId { get; private set; } = null!;

        /// <summary>
        /// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Output("setSubjectKeyId")]
        public Output<bool> SetSubjectKeyId { get; private set; } = null!;

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Output("subject")]
        public Output<Outputs.SelfSignedCertSubject?> Subject { get; private set; } = null!;

        /// <summary>
        /// List of URIs for which a certificate is being requested (i.e. certificate subjects).
        /// </summary>
        [Output("uris")]
        public Output<ImmutableArray<string>> Uris { get; private set; } = null!;

        /// <summary>
        /// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Output("validityEndTime")]
        public Output<string> ValidityEndTime { get; private set; } = null!;

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid for.
        /// </summary>
        [Output("validityPeriodHours")]
        public Output<int> ValidityPeriodHours { get; private set; } = null!;

        /// <summary>
        /// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Output("validityStartTime")]
        public Output<string> ValidityStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a SelfSignedCert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SelfSignedCert(string name, SelfSignedCertArgs args, CustomResourceOptions? options = null)
            : base("tls:index/selfSignedCert:SelfSignedCert", name, args ?? new SelfSignedCertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SelfSignedCert(string name, Input<string> id, SelfSignedCertState? state = null, CustomResourceOptions? options = null)
            : base("tls:index/selfSignedCert:SelfSignedCert", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SelfSignedCert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SelfSignedCert Get(string name, Input<string> id, SelfSignedCertState? state = null, CustomResourceOptions? options = null)
        {
            return new SelfSignedCert(name, id, state, options);
        }
    }

    public sealed class SelfSignedCertArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedUses", required: true)]
        private InputList<string>? _allowedUses;

        /// <summary>
        /// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        /// </summary>
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
            set => _allowedUses = value;
        }

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

        /// <summary>
        /// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
        /// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
        /// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
        /// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
        /// early renewal period. (default: `0`)
        /// </summary>
        [Input("earlyRenewalHours")]
        public Input<int>? EarlyRenewalHours { get; set; }

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
        /// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        /// </summary>
        [Input("isCaCertificate")]
        public Input<bool>? IsCaCertificate { get; set; }

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
        /// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Input("setAuthorityKeyId")]
        public Input<bool>? SetAuthorityKeyId { get; set; }

        /// <summary>
        /// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Input("setSubjectKeyId")]
        public Input<bool>? SetSubjectKeyId { get; set; }

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Input("subject")]
        public Input<Inputs.SelfSignedCertSubjectArgs>? Subject { get; set; }

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

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid for.
        /// </summary>
        [Input("validityPeriodHours", required: true)]
        public Input<int> ValidityPeriodHours { get; set; } = null!;

        public SelfSignedCertArgs()
        {
        }
        public static new SelfSignedCertArgs Empty => new SelfSignedCertArgs();
    }

    public sealed class SelfSignedCertState : global::Pulumi.ResourceArgs
    {
        [Input("allowedUses")]
        private InputList<string>? _allowedUses;

        /// <summary>
        /// List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        /// </summary>
        public InputList<string> AllowedUses
        {
            get => _allowedUses ?? (_allowedUses = new InputList<string>());
            set => _allowedUses = value;
        }

        /// <summary>
        /// Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
        /// [underlying](https://pkg.go.dev/encoding/pem#Encode)
        /// [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
        /// the end of the PEM. In case this disrupts your use case, we recommend using
        /// [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
        /// </summary>
        [Input("certPem")]
        public Input<string>? CertPem { get; set; }

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

        /// <summary>
        /// The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
        /// can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
        /// certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
        /// revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
        /// early renewal period. (default: `0`)
        /// </summary>
        [Input("earlyRenewalHours")]
        public Input<int>? EarlyRenewalHours { get; set; }

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
        /// Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        /// </summary>
        [Input("isCaCertificate")]
        public Input<bool>? IsCaCertificate { get; set; }

        /// <summary>
        /// Name of the algorithm used when generating the private key provided in `private_key_pem`.
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
        /// Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        /// </summary>
        [Input("readyForRenewal")]
        public Input<bool>? ReadyForRenewal { get; set; }

        /// <summary>
        /// Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Input("setAuthorityKeyId")]
        public Input<bool>? SetAuthorityKeyId { get; set; }

        /// <summary>
        /// Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        /// </summary>
        [Input("setSubjectKeyId")]
        public Input<bool>? SetSubjectKeyId { get; set; }

        /// <summary>
        /// The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        /// </summary>
        [Input("subject")]
        public Input<Inputs.SelfSignedCertSubjectGetArgs>? Subject { get; set; }

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

        /// <summary>
        /// The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Input("validityEndTime")]
        public Input<string>? ValidityEndTime { get; set; }

        /// <summary>
        /// Number of hours, after initial issuing, that the certificate will remain valid for.
        /// </summary>
        [Input("validityPeriodHours")]
        public Input<int>? ValidityPeriodHours { get; set; }

        /// <summary>
        /// The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        [Input("validityStartTime")]
        public Input<string>? ValidityStartTime { get; set; }

        public SelfSignedCertState()
        {
        }
        public static new SelfSignedCertState Empty => new SelfSignedCertState();
    }
}
