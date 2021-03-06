// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface CertRequestSubject {
    commonName?: string;
    country?: string;
    locality?: string;
    organization?: string;
    organizationalUnit?: string;
    postalCode?: string;
    province?: string;
    serialNumber?: string;
    streetAddresses?: string[];
}

export interface GetCertificateCertificate {
    isCa: boolean;
    issuer: string;
    notAfter: string;
    notBefore: string;
    publicKeyAlgorithm: string;
    serialNumber: string;
    sha1Fingerprint: string;
    signatureAlgorithm: string;
    subject: string;
    version: number;
}

export interface SelfSignedCertSubject {
    commonName?: string;
    country?: string;
    locality?: string;
    organization?: string;
    organizationalUnit?: string;
    postalCode?: string;
    province?: string;
    serialNumber?: string;
    streetAddresses?: string[];
}
