// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    public static partial class Invokes
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Snapshot.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/netapp_snapshot.html.markdown.
        /// </summary>
        public static Task<GetSnapshotResult> GetSnapshot(GetSnapshotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("azure:netapp/getSnapshot:getSnapshot", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp Account where the NetApp Pool exists.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Snapshot.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Pool where the NetApp Volume exists.
        /// </summary>
        [Input("poolName", required: true)]
        public string PoolName { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Snapshot exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the NetApp Volume where the NetApp Snapshot exists.
        /// </summary>
        [Input("volumeName", required: true)]
        public string VolumeName { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSnapshotResult
    {
        public readonly string AccountName;
        /// <summary>
        /// The Azure Region where the NetApp Snapshot exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string PoolName;
        public readonly string ResourceGroupName;
        public readonly string VolumeName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSnapshotResult(
            string accountName,
            string location,
            string name,
            string poolName,
            string resourceGroupName,
            string volumeName,
            string id)
        {
            AccountName = accountName;
            Location = location;
            Name = name;
            PoolName = poolName;
            ResourceGroupName = resourceGroupName;
            VolumeName = volumeName;
            Id = id;
        }
    }
}
