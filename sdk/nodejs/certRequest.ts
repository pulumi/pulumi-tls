// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as tls from "@pulumi/tls";
 *
 * const example = new tls.CertRequest("example", {
 *     privateKeyPem: fs.readFileSync("private_key.pem", "utf8"),
 *     subject: {
 *         commonName: "example.com",
 *         organization: "ACME Examples, Inc",
 *     },
 * });
 * ```
 */
export class CertRequest extends pulumi.CustomResource {
    /**
     * Get an existing CertRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertRequestState, opts?: pulumi.CustomResourceOptions): CertRequest {
        return new CertRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tls:index/certRequest:CertRequest';

    /**
     * Returns true if the given object is an instance of CertRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertRequest.__pulumiType;
    }

    /**
     * The certificate request data in PEM (RFC 1421).
     */
    public /*out*/ readonly certRequestPem!: pulumi.Output<string>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly dnsNames!: pulumi.Output<string[] | undefined>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the algorithm used when generating the private key provided in `privateKeyPem`.
     */
    public /*out*/ readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * Private key in PEM (RFC 1421) interpolation function.
     */
    public readonly privateKeyPem!: pulumi.Output<string>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    public readonly subject!: pulumi.Output<outputs.CertRequestSubject | undefined>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly uris!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CertRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertRequestArgs | CertRequestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertRequestState | undefined;
            resourceInputs["certRequestPem"] = state ? state.certRequestPem : undefined;
            resourceInputs["dnsNames"] = state ? state.dnsNames : undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            resourceInputs["privateKeyPem"] = state ? state.privateKeyPem : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["uris"] = state ? state.uris : undefined;
        } else {
            const args = argsOrState as CertRequestArgs | undefined;
            if ((!args || args.privateKeyPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKeyPem'");
            }
            resourceInputs["dnsNames"] = args ? args.dnsNames : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["privateKeyPem"] = args?.privateKeyPem ? pulumi.secret(args.privateKeyPem) : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["uris"] = args ? args.uris : undefined;
            resourceInputs["certRequestPem"] = undefined /*out*/;
            resourceInputs["keyAlgorithm"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKeyPem"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CertRequest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertRequest resources.
 */
export interface CertRequestState {
    /**
     * The certificate request data in PEM (RFC 1421).
     */
    certRequestPem?: pulumi.Input<string>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the algorithm used when generating the private key provided in `privateKeyPem`.
     */
    keyAlgorithm?: pulumi.Input<string>;
    /**
     * Private key in PEM (RFC 1421) interpolation function.
     */
    privateKeyPem?: pulumi.Input<string>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    subject?: pulumi.Input<inputs.CertRequestSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a CertRequest resource.
 */
export interface CertRequestArgs {
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Private key in PEM (RFC 1421) interpolation function.
     */
    privateKeyPem: pulumi.Input<string>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    subject?: pulumi.Input<inputs.CertRequestSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
}
