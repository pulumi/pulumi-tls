// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-tls/blob/master/website/docs/r/self_signed_cert.html.markdown.
 */
export class SelfSignedCert extends pulumi.CustomResource {
    /**
     * Get an existing SelfSignedCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SelfSignedCertState, opts?: pulumi.CustomResourceOptions): SelfSignedCert {
        return new SelfSignedCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tls:index/selfSignedCert:SelfSignedCert';

    /**
     * Returns true if the given object is an instance of SelfSignedCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SelfSignedCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SelfSignedCert.__pulumiType;
    }

    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    public readonly allowedUses!: pulumi.Output<string[]>;
    /**
     * The certificate data in PEM format.
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    public readonly dnsNames!: pulumi.Output<string[] | undefined>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    public readonly earlyRenewalHours!: pulumi.Output<number | undefined>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    public readonly isCaCertificate!: pulumi.Output<boolean | undefined>;
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
     * The subject for which a certificate is being requested.
     * This is a nested configuration block whose structure matches the
     * corresponding block for `tls..CertRequest`.
     */
    public readonly subjects!: pulumi.Output<{ commonName?: string, country?: string, locality?: string, organization?: string, organizationalUnit?: string, postalCode?: string, province?: string, serialNumber?: string, streetAddresses?: string[] }[]>;
    /**
     * The time until which the certificate is invalid, as an
     * [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    public /*out*/ readonly validityEndTime!: pulumi.Output<string>;
    /**
     * The number of hours after initial issuing that the
     * certificate will become invalid.
     */
    public readonly validityPeriodHours!: pulumi.Output<number>;
    /**
     * The time after which the certificate is valid, as an
     * [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    public /*out*/ readonly validityStartTime!: pulumi.Output<string>;

    /**
     * Create a SelfSignedCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SelfSignedCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SelfSignedCertArgs | SelfSignedCertState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SelfSignedCertState | undefined;
            inputs["allowedUses"] = state ? state.allowedUses : undefined;
            inputs["certPem"] = state ? state.certPem : undefined;
            inputs["dnsNames"] = state ? state.dnsNames : undefined;
            inputs["earlyRenewalHours"] = state ? state.earlyRenewalHours : undefined;
            inputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            inputs["isCaCertificate"] = state ? state.isCaCertificate : undefined;
            inputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            inputs["privateKeyPem"] = state ? state.privateKeyPem : undefined;
            inputs["subjects"] = state ? state.subjects : undefined;
            inputs["validityEndTime"] = state ? state.validityEndTime : undefined;
            inputs["validityPeriodHours"] = state ? state.validityPeriodHours : undefined;
            inputs["validityStartTime"] = state ? state.validityStartTime : undefined;
        } else {
            const args = argsOrState as SelfSignedCertArgs | undefined;
            if (!args || args.allowedUses === undefined) {
                throw new Error("Missing required property 'allowedUses'");
            }
            if (!args || args.keyAlgorithm === undefined) {
                throw new Error("Missing required property 'keyAlgorithm'");
            }
            if (!args || args.privateKeyPem === undefined) {
                throw new Error("Missing required property 'privateKeyPem'");
            }
            if (!args || args.subjects === undefined) {
                throw new Error("Missing required property 'subjects'");
            }
            if (!args || args.validityPeriodHours === undefined) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            inputs["allowedUses"] = args ? args.allowedUses : undefined;
            inputs["dnsNames"] = args ? args.dnsNames : undefined;
            inputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            inputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            inputs["isCaCertificate"] = args ? args.isCaCertificate : undefined;
            inputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            inputs["privateKeyPem"] = args ? args.privateKeyPem : undefined;
            inputs["subjects"] = args ? args.subjects : undefined;
            inputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            inputs["certPem"] = undefined /*out*/;
            inputs["validityEndTime"] = undefined /*out*/;
            inputs["validityStartTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SelfSignedCert.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SelfSignedCert resources.
 */
export interface SelfSignedCertState {
    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    readonly allowedUses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The certificate data in PEM format.
     */
    readonly certPem?: pulumi.Input<string>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    readonly earlyRenewalHours?: pulumi.Input<number>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    readonly isCaCertificate?: pulumi.Input<boolean>;
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
     * The subject for which a certificate is being requested.
     * This is a nested configuration block whose structure matches the
     * corresponding block for `tls..CertRequest`.
     */
    readonly subjects?: pulumi.Input<pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, postalCode?: pulumi.Input<string>, province?: pulumi.Input<string>, serialNumber?: pulumi.Input<string>, streetAddresses?: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The time until which the certificate is invalid, as an
     * [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    readonly validityEndTime?: pulumi.Input<string>;
    /**
     * The number of hours after initial issuing that the
     * certificate will become invalid.
     */
    readonly validityPeriodHours?: pulumi.Input<number>;
    /**
     * The time after which the certificate is valid, as an
     * [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    readonly validityStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SelfSignedCert resource.
 */
export interface SelfSignedCertArgs {
    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    readonly allowedUses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DNS names for which a certificate is being requested.
     */
    readonly dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    readonly earlyRenewalHours?: pulumi.Input<number>;
    /**
     * List of IP addresses for which a certificate is being requested.
     */
    readonly ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    readonly isCaCertificate?: pulumi.Input<boolean>;
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
     * The subject for which a certificate is being requested.
     * This is a nested configuration block whose structure matches the
     * corresponding block for `tls..CertRequest`.
     */
    readonly subjects: pulumi.Input<pulumi.Input<{ commonName?: pulumi.Input<string>, country?: pulumi.Input<string>, locality?: pulumi.Input<string>, organization?: pulumi.Input<string>, organizationalUnit?: pulumi.Input<string>, postalCode?: pulumi.Input<string>, province?: pulumi.Input<string>, serialNumber?: pulumi.Input<string>, streetAddresses?: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    /**
     * The number of hours after initial issuing that the
     * certificate will become invalid.
     */
    readonly validityPeriodHours: pulumi.Input<number>;
}
