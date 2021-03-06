// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ConfigurationStorePrimaryReadKey struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString *string `pulumi:"connectionString"`
	// The ID of the access key.
	Id *string `pulumi:"id"`
	// The secret of the access key.
	Secret *string `pulumi:"secret"`
}

type ConfigurationStorePrimaryReadKeyInput interface {
	pulumi.Input

	ToConfigurationStorePrimaryReadKeyOutput() ConfigurationStorePrimaryReadKeyOutput
	ToConfigurationStorePrimaryReadKeyOutputWithContext(context.Context) ConfigurationStorePrimaryReadKeyOutput
}

type ConfigurationStorePrimaryReadKeyArgs struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
	// The ID of the access key.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The secret of the access key.
	Secret pulumi.StringPtrInput `pulumi:"secret"`
}

func (ConfigurationStorePrimaryReadKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStorePrimaryReadKey)(nil)).Elem()
}

func (i ConfigurationStorePrimaryReadKeyArgs) ToConfigurationStorePrimaryReadKeyOutput() ConfigurationStorePrimaryReadKeyOutput {
	return i.ToConfigurationStorePrimaryReadKeyOutputWithContext(context.Background())
}

func (i ConfigurationStorePrimaryReadKeyArgs) ToConfigurationStorePrimaryReadKeyOutputWithContext(ctx context.Context) ConfigurationStorePrimaryReadKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStorePrimaryReadKeyOutput)
}

type ConfigurationStorePrimaryReadKeyArrayInput interface {
	pulumi.Input

	ToConfigurationStorePrimaryReadKeyArrayOutput() ConfigurationStorePrimaryReadKeyArrayOutput
	ToConfigurationStorePrimaryReadKeyArrayOutputWithContext(context.Context) ConfigurationStorePrimaryReadKeyArrayOutput
}

type ConfigurationStorePrimaryReadKeyArray []ConfigurationStorePrimaryReadKeyInput

func (ConfigurationStorePrimaryReadKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStorePrimaryReadKey)(nil)).Elem()
}

func (i ConfigurationStorePrimaryReadKeyArray) ToConfigurationStorePrimaryReadKeyArrayOutput() ConfigurationStorePrimaryReadKeyArrayOutput {
	return i.ToConfigurationStorePrimaryReadKeyArrayOutputWithContext(context.Background())
}

func (i ConfigurationStorePrimaryReadKeyArray) ToConfigurationStorePrimaryReadKeyArrayOutputWithContext(ctx context.Context) ConfigurationStorePrimaryReadKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStorePrimaryReadKeyArrayOutput)
}

type ConfigurationStorePrimaryReadKeyOutput struct { *pulumi.OutputState }

func (ConfigurationStorePrimaryReadKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStorePrimaryReadKey)(nil)).Elem()
}

func (o ConfigurationStorePrimaryReadKeyOutput) ToConfigurationStorePrimaryReadKeyOutput() ConfigurationStorePrimaryReadKeyOutput {
	return o
}

func (o ConfigurationStorePrimaryReadKeyOutput) ToConfigurationStorePrimaryReadKeyOutputWithContext(ctx context.Context) ConfigurationStorePrimaryReadKeyOutput {
	return o
}

// The connection string including the endpoint, id and secret.
func (o ConfigurationStorePrimaryReadKeyOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryReadKey) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The ID of the access key.
func (o ConfigurationStorePrimaryReadKeyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryReadKey) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The secret of the access key.
func (o ConfigurationStorePrimaryReadKeyOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryReadKey) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ConfigurationStorePrimaryReadKeyArrayOutput struct { *pulumi.OutputState}

func (ConfigurationStorePrimaryReadKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStorePrimaryReadKey)(nil)).Elem()
}

func (o ConfigurationStorePrimaryReadKeyArrayOutput) ToConfigurationStorePrimaryReadKeyArrayOutput() ConfigurationStorePrimaryReadKeyArrayOutput {
	return o
}

func (o ConfigurationStorePrimaryReadKeyArrayOutput) ToConfigurationStorePrimaryReadKeyArrayOutputWithContext(ctx context.Context) ConfigurationStorePrimaryReadKeyArrayOutput {
	return o
}

func (o ConfigurationStorePrimaryReadKeyArrayOutput) Index(i pulumi.IntInput) ConfigurationStorePrimaryReadKeyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ConfigurationStorePrimaryReadKey {
		return vs[0].([]ConfigurationStorePrimaryReadKey)[vs[1].(int)]
	}).(ConfigurationStorePrimaryReadKeyOutput)
}

type ConfigurationStorePrimaryWriteKey struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString *string `pulumi:"connectionString"`
	// The ID of the access key.
	Id *string `pulumi:"id"`
	// The secret of the access key.
	Secret *string `pulumi:"secret"`
}

