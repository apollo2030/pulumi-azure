// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package privatelink

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Private Link Service.
// 
// > **NOTE** Private Link is currently in Public Preview.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/private_link_service.html.markdown.
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	var rv GetServiceResult
	err := ctx.Invoke("azure:privatelink/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// The name of the private link service.
	Name string `pulumi:"name"`
	// The name of the resource group in which the private link service resides.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getService.
type GetServiceResult struct {
	// The alias is a globally unique name for your private link service which Azure generates for you. Your can use this alias to request a connection to your private link service.
	Alias string `pulumi:"alias"`
	// The list of subscription(s) globally unique identifiers that will be auto approved to use the private link service.
	AutoApprovalSubscriptionIds []string `pulumi:"autoApprovalSubscriptionIds"`
	// Does the Private Link Service support the Proxy Protocol?
	EnableProxyProtocol bool `pulumi:"enableProxyProtocol"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of Standard Load Balancer(SLB) resource IDs. The Private Link service is tied to the frontend IP address of a SLB. All traffic destined for the private link service will reach the frontend of the SLB. You can configure SLB rules to direct this traffic to appropriate backend pools where your applications are running.
	LoadBalancerFrontendIpConfigurationIds []string `pulumi:"loadBalancerFrontendIpConfigurationIds"`
	// The supported Azure location where the resource exists.
	Location string `pulumi:"location"`
	// The name of private link service NAT IP configuration.
	Name string `pulumi:"name"`
	// The `natIpConfiguration` block as defined below.
	NatIpConfiguration GetServiceNatIpConfiguration `pulumi:"natIpConfiguration"`
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The list of subscription(s) globally unique identifiers(GUID) that will be able to see the private link service.
	VisibilitySubscriptionIds []string `pulumi:"visibilitySubscriptionIds"`
}

