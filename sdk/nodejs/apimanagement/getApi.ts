// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to access information about an existing API Management API.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const test = pulumi.output(azure.apimanagement.getApi({
 *     apiManagementName: "search-api-management",
 *     name: "search-api",
 *     resourceGroupName: "search-service",
 *     revision: "2",
 * }));
 * 
 * export const apiManagementApiId = test.apply(test => test.id);
 * ```
 */
export function getApi(args: GetApiArgs, opts?: pulumi.InvokeOptions): Promise<GetApiResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:apimanagement/getApi:getApi", {
        "apiManagementName": args.apiManagementName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "revision": args.revision,
    }, opts);
}

/**
 * A collection of arguments for invoking getApi.
 */
export interface GetApiArgs {
    /**
     * The name of the API Management Service in which the API Management API exists.
     */
    readonly apiManagementName: string;
    /**
     * The name of the API Management API.
     */
    readonly name: string;
    /**
     * The Name of the Resource Group in which the API Management Service exists.
     */
    readonly resourceGroupName: string;
    /**
     * The Revision of the API Management API.
     */
    readonly revision: string;
}

/**
 * A collection of values returned by getApi.
 */
export interface GetApiResult {
    readonly apiManagementName: string;
    /**
     * A description of the API Management API, which may include HTML formatting tags.
     */
    readonly description: string;
    /**
     * The display name of the API.
     */
    readonly displayName: string;
    /**
     * Is this the current API Revision?
     */
    readonly isCurrent: boolean;
    /**
     * Is this API Revision online/accessible via the Gateway?
     */
    readonly isOnline: boolean;
    readonly name: string;
    /**
     * The Path for this API Management API.
     */
    readonly path: string;
    /**
     * A list of protocols the operations in this API can be invoked.
     */
    readonly protocols: string[];
    readonly resourceGroupName: string;
    readonly revision: string;
    /**
     * Absolute URL of the backend service implementing this API.
     */
    readonly serviceUrl: string;
    /**
     * Should this API expose a SOAP frontend, rather than a HTTP frontend?
     */
    readonly soapPassThrough: boolean;
    /**
     * A `subscription_key_parameter_names` block as documented below.
     */
    readonly subscriptionKeyParameterNames: { header: string, query: string }[];
    /**
     * The Version number of this API, if this API is versioned.
     */
    readonly version: string;
    /**
     * The ID of the Version Set which this API is associated with.
     */
    readonly versionSetId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
