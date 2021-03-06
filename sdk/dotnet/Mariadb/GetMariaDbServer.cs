// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MariaDB
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing MariaDB Server.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/mariadb_server.html.markdown.
        /// </summary>
        public static Task<GetMariaDbServerResult> GetMariaDbServer(GetMariaDbServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMariaDbServerResult>("azure:mariadb/getMariaDbServer:getMariaDbServer", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetMariaDbServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the MariaDB Server to retrieve information about.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the MariaDB Server exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMariaDbServerArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetMariaDbServerResult
    {
        /// <summary>
        /// The Administrator Login for the MariaDB Server.
        /// </summary>
        public readonly string AdministratorLogin;
        /// <summary>
        /// The password associated with the `administrator_login` for the MariaDB Server.
        /// </summary>
        public readonly string AdministratorLoginPassword;
        /// <summary>
        /// The FQDN of the MariaDB Server.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// The Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The SKU Name for this MariaDB Server. 
        /// </summary>
        public readonly string SkuName;
        /// <summary>
        /// The SSL being enforced on connections.
        /// </summary>
        public readonly string SslEnforcement;
        /// <summary>
        /// A `storage_profile` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMariaDbServerStorageProfilesResult> StorageProfiles;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// ---
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The version of MariaDB being used.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetMariaDbServerResult(
            string administratorLogin,
            string administratorLoginPassword,
            string fqdn,
            string location,
            string name,
            string resourceGroupName,
            string skuName,
            string sslEnforcement,
            ImmutableArray<Outputs.GetMariaDbServerStorageProfilesResult> storageProfiles,
            ImmutableDictionary<string, string> tags,
            string version,
            string id)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            Fqdn = fqdn;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            SkuName = skuName;
            SslEnforcement = sslEnforcement;
            StorageProfiles = storageProfiles;
            Tags = tags;
            Version = version;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetMariaDbServerStorageProfilesResult
    {
        /// <summary>
        /// Whether autogrow is enabled or disabled for the storage.
        /// </summary>
        public readonly string AutoGrow;
        /// <summary>
        /// Backup retention days for the server.
        /// </summary>
        public readonly int BackupRetentionDays;
        /// <summary>
        /// Whether Geo-redundant is enabled or not for server backup.
        /// </summary>
        public readonly string GeoRedundantBackup;
        /// <summary>
        /// The max storage allowed for a server.
        /// </summary>
        public readonly int StorageMb;

        [OutputConstructor]
        private GetMariaDbServerStorageProfilesResult(
            string autoGrow,
            int backupRetentionDays,
            string geoRedundantBackup,
            int storageMb)
        {
            AutoGrow = autoGrow;
            BackupRetentionDays = backupRetentionDays;
            GeoRedundantBackup = geoRedundantBackup;
            StorageMb = storageMb;
        }
    }
    }
}
