// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Gremlin Database within a Cosmos DB Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cosmosdb_gremlin_database.html.markdown.
type GremlinDatabase struct {
	pulumi.CustomResourceState

	// The name of the CosmosDB Account to create the Gremlin Database within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// Specifies the name of the Cosmos DB Gremlin Database. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Gremlin Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	Throughput pulumi.IntOutput `pulumi:"throughput"`
}

// NewGremlinDatabase registers a new resource with the given unique name, arguments, and options.
func NewGremlinDatabase(ctx *pulumi.Context,
	name string, args *GremlinDatabaseArgs, opts ...pulumi.ResourceOption) (*GremlinDatabase, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GremlinDatabaseArgs{}
	}
	var resource GremlinDatabase
	err := ctx.RegisterResource("azure:cosmosdb/gremlinDatabase:GremlinDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGremlinDatabase gets an existing GremlinDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGremlinDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinDatabaseState, opts ...pulumi.ResourceOption) (*GremlinDatabase, error) {
	var resource GremlinDatabase
	err := ctx.ReadResource("azure:cosmosdb/gremlinDatabase:GremlinDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GremlinDatabase resources.
type gremlinDatabaseState struct {
	// The name of the CosmosDB Account to create the Gremlin Database within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// Specifies the name of the Cosmos DB Gremlin Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Gremlin Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	Throughput *int `pulumi:"throughput"`
}

type GremlinDatabaseState struct {
	// The name of the CosmosDB Account to create the Gremlin Database within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// Specifies the name of the Cosmos DB Gremlin Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	Throughput pulumi.IntPtrInput
}

func (GremlinDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinDatabaseState)(nil)).Elem()
}

type gremlinDatabaseArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Database within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// Specifies the name of the Cosmos DB Gremlin Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Cosmos DB Gremlin Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Throughput *int `pulumi:"throughput"`
}

// The set of arguments for constructing a GremlinDatabase resource.
type GremlinDatabaseArgs struct {
	// The name of the CosmosDB Account to create the Gremlin Database within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// Specifies the name of the Cosmos DB Gremlin Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB Gremlin Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	Throughput pulumi.IntPtrInput
}

func (GremlinDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinDatabaseArgs)(nil)).Elem()
}

