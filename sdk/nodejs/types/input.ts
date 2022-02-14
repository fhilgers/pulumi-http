// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface HeaderArgs {
}

export interface RequestArgs {
    body?: pulumi.Input<string>;
    certificates?: pulumi.Input<pulumi.Input<string>[]>;
    expectedStatusCode?: pulumi.Input<number>;
    header?: pulumi.Input<inputs.HeaderArgs>;
    insecureSkipVerify?: pulumi.Input<boolean>;
    maxRetries?: pulumi.Input<number>;
    method: pulumi.Input<string>;
    retryWaitMax?: pulumi.Input<number>;
    retryWaitMin?: pulumi.Input<number>;
    rootCAs?: pulumi.Input<pulumi.Input<string>[]>;
    serverName?: pulumi.Input<string>;
    url: pulumi.Input<string>;
}
