// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DevSpace Controller.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "westeurope",
 *     name: "acctestRG1",
 * });
 * const testKubernetesCluster = new azure.containerservice.KubernetesCluster("test", {
 *     agentPoolProfile: {
 *         count: 1,
 *         name: "default",
 *         vmSize: "Standard_DS2_v2",
 *     },
 *     dnsPrefix: "acctestaks1",
 *     location: testResourceGroup.location,
 *     name: "acctestaks1",
 *     resourceGroupName: testResourceGroup.name,
 *     servicePrincipal: {
 *         clientId: "00000000-0000-0000-0000-000000000000",
 *         clientSecret: "00000000000000000000000000000000",
 *     },
 * });
 * const testController = new azure.devspace.Controller("test", {
 *     hostSuffix: "suffix",
 *     location: testResourceGroup.location,
 *     name: "acctestdsc1",
 *     resourceGroupName: testResourceGroup.name,
 *     sku: {
 *         name: "S1",
 *         tier: "Standard",
 *     },
 *     tags: {
 *         Environment: "Testing",
 *     },
 *     targetContainerHostCredentialsBase64: testKubernetesCluster.kubeConfigRaw.apply(kubeConfigRaw => Buffer.from(kubeConfigRaw).toString("base64")),
 *     targetContainerHostResourceId: testKubernetesCluster.id,
 * });
 * ```
 */
export class Controller extends pulumi.CustomResource {
    /**
     * Get an existing Controller resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ControllerState, opts?: pulumi.CustomResourceOptions): Controller {
        return new Controller(name, <any>state, { ...opts, id: id });
    }

    /**
     * DNS name for accessing DataPlane services.
     */
    public /*out*/ readonly dataPlaneFqdn: pulumi.Output<string>;
    /**
     * The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
     */
    public readonly hostSuffix: pulumi.Output<string>;
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * A `sku` block as documented below. Changing this forces a new resource to be created.
     */
    public readonly sku: pulumi.Output<{ name: string, tier: string }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;
    /**
     * Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    public readonly targetContainerHostCredentialsBase64: pulumi.Output<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    public readonly targetContainerHostResourceId: pulumi.Output<string>;

    /**
     * Create a Controller resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ControllerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ControllerArgs | ControllerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ControllerState = argsOrState as ControllerState | undefined;
            inputs["dataPlaneFqdn"] = state ? state.dataPlaneFqdn : undefined;
            inputs["hostSuffix"] = state ? state.hostSuffix : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetContainerHostCredentialsBase64"] = state ? state.targetContainerHostCredentialsBase64 : undefined;
            inputs["targetContainerHostResourceId"] = state ? state.targetContainerHostResourceId : undefined;
        } else {
            const args = argsOrState as ControllerArgs | undefined;
            if (!args || args.hostSuffix === undefined) {
                throw new Error("Missing required property 'hostSuffix'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            if (!args || args.targetContainerHostCredentialsBase64 === undefined) {
                throw new Error("Missing required property 'targetContainerHostCredentialsBase64'");
            }
            if (!args || args.targetContainerHostResourceId === undefined) {
                throw new Error("Missing required property 'targetContainerHostResourceId'");
            }
            inputs["hostSuffix"] = args ? args.hostSuffix : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetContainerHostCredentialsBase64"] = args ? args.targetContainerHostCredentialsBase64 : undefined;
            inputs["targetContainerHostResourceId"] = args ? args.targetContainerHostResourceId : undefined;
            inputs["dataPlaneFqdn"] = undefined /*out*/;
        }
        super("azure:devspace/controller:Controller", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Controller resources.
 */
export interface ControllerState {
    /**
     * DNS name for accessing DataPlane services.
     */
    readonly dataPlaneFqdn?: pulumi.Input<string>;
    /**
     * The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly hostSuffix?: pulumi.Input<string>;
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `sku` block as documented below. Changing this forces a new resource to be created.
     */
    readonly sku?: pulumi.Input<{ name: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostCredentialsBase64?: pulumi.Input<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostResourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Controller resource.
 */
export interface ControllerArgs {
    /**
     * The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly hostSuffix: pulumi.Input<string>;
    /**
     * Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `sku` block as documented below. Changing this forces a new resource to be created.
     */
    readonly sku: pulumi.Input<{ name: pulumi.Input<string>, tier: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostCredentialsBase64: pulumi.Input<string>;
    /**
     * The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
     */
    readonly targetContainerHostResourceId: pulumi.Input<string>;
}
