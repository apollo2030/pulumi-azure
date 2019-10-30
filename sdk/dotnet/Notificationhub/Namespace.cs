// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Notificationhub
{
    /// <summary>
    /// Manages a Notification Hub Namespace.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/notification_hub_namespace.html.markdown.
    /// </summary>
    public partial class Namespace : Pulumi.CustomResource
    {
        /// <summary>
        /// Is this Notification Hub Namespace enabled? Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The Azure Region in which this Notification Hub Namespace should be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("namespaceType")]
        public Output<string> NamespaceType { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The ServiceBus Endpoint for this Notification Hub Namespace.
        /// </summary>
        [Output("servicebusEndpoint")]
        public Output<string> ServicebusEndpoint { get; private set; } = null!;

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.NamespaceSku> Sku { get; private set; } = null!;

        /// <summary>
        /// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;


        /// <summary>
        /// Create a Namespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Namespace(string name, NamespaceArgs args, CustomResourceOptions? options = null)
            : base("azure:notificationhub/namespace:Namespace", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Namespace(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
            : base("azure:notificationhub/namespace:Namespace", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Namespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Namespace Get(string name, Input<string> id, NamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Namespace(name, id, state, options);
        }
    }

    public sealed class NamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is this Notification Hub Namespace enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Azure Region in which this Notification Hub Namespace should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceType", required: true)]
        public Input<string> NamespaceType { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.NamespaceSkuArgs>? Sku { get; set; }

        /// <summary>
        /// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        public NamespaceArgs()
        {
        }
    }

    public sealed class NamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is this Notification Hub Namespace enabled? Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The Azure Region in which this Notification Hub Namespace should be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("namespaceType")]
        public Input<string>? NamespaceType { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The ServiceBus Endpoint for this Notification Hub Namespace.
        /// </summary>
        [Input("servicebusEndpoint")]
        public Input<string>? ServicebusEndpoint { get; set; }

        /// <summary>
        /// ) A `sku` block as described below.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.NamespaceSkuGetArgs>? Sku { get; set; }

        /// <summary>
        /// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        public NamespaceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class NamespaceSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public NamespaceSkuArgs()
        {
        }
    }

    public sealed class NamespaceSkuGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public NamespaceSkuGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class NamespaceSku
    {
        /// <summary>
        /// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private NamespaceSku(string name)
        {
            Name = name;
        }
    }
    }
}