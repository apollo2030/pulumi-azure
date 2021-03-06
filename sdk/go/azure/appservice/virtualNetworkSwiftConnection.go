// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VirtualNetworkSwiftConnection struct {
	pulumi.CustomResourceState

	AppServiceId pulumi.StringOutput `pulumi:"appServiceId"`
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewVirtualNetworkSwiftConnection registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkSwiftConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	if args == nil || args.AppServiceId == nil {
		return nil, errors.New("missing required argument 'AppServiceId'")
	}
	if args == nil || args.SubnetId == nil {
		return nil, errors.New("missing required argument 'SubnetId'")
	}
	if args == nil {
		args = &VirtualNetworkSwiftConnectionArgs{}
	}
	var resource VirtualNetworkSwiftConnection
	err := ctx.RegisterResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkSwiftConnection gets an existing VirtualNetworkSwiftConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkSwiftConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkSwiftConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkSwiftConnection, error) {
	var resource VirtualNetworkSwiftConnection
	err := ctx.ReadResource("azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkSwiftConnection resources.
type virtualNetworkSwiftConnectionState struct {
	AppServiceId *string `pulumi:"appServiceId"`
	SubnetId *string `pulumi:"subnetId"`
}

type VirtualNetworkSwiftConnectionState struct {
	AppServiceId pulumi.StringPtrInput
	SubnetId pulumi.StringPtrInput
}

func (VirtualNetworkSwiftConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionState)(nil)).Elem()
}

type virtualNetworkSwiftConnectionArgs struct {
	AppServiceId string `pulumi:"appServiceId"`
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a VirtualNetworkSwiftConnection resource.
type VirtualNetworkSwiftConnectionArgs struct {
	AppServiceId pulumi.StringInput
	SubnetId pulumi.StringInput
}

func (VirtualNetworkSwiftConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkSwiftConnectionArgs)(nil)).Elem()
}

