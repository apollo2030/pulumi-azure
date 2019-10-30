// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Devspace
{
    /// <summary>
    /// Manages a DevSpace Controller.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/devspace_controller.html.markdown.
    /// </summary>
    public partial class Controller : Pulumi.CustomResource
    {
        /// <summary>
        /// DNS name for accessing DataPlane services.
        /// </summary>
        [Output("dataPlaneFqdn")]
        public Output<string> DataPlaneFqdn { get; private set; } = null!;

        /// <summary>
        /// The host suffix for the DevSpace Controller.
        /// </summary>
        [Output("hostSuffix")]
        public Output<string> HostSuffix { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `sku` block as documented below. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ControllerSku> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>> Tags { get; private set; } = null!;

        /// <summary>
        /// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetContainerHostCredentialsBase64")]
        public Output<string> TargetContainerHostCredentialsBase64 { get; private set; } = null!;

        /// <summary>
        /// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("targetContainerHostResourceId")]
        public Output<string> TargetContainerHostResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a Controller resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Controller(string name, ControllerArgs args, CustomResourceOptions? options = null)
            : base("azure:devspace/controller:Controller", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Controller(string name, Input<string> id, ControllerState? state = null, CustomResourceOptions? options = null)
            : base("azure:devspace/controller:Controller", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Controller resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Controller Get(string name, Input<string> id, ControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new Controller(name, id, state, options);
        }
    }

    public sealed class ControllerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `sku` block as documented below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ControllerSkuArgs> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetContainerHostCredentialsBase64", required: true)]
        public Input<string> TargetContainerHostCredentialsBase64 { get; set; } = null!;

        /// <summary>
        /// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetContainerHostResourceId", required: true)]
        public Input<string> TargetContainerHostResourceId { get; set; } = null!;

        public ControllerArgs()
        {
        }
    }

    public sealed class ControllerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS name for accessing DataPlane services.
        /// </summary>
        [Input("dataPlaneFqdn")]
        public Input<string>? DataPlaneFqdn { get; set; }

        /// <summary>
        /// The host suffix for the DevSpace Controller.
        /// </summary>
        [Input("hostSuffix")]
        public Input<string>? HostSuffix { get; set; }

        /// <summary>
        /// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `sku` block as documented below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ControllerSkuGetArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetContainerHostCredentialsBase64")]
        public Input<string>? TargetContainerHostCredentialsBase64 { get; set; }

        /// <summary>
        /// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("targetContainerHostResourceId")]
        public Input<string>? TargetContainerHostResourceId { get; set; }

        public ControllerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ControllerSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public ControllerSkuArgs()
        {
        }
    }

    public sealed class ControllerSkuGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public ControllerSkuGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ControllerSku
    {
        /// <summary>
        /// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;
        public readonly string Tier;

        [OutputConstructor]
        private ControllerSku(
            string name,
            string tier)
        {
            Name = name;
            Tier = tier;
        }
    }
    }
}