type ConfigurationStorePrimaryWriteKeyInput interface {
	pulumi.Input

	ToConfigurationStorePrimaryWriteKeyOutput() ConfigurationStorePrimaryWriteKeyOutput
	ToConfigurationStorePrimaryWriteKeyOutputWithContext(context.Context) ConfigurationStorePrimaryWriteKeyOutput
}

type ConfigurationStorePrimaryWriteKeyArgs struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
	// The ID of the access key.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The secret of the access key.
	Secret pulumi.StringPtrInput `pulumi:"secret"`
}

func (ConfigurationStorePrimaryWriteKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStorePrimaryWriteKey)(nil)).Elem()
}

func (i ConfigurationStorePrimaryWriteKeyArgs) ToConfigurationStorePrimaryWriteKeyOutput() ConfigurationStorePrimaryWriteKeyOutput {
	return i.ToConfigurationStorePrimaryWriteKeyOutputWithContext(context.Background())
}

func (i ConfigurationStorePrimaryWriteKeyArgs) ToConfigurationStorePrimaryWriteKeyOutputWithContext(ctx context.Context) ConfigurationStorePrimaryWriteKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStorePrimaryWriteKeyOutput)
}

type ConfigurationStorePrimaryWriteKeyArrayInput interface {
	pulumi.Input

	ToConfigurationStorePrimaryWriteKeyArrayOutput() ConfigurationStorePrimaryWriteKeyArrayOutput
	ToConfigurationStorePrimaryWriteKeyArrayOutputWithContext(context.Context) ConfigurationStorePrimaryWriteKeyArrayOutput
}

type ConfigurationStorePrimaryWriteKeyArray []ConfigurationStorePrimaryWriteKeyInput

func (ConfigurationStorePrimaryWriteKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStorePrimaryWriteKey)(nil)).Elem()
}

func (i ConfigurationStorePrimaryWriteKeyArray) ToConfigurationStorePrimaryWriteKeyArrayOutput() ConfigurationStorePrimaryWriteKeyArrayOutput {
	return i.ToConfigurationStorePrimaryWriteKeyArrayOutputWithContext(context.Background())
}

func (i ConfigurationStorePrimaryWriteKeyArray) ToConfigurationStorePrimaryWriteKeyArrayOutputWithContext(ctx context.Context) ConfigurationStorePrimaryWriteKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStorePrimaryWriteKeyArrayOutput)
}

type ConfigurationStorePrimaryWriteKeyOutput struct { *pulumi.OutputState }

func (ConfigurationStorePrimaryWriteKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStorePrimaryWriteKey)(nil)).Elem()
}

func (o ConfigurationStorePrimaryWriteKeyOutput) ToConfigurationStorePrimaryWriteKeyOutput() ConfigurationStorePrimaryWriteKeyOutput {
	return o
}

func (o ConfigurationStorePrimaryWriteKeyOutput) ToConfigurationStorePrimaryWriteKeyOutputWithContext(ctx context.Context) ConfigurationStorePrimaryWriteKeyOutput {
	return o
}

// The connection string including the endpoint, id and secret.
func (o ConfigurationStorePrimaryWriteKeyOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryWriteKey) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The ID of the access key.
func (o ConfigurationStorePrimaryWriteKeyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryWriteKey) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The secret of the access key.
func (o ConfigurationStorePrimaryWriteKeyOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStorePrimaryWriteKey) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ConfigurationStorePrimaryWriteKeyArrayOutput struct { *pulumi.OutputState}

func (ConfigurationStorePrimaryWriteKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStorePrimaryWriteKey)(nil)).Elem()
}

func (o ConfigurationStorePrimaryWriteKeyArrayOutput) ToConfigurationStorePrimaryWriteKeyArrayOutput() ConfigurationStorePrimaryWriteKeyArrayOutput {
	return o
}

func (o ConfigurationStorePrimaryWriteKeyArrayOutput) ToConfigurationStorePrimaryWriteKeyArrayOutputWithContext(ctx context.Context) ConfigurationStorePrimaryWriteKeyArrayOutput {
	return o
}

func (o ConfigurationStorePrimaryWriteKeyArrayOutput) Index(i pulumi.IntInput) ConfigurationStorePrimaryWriteKeyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ConfigurationStorePrimaryWriteKey {
		return vs[0].([]ConfigurationStorePrimaryWriteKey)[vs[1].(int)]
	}).(ConfigurationStorePrimaryWriteKeyOutput)
}

type ConfigurationStoreSecondaryReadKey struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString *string `pulumi:"connectionString"`
	// The ID of the access key.
	Id *string `pulumi:"id"`
	// The secret of the access key.
	Secret *string `pulumi:"secret"`
}

