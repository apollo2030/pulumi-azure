// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages attaching a Disk to a Virtual Machine.
// 
// > **NOTE:** Data Disks can be attached either directly on the `azurerm_virtual_machine` resource, or using the `azurerm_virtual_machine_data_disk_attachment` resource - but the two cannot be used together. If both are used against the same Virtual Machine, spurious changes will occur.
// 
// > **Please Note:** only Managed Disks are supported via this separate resource, Unmanaged Disks can be attached using the `storage_data_disk` block in the `azurerm_virtual_machine` resource.
type DataDiskAttachment struct {
	s *pulumi.ResourceState
}

// NewDataDiskAttachment registers a new resource with the given unique name, arguments, and options.
func NewDataDiskAttachment(ctx *pulumi.Context,
	name string, args *DataDiskAttachmentArgs, opts ...pulumi.ResourceOpt) (*DataDiskAttachment, error) {
	if args == nil || args.Caching == nil {
		return nil, errors.New("missing required argument 'Caching'")
	}
	if args == nil || args.Lun == nil {
		return nil, errors.New("missing required argument 'Lun'")
	}
	if args == nil || args.ManagedDiskId == nil {
		return nil, errors.New("missing required argument 'ManagedDiskId'")
	}
	if args == nil || args.VirtualMachineId == nil {
		return nil, errors.New("missing required argument 'VirtualMachineId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["caching"] = nil
		inputs["createOption"] = nil
		inputs["lun"] = nil
		inputs["managedDiskId"] = nil
		inputs["virtualMachineId"] = nil
		inputs["writeAcceleratorEnabled"] = nil
	} else {
		inputs["caching"] = args.Caching
		inputs["createOption"] = args.CreateOption
		inputs["lun"] = args.Lun
		inputs["managedDiskId"] = args.ManagedDiskId
		inputs["virtualMachineId"] = args.VirtualMachineId
		inputs["writeAcceleratorEnabled"] = args.WriteAcceleratorEnabled
	}
	s, err := ctx.RegisterResource("azure:compute/dataDiskAttachment:DataDiskAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataDiskAttachment{s: s}, nil
}

// GetDataDiskAttachment gets an existing DataDiskAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataDiskAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DataDiskAttachmentState, opts ...pulumi.ResourceOpt) (*DataDiskAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["caching"] = state.Caching
		inputs["createOption"] = state.CreateOption
		inputs["lun"] = state.Lun
		inputs["managedDiskId"] = state.ManagedDiskId
		inputs["virtualMachineId"] = state.VirtualMachineId
		inputs["writeAcceleratorEnabled"] = state.WriteAcceleratorEnabled
	}
	s, err := ctx.ReadResource("azure:compute/dataDiskAttachment:DataDiskAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DataDiskAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DataDiskAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DataDiskAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
func (r *DataDiskAttachment) Caching() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["caching"])
}

// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
func (r *DataDiskAttachment) CreateOption() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createOption"])
}

// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
func (r *DataDiskAttachment) Lun() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["lun"])
}

// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
func (r *DataDiskAttachment) ManagedDiskId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managedDiskId"])
}

// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
func (r *DataDiskAttachment) VirtualMachineId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualMachineId"])
}

// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
func (r *DataDiskAttachment) WriteAcceleratorEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["writeAcceleratorEnabled"])
}

// Input properties used for looking up and filtering DataDiskAttachment resources.
type DataDiskAttachmentState struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching interface{}
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption interface{}
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun interface{}
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId interface{}
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId interface{}
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled interface{}
}

// The set of arguments for constructing a DataDiskAttachment resource.
type DataDiskAttachmentArgs struct {
	// Specifies the caching requirements for this Data Disk. Possible values include `None`, `ReadOnly` and `ReadWrite`.
	Caching interface{}
	// The Create Option of the Data Disk, such as `Empty` or `Attach`. Defaults to `Attach`. Changing this forces a new resource to be created.
	CreateOption interface{}
	// The Logical Unit Number of the Data Disk, which needs to be unique within the Virtual Machine. Changing this forces a new resource to be created.
	Lun interface{}
	// The ID of an existing Managed Disk which should be attached. Changing this forces a new resource to be created.
	ManagedDiskId interface{}
	// The ID of the Virtual Machine to which the Data Disk should be attached. Changing this forces a new resource to be created.
	VirtualMachineId interface{}
	// Specifies if Write Accelerator is enabled on the disk. This can only be enabled on `Premium_LRS` managed disks with no caching and [M-Series VMs](https://docs.microsoft.com/en-us/azure/virtual-machines/workloads/sap/how-to-enable-write-accelerator). Defaults to `false`.
	WriteAcceleratorEnabled interface{}
}
