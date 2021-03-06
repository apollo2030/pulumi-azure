// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing API Management Product.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/api_management_product.html.markdown.
func LookupProduct(ctx *pulumi.Context, args *LookupProductArgs, opts ...pulumi.InvokeOption) (*LookupProductResult, error) {
	var rv LookupProductResult
	err := ctx.Invoke("azure:apimanagement/getProduct:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProduct.
type LookupProductArgs struct {
	// The Name of the API Management Service in which this Product exists.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The Identifier for the API Management Product.
	ProductId string `pulumi:"productId"`
	// The Name of the Resource Group in which the API Management Service exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getProduct.
type LookupProductResult struct {
	ApiManagementName string `pulumi:"apiManagementName"`
	// Do subscribers need to be approved prior to being able to use the Product?
	ApprovalRequired bool `pulumi:"approvalRequired"`
	// The description of this Product, which may include HTML formatting tags.
	Description string `pulumi:"description"`
	// The Display Name for this API Management Product.
	DisplayName string `pulumi:"displayName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	ProductId string `pulumi:"productId"`
	// Is this Product Published?
	Published bool `pulumi:"published"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Is a Subscription required to access API's included in this Product?
	SubscriptionRequired bool `pulumi:"subscriptionRequired"`
	// The number of subscriptions a user can have to this Product at the same time.
	SubscriptionsLimit int `pulumi:"subscriptionsLimit"`
	// Any Terms and Conditions for this Product, which must be accepted by Developers before they can begin the Subscription process.
	Terms string `pulumi:"terms"`
}

