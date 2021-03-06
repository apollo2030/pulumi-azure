// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package relay

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type NamespaceSku struct {
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
}

type NamespaceSkuInput interface {
	pulumi.Input

	ToNamespaceSkuOutput() NamespaceSkuOutput
	ToNamespaceSkuOutputWithContext(context.Context) NamespaceSkuOutput
}

type NamespaceSkuArgs struct {
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
}

func (NamespaceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceSku)(nil)).Elem()
}

func (i NamespaceSkuArgs) ToNamespaceSkuOutput() NamespaceSkuOutput {
	return i.ToNamespaceSkuOutputWithContext(context.Background())
}

func (i NamespaceSkuArgs) ToNamespaceSkuOutputWithContext(ctx context.Context) NamespaceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceSkuOutput)
}

func (i NamespaceSkuArgs) ToNamespaceSkuPtrOutput() NamespaceSkuPtrOutput {
	return i.ToNamespaceSkuPtrOutputWithContext(context.Background())
}

func (i NamespaceSkuArgs) ToNamespaceSkuPtrOutputWithContext(ctx context.Context) NamespaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceSkuOutput).ToNamespaceSkuPtrOutputWithContext(ctx)
}

type NamespaceSkuPtrInput interface {
	pulumi.Input

	ToNamespaceSkuPtrOutput() NamespaceSkuPtrOutput
	ToNamespaceSkuPtrOutputWithContext(context.Context) NamespaceSkuPtrOutput
}

type namespaceSkuPtrType NamespaceSkuArgs

func NamespaceSkuPtr(v *NamespaceSkuArgs) NamespaceSkuPtrInput {	return (*namespaceSkuPtrType)(v)
}

func (*namespaceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceSku)(nil)).Elem()
}

func (i *namespaceSkuPtrType) ToNamespaceSkuPtrOutput() NamespaceSkuPtrOutput {
	return i.ToNamespaceSkuPtrOutputWithContext(context.Background())
}

func (i *namespaceSkuPtrType) ToNamespaceSkuPtrOutputWithContext(ctx context.Context) NamespaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceSkuPtrOutput)
}

type NamespaceSkuOutput struct { *pulumi.OutputState }

func (NamespaceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceSku)(nil)).Elem()
}

func (o NamespaceSkuOutput) ToNamespaceSkuOutput() NamespaceSkuOutput {
	return o
}

func (o NamespaceSkuOutput) ToNamespaceSkuOutputWithContext(ctx context.Context) NamespaceSkuOutput {
	return o
}

func (o NamespaceSkuOutput) ToNamespaceSkuPtrOutput() NamespaceSkuPtrOutput {
	return o.ToNamespaceSkuPtrOutputWithContext(context.Background())
}

func (o NamespaceSkuOutput) ToNamespaceSkuPtrOutputWithContext(ctx context.Context) NamespaceSkuPtrOutput {
	return o.ApplyT(func(v NamespaceSku) *NamespaceSku {
		return &v
	}).(NamespaceSkuPtrOutput)
}
// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
func (o NamespaceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v NamespaceSku) string { return v.Name }).(pulumi.StringOutput)
}

type NamespaceSkuPtrOutput struct { *pulumi.OutputState}

func (NamespaceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceSku)(nil)).Elem()
}

func (o NamespaceSkuPtrOutput) ToNamespaceSkuPtrOutput() NamespaceSkuPtrOutput {
	return o
}

func (o NamespaceSkuPtrOutput) ToNamespaceSkuPtrOutputWithContext(ctx context.Context) NamespaceSkuPtrOutput {
	return o
}

func (o NamespaceSkuPtrOutput) Elem() NamespaceSkuOutput {
	return o.ApplyT(func (v *NamespaceSku) NamespaceSku { return *v }).(NamespaceSkuOutput)
}

// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
func (o NamespaceSkuPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v NamespaceSku) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceSkuOutput{})
	pulumi.RegisterOutputType(NamespaceSkuPtrOutput{})
}
