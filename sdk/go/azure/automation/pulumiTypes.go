// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type AccountSku struct {
	// Specifies the name of the Automation Account. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
}

type AccountSkuInput interface {
	pulumi.Input

	ToAccountSkuOutput() AccountSkuOutput
	ToAccountSkuOutputWithContext(context.Context) AccountSkuOutput
}

type AccountSkuArgs struct {
	// Specifies the name of the Automation Account. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (AccountSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountSku)(nil)).Elem()
}

func (i AccountSkuArgs) ToAccountSkuOutput() AccountSkuOutput {
	return i.ToAccountSkuOutputWithContext(context.Background())
}

func (i AccountSkuArgs) ToAccountSkuOutputWithContext(ctx context.Context) AccountSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuOutput)
}

func (i AccountSkuArgs) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return i.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (i AccountSkuArgs) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuOutput).ToAccountSkuPtrOutputWithContext(ctx)
}

type AccountSkuPtrInput interface {
	pulumi.Input

	ToAccountSkuPtrOutput() AccountSkuPtrOutput
	ToAccountSkuPtrOutputWithContext(context.Context) AccountSkuPtrOutput
}

type accountSkuPtrType AccountSkuArgs

func AccountSkuPtr(v *AccountSkuArgs) AccountSkuPtrInput {	return (*accountSkuPtrType)(v)
}

func (*accountSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountSku)(nil)).Elem()
}

func (i *accountSkuPtrType) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return i.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (i *accountSkuPtrType) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountSkuPtrOutput)
}

type AccountSkuOutput struct { *pulumi.OutputState }

func (AccountSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountSku)(nil)).Elem()
}

func (o AccountSkuOutput) ToAccountSkuOutput() AccountSkuOutput {
	return o
}

func (o AccountSkuOutput) ToAccountSkuOutputWithContext(ctx context.Context) AccountSkuOutput {
	return o
}

func (o AccountSkuOutput) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return o.ToAccountSkuPtrOutputWithContext(context.Background())
}

func (o AccountSkuOutput) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return o.ApplyT(func(v AccountSku) *AccountSku {
		return &v
	}).(AccountSkuPtrOutput)
}
// Specifies the name of the Automation Account. Changing this forces a new resource to be created.
func (o AccountSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AccountSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AccountSkuPtrOutput struct { *pulumi.OutputState}

func (AccountSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccountSku)(nil)).Elem()
}

func (o AccountSkuPtrOutput) ToAccountSkuPtrOutput() AccountSkuPtrOutput {
	return o
}

func (o AccountSkuPtrOutput) ToAccountSkuPtrOutputWithContext(ctx context.Context) AccountSkuPtrOutput {
	return o
}

func (o AccountSkuPtrOutput) Elem() AccountSkuOutput {
	return o.ApplyT(func (v *AccountSku) AccountSku { return *v }).(AccountSkuOutput)
}

// Specifies the name of the Automation Account. Changing this forces a new resource to be created.
func (o AccountSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v AccountSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ModuleModuleLink struct {
	Hash *ModuleModuleLinkHash `pulumi:"hash"`
	// The uri of the module content (zip or nupkg).
	Uri string `pulumi:"uri"`
}

type ModuleModuleLinkInput interface {
	pulumi.Input

	ToModuleModuleLinkOutput() ModuleModuleLinkOutput
	ToModuleModuleLinkOutputWithContext(context.Context) ModuleModuleLinkOutput
}

type ModuleModuleLinkArgs struct {
	Hash ModuleModuleLinkHashPtrInput `pulumi:"hash"`
	// The uri of the module content (zip or nupkg).
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (ModuleModuleLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLink)(nil)).Elem()
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkOutput() ModuleModuleLinkOutput {
	return i.ToModuleModuleLinkOutputWithContext(context.Background())
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkOutputWithContext(ctx context.Context) ModuleModuleLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkOutput)
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return i.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (i ModuleModuleLinkArgs) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkOutput).ToModuleModuleLinkPtrOutputWithContext(ctx)
}

