// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the connection status information about an existing Private Endpoint.
// 
// > **NOTE** Private Endpoint is currently in Public Preview.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/private_endpoint_connection.html.markdown.
func LookupEndpointConnection(ctx *pulumi.Context, args *GetEndpointConnectionArgs) (*GetEndpointConnectionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:privatelink/getEndpointConnection:getEndpointConnection", inputs)
	if err != nil {
		return nil, err
	}
	return &GetEndpointConnectionResult{
		Location: outputs["location"],
		Name: outputs["name"],
		PrivateServiceConnections: outputs["privateServiceConnections"],
		ResourceGroupName: outputs["resourceGroupName"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getEndpointConnection.
type GetEndpointConnectionArgs struct {
	// Specifies the Name of the private endpoint.
	Name interface{}
	// Specifies the Name of the Resource Group within which the private endpoint exists.
	ResourceGroupName interface{}
}

// A collection of values returned by getEndpointConnection.
type GetEndpointConnectionResult struct {
	// The supported Azure location where the resource exists.
	Location interface{}
	// The name of the private endpoint.
	Name interface{}
	PrivateServiceConnections interface{}
	ResourceGroupName interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}