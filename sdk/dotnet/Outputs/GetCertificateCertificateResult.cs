// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tls.Outputs
{

    [OutputType]
    public sealed class GetCertificateCertificateResult
    {
        /// <summary>
        /// Certificate data in PEM (RFC 1421).
        /// </summary>
        public readonly string CertPem;
        /// <summary>
        /// `true` if the certificate is of a CA (Certificate Authority).
        /// </summary>
        public readonly bool IsCa;
        /// <summary>
        /// Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        public readonly string NotAfter;
        /// <summary>
        /// The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        /// </summary>
        public readonly string NotBefore;
        /// <summary>
        /// The key algorithm used to create the certificate.
        /// </summary>
        public readonly string PublicKeyAlgorithm;
        /// <summary>
        /// Number that uniquely identifies the certificate with the CA's system.
        /// The `format` function can be used to convert this *base 10* number into other bases, such as hex.
        /// </summary>
        public readonly string SerialNumber;
        /// <summary>
        /// The SHA1 fingerprint of the public key of the certificate.
        /// </summary>
        public readonly string Sha1Fingerprint;
        /// <summary>
        /// The algorithm used to sign the certificate.
        /// </summary>
        public readonly string SignatureAlgorithm;
        /// <summary>
        /// The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        /// </summary>
        public readonly string Subject;
        /// <summary>
        /// The version the certificate is in.
        /// </summary>
        public readonly int Version;

        [OutputConstructor]
        private GetCertificateCertificateResult(
            string certPem,

            bool isCa,

            string issuer,

            string notAfter,

            string notBefore,

            string publicKeyAlgorithm,

            string serialNumber,

            string sha1Fingerprint,

            string signatureAlgorithm,

            string subject,

            int version)
        {
            CertPem = certPem;
            IsCa = isCa;
            Issuer = issuer;
            NotAfter = notAfter;
            NotBefore = notBefore;
            PublicKeyAlgorithm = publicKeyAlgorithm;
            SerialNumber = serialNumber;
            Sha1Fingerprint = sha1Fingerprint;
            SignatureAlgorithm = signatureAlgorithm;
            Subject = subject;
            Version = version;
        }
    }
}
