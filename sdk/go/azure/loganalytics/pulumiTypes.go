// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package loganalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type LinkedServiceLinkedServiceProperties struct {
	// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
	ResourceId string `pulumi:"resourceId"`
}

type LinkedServiceLinkedServicePropertiesInput interface {
	pulumi.Input

	ToLinkedServiceLinkedServicePropertiesOutput() LinkedServiceLinkedServicePropertiesOutput
	ToLinkedServiceLinkedServicePropertiesOutputWithContext(context.Context) LinkedServiceLinkedServicePropertiesOutput
}

type LinkedServiceLinkedServicePropertiesArgs struct {
	// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
}

func (LinkedServiceLinkedServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkedServiceProperties)(nil)).Elem()
}

func (i LinkedServiceLinkedServicePropertiesArgs) ToLinkedServiceLinkedServicePropertiesOutput() LinkedServiceLinkedServicePropertiesOutput {
	return i.ToLinkedServiceLinkedServicePropertiesOutputWithContext(context.Background())
}

func (i LinkedServiceLinkedServicePropertiesArgs) ToLinkedServiceLinkedServicePropertiesOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceLinkedServicePropertiesOutput)
}

func (i LinkedServiceLinkedServicePropertiesArgs) ToLinkedServiceLinkedServicePropertiesPtrOutput() LinkedServiceLinkedServicePropertiesPtrOutput {
	return i.ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(context.Background())
}

func (i LinkedServiceLinkedServicePropertiesArgs) ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceLinkedServicePropertiesOutput).ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(ctx)
}

type LinkedServiceLinkedServicePropertiesPtrInput interface {
	pulumi.Input

	ToLinkedServiceLinkedServicePropertiesPtrOutput() LinkedServiceLinkedServicePropertiesPtrOutput
	ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(context.Context) LinkedServiceLinkedServicePropertiesPtrOutput
}

type linkedServiceLinkedServicePropertiesPtrType LinkedServiceLinkedServicePropertiesArgs

func LinkedServiceLinkedServicePropertiesPtr(v *LinkedServiceLinkedServicePropertiesArgs) LinkedServiceLinkedServicePropertiesPtrInput {	return (*linkedServiceLinkedServicePropertiesPtrType)(v)
}

func (*linkedServiceLinkedServicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceLinkedServiceProperties)(nil)).Elem()
}

func (i *linkedServiceLinkedServicePropertiesPtrType) ToLinkedServiceLinkedServicePropertiesPtrOutput() LinkedServiceLinkedServicePropertiesPtrOutput {
	return i.ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(context.Background())
}

func (i *linkedServiceLinkedServicePropertiesPtrType) ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceLinkedServicePropertiesPtrOutput)
}

type LinkedServiceLinkedServicePropertiesOutput struct { *pulumi.OutputState }

func (LinkedServiceLinkedServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceLinkedServiceProperties)(nil)).Elem()
}

func (o LinkedServiceLinkedServicePropertiesOutput) ToLinkedServiceLinkedServicePropertiesOutput() LinkedServiceLinkedServicePropertiesOutput {
	return o
}

func (o LinkedServiceLinkedServicePropertiesOutput) ToLinkedServiceLinkedServicePropertiesOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesOutput {
	return o
}

func (o LinkedServiceLinkedServicePropertiesOutput) ToLinkedServiceLinkedServicePropertiesPtrOutput() LinkedServiceLinkedServicePropertiesPtrOutput {
	return o.ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(context.Background())
}

func (o LinkedServiceLinkedServicePropertiesOutput) ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesPtrOutput {
	return o.ApplyT(func(v LinkedServiceLinkedServiceProperties) *LinkedServiceLinkedServiceProperties {
		return &v
	}).(LinkedServiceLinkedServicePropertiesPtrOutput)
}
// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
func (o LinkedServiceLinkedServicePropertiesOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func (v LinkedServiceLinkedServiceProperties) string { return v.ResourceId }).(pulumi.StringOutput)
}

type LinkedServiceLinkedServicePropertiesPtrOutput struct { *pulumi.OutputState}

func (LinkedServiceLinkedServicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceLinkedServiceProperties)(nil)).Elem()
}

func (o LinkedServiceLinkedServicePropertiesPtrOutput) ToLinkedServiceLinkedServicePropertiesPtrOutput() LinkedServiceLinkedServicePropertiesPtrOutput {
	return o
}

func (o LinkedServiceLinkedServicePropertiesPtrOutput) ToLinkedServiceLinkedServicePropertiesPtrOutputWithContext(ctx context.Context) LinkedServiceLinkedServicePropertiesPtrOutput {
	return o
}

func (o LinkedServiceLinkedServicePropertiesPtrOutput) Elem() LinkedServiceLinkedServicePropertiesOutput {
	return o.ApplyT(func (v *LinkedServiceLinkedServiceProperties) LinkedServiceLinkedServiceProperties { return *v }).(LinkedServiceLinkedServicePropertiesOutput)
}

// The resource id of the resource that will be linked to the workspace. This field has been deprecated in favour of the top-level `resourceId` field and will be removed in v2.0 of the AzureRM Provider.
func (o LinkedServiceLinkedServicePropertiesPtrOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func (v LinkedServiceLinkedServiceProperties) string { return v.ResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceLinkedServicePropertiesOutput{})
	pulumi.RegisterOutputType(LinkedServiceLinkedServicePropertiesPtrOutput{})
}