type ModuleModuleLinkPtrInput interface {
	pulumi.Input

	ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput
	ToModuleModuleLinkPtrOutputWithContext(context.Context) ModuleModuleLinkPtrOutput
}

type moduleModuleLinkPtrType ModuleModuleLinkArgs

func ModuleModuleLinkPtr(v *ModuleModuleLinkArgs) ModuleModuleLinkPtrInput {	return (*moduleModuleLinkPtrType)(v)
}

func (*moduleModuleLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLink)(nil)).Elem()
}

func (i *moduleModuleLinkPtrType) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return i.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (i *moduleModuleLinkPtrType) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkPtrOutput)
}

type ModuleModuleLinkOutput struct { *pulumi.OutputState }

func (ModuleModuleLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLink)(nil)).Elem()
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkOutput() ModuleModuleLinkOutput {
	return o
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkOutputWithContext(ctx context.Context) ModuleModuleLinkOutput {
	return o
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return o.ToModuleModuleLinkPtrOutputWithContext(context.Background())
}

func (o ModuleModuleLinkOutput) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return o.ApplyT(func(v ModuleModuleLink) *ModuleModuleLink {
		return &v
	}).(ModuleModuleLinkPtrOutput)
}
func (o ModuleModuleLinkOutput) Hash() ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func (v ModuleModuleLink) *ModuleModuleLinkHash { return v.Hash }).(ModuleModuleLinkHashPtrOutput)
}

// The uri of the module content (zip or nupkg).
func (o ModuleModuleLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLink) string { return v.Uri }).(pulumi.StringOutput)
}

type ModuleModuleLinkPtrOutput struct { *pulumi.OutputState}

func (ModuleModuleLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLink)(nil)).Elem()
}

func (o ModuleModuleLinkPtrOutput) ToModuleModuleLinkPtrOutput() ModuleModuleLinkPtrOutput {
	return o
}

func (o ModuleModuleLinkPtrOutput) ToModuleModuleLinkPtrOutputWithContext(ctx context.Context) ModuleModuleLinkPtrOutput {
	return o
}

func (o ModuleModuleLinkPtrOutput) Elem() ModuleModuleLinkOutput {
	return o.ApplyT(func (v *ModuleModuleLink) ModuleModuleLink { return *v }).(ModuleModuleLinkOutput)
}

func (o ModuleModuleLinkPtrOutput) Hash() ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func (v ModuleModuleLink) *ModuleModuleLinkHash { return v.Hash }).(ModuleModuleLinkHashPtrOutput)
}

// The uri of the module content (zip or nupkg).
func (o ModuleModuleLinkPtrOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLink) string { return v.Uri }).(pulumi.StringOutput)
}

type ModuleModuleLinkHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value string `pulumi:"value"`
}

type ModuleModuleLinkHashInput interface {
	pulumi.Input

	ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput
	ToModuleModuleLinkHashOutputWithContext(context.Context) ModuleModuleLinkHashOutput
}

type ModuleModuleLinkHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ModuleModuleLinkHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLinkHash)(nil)).Elem()
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput {
	return i.ToModuleModuleLinkHashOutputWithContext(context.Background())
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashOutputWithContext(ctx context.Context) ModuleModuleLinkHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashOutput)
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return i.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (i ModuleModuleLinkHashArgs) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashOutput).ToModuleModuleLinkHashPtrOutputWithContext(ctx)
}

type ModuleModuleLinkHashPtrInput interface {
	pulumi.Input

	ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput
	ToModuleModuleLinkHashPtrOutputWithContext(context.Context) ModuleModuleLinkHashPtrOutput
}

type moduleModuleLinkHashPtrType ModuleModuleLinkHashArgs

func ModuleModuleLinkHashPtr(v *ModuleModuleLinkHashArgs) ModuleModuleLinkHashPtrInput {	return (*moduleModuleLinkHashPtrType)(v)
}

func (*moduleModuleLinkHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLinkHash)(nil)).Elem()
}

