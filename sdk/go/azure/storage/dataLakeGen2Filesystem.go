// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Data Lake Gen2 File System within an Azure Storage Account.
// 
// > **NOTE:** This Resource requires using Azure Active Directory to connect to Azure Storage, which in turn requires the `Storage` specific roles - which are not granted by default.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_data_lake_gen2_filesystem.html.markdown.
type DataLakeGen2Filesystem struct {
	pulumi.CustomResourceState

	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
}

// NewDataLakeGen2Filesystem registers a new resource with the given unique name, arguments, and options.
func NewDataLakeGen2Filesystem(ctx *pulumi.Context,
	name string, args *DataLakeGen2FilesystemArgs, opts ...pulumi.ResourceOption) (*DataLakeGen2Filesystem, error) {
	if args == nil || args.StorageAccountId == nil {
		return nil, errors.New("missing required argument 'StorageAccountId'")
	}
	if args == nil {
		args = &DataLakeGen2FilesystemArgs{}
	}
	var resource DataLakeGen2Filesystem
	err := ctx.RegisterResource("azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLakeGen2Filesystem gets an existing DataLakeGen2Filesystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLakeGen2Filesystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLakeGen2FilesystemState, opts ...pulumi.ResourceOption) (*DataLakeGen2Filesystem, error) {
	var resource DataLakeGen2Filesystem
	err := ctx.ReadResource("azure:storage/dataLakeGen2Filesystem:DataLakeGen2Filesystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLakeGen2Filesystem resources.
type dataLakeGen2FilesystemState struct {
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties map[string]string `pulumi:"properties"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
}

type DataLakeGen2FilesystemState struct {
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties pulumi.StringMapInput
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringPtrInput
}

func (DataLakeGen2FilesystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeGen2FilesystemState)(nil)).Elem()
}

type dataLakeGen2FilesystemArgs struct {
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties map[string]string `pulumi:"properties"`
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// The set of arguments for constructing a DataLakeGen2Filesystem resource.
type DataLakeGen2FilesystemArgs struct {
	// The name of the Data Lake Gen2 File System which should be created within the Storage Account. Must be unique within the storage account the queue is located. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A mapping of Key to Base64-Encoded Values which should be assigned to this Data Lake Gen2 File System.
	Properties pulumi.StringMapInput
	// Specifies the ID of the Storage Account in which the Data Lake Gen2 File System should exist. Changing this forces a new resource to be created.
	StorageAccountId pulumi.StringInput
}

func (DataLakeGen2FilesystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeGen2FilesystemArgs)(nil)).Elem()
}

