// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devspace

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a DevSpace Controller.
type Controller struct {
	s *pulumi.ResourceState
}

// NewController registers a new resource with the given unique name, arguments, and options.
func NewController(ctx *pulumi.Context,
	name string, args *ControllerArgs, opts ...pulumi.ResourceOpt) (*Controller, error) {
	if args == nil || args.HostSuffix == nil {
		return nil, errors.New("missing required argument 'HostSuffix'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil || args.TargetContainerHostCredentialsBase64 == nil {
		return nil, errors.New("missing required argument 'TargetContainerHostCredentialsBase64'")
	}
	if args == nil || args.TargetContainerHostResourceId == nil {
		return nil, errors.New("missing required argument 'TargetContainerHostResourceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["hostSuffix"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
		inputs["targetContainerHostCredentialsBase64"] = nil
		inputs["targetContainerHostResourceId"] = nil
	} else {
		inputs["hostSuffix"] = args.HostSuffix
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
		inputs["targetContainerHostCredentialsBase64"] = args.TargetContainerHostCredentialsBase64
		inputs["targetContainerHostResourceId"] = args.TargetContainerHostResourceId
	}
	inputs["dataPlaneFqdn"] = nil
	s, err := ctx.RegisterResource("azure:devspace/controller:Controller", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Controller{s: s}, nil
}

// GetController gets an existing Controller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetController(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ControllerState, opts ...pulumi.ResourceOpt) (*Controller, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dataPlaneFqdn"] = state.DataPlaneFqdn
		inputs["hostSuffix"] = state.HostSuffix
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["targetContainerHostCredentialsBase64"] = state.TargetContainerHostCredentialsBase64
		inputs["targetContainerHostResourceId"] = state.TargetContainerHostResourceId
	}
	s, err := ctx.ReadResource("azure:devspace/controller:Controller", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Controller{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Controller) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Controller) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// DNS name for accessing DataPlane services.
func (r *Controller) DataPlaneFqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dataPlaneFqdn"])
}

// The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
func (r *Controller) HostSuffix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostSuffix"])
}

// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
func (r *Controller) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
func (r *Controller) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
func (r *Controller) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A `sku` block as documented below. Changing this forces a new resource to be created.
func (r *Controller) Sku() *pulumi.Output {
	return r.s.State["sku"]
}

// A mapping of tags to assign to the resource.
func (r *Controller) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
func (r *Controller) TargetContainerHostCredentialsBase64() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetContainerHostCredentialsBase64"])
}

// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
func (r *Controller) TargetContainerHostResourceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetContainerHostResourceId"])
}

// Input properties used for looking up and filtering Controller resources.
type ControllerState struct {
	// DNS name for accessing DataPlane services.
	DataPlaneFqdn interface{}
	// The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
	HostSuffix interface{}
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `sku` block as documented below. Changing this forces a new resource to be created.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 interface{}
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId interface{}
}

// The set of arguments for constructing a Controller resource.
type ControllerArgs struct {
	// The host suffix for the DevSpace Controller. Changing this forces a new resource to be created.
	HostSuffix interface{}
	// Specifies the supported location where the DevSpace Controller should exist. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the DevSpace Controller. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the resource group under which the DevSpace Controller resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A `sku` block as documented below. Changing this forces a new resource to be created.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Base64 encoding of `kube_config_raw` of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostCredentialsBase64 interface{}
	// The resource id of Azure Kubernetes Service cluster. Changing this forces a new resource to be created.
	TargetContainerHostResourceId interface{}
}
