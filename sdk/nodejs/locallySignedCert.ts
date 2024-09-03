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
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    public readonly allowedUses!: pulumi.Output<string[]>;
    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public readonly caCertPem!: pulumi.Output<string>;
    /**
     * Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`.
     */
    public /*out*/ readonly caKeyAlgorithm!: pulumi.Output<string>;
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public readonly caPrivateKeyPem!: pulumi.Output<string>;
    /**
     * Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    /**
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    public readonly certRequestPem!: pulumi.Output<string>;
    public readonly earlyRenewalHours!: pulumi.Output<number>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    public readonly isCaCertificate!: pulumi.Output<boolean>;
    /**
     * Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
     */
    public /*out*/ readonly readyForRenewal!: pulumi.Output<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setSubjectKeyId!: pulumi.Output<boolean>;
    /**
     * The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    public /*out*/ readonly validityEndTime!: pulumi.Output<string>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    public readonly validityPeriodHours!: pulumi.Output<number>;
    /**
     * The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocallySignedCertState | undefined;
            resourceInputs["allowedUses"] = state ? state.allowedUses : undefined;
            resourceInputs["caCertPem"] = state ? state.caCertPem : undefined;
            resourceInputs["caKeyAlgorithm"] = state ? state.caKeyAlgorithm : undefined;
            resourceInputs["caPrivateKeyPem"] = state ? state.caPrivateKeyPem : undefined;
            resourceInputs["certPem"] = state ? state.certPem : undefined;
            resourceInputs["certRequestPem"] = state ? state.certRequestPem : undefined;
            resourceInputs["earlyRenewalHours"] = state ? state.earlyRenewalHours : undefined;
            resourceInputs["isCaCertificate"] = state ? state.isCaCertificate : undefined;
            resourceInputs["readyForRenewal"] = state ? state.readyForRenewal : undefined;
            resourceInputs["setSubjectKeyId"] = state ? state.setSubjectKeyId : undefined;
            resourceInputs["validityEndTime"] = state ? state.validityEndTime : undefined;
            resourceInputs["validityPeriodHours"] = state ? state.validityPeriodHours : undefined;
            resourceInputs["validityStartTime"] = state ? state.validityStartTime : undefined;
        } else {
            const args = argsOrState as LocallySignedCertArgs | undefined;
            if ((!args || args.allowedUses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedUses'");
            }
            if ((!args || args.caCertPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caCertPem'");
            }
            if ((!args || args.caPrivateKeyPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'caPrivateKeyPem'");
            }
            if ((!args || args.certRequestPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certRequestPem'");
            }
            if ((!args || args.validityPeriodHours === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            resourceInputs["allowedUses"] = args ? args.allowedUses : undefined;
            resourceInputs["caCertPem"] = args ? args.caCertPem : undefined;
            resourceInputs["caPrivateKeyPem"] = args?.caPrivateKeyPem ? pulumi.secret(args.caPrivateKeyPem) : undefined;
            resourceInputs["certRequestPem"] = args ? args.certRequestPem : undefined;
            resourceInputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            resourceInputs["isCaCertificate"] = args ? args.isCaCertificate : undefined;
            resourceInputs["setSubjectKeyId"] = args ? args.setSubjectKeyId : undefined;
            resourceInputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            resourceInputs["caKeyAlgorithm"] = undefined /*out*/;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["readyForRenewal"] = undefined /*out*/;
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["caPrivateKeyPem"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LocallySignedCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocallySignedCert resources.
 */
export interface LocallySignedCertState {
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    allowedUses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caCertPem?: pulumi.Input<string>;
    /**
     * Name of the algorithm used when generating the private key provided in `caPrivateKeyPem`.
     */
    caKeyAlgorithm?: pulumi.Input<string>;
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caPrivateKeyPem?: pulumi.Input<string>;
    /**
     * Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
     */
    certPem?: pulumi.Input<string>;
    /**
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    certRequestPem?: pulumi.Input<string>;
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    isCaCertificate?: pulumi.Input<boolean>;
    /**
     * Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
     */
    readyForRenewal?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    validityEndTime?: pulumi.Input<string>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    validityPeriodHours?: pulumi.Input<number>;
    /**
     * The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
     */
    validityStartTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocallySignedCert resource.
 */
export interface LocallySignedCertArgs {
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    allowedUses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caCertPem: pulumi.Input<string>;
    /**
     * Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    caPrivateKeyPem: pulumi.Input<string>;
    /**
     * Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
     */
    certRequestPem: pulumi.Input<string>;
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    isCaCertificate?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    validityPeriodHours: pulumi.Input<number>;
}