type ConfigurationStoreSecondaryReadKeyInput interface {
	pulumi.Input

	ToConfigurationStoreSecondaryReadKeyOutput() ConfigurationStoreSecondaryReadKeyOutput
	ToConfigurationStoreSecondaryReadKeyOutputWithContext(context.Context) ConfigurationStoreSecondaryReadKeyOutput
}

type ConfigurationStoreSecondaryReadKeyArgs struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
	// The ID of the access key.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The secret of the access key.
	Secret pulumi.StringPtrInput `pulumi:"secret"`
}

func (ConfigurationStoreSecondaryReadKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStoreSecondaryReadKey)(nil)).Elem()
}

func (i ConfigurationStoreSecondaryReadKeyArgs) ToConfigurationStoreSecondaryReadKeyOutput() ConfigurationStoreSecondaryReadKeyOutput {
	return i.ToConfigurationStoreSecondaryReadKeyOutputWithContext(context.Background())
}

func (i ConfigurationStoreSecondaryReadKeyArgs) ToConfigurationStoreSecondaryReadKeyOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryReadKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreSecondaryReadKeyOutput)
}

type ConfigurationStoreSecondaryReadKeyArrayInput interface {
	pulumi.Input

	ToConfigurationStoreSecondaryReadKeyArrayOutput() ConfigurationStoreSecondaryReadKeyArrayOutput
	ToConfigurationStoreSecondaryReadKeyArrayOutputWithContext(context.Context) ConfigurationStoreSecondaryReadKeyArrayOutput
}

type ConfigurationStoreSecondaryReadKeyArray []ConfigurationStoreSecondaryReadKeyInput

func (ConfigurationStoreSecondaryReadKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStoreSecondaryReadKey)(nil)).Elem()
}

func (i ConfigurationStoreSecondaryReadKeyArray) ToConfigurationStoreSecondaryReadKeyArrayOutput() ConfigurationStoreSecondaryReadKeyArrayOutput {
	return i.ToConfigurationStoreSecondaryReadKeyArrayOutputWithContext(context.Background())
}

func (i ConfigurationStoreSecondaryReadKeyArray) ToConfigurationStoreSecondaryReadKeyArrayOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryReadKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreSecondaryReadKeyArrayOutput)
}

type ConfigurationStoreSecondaryReadKeyOutput struct { *pulumi.OutputState }

func (ConfigurationStoreSecondaryReadKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStoreSecondaryReadKey)(nil)).Elem()
}

func (o ConfigurationStoreSecondaryReadKeyOutput) ToConfigurationStoreSecondaryReadKeyOutput() ConfigurationStoreSecondaryReadKeyOutput {
	return o
}

func (o ConfigurationStoreSecondaryReadKeyOutput) ToConfigurationStoreSecondaryReadKeyOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryReadKeyOutput {
	return o
}

// The connection string including the endpoint, id and secret.
func (o ConfigurationStoreSecondaryReadKeyOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryReadKey) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The ID of the access key.
func (o ConfigurationStoreSecondaryReadKeyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryReadKey) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The secret of the access key.
func (o ConfigurationStoreSecondaryReadKeyOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryReadKey) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ConfigurationStoreSecondaryReadKeyArrayOutput struct { *pulumi.OutputState}

func (ConfigurationStoreSecondaryReadKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStoreSecondaryReadKey)(nil)).Elem()
}

func (o ConfigurationStoreSecondaryReadKeyArrayOutput) ToConfigurationStoreSecondaryReadKeyArrayOutput() ConfigurationStoreSecondaryReadKeyArrayOutput {
	return o
}

func (o ConfigurationStoreSecondaryReadKeyArrayOutput) ToConfigurationStoreSecondaryReadKeyArrayOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryReadKeyArrayOutput {
	return o
}

func (o ConfigurationStoreSecondaryReadKeyArrayOutput) Index(i pulumi.IntInput) ConfigurationStoreSecondaryReadKeyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ConfigurationStoreSecondaryReadKey {
		return vs[0].([]ConfigurationStoreSecondaryReadKey)[vs[1].(int)]
	}).(ConfigurationStoreSecondaryReadKeyOutput)
}

type ConfigurationStoreSecondaryWriteKey struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString *string `pulumi:"connectionString"`
	// The ID of the access key.
	Id *string `pulumi:"id"`
	// The secret of the access key.
	Secret *string `pulumi:"secret"`
}

type ConfigurationStoreSecondaryWriteKeyInput interface {
	pulumi.Input

	ToConfigurationStoreSecondaryWriteKeyOutput() ConfigurationStoreSecondaryWriteKeyOutput
	ToConfigurationStoreSecondaryWriteKeyOutputWithContext(context.Context) ConfigurationStoreSecondaryWriteKeyOutput
}

