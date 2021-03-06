// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an ExpressRoute Circuit Authorization.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/express_route_circuit_authorization.html.markdown.
type ExpressRouteCircuitAuthorization struct {
	pulumi.CustomResourceState

	// The Authorization Key.
	AuthorizationKey pulumi.StringOutput `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus pulumi.StringOutput `pulumi:"authorizationUseStatus"`
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringOutput `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewExpressRouteCircuitAuthorization registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	if args == nil || args.ExpressRouteCircuitName == nil {
		return nil, errors.New("missing required argument 'ExpressRouteCircuitName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ExpressRouteCircuitAuthorizationArgs{}
	}
	var resource ExpressRouteCircuitAuthorization
	err := ctx.RegisterResource("azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCircuitAuthorization gets an existing ExpressRouteCircuitAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	var resource ExpressRouteCircuitAuthorization
	err := ctx.ReadResource("azure:network/expressRouteCircuitAuthorization:ExpressRouteCircuitAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCircuitAuthorization resources.
type expressRouteCircuitAuthorizationState struct {
	// The Authorization Key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName *string `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ExpressRouteCircuitAuthorizationState struct {
	// The Authorization Key.
	AuthorizationKey pulumi.StringPtrInput
	// The authorization use status.
	AuthorizationUseStatus pulumi.StringPtrInput
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringPtrInput
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ExpressRouteCircuitAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationState)(nil)).Elem()
}

type expressRouteCircuitAuthorizationArgs struct {
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName string `pulumi:"expressRouteCircuitName"`
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRouteCircuitAuthorization resource.
type ExpressRouteCircuitAuthorizationArgs struct {
	// The name of the Express Route Circuit in which to create the Authorization.
	ExpressRouteCircuitName pulumi.StringInput
	// The name of the ExpressRoute circuit. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the ExpressRoute circuit. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRouteCircuitAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationArgs)(nil)).Elem()
}

