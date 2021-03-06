// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Linked Service (connection) between PostgreSQL and Azure Data Factory.
// 
// > **Note:** All arguments including the connectionString will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_factory_linked_service_postgresql.html.markdown.
type LinkedServicePostgresql struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServicePostgresql registers a new resource with the given unique name, arguments, and options.
func NewLinkedServicePostgresql(ctx *pulumi.Context,
	name string, args *LinkedServicePostgresqlArgs, opts ...pulumi.ResourceOption) (*LinkedServicePostgresql, error) {
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
		args = &LinkedServicePostgresqlArgs{}
	}
	var resource LinkedServicePostgresql
	err := ctx.RegisterResource("azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServicePostgresql gets an existing LinkedServicePostgresql resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServicePostgresql(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServicePostgresqlState, opts ...pulumi.ResourceOption) (*LinkedServicePostgresql, error) {
	var resource LinkedServicePostgresql
	err := ctx.ReadResource("azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServicePostgresql resources.
type linkedServicePostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServicePostgresqlState struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServicePostgresqlState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServicePostgresqlState)(nil)).Elem()
}

type linkedServicePostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service PostgreSQL.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServicePostgresql resource.
type LinkedServicePostgresqlArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service PostgreSQL.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service PostgreSQL.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with PostgreSQL.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service PostgreSQL.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service PostgreSQL.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service PostgreSQL. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service PostgreSQL.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service PostgreSQL. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServicePostgresqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServicePostgresqlArgs)(nil)).Elem()
}

