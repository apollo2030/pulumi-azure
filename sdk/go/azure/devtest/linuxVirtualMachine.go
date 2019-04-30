// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devtest

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Linux Virtual Machine within a Dev Test Lab.
type LinuxVirtualMachine struct {
	s *pulumi.ResourceState
}

// NewLinuxVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewLinuxVirtualMachine(ctx *pulumi.Context,
	name string, args *LinuxVirtualMachineArgs, opts ...pulumi.ResourceOpt) (*LinuxVirtualMachine, error) {
	if args == nil || args.GalleryImageReference == nil {
		return nil, errors.New("missing required argument 'GalleryImageReference'")
	}
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.LabSubnetName == nil {
		return nil, errors.New("missing required argument 'LabSubnetName'")
	}
	if args == nil || args.LabVirtualNetworkId == nil {
		return nil, errors.New("missing required argument 'LabVirtualNetworkId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil || args.StorageType == nil {
		return nil, errors.New("missing required argument 'StorageType'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowClaim"] = nil
		inputs["disallowPublicIpAddress"] = nil
		inputs["galleryImageReference"] = nil
		inputs["inboundNatRules"] = nil
		inputs["labName"] = nil
		inputs["labSubnetName"] = nil
		inputs["labVirtualNetworkId"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["notes"] = nil
		inputs["password"] = nil
		inputs["resourceGroupName"] = nil
		inputs["size"] = nil
		inputs["sshKey"] = nil
		inputs["storageType"] = nil
		inputs["tags"] = nil
		inputs["username"] = nil
	} else {
		inputs["allowClaim"] = args.AllowClaim
		inputs["disallowPublicIpAddress"] = args.DisallowPublicIpAddress
		inputs["galleryImageReference"] = args.GalleryImageReference
		inputs["inboundNatRules"] = args.InboundNatRules
		inputs["labName"] = args.LabName
		inputs["labSubnetName"] = args.LabSubnetName
		inputs["labVirtualNetworkId"] = args.LabVirtualNetworkId
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["notes"] = args.Notes
		inputs["password"] = args.Password
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["size"] = args.Size
		inputs["sshKey"] = args.SshKey
		inputs["storageType"] = args.StorageType
		inputs["tags"] = args.Tags
		inputs["username"] = args.Username
	}
	inputs["fqdn"] = nil
	inputs["uniqueIdentifier"] = nil
	s, err := ctx.RegisterResource("azure:devtest/linuxVirtualMachine:LinuxVirtualMachine", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinuxVirtualMachine{s: s}, nil
}

// GetLinuxVirtualMachine gets an existing LinuxVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinuxVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.ID, state *LinuxVirtualMachineState, opts ...pulumi.ResourceOpt) (*LinuxVirtualMachine, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowClaim"] = state.AllowClaim
		inputs["disallowPublicIpAddress"] = state.DisallowPublicIpAddress
		inputs["fqdn"] = state.Fqdn
		inputs["galleryImageReference"] = state.GalleryImageReference
		inputs["inboundNatRules"] = state.InboundNatRules
		inputs["labName"] = state.LabName
		inputs["labSubnetName"] = state.LabSubnetName
		inputs["labVirtualNetworkId"] = state.LabVirtualNetworkId
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["notes"] = state.Notes
		inputs["password"] = state.Password
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["size"] = state.Size
		inputs["sshKey"] = state.SshKey
		inputs["storageType"] = state.StorageType
		inputs["tags"] = state.Tags
		inputs["uniqueIdentifier"] = state.UniqueIdentifier
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("azure:devtest/linuxVirtualMachine:LinuxVirtualMachine", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &LinuxVirtualMachine{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *LinuxVirtualMachine) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *LinuxVirtualMachine) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Can this Virtual Machine be claimed by users? Defaults to `true`.
func (r *LinuxVirtualMachine) AllowClaim() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowClaim"])
}

// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) DisallowPublicIpAddress() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disallowPublicIpAddress"])
}

// The FQDN of the Virtual Machine.
func (r *LinuxVirtualMachine) Fqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fqdn"])
}

// A `gallery_image_reference` block as defined below.
func (r *LinuxVirtualMachine) GalleryImageReference() *pulumi.Output {
	return r.s.State["galleryImageReference"]
}

// One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) InboundNatRules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["inboundNatRules"])
}

// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) LabName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labName"])
}

// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) LabSubnetName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labSubnetName"])
}

// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) LabVirtualNetworkId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labVirtualNetworkId"])
}

// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Any notes about the Virtual Machine.
func (r *LinuxVirtualMachine) Notes() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["notes"])
}

// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) Size() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["size"])
}

// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) SshKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshKey"])
}

// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
func (r *LinuxVirtualMachine) StorageType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["storageType"])
}

// A mapping of tags to assign to the resource.
func (r *LinuxVirtualMachine) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The unique immutable identifier of the Virtual Machine.
func (r *LinuxVirtualMachine) UniqueIdentifier() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["uniqueIdentifier"])
}

// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
func (r *LinuxVirtualMachine) Username() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering LinuxVirtualMachine resources.
type LinuxVirtualMachineState struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim interface{}
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress interface{}
	// The FQDN of the Virtual Machine.
	Fqdn interface{}
	// A `gallery_image_reference` block as defined below.
	GalleryImageReference interface{}
	// One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules interface{}
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName interface{}
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName interface{}
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId interface{}
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name interface{}
	// Any notes about the Virtual Machine.
	Notes interface{}
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password interface{}
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size interface{}
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey interface{}
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The unique immutable identifier of the Virtual Machine.
	UniqueIdentifier interface{}
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username interface{}
}

// The set of arguments for constructing a LinuxVirtualMachine resource.
type LinuxVirtualMachineArgs struct {
	// Can this Virtual Machine be claimed by users? Defaults to `true`.
	AllowClaim interface{}
	// Should the Virtual Machine be created without a Public IP Address? Changing this forces a new resource to be created.
	DisallowPublicIpAddress interface{}
	// A `gallery_image_reference` block as defined below.
	GalleryImageReference interface{}
	// One or more `inbound_nat_rule` blocks as defined below. Changing this forces a new resource to be created.
	InboundNatRules interface{}
	// Specifies the name of the Dev Test Lab in which the Virtual Machine should be created. Changing this forces a new resource to be created.
	LabName interface{}
	// The name of a Subnet within the Dev Test Virtual Network where this machine should exist. Changing this forces a new resource to be created.
	LabSubnetName interface{}
	// The ID of the Dev Test Virtual Network where this Virtual Machine should be created. Changing this forces a new resource to be created.
	LabVirtualNetworkId interface{}
	// Specifies the supported Azure location where the Dev Test Lab exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Dev Test Machine. Changing this forces a new resource to be created.
	Name interface{}
	// Any notes about the Virtual Machine.
	Notes interface{}
	// The Password associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	Password interface{}
	// The name of the resource group in which the Dev Test Lab resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The Machine Size to use for this Virtual Machine, such as `Standard_F2`. Changing this forces a new resource to be created.
	Size interface{}
	// The SSH Key associated with the `username` used to login to this Virtual Machine. Changing this forces a new resource to be created.
	SshKey interface{}
	// The type of Storage to use on this Virtual Machine. Possible values are `Standard` and `Premium`.
	StorageType interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The Username associated with the local administrator on this Virtual Machine. Changing this forces a new resource to be created.
	Username interface{}
}
