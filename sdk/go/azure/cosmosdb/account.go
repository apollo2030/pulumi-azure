// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a CosmosDB (formally DocumentDB) Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/cosmosdb_account.html.markdown.
type Account struct {
	pulumi.CustomResourceState

	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayOutput `pulumi:"capabilities"`
	// A list of connection strings available for this CosmosDB account. If the kind is `GlobalDocumentDB`, this will be empty.
	ConnectionStrings pulumi.StringArrayOutput `pulumi:"connectionStrings"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyOutput `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrOutput `pulumi:"enableAutomaticFailover"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrOutput `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	FailoverPolicies AccountFailoverPolicyArrayOutput `pulumi:"failoverPolicies"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayOutput `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrOutput `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrOutput `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the Azure region to host replicated data.
	Location pulumi.StringOutput `pulumi:"location"`
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`,`EnableMongo`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringOutput `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey pulumi.StringOutput `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey pulumi.StringOutput `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayOutput `pulumi:"readEndpoints"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey pulumi.StringOutput `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey pulumi.StringOutput `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayOutput `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayOutput `pulumi:"writeEndpoints"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.ConsistencyPolicy == nil {
		return nil, errors.New("missing required argument 'ConsistencyPolicy'")
	}
	if args == nil || args.OfferType == nil {
		return nil, errors.New("missing required argument 'OfferType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("azure:cosmosdb/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:cosmosdb/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities []AccountCapability `pulumi:"capabilities"`
	// A list of connection strings available for this CosmosDB account. If the kind is `GlobalDocumentDB`, this will be empty.
	ConnectionStrings []string `pulumi:"connectionStrings"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy *AccountConsistencyPolicy `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// The endpoint used to connect to the CosmosDB account.
	Endpoint *string `pulumi:"endpoint"`
	FailoverPolicies []AccountFailoverPolicy `pulumi:"failoverPolicies"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations []AccountGeoLocation `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter *string `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// The name of the Azure region to host replicated data.
	Location *string `pulumi:"location"`
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`,`EnableMongo`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name *string `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType *string `pulumi:"offerType"`
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey *string `pulumi:"primaryMasterKey"`
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey *string `pulumi:"primaryReadonlyMasterKey"`
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints []string `pulumi:"readEndpoints"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey *string `pulumi:"secondaryMasterKey"`
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey *string `pulumi:"secondaryReadonlyMasterKey"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules []AccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints []string `pulumi:"writeEndpoints"`
}

type AccountState struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayInput
	// A list of connection strings available for this CosmosDB account. If the kind is `GlobalDocumentDB`, this will be empty.
	ConnectionStrings pulumi.StringArrayInput
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyPtrInput
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// The endpoint used to connect to the CosmosDB account.
	Endpoint pulumi.StringPtrInput
	FailoverPolicies AccountFailoverPolicyArrayInput
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayInput
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrInput
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// The name of the Azure region to host replicated data.
	Location pulumi.StringPtrInput
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`,`EnableMongo`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringPtrInput
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringPtrInput
	// The Primary master key for the CosmosDB Account.
	PrimaryMasterKey pulumi.StringPtrInput
	// The Primary read-only master Key for the CosmosDB Account.
	PrimaryReadonlyMasterKey pulumi.StringPtrInput
	// A list of read endpoints available for this CosmosDB account.
	ReadEndpoints pulumi.StringArrayInput
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary master key for the CosmosDB Account.
	SecondaryMasterKey pulumi.StringPtrInput
	// The Secondary read-only master key for the CosmosDB Account.
	SecondaryReadonlyMasterKey pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayInput
	// A list of write endpoints available for this CosmosDB account.
	WriteEndpoints pulumi.StringArrayInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities []AccountCapability `pulumi:"capabilities"`
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicy `pulumi:"consistencyPolicy"`
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	FailoverPolicies []AccountFailoverPolicy `pulumi:"failoverPolicies"`
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations []AccountGeoLocation `pulumi:"geoLocations"`
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter *string `pulumi:"ipRangeFilter"`
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind *string `pulumi:"kind"`
	// The name of the Azure region to host replicated data.
	Location *string `pulumi:"location"`
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`,`EnableMongo`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name *string `pulumi:"name"`
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType string `pulumi:"offerType"`
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules []AccountVirtualNetworkRule `pulumi:"virtualNetworkRules"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Capabilities AccountCapabilityArrayInput
	// Specifies a `consistencyPolicy` resource, used to define the consistency policy for this CosmosDB account.
	ConsistencyPolicy AccountConsistencyPolicyInput
	// Enable automatic fail over for this Cosmos DB account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enable multi-master support for this Cosmos DB account.
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	FailoverPolicies AccountFailoverPolicyArrayInput
	// Specifies a `geoLocation` resource, used to define where data should be replicated with the `failoverPriority` 0 specifying the primary location.
	GeoLocations AccountGeoLocationArrayInput
	// CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
	IpRangeFilter pulumi.StringPtrInput
	// Enables virtual network filtering for this Cosmos DB account.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
	Kind pulumi.StringPtrInput
	// The name of the Azure region to host replicated data.
	Location pulumi.StringPtrInput
	// The capability to enable - Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`,`EnableMongo`, `EnableTable`, `MongoDBv3.4`, and `mongoEnableDocLevelTTL`.
	Name pulumi.StringPtrInput
	// Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
	OfferType pulumi.StringInput
	// The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies a `virtualNetworkRules` resource, used to define which subnets are allowed to access this CosmosDB account.
	VirtualNetworkRules AccountVirtualNetworkRuleArrayInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

