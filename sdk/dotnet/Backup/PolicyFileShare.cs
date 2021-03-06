// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup
{
    /// <summary>
    /// Manages an Azure File Share Backup Policy within a Recovery Services vault.
    /// 
    /// &gt; **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_policy_file_share.html.markdown.
    /// </summary>
    public partial class PolicyFileShare : Pulumi.CustomResource
    {
        /// <summary>
        /// Configures the Policy backup frequency and times as documented in the `backup` block below.
        /// </summary>
        [Output("backup")]
        public Output<Outputs.PolicyFileShareBackup> Backup { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Output("recoveryVaultName")]
        public Output<string> RecoveryVaultName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below.
        /// </summary>
        [Output("retentionDaily")]
        public Output<Outputs.PolicyFileShareRetentionDaily> RetentionDaily { get; private set; } = null!;

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyFileShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyFileShare(string name, PolicyFileShareArgs args, CustomResourceOptions? options = null)
            : base("azure:backup/policyFileShare:PolicyFileShare", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private PolicyFileShare(string name, Input<string> id, PolicyFileShareState? state = null, CustomResourceOptions? options = null)
            : base("azure:backup/policyFileShare:PolicyFileShare", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyFileShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyFileShare Get(string name, Input<string> id, PolicyFileShareState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyFileShare(name, id, state, options);
        }
    }

    public sealed class PolicyFileShareArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the Policy backup frequency and times as documented in the `backup` block below.
        /// </summary>
        [Input("backup", required: true)]
        public Input<Inputs.PolicyFileShareBackupArgs> Backup { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below.
        /// </summary>
        [Input("retentionDaily", required: true)]
        public Input<Inputs.PolicyFileShareRetentionDailyArgs> RetentionDaily { get; set; } = null!;

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public PolicyFileShareArgs()
        {
        }
    }

    public sealed class PolicyFileShareState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the Policy backup frequency and times as documented in the `backup` block below.
        /// </summary>
        [Input("backup")]
        public Input<Inputs.PolicyFileShareBackupGetArgs>? Backup { get; set; }

        /// <summary>
        /// Specifies the name of the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName")]
        public Input<string>? RecoveryVaultName { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Configures the policy daily retention as documented in the `retention_daily` block below.
        /// </summary>
        [Input("retentionDaily")]
        public Input<Inputs.PolicyFileShareRetentionDailyGetArgs>? RetentionDaily { get; set; }

        /// <summary>
        /// Specifies the timezone. Defaults to `UTC`
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public PolicyFileShareState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class PolicyFileShareBackupArgs : Pulumi.ResourceArgs
    {
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        public PolicyFileShareBackupArgs()
        {
        }
    }

    public sealed class PolicyFileShareBackupGetArgs : Pulumi.ResourceArgs
    {
        [Input("frequency", required: true)]
        public Input<string> Frequency { get; set; } = null!;

        [Input("time", required: true)]
        public Input<string> Time { get; set; } = null!;

        public PolicyFileShareBackupGetArgs()
        {
        }
    }

    public sealed class PolicyFileShareRetentionDailyArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public PolicyFileShareRetentionDailyArgs()
        {
        }
    }

    public sealed class PolicyFileShareRetentionDailyGetArgs : Pulumi.ResourceArgs
    {
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public PolicyFileShareRetentionDailyGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class PolicyFileShareBackup
    {
        public readonly string Frequency;
        public readonly string Time;

        [OutputConstructor]
        private PolicyFileShareBackup(
            string frequency,
            string time)
        {
            Frequency = frequency;
            Time = time;
        }
    }

    [OutputType]
    public sealed class PolicyFileShareRetentionDaily
    {
        public readonly int Count;

        [OutputConstructor]
        private PolicyFileShareRetentionDaily(int count)
        {
            Count = count;
        }
    }
    }
}
