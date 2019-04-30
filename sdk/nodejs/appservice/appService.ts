// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Service (within an App Service Plan).
 * 
 * > **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azurerm_app_service` resource will be overwritten when promoting a Slot using the `azurerm_app_service_active_slot` resource.
 * 
 * ## Example Usage
 * 
 * This example provisions a Windows App Service. Other examples of the `azurerm_app_service` resource can be found in [the `./examples/app-service` directory within the Github Repository](https://github.com/terraform-providers/terraform-provider-azurerm/tree/master/examples/app-service)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "example-resources",
 * });
 * const testPlan = new azure.appservice.Plan("test", {
 *     location: testResourceGroup.location,
 *     name: "example-appserviceplan",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         size: "S1",
 *         tier: "Standard",
 *     },
 * });
 * const testAppService = new azure.appservice.AppService("test", {
 *     appServicePlanId: testPlan.id,
 *     appSettings: {
 *         SOME_KEY: "some-value",
 *     },
 *     connectionStrings: [{
 *         name: "Database",
 *         type: "SQLServer",
 *         value: "Server=some-server.mydomain.com;Integrated Security=SSPI",
 *     }],
 *     location: testResourceGroup.location,
 *     name: "example-app-service",
 *     resourceGroupName: testResourceGroup.name,
 *     siteConfig: {
 *         dotnetFrameworkVersion: "v4.0",
 *         scmType: "LocalGit",
 *     },
 * });
 * ```
 */
