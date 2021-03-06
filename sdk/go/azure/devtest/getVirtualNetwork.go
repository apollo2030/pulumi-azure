// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package devtest

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Dev Test Lab Virtual Network.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/dev_test_virtual_network.html.markdown.
func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure:devtest/getVirtualNetwork:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualNetwork.
type LookupVirtualNetworkArgs struct {
	// Specifies the name of the Dev Test Lab.
	LabName string `pulumi:"labName"`
	// Specifies the name of the Virtual Network.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group that contains the Virtual Network.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getVirtualNetwork.
type LookupVirtualNetworkResult struct {
	// The list of subnets enabled for the virtual network as defined below.
	AllowedSubnets []GetVirtualNetworkAllowedSubnet `pulumi:"allowedSubnets"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	LabName string `pulumi:"labName"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of permission overrides for the subnets as defined below.
	SubnetOverrides GetVirtualNetworkSubnetOverrides `pulumi:"subnetOverrides"`
	// The unique immutable identifier of the virtual network.
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
}

