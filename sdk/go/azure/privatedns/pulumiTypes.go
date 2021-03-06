// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package privatedns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type LinkEndpointPrivateServiceConnection struct {
	// Does the Private Link Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
	IsManualConnection bool `pulumi:"isManualConnection"`
	// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	// The ID of the Private Link Enabled Remote Resource which this Private Link Endpoint should be connected to. Changing this forces a new resource to be created.
	PrivateConnectionResourceId string `pulumi:"privateConnectionResourceId"`
	// A message passed to the owner of the remote resource when the private link endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of `140` characters in length. Only valid if `isManualConnection` is set to `true`.
	RequestMessage *string `pulumi:"requestMessage"`
	// A list of subresource names which the Private Link Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceNames []string `pulumi:"subresourceNames"`
}

type LinkEndpointPrivateServiceConnectionInput interface {
	pulumi.Input

	ToLinkEndpointPrivateServiceConnectionOutput() LinkEndpointPrivateServiceConnectionOutput
	ToLinkEndpointPrivateServiceConnectionOutputWithContext(context.Context) LinkEndpointPrivateServiceConnectionOutput
}

type LinkEndpointPrivateServiceConnectionArgs struct {
	// Does the Private Link Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
	IsManualConnection pulumi.BoolInput `pulumi:"isManualConnection"`
	// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of the Private Link Enabled Remote Resource which this Private Link Endpoint should be connected to. Changing this forces a new resource to be created.
	PrivateConnectionResourceId pulumi.StringInput `pulumi:"privateConnectionResourceId"`
	// A message passed to the owner of the remote resource when the private link endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of `140` characters in length. Only valid if `isManualConnection` is set to `true`.
	RequestMessage pulumi.StringPtrInput `pulumi:"requestMessage"`
	// A list of subresource names which the Private Link Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceNames pulumi.StringArrayInput `pulumi:"subresourceNames"`
}

func (LinkEndpointPrivateServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkEndpointPrivateServiceConnection)(nil)).Elem()
}

func (i LinkEndpointPrivateServiceConnectionArgs) ToLinkEndpointPrivateServiceConnectionOutput() LinkEndpointPrivateServiceConnectionOutput {
	return i.ToLinkEndpointPrivateServiceConnectionOutputWithContext(context.Background())
}

func (i LinkEndpointPrivateServiceConnectionArgs) ToLinkEndpointPrivateServiceConnectionOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkEndpointPrivateServiceConnectionOutput)
}

func (i LinkEndpointPrivateServiceConnectionArgs) ToLinkEndpointPrivateServiceConnectionPtrOutput() LinkEndpointPrivateServiceConnectionPtrOutput {
	return i.ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(context.Background())
}

func (i LinkEndpointPrivateServiceConnectionArgs) ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkEndpointPrivateServiceConnectionOutput).ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(ctx)
}

type LinkEndpointPrivateServiceConnectionPtrInput interface {
	pulumi.Input

	ToLinkEndpointPrivateServiceConnectionPtrOutput() LinkEndpointPrivateServiceConnectionPtrOutput
	ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(context.Context) LinkEndpointPrivateServiceConnectionPtrOutput
}

type linkEndpointPrivateServiceConnectionPtrType LinkEndpointPrivateServiceConnectionArgs

func LinkEndpointPrivateServiceConnectionPtr(v *LinkEndpointPrivateServiceConnectionArgs) LinkEndpointPrivateServiceConnectionPtrInput {	return (*linkEndpointPrivateServiceConnectionPtrType)(v)
}

func (*linkEndpointPrivateServiceConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkEndpointPrivateServiceConnection)(nil)).Elem()
}

func (i *linkEndpointPrivateServiceConnectionPtrType) ToLinkEndpointPrivateServiceConnectionPtrOutput() LinkEndpointPrivateServiceConnectionPtrOutput {
	return i.ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(context.Background())
}

func (i *linkEndpointPrivateServiceConnectionPtrType) ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkEndpointPrivateServiceConnectionPtrOutput)
}

type LinkEndpointPrivateServiceConnectionOutput struct { *pulumi.OutputState }

func (LinkEndpointPrivateServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkEndpointPrivateServiceConnection)(nil)).Elem()
}

func (o LinkEndpointPrivateServiceConnectionOutput) ToLinkEndpointPrivateServiceConnectionOutput() LinkEndpointPrivateServiceConnectionOutput {
	return o
}

