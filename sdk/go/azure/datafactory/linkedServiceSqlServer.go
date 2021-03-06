// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Linked Service (connection) between a SQL Server and Azure Data Factory.
// 
// > **Note:** All arguments including the client secret will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_factory_linked_service_sql_server.html.markdown.
type LinkedServiceSqlServer struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceSqlServer registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, args *LinkedServiceSqlServerArgs, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	if args == nil || args.ConnectionString == nil {
		return nil, errors.New("missing required argument 'ConnectionString'")
	}
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LinkedServiceSqlServerArgs{}
	}
	var resource LinkedServiceSqlServer
	err := ctx.RegisterResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceSqlServer gets an existing LinkedServiceSqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceSqlServerState, opts ...pulumi.ResourceOption) (*LinkedServiceSqlServer, error) {
	var resource LinkedServiceSqlServer
	err := ctx.ReadResource("azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceSqlServer resources.
type linkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceSqlServerState struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceSqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerState)(nil)).Elem()
}

type linkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service SQL Server.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceSqlServer resource.
type LinkedServiceSqlServerArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service SQL Server.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service SQL Server.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with the SQL Server.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service SQL Server.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service SQL Server.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service SQL Server. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service SQL Server.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service SQL Server. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceSqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceSqlServerArgs)(nil)).Elem()
}

