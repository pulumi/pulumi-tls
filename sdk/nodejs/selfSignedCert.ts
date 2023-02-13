// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SelfSignedCert extends pulumi.CustomResource {
    /**
     * Get an existing SelfSignedCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    public readonly allowedUses!: pulumi.Output<string[]>;
    /**
     * Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     */
    public /*out*/ readonly certPem!: pulumi.Output<string>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly dnsNames!: pulumi.Output<string[] | undefined>;
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     */
    public readonly earlyRenewalHours!: pulumi.Output<number>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    public readonly isCaCertificate!: pulumi.Output<boolean>;
    /**
     * Name of the algorithm used when generating the private key provided in `privateKeyPem`.
     */
    public /*out*/ readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     */
    public readonly privateKeyPem!: pulumi.Output<string>;
    /**
     * Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
     */
    public /*out*/ readonly readyForRenewal!: pulumi.Output<boolean>;
    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setAuthorityKeyId!: pulumi.Output<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    public readonly setSubjectKeyId!: pulumi.Output<boolean>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    public readonly subject!: pulumi.Output<outputs.SelfSignedCertSubject | undefined>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    public readonly uris!: pulumi.Output<string[] | undefined>;
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
     * Create a SelfSignedCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SelfSignedCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SelfSignedCertArgs | SelfSignedCertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SelfSignedCertState | undefined;
            resourceInputs["allowedUses"] = state ? state.allowedUses : undefined;
            resourceInputs["certPem"] = state ? state.certPem : undefined;
            resourceInputs["dnsNames"] = state ? state.dnsNames : undefined;
            resourceInputs["earlyRenewalHours"] = state ? state.earlyRenewalHours : undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["isCaCertificate"] = state ? state.isCaCertificate : undefined;
            resourceInputs["keyAlgorithm"] = state ? state.keyAlgorithm : undefined;
            resourceInputs["privateKeyPem"] = state ? state.privateKeyPem : undefined;
            resourceInputs["readyForRenewal"] = state ? state.readyForRenewal : undefined;
            resourceInputs["setAuthorityKeyId"] = state ? state.setAuthorityKeyId : undefined;
            resourceInputs["setSubjectKeyId"] = state ? state.setSubjectKeyId : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
            resourceInputs["uris"] = state ? state.uris : undefined;
            resourceInputs["validityEndTime"] = state ? state.validityEndTime : undefined;
            resourceInputs["validityPeriodHours"] = state ? state.validityPeriodHours : undefined;
            resourceInputs["validityStartTime"] = state ? state.validityStartTime : undefined;
        } else {
            const args = argsOrState as SelfSignedCertArgs | undefined;
            if ((!args || args.allowedUses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedUses'");
            }
            if ((!args || args.privateKeyPem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKeyPem'");
            }
            if ((!args || args.validityPeriodHours === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validityPeriodHours'");
            }
            resourceInputs["allowedUses"] = args ? args.allowedUses : undefined;
            resourceInputs["dnsNames"] = args ? args.dnsNames : undefined;
            resourceInputs["earlyRenewalHours"] = args ? args.earlyRenewalHours : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["isCaCertificate"] = args ? args.isCaCertificate : undefined;
            resourceInputs["privateKeyPem"] = args?.privateKeyPem ? pulumi.secret(args.privateKeyPem) : undefined;
            resourceInputs["setAuthorityKeyId"] = args ? args.setAuthorityKeyId : undefined;
            resourceInputs["setSubjectKeyId"] = args ? args.setSubjectKeyId : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["uris"] = args ? args.uris : undefined;
            resourceInputs["validityPeriodHours"] = args ? args.validityPeriodHours : undefined;
            resourceInputs["certPem"] = undefined /*out*/;
            resourceInputs["keyAlgorithm"] = undefined /*out*/;
            resourceInputs["readyForRenewal"] = undefined /*out*/;
            resourceInputs["validityEndTime"] = undefined /*out*/;
            resourceInputs["validityStartTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["privateKeyPem"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SelfSignedCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SelfSignedCert resources.
 */
export interface SelfSignedCertState {
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    allowedUses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the
     * [underlying](https://pkg.go.dev/encoding/pem#Encode)
     * [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\n` at
     * the end of the PEM. In case this disrupts your use case, we recommend using
     * [`trimspace()`](https://www.terraform.io/language/functions/trimspace).
     */
    certPem?: pulumi.Input<string>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     */
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    isCaCertificate?: pulumi.Input<boolean>;
    /**
     * Name of the algorithm used when generating the private key provided in `privateKeyPem`.
     */
    keyAlgorithm?: pulumi.Input<string>;
    /**
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     */
    privateKeyPem?: pulumi.Input<string>;
    /**
     * Is the certificate either expired (i.e. beyond the `validityPeriodHours`) or ready for an early renewal (i.e. within the `earlyRenewalHours`)?
     */
    readyForRenewal?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setAuthorityKeyId?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    subject?: pulumi.Input<inputs.SelfSignedCertSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
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
 * The set of arguments for constructing a SelfSignedCert resource.
 */
export interface SelfSignedCertArgs {
    /**
     * List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `anyExtended`, `certSigning`, `clientAuth`, `codeSigning`, `contentCommitment`, `crlSigning`, `dataEncipherment`, `decipherOnly`, `digitalSignature`, `emailProtection`, `encipherOnly`, `ipsecEndSystem`, `ipsecTunnel`, `ipsecUser`, `keyAgreement`, `keyEncipherment`, `microsoftCommercialCodeSigning`, `microsoftKernelCodeSigning`, `microsoftServerGatedCrypto`, `netscapeServerGatedCrypto`, `ocspSigning`, `serverAuth`, `timestamping`.
     */
    allowedUses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DNS names for which a certificate is being requested (i.e. certificate subjects).
     */
    dnsNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource will consider the certificate to have expired the given number of hours before its actual expiry time. This
     * can be useful to deploy an updated certificate in advance of the expiration of the current certificate. However, the old
     * certificate remains valid until its true expiration time, since this resource does not (and cannot) support certificate
     * revocation. Also, this advance update can only be performed should the Terraform configuration be applied during the
     * early renewal period. (default: `0`)
     */
    earlyRenewalHours?: pulumi.Input<number>;
    /**
     * List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
     */
    isCaCertificate?: pulumi.Input<boolean>;
    /**
     * Private key in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format, that the certificate will belong
     * to. This can be read from a separate file using the [`file`](https://www.terraform.io/language/functions/file)
     * interpolation function. Only an irreversible secure hash of the private key will be stored in the Terraform state.
     */
    privateKeyPem: pulumi.Input<string>;
    /**
     * Should the generated certificate include an [authority key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.1): for self-signed certificates this is the same value as the [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setAuthorityKeyId?: pulumi.Input<boolean>;
    /**
     * Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
     */
    setSubjectKeyId?: pulumi.Input<boolean>;
    /**
     * The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
     */
    subject?: pulumi.Input<inputs.SelfSignedCertSubject>;
    /**
     * List of URIs for which a certificate is being requested (i.e. certificate subjects).
     */
    uris?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of hours, after initial issuing, that the certificate will remain valid for.
     */
    validityPeriodHours: pulumi.Input<number>;
}