func (o LinkEndpointPrivateServiceConnectionOutput) ToLinkEndpointPrivateServiceConnectionOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionOutput {
	return o
}

func (o LinkEndpointPrivateServiceConnectionOutput) ToLinkEndpointPrivateServiceConnectionPtrOutput() LinkEndpointPrivateServiceConnectionPtrOutput {
	return o.ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(context.Background())
}

func (o LinkEndpointPrivateServiceConnectionOutput) ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionPtrOutput {
	return o.ApplyT(func(v LinkEndpointPrivateServiceConnection) *LinkEndpointPrivateServiceConnection {
		return &v
	}).(LinkEndpointPrivateServiceConnectionPtrOutput)
}
// Does the Private Link Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionOutput) IsManualConnection() pulumi.BoolOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) bool { return v.IsManualConnection }).(pulumi.BoolOutput)
}

// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Private Link Enabled Remote Resource which this Private Link Endpoint should be connected to. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionOutput) PrivateConnectionResourceId() pulumi.StringOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) string { return v.PrivateConnectionResourceId }).(pulumi.StringOutput)
}

// A message passed to the owner of the remote resource when the private link endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of `140` characters in length. Only valid if `isManualConnection` is set to `true`.
func (o LinkEndpointPrivateServiceConnectionOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// A list of subresource names which the Private Link Endpoint is able to connect to. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionOutput) SubresourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) []string { return v.SubresourceNames }).(pulumi.StringArrayOutput)
}

type LinkEndpointPrivateServiceConnectionPtrOutput struct { *pulumi.OutputState}

func (LinkEndpointPrivateServiceConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkEndpointPrivateServiceConnection)(nil)).Elem()
}

func (o LinkEndpointPrivateServiceConnectionPtrOutput) ToLinkEndpointPrivateServiceConnectionPtrOutput() LinkEndpointPrivateServiceConnectionPtrOutput {
	return o
}

func (o LinkEndpointPrivateServiceConnectionPtrOutput) ToLinkEndpointPrivateServiceConnectionPtrOutputWithContext(ctx context.Context) LinkEndpointPrivateServiceConnectionPtrOutput {
	return o
}

func (o LinkEndpointPrivateServiceConnectionPtrOutput) Elem() LinkEndpointPrivateServiceConnectionOutput {
	return o.ApplyT(func (v *LinkEndpointPrivateServiceConnection) LinkEndpointPrivateServiceConnection { return *v }).(LinkEndpointPrivateServiceConnectionOutput)
}

// Does the Private Link Endpoint require Manual Approval from the remote resource owner? Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionPtrOutput) IsManualConnection() pulumi.BoolOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) bool { return v.IsManualConnection }).(pulumi.BoolOutput)
}

// Specifies the Name of the Private Service Connection. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionPtrOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Private Link Enabled Remote Resource which this Private Link Endpoint should be connected to. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionPtrOutput) PrivateConnectionResourceId() pulumi.StringOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) string { return v.PrivateConnectionResourceId }).(pulumi.StringOutput)
}

// A message passed to the owner of the remote resource when the private link endpoint attempts to establish the connection to the remote resource. The request message can be a maximum of `140` characters in length. Only valid if `isManualConnection` is set to `true`.
func (o LinkEndpointPrivateServiceConnectionPtrOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// A list of subresource names which the Private Link Endpoint is able to connect to. Changing this forces a new resource to be created.
func (o LinkEndpointPrivateServiceConnectionPtrOutput) SubresourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func (v LinkEndpointPrivateServiceConnection) []string { return v.SubresourceNames }).(pulumi.StringArrayOutput)
}

type LinkServiceNatIpConfiguration struct {
	// Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	Primary bool `pulumi:"primary"`
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	PrivateIpAddressVersion *string `pulumi:"privateIpAddressVersion"`
	SubnetId string `pulumi:"subnetId"`
}

type LinkServiceNatIpConfigurationInput interface {
	pulumi.Input

	ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput
	ToLinkServiceNatIpConfigurationOutputWithContext(context.Context) LinkServiceNatIpConfigurationOutput
}

