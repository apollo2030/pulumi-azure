// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages network rules inside of a Azure Storage Account.
// 
// > **NOTE:** Network Rules can be defined either directly on the `storage.Account` resource, or using the `storage.AccountNetworkRules` resource - but the two cannot be used together. Spurious changes will occur if both are used against the same Storage Account.
// 
// > **NOTE:** Only one `storage.AccountNetworkRules` can be tied to an `storage.Account`. Spurious changes will occur if more than `storage.AccountNetworkRules` is tied to the same `storage.Account`.
// 
// > **NOTE:** Deleting this resource updates the storage account back to the default values it had when the storage account was created.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_account_network_rules.html.markdown.
type AccountNetworkRules struct {
	pulumi.CustomResourceState

	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
	Bypasses pulumi.StringArrayOutput `pulumi:"bypasses"`
	// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
	DefaultAction pulumi.StringOutput `pulumi:"defaultAction"`
	// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
	IpRules pulumi.StringArrayOutput `pulumi:"ipRules"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// A list of virtual network subnet ids to to secure the storage account.
	VirtualNetworkSubnetIds pulumi.StringArrayOutput `pulumi:"virtualNetworkSubnetIds"`
}

// NewAccountNetworkRules registers a new resource with the given unique name, arguments, and options.
func NewAccountNetworkRules(ctx *pulumi.Context,
	name string, args *AccountNetworkRulesArgs, opts ...pulumi.ResourceOption) (*AccountNetworkRules, error) {
	if args == nil || args.DefaultAction == nil {
		return nil, errors.New("missing required argument 'DefaultAction'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil {
		args = &AccountNetworkRulesArgs{}
	}
	var resource AccountNetworkRules
	err := ctx.RegisterResource("azure:storage/accountNetworkRules:AccountNetworkRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccountNetworkRules gets an existing AccountNetworkRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccountNetworkRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountNetworkRulesState, opts ...pulumi.ResourceOption) (*AccountNetworkRules, error) {
	var resource AccountNetworkRules
	err := ctx.ReadResource("azure:storage/accountNetworkRules:AccountNetworkRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccountNetworkRules resources.
type accountNetworkRulesState struct {
	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
	Bypasses []string `pulumi:"bypasses"`
	// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
	DefaultAction *string `pulumi:"defaultAction"`
	// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
	IpRules []string `pulumi:"ipRules"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	StorageAccountName *string `pulumi:"storageAccountName"`
	// A list of virtual network subnet ids to to secure the storage account.
	VirtualNetworkSubnetIds []string `pulumi:"virtualNetworkSubnetIds"`
}

type AccountNetworkRulesState struct {
	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
	Bypasses pulumi.StringArrayInput
	// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
	DefaultAction pulumi.StringPtrInput
	// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
	IpRules pulumi.StringArrayInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	StorageAccountName pulumi.StringPtrInput
	// A list of virtual network subnet ids to to secure the storage account.
	VirtualNetworkSubnetIds pulumi.StringArrayInput
}

func (AccountNetworkRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountNetworkRulesState)(nil)).Elem()
}

type accountNetworkRulesArgs struct {
	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
	Bypasses []string `pulumi:"bypasses"`
	// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
	DefaultAction string `pulumi:"defaultAction"`
	// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
	IpRules []string `pulumi:"ipRules"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	StorageAccountName string `pulumi:"storageAccountName"`
	// A list of virtual network subnet ids to to secure the storage account.
	VirtualNetworkSubnetIds []string `pulumi:"virtualNetworkSubnetIds"`
}

// The set of arguments for constructing a AccountNetworkRules resource.
type AccountNetworkRulesArgs struct {
	// Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Valid options are any combination of `Logging`, `Metrics`, `AzureServices`, or `None`.
	Bypasses pulumi.StringArrayInput
	// Specifies the default action of allow or deny when no other rules match. Valid options are `Deny` or `Allow`.
	DefaultAction pulumi.StringInput
	// List of public IP or IP ranges in CIDR Format. Only IPV4 addresses are allowed. Private IP address ranges (as defined in [RFC 1918](https://tools.ietf.org/html/rfc1918#section-3)) are not allowed.
	IpRules pulumi.StringArrayInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	StorageAccountName pulumi.StringInput
	// A list of virtual network subnet ids to to secure the storage account.
	VirtualNetworkSubnetIds pulumi.StringArrayInput
}

func (AccountNetworkRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountNetworkRulesArgs)(nil)).Elem()
}

