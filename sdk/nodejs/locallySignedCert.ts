// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LocallySignedCert extends pulumi.CustomResource {
    /**
     * Get an existing LocallySignedCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocallySignedCertState, opts?: pulumi.CustomResourceOptions): LocallySignedCert {
        return new LocallySignedCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tls:index/locallySignedCert:LocallySignedCert';

    /**
     * Returns true if the given object is an instance of LocallySignedCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocallySignedCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocallySignedCert.__pulumiType;
    }

    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    public readonly allowedUses!: pulumi.Output<string[]>;
    /**
     * PEM-encoded certificate data for the CA.
     */
    public readonly caCertPem!: pulumi.Output<string>;
    /**
     * The name of the algorithm for the key provided
     * in `caPrivateKeyPem`.
     */
    public readonly caKeyAlgorithm!: pulumi.Output<string>;
    /**
     * PEM-encoded private key data for the CA.
     * This can be read from a separate file using the ``file`` interpolation
     * function.
     */
    public readonly caPrivateKeyPem!: pulumi.Output<string>;
    /**
     * The certificate data in PEM format.
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    /**
     * PEM-encoded request certificate data.
     */
    public readonly certRequestPem!: pulumi.Output<string>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    public readonly earlyRenewalHours!: pulumi.Output<number | undefined>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    public readonly isCaCertificate!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly readyForRenewal!: pulumi.Output<boolean>;
    /**
     * If `true`, the certificate will include
     * the subject key identifier. Defaults to `false`, in which case the subject
     * key identifier is not set at all.
     */
    public readonly setSubjectKeyId!: pulumi.Output<boolean | undefined>;
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
     * Create a LocallySignedCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocallySignedCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocallySignedCertArgs | LocallySignedCertState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LocallySignedCertState | undefined;
            inputs["allowedUses"] = state ? state.allowedUses : undefined;
            inputs["caCertPem"] = state ? state.caCertPem : undefined;
            inputs["caKeyAlgorithm"] = state ? state.caKeyAlgorithm : undefined;
            inputs["caPrivateKeyPem"] = state ? state.caPrivateKeyPem : undefined;
            inputs["certPem"] = state ? state.certPem : undefined;
            inputs["certRequestPem"] = state ? state.certRequestPem : undefined;
            inputs["earlyRenewalHours"] = state ? state.earlyRenewalHours : undefined;
            inputs["isCaCertificate"] = state ? state.isCaCertificate : undefined;
            inputs["readyForRenewal"] = state ? state.readyForRenewal : undefined;
            inputs["setSubjectKeyId"] = state ? state.setSubjectKeyId : undefined;
            inputs["validityEndTime"] = state ? state.validityEndTime : undefined;
            inputs["validityPeriodHours"] = state ? state.validityPeriodHours : undefined;
            inputs["validityStartTime"] = state ? state.validityStartTime : undefined;
        } else {
            const args = argsOrState as LocallySignedCertArgs | undefined;
            if (!args || args.allowedUses === undefined) {
                throw new Error("Missing required property 'allowedUses'");
            }
            if (!args || args.caCertPem === undefined) {
                throw new Error("Missing required property 'caCertPem'");
            }
            if (!args || args.caKeyAlgorithm === undefined) {
                throw new Error("Missing required property 'caKeyAlgorithm'");
            }
            if (!args || args.caPrivateKeyPem === undefined) {
                throw new Error("Missing required property 'caPrivateKeyPem'");
            }
            if (!args || args.certRequestPem === undefined) {
                throw new Error("Missing required property 'certRequestPem'");
            }
            if (!args || args.validityPeriodHours === undefined) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            inputs["allowedUses"] = args ? args.allowedUses : undefined;
            inputs["caCertPem"] = args ? args.caCertPem : undefined;
            inputs["caKeyAlgorithm"] = args ? args.caKeyAlgorithm : undefined;
            inputs["caPrivateKeyPem"] = args ? args.caPrivateKeyPem : undefined;
            inputs["certRequestPem"] = args ? args.certRequestPem : undefined;
            inputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            inputs["isCaCertificate"] = args ? args.isCaCertificate : undefined;
            inputs["setSubjectKeyId"] = args ? args.setSubjectKeyId : undefined;
            inputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            inputs["certPem"] = undefined /*out*/;
            inputs["readyForRenewal"] = undefined /*out*/;
            inputs["validityEndTime"] = undefined /*out*/;
            inputs["validityStartTime"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LocallySignedCert.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocallySignedCert resources.
 */
export interface LocallySignedCertState {
    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    readonly allowedUses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * PEM-encoded certificate data for the CA.
     */
    readonly caCertPem?: pulumi.Input<string>;
    /**
     * The name of the algorithm for the key provided
     * in `caPrivateKeyPem`.
     */
    readonly caKeyAlgorithm?: pulumi.Input<string>;
    /**
     * PEM-encoded private key data for the CA.
     * This can be read from a separate file using the ``file`` interpolation
     * function.
     */
    readonly caPrivateKeyPem?: pulumi.Input<string>;
    /**
     * The certificate data in PEM format.
     */
    readonly certPem?: pulumi.Input<string>;
    /**
     * PEM-encoded request certificate data.
     */
    readonly certRequestPem?: pulumi.Input<string>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    readonly earlyRenewalHours?: pulumi.Input<number>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    readonly isCaCertificate?: pulumi.Input<boolean>;
    readonly readyForRenewal?: pulumi.Input<boolean>;
    /**
     * If `true`, the certificate will include
     * the subject key identifier. Defaults to `false`, in which case the subject
     * key identifier is not set at all.
     */
    readonly setSubjectKeyId?: pulumi.Input<boolean>;
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
 * The set of arguments for constructing a LocallySignedCert resource.
 */
export interface LocallySignedCertArgs {
    /**
     * List of keywords each describing a use that is permitted
     * for the issued certificate. The valid keywords are listed below.
     */
    readonly allowedUses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * PEM-encoded certificate data for the CA.
     */
    readonly caCertPem: pulumi.Input<string>;
    /**
     * The name of the algorithm for the key provided
     * in `caPrivateKeyPem`.
     */
    readonly caKeyAlgorithm: pulumi.Input<string>;
    /**
     * PEM-encoded private key data for the CA.
     * This can be read from a separate file using the ``file`` interpolation
     * function.
     */
    readonly caPrivateKeyPem: pulumi.Input<string>;
    /**
     * PEM-encoded request certificate data.
     */
    readonly certRequestPem: pulumi.Input<string>;
    /**
     * Number of hours before the certificates expiry when a new certificate will be generated
     */
    readonly earlyRenewalHours?: pulumi.Input<number>;
    /**
     * Boolean controlling whether the CA flag will be set in the
     * generated certificate. Defaults to `false`, meaning that the certificate does not represent
     * a certificate authority.
     */
    readonly isCaCertificate?: pulumi.Input<boolean>;
    /**
     * If `true`, the certificate will include
     * the subject key identifier. Defaults to `false`, in which case the subject
     * key identifier is not set at all.
     */
    readonly setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * The number of hours after initial issuing that the
     * certificate will become invalid.
     */
    readonly validityPeriodHours: pulumi.Input<number>;
}