type LinkServiceNatIpConfigurationArgs struct {
	// Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	Primary pulumi.BoolInput `pulumi:"primary"`
	PrivateIpAddress pulumi.StringPtrInput `pulumi:"privateIpAddress"`
	PrivateIpAddressVersion pulumi.StringPtrInput `pulumi:"privateIpAddressVersion"`
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (LinkServiceNatIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (i LinkServiceNatIpConfigurationArgs) ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput {
	return i.ToLinkServiceNatIpConfigurationOutputWithContext(context.Background())
}

func (i LinkServiceNatIpConfigurationArgs) ToLinkServiceNatIpConfigurationOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkServiceNatIpConfigurationOutput)
}

type LinkServiceNatIpConfigurationArrayInput interface {
	pulumi.Input

	ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput
	ToLinkServiceNatIpConfigurationArrayOutputWithContext(context.Context) LinkServiceNatIpConfigurationArrayOutput
}

type LinkServiceNatIpConfigurationArray []LinkServiceNatIpConfigurationInput

func (LinkServiceNatIpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (i LinkServiceNatIpConfigurationArray) ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput {
	return i.ToLinkServiceNatIpConfigurationArrayOutputWithContext(context.Background())
}

func (i LinkServiceNatIpConfigurationArray) ToLinkServiceNatIpConfigurationArrayOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkServiceNatIpConfigurationArrayOutput)
}

type LinkServiceNatIpConfigurationOutput struct { *pulumi.OutputState }

func (LinkServiceNatIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (o LinkServiceNatIpConfigurationOutput) ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput {
	return o
}

func (o LinkServiceNatIpConfigurationOutput) ToLinkServiceNatIpConfigurationOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationOutput {
	return o
}

// Specifies the name of this Private Link Service. Changing this forces a new resource to be created.
func (o LinkServiceNatIpConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v LinkServiceNatIpConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o LinkServiceNatIpConfigurationOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func (v LinkServiceNatIpConfiguration) bool { return v.Primary }).(pulumi.BoolOutput)
}

func (o LinkServiceNatIpConfigurationOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LinkServiceNatIpConfiguration) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o LinkServiceNatIpConfigurationOutput) PrivateIpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func (v LinkServiceNatIpConfiguration) *string { return v.PrivateIpAddressVersion }).(pulumi.StringPtrOutput)
}

func (o LinkServiceNatIpConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func (v LinkServiceNatIpConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type LinkServiceNatIpConfigurationArrayOutput struct { *pulumi.OutputState}

func (LinkServiceNatIpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (o LinkServiceNatIpConfigurationArrayOutput) ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput {
	return o
}

func (o LinkServiceNatIpConfigurationArrayOutput) ToLinkServiceNatIpConfigurationArrayOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationArrayOutput {
	return o
}

func (o LinkServiceNatIpConfigurationArrayOutput) Index(i pulumi.IntInput) LinkServiceNatIpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) LinkServiceNatIpConfiguration {
		return vs[0].([]LinkServiceNatIpConfiguration)[vs[1].(int)]
	}).(LinkServiceNatIpConfigurationOutput)
}

type MxRecordRecord struct {
	Exchange string `pulumi:"exchange"`
	Preference int `pulumi:"preference"`
}

type MxRecordRecordInput interface {
	pulumi.Input

	ToMxRecordRecordOutput() MxRecordRecordOutput
	ToMxRecordRecordOutputWithContext(context.Context) MxRecordRecordOutput
}

type MxRecordRecordArgs struct {
	Exchange pulumi.StringInput `pulumi:"exchange"`
	Preference pulumi.IntInput `pulumi:"preference"`
}

func (MxRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return i.ToMxRecordRecordOutputWithContext(context.Background())
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordOutput)
}

type MxRecordRecordArrayInput interface {
	pulumi.Input

	ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput
	ToMxRecordRecordArrayOutputWithContext(context.Context) MxRecordRecordArrayOutput
}

type MxRecordRecordArray []MxRecordRecordInput

func (MxRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return i.ToMxRecordRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordArrayOutput)
}

type MxRecordRecordOutput struct { *pulumi.OutputState }

func (MxRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return o
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return o
}

func (o MxRecordRecordOutput) Exchange() pulumi.StringOutput {
	return o.ApplyT(func (v MxRecordRecord) string { return v.Exchange }).(pulumi.StringOutput)
}

