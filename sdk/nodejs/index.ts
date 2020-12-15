// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./certRequest";
export * from "./getCertificate";
export * from "./getPublicKey";
export * from "./locallySignedCert";
export * from "./privateKey";
export * from "./provider";
export * from "./selfSignedCert";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

// Import resources to register:
import { CertRequest } from "./certRequest";
import { LocallySignedCert } from "./locallySignedCert";
import { PrivateKey } from "./privateKey";
import { SelfSignedCert } from "./selfSignedCert";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tls:index/certRequest:CertRequest":
                return new CertRequest(name, <any>undefined, { urn })
            case "tls:index/locallySignedCert:LocallySignedCert":
                return new LocallySignedCert(name, <any>undefined, { urn })
            case "tls:index/privateKey:PrivateKey":
                return new PrivateKey(name, <any>undefined, { urn })
            case "tls:index/selfSignedCert:SelfSignedCert":
                return new SelfSignedCert(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tls", "index/certRequest", _module)
pulumi.runtime.registerResourceModule("tls", "index/locallySignedCert", _module)
pulumi.runtime.registerResourceModule("tls", "index/privateKey", _module)
pulumi.runtime.registerResourceModule("tls", "index/selfSignedCert", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("tls", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:tls") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
