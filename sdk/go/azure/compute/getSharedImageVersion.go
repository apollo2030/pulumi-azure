// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
// 
// > **NOTE** Shared Image Galleries are currently in Public Preview. You can find more information, including [how to register for the Public Preview here](https://azure.microsoft.com/en-gb/blog/announcing-the-public-preview-of-shared-image-gallery/).
func LookupSharedImageVersion(ctx *pulumi.Context, args *GetSharedImageVersionArgs) (*GetSharedImageVersionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["galleryName"] = args.GalleryName
		inputs["imageName"] = args.ImageName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
	}
	outputs, err := ctx.Invoke("azure:compute/getSharedImageVersion:getSharedImageVersion", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSharedImageVersionResult{
		ExcludeFromLatest: outputs["excludeFromLatest"],
		Location: outputs["location"],
		ManagedImageId: outputs["managedImageId"],
		Tags: outputs["tags"],
		TargetRegions: outputs["targetRegions"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSharedImageVersion.
type GetSharedImageVersionArgs struct {
	// The name of the Shared Image in which the Shared Image exists.
	GalleryName interface{}
	// The name of the Shared Image in which this Version exists.
	ImageName interface{}
	// The name of the Image Version.
	Name interface{}
	// The name of the Resource Group in which the Shared Image Gallery exists.
	ResourceGroupName interface{}
}

// A collection of values returned by getSharedImageVersion.
type GetSharedImageVersionResult struct {
	// Is this Image Version excluded from the `latest` filter?
	ExcludeFromLatest interface{}
	// The supported Azure location where the Shared Image Gallery exists.
	Location interface{}
	// The ID of the Managed Image which was the source of this Shared Image Version.
	ManagedImageId interface{}
	// A mapping of tags assigned to the Shared Image.
	Tags interface{}
	// One or more `target_region` blocks as documented below.
	TargetRegions interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
