// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a Service Fabric Cluster.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 * });
 * const testCluster = new azure.servicefabric.Cluster("test", {
 *     clusterCodeVersion: "6.3.176.9494",
 *     location: testResourceGroup.location,
 *     managementEndpoint: "https://example:80",
 *     nodeTypes: [{
 *         clientEndpointPort: 2020,
 *         httpEndpointPort: 80,
 *         instanceCount: 3,
 *         isPrimary: true,
 *         name: "first",
 *     }],
 *     reliabilityLevel: "Bronze",
 *     resourceGroupName: testResourceGroup.name,
 *     upgradeMode: "Manual",
 *     vmImage: "Windows",
 * });
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /**
     * A List of one or more features which should be enabled, such as `DnsService`.
     */
    public readonly addOnFeatures: pulumi.Output<string[] | undefined>;
    /**
     * An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly azureActiveDirectory: pulumi.Output<{ clientApplicationId: string, clusterApplicationId: string, tenantId: string } | undefined>;
    /**
     * A `certificate` block as defined below.
     */
    public readonly certificate: pulumi.Output<{ thumbprint: string, thumbprintSecondary?: string, x509StoreName: string } | undefined>;
    /**
     * One or two `client_certificate_thumbprint` blocks as defined below.
     */
    public readonly clientCertificateThumbprints: pulumi.Output<{ isAdmin: boolean, thumbprint: string }[] | undefined>;
    /**
     * Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
     */
    public readonly clusterCodeVersion: pulumi.Output<string>;
    /**
     * The Cluster Endpoint for this Service Fabric Cluster.
     */
    public /*out*/ readonly clusterEndpoint: pulumi.Output<string>;
    /**
     * A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly diagnosticsConfig: pulumi.Output<{ blobEndpoint: string, protectedAccountKeyName: string, queueEndpoint: string, storageAccountName: string, tableEndpoint: string } | undefined>;
    /**
     * One or more `fabric_settings` blocks as defined below.
     */
    public readonly fabricSettings: pulumi.Output<{ name: string, parameters?: {[key: string]: any} }[] | undefined>;
    /**
     * Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
     */
    public readonly location: pulumi.Output<string>;
    /**
     * Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
     */
    public readonly managementEndpoint: pulumi.Output<string>;
    /**
     * The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * One or more `node_type` blocks as defined below.
     */
    public readonly nodeTypes: pulumi.Output<{ applicationPorts: { endPort: number, startPort: number }, clientEndpointPort: number, durabilityLevel?: string, ephemeralPorts: { endPort: number, startPort: number }, httpEndpointPort: number, instanceCount: number, isPrimary: boolean, name: string, reverseProxyEndpointPort?: number }[]>;
    /**
     * Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
     */
    public readonly reliabilityLevel: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName: pulumi.Output<string>;
    /**
     * A `reverse_proxy_certificate` block as defined below.
     */
    public readonly reverseProxyCertificate: pulumi.Output<{ thumbprint: string, thumbprintSecondary?: string, x509StoreName: string } | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any}>;
    /**
     * Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
     */
    public readonly upgradeMode: pulumi.Output<string>;
    /**
     * Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
     */
    public readonly vmImage: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ClusterState = argsOrState as ClusterState | undefined;
            inputs["addOnFeatures"] = state ? state.addOnFeatures : undefined;
            inputs["azureActiveDirectory"] = state ? state.azureActiveDirectory : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["clientCertificateThumbprints"] = state ? state.clientCertificateThumbprints : undefined;
            inputs["clusterCodeVersion"] = state ? state.clusterCodeVersion : undefined;
            inputs["clusterEndpoint"] = state ? state.clusterEndpoint : undefined;
            inputs["diagnosticsConfig"] = state ? state.diagnosticsConfig : undefined;
            inputs["fabricSettings"] = state ? state.fabricSettings : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managementEndpoint"] = state ? state.managementEndpoint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeTypes"] = state ? state.nodeTypes : undefined;
            inputs["reliabilityLevel"] = state ? state.reliabilityLevel : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["reverseProxyCertificate"] = state ? state.reverseProxyCertificate : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["upgradeMode"] = state ? state.upgradeMode : undefined;
            inputs["vmImage"] = state ? state.vmImage : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.managementEndpoint === undefined) {
                throw new Error("Missing required property 'managementEndpoint'");
            }
            if (!args || args.nodeTypes === undefined) {
                throw new Error("Missing required property 'nodeTypes'");
            }
            if (!args || args.reliabilityLevel === undefined) {
                throw new Error("Missing required property 'reliabilityLevel'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.upgradeMode === undefined) {
                throw new Error("Missing required property 'upgradeMode'");
            }
            if (!args || args.vmImage === undefined) {
                throw new Error("Missing required property 'vmImage'");
            }
            inputs["addOnFeatures"] = args ? args.addOnFeatures : undefined;
            inputs["azureActiveDirectory"] = args ? args.azureActiveDirectory : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["clientCertificateThumbprints"] = args ? args.clientCertificateThumbprints : undefined;
            inputs["clusterCodeVersion"] = args ? args.clusterCodeVersion : undefined;
            inputs["diagnosticsConfig"] = args ? args.diagnosticsConfig : undefined;
            inputs["fabricSettings"] = args ? args.fabricSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managementEndpoint"] = args ? args.managementEndpoint : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeTypes"] = args ? args.nodeTypes : undefined;
            inputs["reliabilityLevel"] = args ? args.reliabilityLevel : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["reverseProxyCertificate"] = args ? args.reverseProxyCertificate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradeMode"] = args ? args.upgradeMode : undefined;
            inputs["vmImage"] = args ? args.vmImage : undefined;
            inputs["clusterEndpoint"] = undefined /*out*/;
        }
        super("azure:servicefabric/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * A List of one or more features which should be enabled, such as `DnsService`.
     */
    readonly addOnFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
     */
    readonly azureActiveDirectory?: pulumi.Input<{ clientApplicationId: pulumi.Input<string>, clusterApplicationId: pulumi.Input<string>, tenantId: pulumi.Input<string> }>;
    /**
     * A `certificate` block as defined below.
     */
    readonly certificate?: pulumi.Input<{ thumbprint: pulumi.Input<string>, thumbprintSecondary?: pulumi.Input<string>, x509StoreName: pulumi.Input<string> }>;
    /**
     * One or two `client_certificate_thumbprint` blocks as defined below.
     */
    readonly clientCertificateThumbprints?: pulumi.Input<pulumi.Input<{ isAdmin: pulumi.Input<boolean>, thumbprint: pulumi.Input<string> }>[]>;
    /**
     * Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
     */
    readonly clusterCodeVersion?: pulumi.Input<string>;
    /**
     * The Cluster Endpoint for this Service Fabric Cluster.
     */
    readonly clusterEndpoint?: pulumi.Input<string>;
    /**
     * A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
     */
    readonly diagnosticsConfig?: pulumi.Input<{ blobEndpoint: pulumi.Input<string>, protectedAccountKeyName: pulumi.Input<string>, queueEndpoint: pulumi.Input<string>, storageAccountName: pulumi.Input<string>, tableEndpoint: pulumi.Input<string> }>;
    /**
     * One or more `fabric_settings` blocks as defined below.
     */
    readonly fabricSettings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, parameters?: pulumi.Input<{[key: string]: any}> }>[]>;
    /**
     * Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
     */
    readonly managementEndpoint?: pulumi.Input<string>;
    /**
     * The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `node_type` blocks as defined below.
     */
    readonly nodeTypes?: pulumi.Input<pulumi.Input<{ applicationPorts?: pulumi.Input<{ endPort: pulumi.Input<number>, startPort: pulumi.Input<number> }>, clientEndpointPort: pulumi.Input<number>, durabilityLevel?: pulumi.Input<string>, ephemeralPorts?: pulumi.Input<{ endPort: pulumi.Input<number>, startPort: pulumi.Input<number> }>, httpEndpointPort: pulumi.Input<number>, instanceCount: pulumi.Input<number>, isPrimary: pulumi.Input<boolean>, name: pulumi.Input<string>, reverseProxyEndpointPort?: pulumi.Input<number> }>[]>;
    /**
     * Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
     */
    readonly reliabilityLevel?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `reverse_proxy_certificate` block as defined below.
     */
    readonly reverseProxyCertificate?: pulumi.Input<{ thumbprint: pulumi.Input<string>, thumbprintSecondary?: pulumi.Input<string>, x509StoreName: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
     */
    readonly upgradeMode?: pulumi.Input<string>;
    /**
     * Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
     */
    readonly vmImage?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * A List of one or more features which should be enabled, such as `DnsService`.
     */
    readonly addOnFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An `azure_active_directory` block as defined below. Changing this forces a new resource to be created.
     */
    readonly azureActiveDirectory?: pulumi.Input<{ clientApplicationId: pulumi.Input<string>, clusterApplicationId: pulumi.Input<string>, tenantId: pulumi.Input<string> }>;
    /**
     * A `certificate` block as defined below.
     */
    readonly certificate?: pulumi.Input<{ thumbprint: pulumi.Input<string>, thumbprintSecondary?: pulumi.Input<string>, x509StoreName: pulumi.Input<string> }>;
    /**
     * One or two `client_certificate_thumbprint` blocks as defined below.
     */
    readonly clientCertificateThumbprints?: pulumi.Input<pulumi.Input<{ isAdmin: pulumi.Input<boolean>, thumbprint: pulumi.Input<string> }>[]>;
    /**
     * Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
     */
    readonly clusterCodeVersion?: pulumi.Input<string>;
    /**
     * A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
     */
    readonly diagnosticsConfig?: pulumi.Input<{ blobEndpoint: pulumi.Input<string>, protectedAccountKeyName: pulumi.Input<string>, queueEndpoint: pulumi.Input<string>, storageAccountName: pulumi.Input<string>, tableEndpoint: pulumi.Input<string> }>;
    /**
     * One or more `fabric_settings` blocks as defined below.
     */
    readonly fabricSettings?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, parameters?: pulumi.Input<{[key: string]: any}> }>[]>;
    /**
     * Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
     */
    readonly managementEndpoint: pulumi.Input<string>;
    /**
     * The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more `node_type` blocks as defined below.
     */
    readonly nodeTypes: pulumi.Input<pulumi.Input<{ applicationPorts?: pulumi.Input<{ endPort: pulumi.Input<number>, startPort: pulumi.Input<number> }>, clientEndpointPort: pulumi.Input<number>, durabilityLevel?: pulumi.Input<string>, ephemeralPorts?: pulumi.Input<{ endPort: pulumi.Input<number>, startPort: pulumi.Input<number> }>, httpEndpointPort: pulumi.Input<number>, instanceCount: pulumi.Input<number>, isPrimary: pulumi.Input<boolean>, name: pulumi.Input<string>, reverseProxyEndpointPort?: pulumi.Input<number> }>[]>;
    /**
     * Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
     */
    readonly reliabilityLevel: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `reverse_proxy_certificate` block as defined below.
     */
    readonly reverseProxyCertificate?: pulumi.Input<{ thumbprint: pulumi.Input<string>, thumbprintSecondary?: pulumi.Input<string>, x509StoreName: pulumi.Input<string> }>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
     */
    readonly upgradeMode: pulumi.Input<string>;
    /**
     * Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
     */
    readonly vmImage: pulumi.Input<string>;
}