func (i *moduleModuleLinkHashPtrType) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return i.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (i *moduleModuleLinkHashPtrType) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleModuleLinkHashPtrOutput)
}

type ModuleModuleLinkHashOutput struct { *pulumi.OutputState }

func (ModuleModuleLinkHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleModuleLinkHash)(nil)).Elem()
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashOutput() ModuleModuleLinkHashOutput {
	return o
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashOutputWithContext(ctx context.Context) ModuleModuleLinkHashOutput {
	return o
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return o.ToModuleModuleLinkHashPtrOutputWithContext(context.Background())
}

func (o ModuleModuleLinkHashOutput) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return o.ApplyT(func(v ModuleModuleLinkHash) *ModuleModuleLinkHash {
		return &v
	}).(ModuleModuleLinkHashPtrOutput)
}
func (o ModuleModuleLinkHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ModuleModuleLinkHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type ModuleModuleLinkHashPtrOutput struct { *pulumi.OutputState}

func (ModuleModuleLinkHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleModuleLinkHash)(nil)).Elem()
}

func (o ModuleModuleLinkHashPtrOutput) ToModuleModuleLinkHashPtrOutput() ModuleModuleLinkHashPtrOutput {
	return o
}

func (o ModuleModuleLinkHashPtrOutput) ToModuleModuleLinkHashPtrOutputWithContext(ctx context.Context) ModuleModuleLinkHashPtrOutput {
	return o
}

func (o ModuleModuleLinkHashPtrOutput) Elem() ModuleModuleLinkHashOutput {
	return o.ApplyT(func (v *ModuleModuleLinkHash) ModuleModuleLinkHash { return *v }).(ModuleModuleLinkHashOutput)
}

func (o ModuleModuleLinkHashPtrOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ModuleModuleLinkHashPtrOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v ModuleModuleLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type RunBookPublishContentLink struct {
	Hash *RunBookPublishContentLinkHash `pulumi:"hash"`
	// The uri of the runbook content.
	Uri string `pulumi:"uri"`
	Version *string `pulumi:"version"`
}

type RunBookPublishContentLinkInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput
	ToRunBookPublishContentLinkOutputWithContext(context.Context) RunBookPublishContentLinkOutput
}