type ConfigurationStoreSecondaryWriteKeyArgs struct {
	// The connection string including the endpoint, id and secret.
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
	// The ID of the access key.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The secret of the access key.
	Secret pulumi.StringPtrInput `pulumi:"secret"`
}

func (ConfigurationStoreSecondaryWriteKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStoreSecondaryWriteKey)(nil)).Elem()
}

func (i ConfigurationStoreSecondaryWriteKeyArgs) ToConfigurationStoreSecondaryWriteKeyOutput() ConfigurationStoreSecondaryWriteKeyOutput {
	return i.ToConfigurationStoreSecondaryWriteKeyOutputWithContext(context.Background())
}

func (i ConfigurationStoreSecondaryWriteKeyArgs) ToConfigurationStoreSecondaryWriteKeyOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryWriteKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreSecondaryWriteKeyOutput)
}

type ConfigurationStoreSecondaryWriteKeyArrayInput interface {
	pulumi.Input

	ToConfigurationStoreSecondaryWriteKeyArrayOutput() ConfigurationStoreSecondaryWriteKeyArrayOutput
	ToConfigurationStoreSecondaryWriteKeyArrayOutputWithContext(context.Context) ConfigurationStoreSecondaryWriteKeyArrayOutput
}

type ConfigurationStoreSecondaryWriteKeyArray []ConfigurationStoreSecondaryWriteKeyInput

func (ConfigurationStoreSecondaryWriteKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStoreSecondaryWriteKey)(nil)).Elem()
}

func (i ConfigurationStoreSecondaryWriteKeyArray) ToConfigurationStoreSecondaryWriteKeyArrayOutput() ConfigurationStoreSecondaryWriteKeyArrayOutput {
	return i.ToConfigurationStoreSecondaryWriteKeyArrayOutputWithContext(context.Background())
}

func (i ConfigurationStoreSecondaryWriteKeyArray) ToConfigurationStoreSecondaryWriteKeyArrayOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryWriteKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreSecondaryWriteKeyArrayOutput)
}

type ConfigurationStoreSecondaryWriteKeyOutput struct { *pulumi.OutputState }

func (ConfigurationStoreSecondaryWriteKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationStoreSecondaryWriteKey)(nil)).Elem()
}

func (o ConfigurationStoreSecondaryWriteKeyOutput) ToConfigurationStoreSecondaryWriteKeyOutput() ConfigurationStoreSecondaryWriteKeyOutput {
	return o
}

func (o ConfigurationStoreSecondaryWriteKeyOutput) ToConfigurationStoreSecondaryWriteKeyOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryWriteKeyOutput {
	return o
}

// The connection string including the endpoint, id and secret.
func (o ConfigurationStoreSecondaryWriteKeyOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryWriteKey) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The ID of the access key.
func (o ConfigurationStoreSecondaryWriteKeyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryWriteKey) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The secret of the access key.
func (o ConfigurationStoreSecondaryWriteKeyOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ConfigurationStoreSecondaryWriteKey) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ConfigurationStoreSecondaryWriteKeyArrayOutput struct { *pulumi.OutputState}

func (ConfigurationStoreSecondaryWriteKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationStoreSecondaryWriteKey)(nil)).Elem()
}

func (o ConfigurationStoreSecondaryWriteKeyArrayOutput) ToConfigurationStoreSecondaryWriteKeyArrayOutput() ConfigurationStoreSecondaryWriteKeyArrayOutput {
	return o
}

func (o ConfigurationStoreSecondaryWriteKeyArrayOutput) ToConfigurationStoreSecondaryWriteKeyArrayOutputWithContext(ctx context.Context) ConfigurationStoreSecondaryWriteKeyArrayOutput {
	return o
}

func (o ConfigurationStoreSecondaryWriteKeyArrayOutput) Index(i pulumi.IntInput) ConfigurationStoreSecondaryWriteKeyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ConfigurationStoreSecondaryWriteKey {
		return vs[0].([]ConfigurationStoreSecondaryWriteKey)[vs[1].(int)]
	}).(ConfigurationStoreSecondaryWriteKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationStorePrimaryReadKeyOutput{})
	pulumi.RegisterOutputType(ConfigurationStorePrimaryReadKeyArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationStorePrimaryWriteKeyOutput{})
	pulumi.RegisterOutputType(ConfigurationStorePrimaryWriteKeyArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationStoreSecondaryReadKeyOutput{})
	pulumi.RegisterOutputType(ConfigurationStoreSecondaryReadKeyArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationStoreSecondaryWriteKeyOutput{})
	pulumi.RegisterOutputType(ConfigurationStoreSecondaryWriteKeyArrayOutput{})
}
