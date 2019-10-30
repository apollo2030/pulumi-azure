// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Logicapps
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Logic App Workflow.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/logic_app_workflow.html.markdown.
        /// </summary>
        public static Task<GetWorkflowResult> GetWorkflow(GetWorkflowArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkflowResult>("azure:logicapps/getWorkflow:getWorkflow", args, options.WithVersion());
    }

    public sealed class GetWorkflowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Logic App Workflow.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Logic App Workflow exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetWorkflowArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetWorkflowResult
    {
        /// <summary>
        /// The Access Endpoint for the Logic App Workflow
        /// </summary>
        public readonly string AccessEndpoint;
        /// <summary>
        /// The Azure location where the Logic App Workflow exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// A map of Key-Value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Parameters;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The Schema used for this Logic App Workflow.
        /// </summary>
        public readonly string WorkflowSchema;
        /// <summary>
        /// The version of the Schema used for this Logic App Workflow. Defaults to `1.0.0.0`.
        /// </summary>
        public readonly string WorkflowVersion;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetWorkflowResult(
            string accessEndpoint,
            string location,
            string name,
            ImmutableDictionary<string, string> parameters,
            string resourceGroupName,
            ImmutableDictionary<string, string> tags,
            string workflowSchema,
            string workflowVersion,
            string id)
        {
            AccessEndpoint = accessEndpoint;
            Location = location;
            Name = name;
            Parameters = parameters;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            WorkflowSchema = workflowSchema;
            WorkflowVersion = workflowVersion;
            Id = id;
        }
    }
}