type RunBookPublishContentLinkArgs struct {
	Hash RunBookPublishContentLinkHashPtrInput `pulumi:"hash"`
	// The uri of the runbook content.
	Uri pulumi.StringInput `pulumi:"uri"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (RunBookPublishContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLink)(nil)).Elem()
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput {
	return i.ToRunBookPublishContentLinkOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkOutputWithContext(ctx context.Context) RunBookPublishContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkOutput)
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return i.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkArgs) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkOutput).ToRunBookPublishContentLinkPtrOutputWithContext(ctx)
}

type RunBookPublishContentLinkPtrInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput
	ToRunBookPublishContentLinkPtrOutputWithContext(context.Context) RunBookPublishContentLinkPtrOutput
}

type runBookPublishContentLinkPtrType RunBookPublishContentLinkArgs

func RunBookPublishContentLinkPtr(v *RunBookPublishContentLinkArgs) RunBookPublishContentLinkPtrInput {	return (*runBookPublishContentLinkPtrType)(v)
}

func (*runBookPublishContentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLink)(nil)).Elem()
}

func (i *runBookPublishContentLinkPtrType) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return i.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (i *runBookPublishContentLinkPtrType) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkPtrOutput)
}

type RunBookPublishContentLinkOutput struct { *pulumi.OutputState }

func (RunBookPublishContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLink)(nil)).Elem()
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkOutput() RunBookPublishContentLinkOutput {
	return o
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkOutputWithContext(ctx context.Context) RunBookPublishContentLinkOutput {
	return o
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return o.ToRunBookPublishContentLinkPtrOutputWithContext(context.Background())
}

func (o RunBookPublishContentLinkOutput) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLink) *RunBookPublishContentLink {
		return &v
	}).(RunBookPublishContentLinkPtrOutput)
}
func (o RunBookPublishContentLinkOutput) Hash() RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) *RunBookPublishContentLinkHash { return v.Hash }).(RunBookPublishContentLinkHashPtrOutput)
}

// The uri of the runbook content.
func (o RunBookPublishContentLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) string { return v.Uri }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type RunBookPublishContentLinkPtrOutput struct { *pulumi.OutputState}

func (RunBookPublishContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLink)(nil)).Elem()
}

func (o RunBookPublishContentLinkPtrOutput) ToRunBookPublishContentLinkPtrOutput() RunBookPublishContentLinkPtrOutput {
	return o
}

func (o RunBookPublishContentLinkPtrOutput) ToRunBookPublishContentLinkPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkPtrOutput {
	return o
}

func (o RunBookPublishContentLinkPtrOutput) Elem() RunBookPublishContentLinkOutput {
	return o.ApplyT(func (v *RunBookPublishContentLink) RunBookPublishContentLink { return *v }).(RunBookPublishContentLinkOutput)
}

func (o RunBookPublishContentLinkPtrOutput) Hash() RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) *RunBookPublishContentLinkHash { return v.Hash }).(RunBookPublishContentLinkHashPtrOutput)
}

// The uri of the runbook content.
func (o RunBookPublishContentLinkPtrOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) string { return v.Uri }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func (v RunBookPublishContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type RunBookPublishContentLinkHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value string `pulumi:"value"`
}

type RunBookPublishContentLinkHashInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput
	ToRunBookPublishContentLinkHashOutputWithContext(context.Context) RunBookPublishContentLinkHashOutput
}

type RunBookPublishContentLinkHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (RunBookPublishContentLinkHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLinkHash)(nil)).Elem()
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput {
	return i.ToRunBookPublishContentLinkHashOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashOutput)
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return i.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (i RunBookPublishContentLinkHashArgs) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashOutput).ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx)
}

type RunBookPublishContentLinkHashPtrInput interface {
	pulumi.Input

	ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput
	ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Context) RunBookPublishContentLinkHashPtrOutput
}

type runBookPublishContentLinkHashPtrType RunBookPublishContentLinkHashArgs

func RunBookPublishContentLinkHashPtr(v *RunBookPublishContentLinkHashArgs) RunBookPublishContentLinkHashPtrInput {	return (*runBookPublishContentLinkHashPtrType)(v)
}

func (*runBookPublishContentLinkHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLinkHash)(nil)).Elem()
}

func (i *runBookPublishContentLinkHashPtrType) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return i.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (i *runBookPublishContentLinkHashPtrType) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunBookPublishContentLinkHashPtrOutput)
}

type RunBookPublishContentLinkHashOutput struct { *pulumi.OutputState }

func (RunBookPublishContentLinkHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunBookPublishContentLinkHash)(nil)).Elem()
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashOutput() RunBookPublishContentLinkHashOutput {
	return o
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashOutput {
	return o
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return o.ToRunBookPublishContentLinkHashPtrOutputWithContext(context.Background())
}

func (o RunBookPublishContentLinkHashOutput) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return o.ApplyT(func(v RunBookPublishContentLinkHash) *RunBookPublishContentLinkHash {
		return &v
	}).(RunBookPublishContentLinkHashPtrOutput)
}
func (o RunBookPublishContentLinkHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type RunBookPublishContentLinkHashPtrOutput struct { *pulumi.OutputState}

func (RunBookPublishContentLinkHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunBookPublishContentLinkHash)(nil)).Elem()
}

func (o RunBookPublishContentLinkHashPtrOutput) ToRunBookPublishContentLinkHashPtrOutput() RunBookPublishContentLinkHashPtrOutput {
	return o
}

func (o RunBookPublishContentLinkHashPtrOutput) ToRunBookPublishContentLinkHashPtrOutputWithContext(ctx context.Context) RunBookPublishContentLinkHashPtrOutput {
	return o
}

func (o RunBookPublishContentLinkHashPtrOutput) Elem() RunBookPublishContentLinkHashOutput {
	return o.ApplyT(func (v *RunBookPublishContentLinkHash) RunBookPublishContentLinkHash { return *v }).(RunBookPublishContentLinkHashOutput)
}

func (o RunBookPublishContentLinkHashPtrOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLinkHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o RunBookPublishContentLinkHashPtrOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v RunBookPublishContentLinkHash) string { return v.Value }).(pulumi.StringOutput)
}

type ScheduleMonthlyOccurrence struct {
	Day string `pulumi:"day"`
	Occurrence int `pulumi:"occurrence"`
}

type ScheduleMonthlyOccurrenceInput interface {
	pulumi.Input

	ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput
	ToScheduleMonthlyOccurrenceOutputWithContext(context.Context) ScheduleMonthlyOccurrenceOutput
}

type ScheduleMonthlyOccurrenceArgs struct {
	Day pulumi.StringInput `pulumi:"day"`
	Occurrence pulumi.IntInput `pulumi:"occurrence"`
}

func (ScheduleMonthlyOccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i ScheduleMonthlyOccurrenceArgs) ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput {
	return i.ToScheduleMonthlyOccurrenceOutputWithContext(context.Background())
}

func (i ScheduleMonthlyOccurrenceArgs) ToScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMonthlyOccurrenceOutput)
}

type ScheduleMonthlyOccurrenceArrayInput interface {
	pulumi.Input

	ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput
	ToScheduleMonthlyOccurrenceArrayOutputWithContext(context.Context) ScheduleMonthlyOccurrenceArrayOutput
}

type ScheduleMonthlyOccurrenceArray []ScheduleMonthlyOccurrenceInput

func (ScheduleMonthlyOccurrenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i ScheduleMonthlyOccurrenceArray) ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput {
	return i.ToScheduleMonthlyOccurrenceArrayOutputWithContext(context.Background())
}

func (i ScheduleMonthlyOccurrenceArray) ToScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMonthlyOccurrenceArrayOutput)
}

type ScheduleMonthlyOccurrenceOutput struct { *pulumi.OutputState }

func (ScheduleMonthlyOccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o ScheduleMonthlyOccurrenceOutput) ToScheduleMonthlyOccurrenceOutput() ScheduleMonthlyOccurrenceOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceOutput) ToScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceOutput) Day() pulumi.StringOutput {
	return o.ApplyT(func (v ScheduleMonthlyOccurrence) string { return v.Day }).(pulumi.StringOutput)
}

func (o ScheduleMonthlyOccurrenceOutput) Occurrence() pulumi.IntOutput {
	return o.ApplyT(func (v ScheduleMonthlyOccurrence) int { return v.Occurrence }).(pulumi.IntOutput)
}

type ScheduleMonthlyOccurrenceArrayOutput struct { *pulumi.OutputState}

func (ScheduleMonthlyOccurrenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o ScheduleMonthlyOccurrenceArrayOutput) ToScheduleMonthlyOccurrenceArrayOutput() ScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceArrayOutput) ToScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) ScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o ScheduleMonthlyOccurrenceArrayOutput) Index(i pulumi.IntInput) ScheduleMonthlyOccurrenceOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ScheduleMonthlyOccurrence {
		return vs[0].([]ScheduleMonthlyOccurrence)[vs[1].(int)]
	}).(ScheduleMonthlyOccurrenceOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountSkuOutput{})
	pulumi.RegisterOutputType(AccountSkuPtrOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkPtrOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkHashOutput{})
	pulumi.RegisterOutputType(ModuleModuleLinkHashPtrOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkPtrOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkHashOutput{})
	pulumi.RegisterOutputType(RunBookPublishContentLinkHashPtrOutput{})
	pulumi.RegisterOutputType(ScheduleMonthlyOccurrenceOutput{})
	pulumi.RegisterOutputType(ScheduleMonthlyOccurrenceArrayOutput{})
}
