// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls
{
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source to get information, such as SHA1 fingerprint or serial number, about the TLS certificates that
        /// protect an HTTPS website. Note that the certificate chain isn't verified.
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("tls:index/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL of the website to get the certificates from.
        /// </summary>
        [Input("url", required: true)]
        public string Url { get; set; } = null!;

        /// <summary>
        /// Whether to verify the certificate chain while parsing it or not
        /// </summary>
        [Input("verifyChain")]
        public bool? VerifyChain { get; set; }

        public GetCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The certificates protecting the site, with the root of the chain first.
        /// * `certificates.#.not_after` - The time until which the certificate is invalid, as an
        /// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// * `certificates.#.not_before` - The time after which the certificate is valid, as an
        /// [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// * `certificates.#.is_ca` - `true` if this certificate is a ca certificate.
        /// * `certificates.#.issuer` - Who verified and signed the certificate, roughly following
        /// [RFC2253](https://tools.ietf.org/html/rfc2253).
        /// * `certificates.#.public_key_algorithm` - The algorithm used to create the certificate.
        /// * `certificates.#.serial_number` - Number that uniquely identifies the certificate with the CA's system. The `format`
        /// function can be used to convert this base 10 number into other bases, such as hex.
        /// * `certificates.#.sha1_fingerprint` - The SHA1 fingerprint of the public key of the certificate.
        /// * `certificates.#.signature_algorithm` - The algorithm used to sign the certificate.
        /// * `certificates.#.subject` - The entity the certificate belongs to, roughly following
        /// [RFC2253](https://tools.ietf.org/html/rfc2253).
        /// * `certificates.#.version` - The version the certificate is in.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCertificateCertificateResult> Certificates;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Url;
        public readonly bool? VerifyChain;

        [OutputConstructor]
        private GetCertificateResult(
            ImmutableArray<Outputs.GetCertificateCertificateResult> certificates,

            string id,

            string url,

            bool? verifyChain)
        {
            Certificates = certificates;
            Id = id;
            Url = url;
            VerifyChain = verifyChain;
        }
    }
}
