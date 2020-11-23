// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface CertRequestSubject {
    commonName?: pulumi.Input<string>;
    country?: pulumi.Input<string>;
    locality?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    organizationalUnit?: pulumi.Input<string>;
    postalCode?: pulumi.Input<string>;
    province?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface SelfSignedCertSubject {
    commonName?: pulumi.Input<string>;
    country?: pulumi.Input<string>;
    locality?: pulumi.Input<string>;
    organization?: pulumi.Input<string>;
    organizationalUnit?: pulumi.Input<string>;
    postalCode?: pulumi.Input<string>;
    province?: pulumi.Input<string>;
    serialNumber?: pulumi.Input<string>;
    streetAddresses?: pulumi.Input<pulumi.Input<string>[]>;
}
