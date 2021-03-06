// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Disaster Recovery Config for an Event Hub Namespace.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub_namespace_disaster_recovery_config.html.markdown.
type EventhubNamespaceDisasterRecoveryConfig struct {
	pulumi.CustomResourceState

	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrOutput `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringOutput `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEventhubNamespaceDisasterRecoveryConfig registers a new resource with the given unique name, arguments, and options.
func NewEventhubNamespaceDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, args *EventhubNamespaceDisasterRecoveryConfigArgs, opts ...pulumi.ResourceOption) (*EventhubNamespaceDisasterRecoveryConfig, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.PartnerNamespaceId == nil {
		return nil, errors.New("missing required argument 'PartnerNamespaceId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventhubNamespaceDisasterRecoveryConfigArgs{}
	}
	var resource EventhubNamespaceDisasterRecoveryConfig
	err := ctx.RegisterResource("azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventhubNamespaceDisasterRecoveryConfig gets an existing EventhubNamespaceDisasterRecoveryConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventhubNamespaceDisasterRecoveryConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventhubNamespaceDisasterRecoveryConfigState, opts ...pulumi.ResourceOption) (*EventhubNamespaceDisasterRecoveryConfig, error) {
	var resource EventhubNamespaceDisasterRecoveryConfig
	err := ctx.ReadResource("azure:eventhub/eventhubNamespaceDisasterRecoveryConfig:EventhubNamespaceDisasterRecoveryConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventhubNamespaceDisasterRecoveryConfig resources.
type eventhubNamespaceDisasterRecoveryConfigState struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName *string `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId *string `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EventhubNamespaceDisasterRecoveryConfigState struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrInput
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringPtrInput
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (EventhubNamespaceDisasterRecoveryConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubNamespaceDisasterRecoveryConfigState)(nil)).Elem()
}

type eventhubNamespaceDisasterRecoveryConfigArgs struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName *string `pulumi:"alternateName"`
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId string `pulumi:"partnerNamespaceId"`
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EventhubNamespaceDisasterRecoveryConfig resource.
type EventhubNamespaceDisasterRecoveryConfigArgs struct {
	// An alternate name to use when the Disaster Recovery Config's name is the same as the replicated namespace's name.
	AlternateName pulumi.StringPtrInput
	// Specifies the name of the Disaster Recovery Config. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the primary EventHub Namespace to replicate. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The ID of the EventHub Namespace to replicate to.
	PartnerNamespaceId pulumi.StringInput
	// The name of the resource group in which the Disaster Recovery Config exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (EventhubNamespaceDisasterRecoveryConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubNamespaceDisasterRecoveryConfigArgs)(nil)).Elem()
}

