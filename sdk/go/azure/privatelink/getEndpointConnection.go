// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package privatelink

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the connection status information about an existing Private Endpoint Connection.
// 
// > **NOTE** Private Endpoint is currently in Public Preview.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/private_endpoint_connection.html.markdown.
func GetEndpointConnection(ctx *pulumi.Context, args *GetEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*GetEndpointConnectionResult, error) {
	var rv GetEndpointConnectionResult
	err := ctx.Invoke("azure:privatelink/getEndpointConnection:getEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointConnection.
type GetEndpointConnectionArgs struct {
	// Specifies the Name of the private endpoint.
	Name string `pulumi:"name"`
	// Specifies the Name of the Resource Group within which the private endpoint exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getEndpointConnection.
type GetEndpointConnectionResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The name of the private endpoint.
	Name string `pulumi:"name"`
	PrivateServiceConnections []GetEndpointConnectionPrivateServiceConnection `pulumi:"privateServiceConnections"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

