// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package maps

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Azure Maps Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/maps_account.html.markdown.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure:maps/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// Specifies the name of the Maps Account.
	Name string `pulumi:"name"`
	// Specifies the name of the Resource Group in which the Maps Account is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Tags map[string]string `pulumi:"tags"`
}


// A collection of values returned by getAccount.
type LookupAccountResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs.
	PrimaryAccessKey string `pulumi:"primaryAccessKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The primary key used to authenticate and authorize access to the Maps REST APIs. The second key is given to provide seamless key regeneration.
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
	// The sku of the Azure Maps Account.
	SkuName string `pulumi:"skuName"`
	Tags map[string]string `pulumi:"tags"`
	// A unique identifier for the Maps Account.
	XMsClientId string `pulumi:"xMsClientId"`
}