export class AppService extends pulumi.CustomResource {
    /**
     * Get an existing AppService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppServiceState, opts?: pulumi.CustomResourceOptions): AppService {
        return new AppService(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ID of the App Service Plan within which to create this App Service.
     */
    public readonly appServicePlanId: pulumi.Output<string>;
    /**
     * A key-value pair of App Settings.
     */
    public readonly appSettings: pulumi.Output<{[key: string]: any}>;
    /**
     * Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
     */
    public readonly clientAffinityEnabled: pulumi.Output<boolean>;
    /**
     * Does the App Service require client certificates for incoming requests? Defaults to `false`.
     */
    public readonly clientCertEnabled: pulumi.Output<boolean | undefined>;
    /**
     * One or more `connection_string` blocks as defined below.
     */
    public readonly connectionStrings: pulumi.Output<{ name: string, type: string, value: string }[]>;
    /**
     * The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
     */
    public /*out*/ readonly defaultSiteHostname: pulumi.Output<string>;
    /**
     * Is the App Service Enabled?
     */
    public readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * Can the App Service only be accessed via HTTPS? Defaults to `false`.
     */
    public readonly httpsOnly: pulumi.Output<boolean | undefined>;
    /**
     * A Managed Service Identity block as defined below.
     */
    public readonly identity: pulumi.Output<{ principalId: string, tenantId: string, type: string }>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    public /*out*/ readonly outboundIpAddresses: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
     */
    public /*out*/ readonly possibleOutboundIpAddresses: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the App Service.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * A `site_config` block as defined below.
     */
    public readonly siteConfig: pulumi.Output<{ alwaysOn?: boolean, appCommandLine?: string, cors: { allowedOrigins: string[], supportCredentials?: boolean }, defaultDocuments?: string[], dotnetFrameworkVersion?: string, ftpsState: string, http2Enabled?: boolean, ipRestrictions: { ipAddress: string, subnetMask?: string }[], javaContainer?: string, javaContainerVersion?: string, javaVersion?: string, linuxFxVersion: string, localMysqlEnabled: boolean, managedPipelineMode: string, minTlsVersion: string, phpVersion?: string, pythonVersion?: string, remoteDebuggingEnabled?: boolean, remoteDebuggingVersion: string, scmType?: string, use32BitWorkerProcess?: boolean, virtualNetworkName?: string, websocketsEnabled: boolean }>;
    /**
     * A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    public /*out*/ readonly siteCredential: pulumi.Output<{ password: string, username: string }>;
    /**
     * A `source_control` block as defined below, which contains the Source Control information when `scm_type` is set to `LocalGit`.
     */
    public /*out*/ readonly sourceControl: pulumi.Output<{ branch: string, repoUrl: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a AppService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppServiceArgs | AppServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AppServiceState = argsOrState as AppServiceState | undefined;
            inputs["appServicePlanId"] = state ? state.appServicePlanId : undefined;
            inputs["appSettings"] = state ? state.appSettings : undefined;
            inputs["clientAffinityEnabled"] = state ? state.clientAffinityEnabled : undefined;
            inputs["clientCertEnabled"] = state ? state.clientCertEnabled : undefined;
            inputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            inputs["defaultSiteHostname"] = state ? state.defaultSiteHostname : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["httpsOnly"] = state ? state.httpsOnly : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outboundIpAddresses"] = state ? state.outboundIpAddresses : undefined;
            inputs["possibleOutboundIpAddresses"] = state ? state.possibleOutboundIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["siteConfig"] = state ? state.siteConfig : undefined;
            inputs["siteCredential"] = state ? state.siteCredential : undefined;
            inputs["sourceControl"] = state ? state.sourceControl : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as AppServiceArgs | undefined;
            if (!args || args.appServicePlanId === undefined) {
                throw new Error("Missing required property 'appServicePlanId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appServicePlanId"] = args ? args.appServicePlanId : undefined;
            inputs["appSettings"] = args ? args.appSettings : undefined;
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["clientCertEnabled"] = args ? args.clientCertEnabled : undefined;
            inputs["connectionStrings"] = args ? args.connectionStrings : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["defaultSiteHostname"] = undefined /*out*/;
            inputs["outboundIpAddresses"] = undefined /*out*/;
            inputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            inputs["siteCredential"] = undefined /*out*/;
            inputs["sourceControl"] = undefined /*out*/;
        }
        super("azure:appservice/appService:AppService", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppService resources.
 */
export interface AppServiceState {
    /**
     * The ID of the App Service Plan within which to create this App Service.
     */
    readonly appServicePlanId?: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: any}>;
    /**
     * Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * Does the App Service require client certificates for incoming requests? Defaults to `false`.
     */
    readonly clientCertEnabled?: pulumi.Input<boolean>;
    /**
     * One or more `connection_string` blocks as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, type: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
     */
    readonly defaultSiteHostname?: pulumi.Input<string>;
    /**
     * Is the App Service Enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the App Service only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<{ principalId?: pulumi.Input<string>, tenantId?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    readonly outboundIpAddresses?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outbound_ip_addresses`.
     */
    readonly possibleOutboundIpAddresses?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `site_config` block as defined below.
     */
    readonly siteConfig?: pulumi.Input<{ alwaysOn?: pulumi.Input<boolean>, appCommandLine?: pulumi.Input<string>, cors?: pulumi.Input<{ allowedOrigins: pulumi.Input<pulumi.Input<string>[]>, supportCredentials?: pulumi.Input<boolean> }>, defaultDocuments?: pulumi.Input<pulumi.Input<string>[]>, dotnetFrameworkVersion?: pulumi.Input<string>, ftpsState?: pulumi.Input<string>, http2Enabled?: pulumi.Input<boolean>, ipRestrictions?: pulumi.Input<pulumi.Input<{ ipAddress: pulumi.Input<string>, subnetMask?: pulumi.Input<string> }>[]>, javaContainer?: pulumi.Input<string>, javaContainerVersion?: pulumi.Input<string>, javaVersion?: pulumi.Input<string>, linuxFxVersion?: pulumi.Input<string>, localMysqlEnabled?: pulumi.Input<boolean>, managedPipelineMode?: pulumi.Input<string>, minTlsVersion?: pulumi.Input<string>, phpVersion?: pulumi.Input<string>, pythonVersion?: pulumi.Input<string>, remoteDebuggingEnabled?: pulumi.Input<boolean>, remoteDebuggingVersion?: pulumi.Input<string>, scmType?: pulumi.Input<string>, use32BitWorkerProcess?: pulumi.Input<boolean>, virtualNetworkName?: pulumi.Input<string>, websocketsEnabled?: pulumi.Input<boolean> }>;
    /**
     * A `site_credential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    readonly siteCredential?: pulumi.Input<{ password?: pulumi.Input<string>, username?: pulumi.Input<string> }>;
    /**
     * A `source_control` block as defined below, which contains the Source Control information when `scm_type` is set to `LocalGit`.
     */
    readonly sourceControl?: pulumi.Input<{ branch?: pulumi.Input<string>, repoUrl?: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a AppService resource.
 */
export interface AppServiceArgs {
    /**
     * The ID of the App Service Plan within which to create this App Service.
     */
    readonly appServicePlanId: pulumi.Input<string>;
    /**
     * A key-value pair of App Settings.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: any}>;
    /**
     * Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * Does the App Service require client certificates for incoming requests? Defaults to `false`.
     */
    readonly clientCertEnabled?: pulumi.Input<boolean>;
    /**
     * One or more `connection_string` blocks as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, type: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * Is the App Service Enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the App Service only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * A Managed Service Identity block as defined below.
     */
    readonly identity?: pulumi.Input<{ principalId?: pulumi.Input<string>, tenantId?: pulumi.Input<string>, type: pulumi.Input<string> }>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the App Service. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the App Service.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `site_config` block as defined below.
     */
    readonly siteConfig?: pulumi.Input<{ alwaysOn?: pulumi.Input<boolean>, appCommandLine?: pulumi.Input<string>, cors?: pulumi.Input<{ allowedOrigins: pulumi.Input<pulumi.Input<string>[]>, supportCredentials?: pulumi.Input<boolean> }>, defaultDocuments?: pulumi.Input<pulumi.Input<string>[]>, dotnetFrameworkVersion?: pulumi.Input<string>, ftpsState?: pulumi.Input<string>, http2Enabled?: pulumi.Input<boolean>, ipRestrictions?: pulumi.Input<pulumi.Input<{ ipAddress: pulumi.Input<string>, subnetMask?: pulumi.Input<string> }>[]>, javaContainer?: pulumi.Input<string>, javaContainerVersion?: pulumi.Input<string>, javaVersion?: pulumi.Input<string>, linuxFxVersion?: pulumi.Input<string>, localMysqlEnabled?: pulumi.Input<boolean>, managedPipelineMode?: pulumi.Input<string>, minTlsVersion?: pulumi.Input<string>, phpVersion?: pulumi.Input<string>, pythonVersion?: pulumi.Input<string>, remoteDebuggingEnabled?: pulumi.Input<boolean>, remoteDebuggingVersion?: pulumi.Input<string>, scmType?: pulumi.Input<string>, use32BitWorkerProcess?: pulumi.Input<boolean>, virtualNetworkName?: pulumi.Input<string>, websocketsEnabled?: pulumi.Input<boolean> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