func (o MxRecordRecordOutput) Preference() pulumi.IntOutput {
	return o.ApplyT(func (v MxRecordRecord) int { return v.Preference }).(pulumi.IntOutput)
}

type MxRecordRecordArrayOutput struct { *pulumi.OutputState}

func (MxRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) Index(i pulumi.IntInput) MxRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) MxRecordRecord {
		return vs[0].([]MxRecordRecord)[vs[1].(int)]
	}).(MxRecordRecordOutput)
}

type SRVRecordRecord struct {
	Port int `pulumi:"port"`
	Priority int `pulumi:"priority"`
	Target string `pulumi:"target"`
	Weight int `pulumi:"weight"`
}

type SRVRecordRecordInput interface {
	pulumi.Input

	ToSRVRecordRecordOutput() SRVRecordRecordOutput
	ToSRVRecordRecordOutputWithContext(context.Context) SRVRecordRecordOutput
}

type SRVRecordRecordArgs struct {
	Port pulumi.IntInput `pulumi:"port"`
	Priority pulumi.IntInput `pulumi:"priority"`
	Target pulumi.StringInput `pulumi:"target"`
	Weight pulumi.IntInput `pulumi:"weight"`
}

func (SRVRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecordRecord)(nil)).Elem()
}

func (i SRVRecordRecordArgs) ToSRVRecordRecordOutput() SRVRecordRecordOutput {
	return i.ToSRVRecordRecordOutputWithContext(context.Background())
}

func (i SRVRecordRecordArgs) ToSRVRecordRecordOutputWithContext(ctx context.Context) SRVRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SRVRecordRecordOutput)
}

type SRVRecordRecordArrayInput interface {
	pulumi.Input

	ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput
	ToSRVRecordRecordArrayOutputWithContext(context.Context) SRVRecordRecordArrayOutput
}

type SRVRecordRecordArray []SRVRecordRecordInput

func (SRVRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SRVRecordRecord)(nil)).Elem()
}

func (i SRVRecordRecordArray) ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput {
	return i.ToSRVRecordRecordArrayOutputWithContext(context.Background())
}

func (i SRVRecordRecordArray) ToSRVRecordRecordArrayOutputWithContext(ctx context.Context) SRVRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SRVRecordRecordArrayOutput)
}

type SRVRecordRecordOutput struct { *pulumi.OutputState }

func (SRVRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecordRecord)(nil)).Elem()
}

func (o SRVRecordRecordOutput) ToSRVRecordRecordOutput() SRVRecordRecordOutput {
	return o
}

func (o SRVRecordRecordOutput) ToSRVRecordRecordOutputWithContext(ctx context.Context) SRVRecordRecordOutput {
	return o
}

func (o SRVRecordRecordOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func (v SRVRecordRecord) int { return v.Port }).(pulumi.IntOutput)
}

func (o SRVRecordRecordOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func (v SRVRecordRecord) int { return v.Priority }).(pulumi.IntOutput)
}

func (o SRVRecordRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func (v SRVRecordRecord) string { return v.Target }).(pulumi.StringOutput)
}

func (o SRVRecordRecordOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func (v SRVRecordRecord) int { return v.Weight }).(pulumi.IntOutput)
}

type SRVRecordRecordArrayOutput struct { *pulumi.OutputState}

func (SRVRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SRVRecordRecord)(nil)).Elem()
}

func (o SRVRecordRecordArrayOutput) ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput {
	return o
}

func (o SRVRecordRecordArrayOutput) ToSRVRecordRecordArrayOutputWithContext(ctx context.Context) SRVRecordRecordArrayOutput {
	return o
}

func (o SRVRecordRecordArrayOutput) Index(i pulumi.IntInput) SRVRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) SRVRecordRecord {
		return vs[0].([]SRVRecordRecord)[vs[1].(int)]
	}).(SRVRecordRecordOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkEndpointPrivateServiceConnectionOutput{})
	pulumi.RegisterOutputType(LinkEndpointPrivateServiceConnectionPtrOutput{})
	pulumi.RegisterOutputType(LinkServiceNatIpConfigurationOutput{})
	pulumi.RegisterOutputType(LinkServiceNatIpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(MxRecordRecordOutput{})
	pulumi.RegisterOutputType(MxRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(SRVRecordRecordOutput{})
	pulumi.RegisterOutputType(SRVRecordRecordArrayOutput{})
}
