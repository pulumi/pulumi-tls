// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

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
     * The certificate request data in PEM format.
     */
    public /*out*/ readonly certRequestPem!: pulumi.Output<string>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    public readonly dnsNames!: pulumi.Output<string[] | undefined>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the algorithm for the key provided
     * in `privateKeyPem`.
     */
    public readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * PEM-encoded private key that the certificate will belong to
     */
    public readonly privateKeyPem!: pulumi.Output<string>;
    /**
     * The subject for which a certificate is being requested. This is
     * a nested configuration block whose structure is described below.
     */
    public readonly subjects!: pulumi.Output<outputs.CertRequestSubject[]>;
    /**
     * List of URIs for which a certificate is being requested.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CertRequestState | undefined;
            inputs["certRequestPem"] = state ? state.certRequestPem : undefined;
            inputs["dnsNames"] = state ? state.dnsNames : undefined;
            inputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            inputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            inputs["privateKeyPem"] = state ? state.privateKeyPem : undefined;
            inputs["subjects"] = state ? state.subjects : undefined;
            inputs["uris"] = state ? state.uris : undefined;
        } else {
            const args = argsOrState as CertRequestArgs | undefined;
            if ((!args || args.keyAlgorithm === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'keyAlgorithm'");
            }
            if ((!args || args.privateKeyPem === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'privateKeyPem'");
            }
            if ((!args || args.subjects === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'subjects'");
            }
            inputs["dnsNames"] = args ? args.dnsNames : undefined;
            inputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            inputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            inputs["privateKeyPem"] = args ? args.privateKeyPem : undefined;
            inputs["subjects"] = args ? args.subjects : undefined;
            inputs["uris"] = args ? args.uris : undefined;
            inputs["certRequestPem"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CertRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertRequest resources.
 */
export interface CertRequestState {
    /**
     * The certificate request data in PEM format.
     */
    readonly certRequestPem?: pulumi.Input<string>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the algorithm for the key provided
     * in `privateKeyPem`.
     */
    readonly keyAlgorithm?: pulumi.Input<string>;
    /**
     * PEM-encoded private key that the certificate will belong to
     */
    readonly privateKeyPem?: pulumi.Input<string>;
    /**
     * The subject for which a certificate is being requested. This is
     * a nested configuration block whose structure is described below.
     */
    readonly subjects?: pulumi.Input<pulumi.Input<inputs.CertRequestSubject>[]>;
    /**
     * List of URIs for which a certificate is being requested.
     */
    readonly uris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a CertRequest resource.
 */
export interface CertRequestArgs {
    /**
     * List of DNS names for which a certificate is being requested.
     */
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the algorithm for the key provided
     * in `privateKeyPem`.
     */
    readonly keyAlgorithm: pulumi.Input<string>;
    /**
     * PEM-encoded private key that the certificate will belong to
     */
    readonly privateKeyPem: pulumi.Input<string>;
    /**
     * The subject for which a certificate is being requested. This is
     * a nested configuration block whose structure is described below.
     */
    readonly subjects: pulumi.Input<pulumi.Input<inputs.CertRequestSubject>[]>;
    /**
     * List of URIs for which a certificate is being requested.
     */
    readonly uris?: pulumi.Input<pulumi.Input<string>[]>;
}
