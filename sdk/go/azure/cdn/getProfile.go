// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package cdn

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing CDN Profile.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/cdn_profile.html.markdown.
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	var rv LookupProfileResult
	err := ctx.Invoke("azure:cdn/getProfile:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProfile.
type LookupProfileArgs struct {
	// The name of the CDN Profile.
	Name string `pulumi:"name"`
	// The name of the resource group in which the CDN Profile exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getProfile.
type LookupProfileResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Azure Region where the resource exists.
	Location string `pulumi:"location"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The pricing related information of current CDN profile.
	Sku string `pulumi:"sku"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

