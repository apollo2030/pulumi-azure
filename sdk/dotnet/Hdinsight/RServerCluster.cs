// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Hdinsight
{
    /// <summary>
    /// Manages a HDInsight RServer Cluster.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_rserver_cluster.html.markdown.
    /// </summary>
    public partial class RServerCluster : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("clusterVersion")]
        public Output<string> ClusterVersion { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for the Edge Node of the HDInsight RServer Cluster.
        /// </summary>
        [Output("edgeSshEndpoint")]
        public Output<string> EdgeSshEndpoint { get; private set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Output("gateway")]
        public Output<Outputs.RServerClusterGateway> Gateway { get; private set; } = null!;

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight RServer Cluster.
        /// </summary>
        [Output("httpsEndpoint")]
        public Output<string> HttpsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Output("roles")]
        public Output<Outputs.RServerClusterRoles> Roles { get; private set; } = null!;

        /// <summary>
        /// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
        /// </summary>
        [Output("rstudio")]
        public Output<bool> Rstudio { get; private set; } = null!;

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight RServer Cluster.
        /// </summary>
        [Output("sshEndpoint")]
        public Output<string> SshEndpoint { get; private set; } = null!;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        [Output("storageAccounts")]
        public Output<ImmutableArray<Outputs.RServerClusterStorageAccounts>> StorageAccounts { get; private set; } = null!;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight RServer Cluster.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a RServerCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RServerCluster(string name, RServerClusterArgs args, CustomResourceOptions? options = null)
            : base("azure:hdinsight/rServerCluster:RServerCluster", name, args, MakeResourceOptions(options, ""))
        {
        }

        private RServerCluster(string name, Input<string> id, RServerClusterState? state = null, CustomResourceOptions? options = null)
            : base("azure:hdinsight/rServerCluster:RServerCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RServerCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RServerCluster Get(string name, Input<string> id, RServerClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new RServerCluster(name, id, state, options);
        }
    }

    public sealed class RServerClusterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion", required: true)]
        public Input<string> ClusterVersion { get; set; } = null!;

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway", required: true)]
        public Input<Inputs.RServerClusterGatewayArgs> Gateway { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles", required: true)]
        public Input<Inputs.RServerClusterRolesArgs> Roles { get; set; } = null!;

        /// <summary>
        /// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
        /// </summary>
        [Input("rstudio", required: true)]
        public Input<bool> Rstudio { get; set; } = null!;

        [Input("storageAccounts", required: true)]
        private InputList<Inputs.RServerClusterStorageAccountsArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.RServerClusterStorageAccountsArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.RServerClusterStorageAccountsArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight RServer Cluster.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public RServerClusterArgs()
        {
        }
    }

    public sealed class RServerClusterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("clusterVersion")]
        public Input<string>? ClusterVersion { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for the Edge Node of the HDInsight RServer Cluster.
        /// </summary>
        [Input("edgeSshEndpoint")]
        public Input<string>? EdgeSshEndpoint { get; set; }

        /// <summary>
        /// A `gateway` block as defined below.
        /// </summary>
        [Input("gateway")]
        public Input<Inputs.RServerClusterGatewayGetArgs>? Gateway { get; set; }

        /// <summary>
        /// The HTTPS Connectivity Endpoint for this HDInsight RServer Cluster.
        /// </summary>
        [Input("httpsEndpoint")]
        public Input<string>? HttpsEndpoint { get; set; }

        /// <summary>
        /// Specifies the Azure Region which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name for this HDInsight RServer Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group in which this HDInsight RServer Cluster should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// A `roles` block as defined below.
        /// </summary>
        [Input("roles")]
        public Input<Inputs.RServerClusterRolesGetArgs>? Roles { get; set; }

        /// <summary>
        /// Should R Studio community edition for RServer be installed? Changing this forces a new resource to be created.
        /// </summary>
        [Input("rstudio")]
        public Input<bool>? Rstudio { get; set; }

        /// <summary>
        /// The SSH Connectivity Endpoint for this HDInsight RServer Cluster.
        /// </summary>
        [Input("sshEndpoint")]
        public Input<string>? SshEndpoint { get; set; }

        [Input("storageAccounts")]
        private InputList<Inputs.RServerClusterStorageAccountsGetArgs>? _storageAccounts;

        /// <summary>
        /// One or more `storage_account` block as defined below.
        /// </summary>
        public InputList<Inputs.RServerClusterStorageAccountsGetArgs> StorageAccounts
        {
            get => _storageAccounts ?? (_storageAccounts = new InputList<Inputs.RServerClusterStorageAccountsGetArgs>());
            set => _storageAccounts = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of Tags which should be assigned to this HDInsight RServer Cluster.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the Tier which should be used for this HDInsight RServer Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public RServerClusterState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class RServerClusterGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public RServerClusterGatewayArgs()
        {
        }
    }

    public sealed class RServerClusterGatewayGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public RServerClusterGatewayGetArgs()
        {
        }
    }

    public sealed class RServerClusterRolesArgs : Pulumi.ResourceArgs
    {
        [Input("edgeNode", required: true)]
        public Input<RServerClusterRolesEdgeNodeArgs> EdgeNode { get; set; } = null!;

        [Input("headNode", required: true)]
        public Input<RServerClusterRolesHeadNodeArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<RServerClusterRolesWorkerNodeArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<RServerClusterRolesZookeeperNodeArgs> ZookeeperNode { get; set; } = null!;

        public RServerClusterRolesArgs()
        {
        }
    }

    public sealed class RServerClusterRolesEdgeNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesEdgeNodeArgs()
        {
        }
    }

    public sealed class RServerClusterRolesEdgeNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesEdgeNodeGetArgs()
        {
        }
    }

    public sealed class RServerClusterRolesGetArgs : Pulumi.ResourceArgs
    {
        [Input("edgeNode", required: true)]
        public Input<RServerClusterRolesEdgeNodeGetArgs> EdgeNode { get; set; } = null!;

        [Input("headNode", required: true)]
        public Input<RServerClusterRolesHeadNodeGetArgs> HeadNode { get; set; } = null!;

        [Input("workerNode", required: true)]
        public Input<RServerClusterRolesWorkerNodeGetArgs> WorkerNode { get; set; } = null!;

        [Input("zookeeperNode", required: true)]
        public Input<RServerClusterRolesZookeeperNodeGetArgs> ZookeeperNode { get; set; } = null!;

        public RServerClusterRolesGetArgs()
        {
        }
    }

    public sealed class RServerClusterRolesHeadNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesHeadNodeArgs()
        {
        }
    }

    public sealed class RServerClusterRolesHeadNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesHeadNodeGetArgs()
        {
        }
    }

    public sealed class RServerClusterRolesWorkerNodeArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesWorkerNodeArgs()
        {
        }
    }

    public sealed class RServerClusterRolesWorkerNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("targetInstanceCount", required: true)]
        public Input<int> TargetInstanceCount { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesWorkerNodeGetArgs()
        {
        }
    }

    public sealed class RServerClusterRolesZookeeperNodeArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesZookeeperNodeArgs()
        {
        }
    }

    public sealed class RServerClusterRolesZookeeperNodeGetArgs : Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("sshKeys")]
        private InputList<string>? _sshKeys;
        public InputList<string> SshKeys
        {
            get => _sshKeys ?? (_sshKeys = new InputList<string>());
            set => _sshKeys = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        [Input("virtualNetworkId")]
        public Input<string>? VirtualNetworkId { get; set; }

        [Input("vmSize", required: true)]
        public Input<string> VmSize { get; set; } = null!;

        public RServerClusterRolesZookeeperNodeGetArgs()
        {
        }
    }

    public sealed class RServerClusterStorageAccountsArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public RServerClusterStorageAccountsArgs()
        {
        }
    }

    public sealed class RServerClusterStorageAccountsGetArgs : Pulumi.ResourceArgs
    {
        [Input("isDefault", required: true)]
        public Input<bool> IsDefault { get; set; } = null!;

        [Input("storageAccountKey", required: true)]
        public Input<string> StorageAccountKey { get; set; } = null!;

        [Input("storageContainerId", required: true)]
        public Input<string> StorageContainerId { get; set; } = null!;

        public RServerClusterStorageAccountsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class RServerClusterGateway
    {
        public readonly bool Enabled;
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private RServerClusterGateway(
            bool enabled,
            string password,
            string username)
        {
            Enabled = enabled;
            Password = password;
            Username = username;
        }
    }

    [OutputType]
    public sealed class RServerClusterRoles
    {
        public readonly RServerClusterRolesEdgeNode EdgeNode;
        public readonly RServerClusterRolesHeadNode HeadNode;
        public readonly RServerClusterRolesWorkerNode WorkerNode;
        public readonly RServerClusterRolesZookeeperNode ZookeeperNode;

        [OutputConstructor]
        private RServerClusterRoles(
            RServerClusterRolesEdgeNode edgeNode,
            RServerClusterRolesHeadNode headNode,
            RServerClusterRolesWorkerNode workerNode,
            RServerClusterRolesZookeeperNode zookeeperNode)
        {
            EdgeNode = edgeNode;
            HeadNode = headNode;
            WorkerNode = workerNode;
            ZookeeperNode = zookeeperNode;
        }
    }

    [OutputType]
    public sealed class RServerClusterRolesEdgeNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private RServerClusterRolesEdgeNode(
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class RServerClusterRolesHeadNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private RServerClusterRolesHeadNode(
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class RServerClusterRolesWorkerNode
    {
        public readonly int? MinInstanceCount;
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly int TargetInstanceCount;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private RServerClusterRolesWorkerNode(
            int? minInstanceCount,
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            int targetInstanceCount,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            MinInstanceCount = minInstanceCount;
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            TargetInstanceCount = targetInstanceCount;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class RServerClusterRolesZookeeperNode
    {
        public readonly string? Password;
        public readonly ImmutableArray<string> SshKeys;
        public readonly string? SubnetId;
        public readonly string Username;
        public readonly string? VirtualNetworkId;
        public readonly string VmSize;

        [OutputConstructor]
        private RServerClusterRolesZookeeperNode(
            string? password,
            ImmutableArray<string> sshKeys,
            string? subnetId,
            string username,
            string? virtualNetworkId,
            string vmSize)
        {
            Password = password;
            SshKeys = sshKeys;
            SubnetId = subnetId;
            Username = username;
            VirtualNetworkId = virtualNetworkId;
            VmSize = vmSize;
        }
    }

    [OutputType]
    public sealed class RServerClusterStorageAccounts
    {
        public readonly bool IsDefault;
        public readonly string StorageAccountKey;
        public readonly string StorageContainerId;

        [OutputConstructor]
        private RServerClusterStorageAccounts(
            bool isDefault,
            string storageAccountKey,
            string storageContainerId)
        {
            IsDefault = isDefault;
            StorageAccountKey = storageAccountKey;
            StorageContainerId = storageContainerId;
        }
    }
    }
}