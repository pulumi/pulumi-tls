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
        public readonly bool IsCa;
        public readonly string Issuer;
        public readonly string NotAfter;
        public readonly string NotBefore;
        public readonly string PublicKeyAlgorithm;
        public readonly string SerialNumber;
        public readonly string Sha1Fingerprint;
        public readonly string SignatureAlgorithm;
        public readonly string Subject;
        public readonly int Version;

        [OutputConstructor]
        private GetCertificateCertificateResult(
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
