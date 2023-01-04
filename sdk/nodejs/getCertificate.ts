// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getCertificate(args?: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tls:index/getCertificate:getCertificate", {
        "content": args.content,
        "url": args.url,
        "verifyChain": args.verifyChain,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. Cannot be used with `url`.
     */
    content?: string;
    /**
     * The URL of the website to get the certificates from. Cannot be used with `content`.
     */
    url?: string;
    /**
     * Whether to verify the certificate chain while parsing it or not (default: `true`). Cannot be used with `content`.
     */
    verifyChain?: boolean;
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    /**
     * The certificates protecting the site, with the root of the chain first.
     */
    readonly certificates: outputs.GetCertificateCertificate[];
    /**
     * The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. Cannot be used with `url`.
     */
    readonly content?: string;
    /**
     * Unique identifier of this data source: hashing of the certificates in the chain.
     */
    readonly id: string;
    /**
     * The URL of the website to get the certificates from. Cannot be used with `content`.
     */
    readonly url?: string;
    /**
     * Whether to verify the certificate chain while parsing it or not (default: `true`). Cannot be used with `content`.
     */
    readonly verifyChain?: boolean;
}
export function getCertificateOutput(args?: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply((a: any) => getCertificate(a, opts))
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateOutputArgs {
    /**
     * The content of the certificate in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. Cannot be used with `url`.
     */
    content?: pulumi.Input<string>;
    /**
     * The URL of the website to get the certificates from. Cannot be used with `content`.
     */
    url?: pulumi.Input<string>;
    /**
     * Whether to verify the certificate chain while parsing it or not (default: `true`). Cannot be used with `content`.
     */
    verifyChain?: pulumi.Input<boolean>;
}
