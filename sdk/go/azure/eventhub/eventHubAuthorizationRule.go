// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Event Hubs authorization Rule within an Event Hub.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub_authorization_rule_legacy.html.markdown.
type EventHubAuthorizationRule struct {
	pulumi.CustomResourceState

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringOutput `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	Location pulumi.StringOutput `pulumi:"location"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewEventHubAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *EventHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	if args == nil || args.EventhubName == nil {
		return nil, errors.New("missing required argument 'EventhubName'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventHubAuthorizationRuleArgs{}
	}
	var resource EventHubAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubAuthorizationRule gets an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	var resource EventHubAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/eventHubAuthorizationRule:EventHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubAuthorizationRule resources.
type eventHubAuthorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName *string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	Location *string `pulumi:"location"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type EventHubAuthorizationRuleState struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the Event Hubs Authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The Secondary Key for the Event Hubs Authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleState)(nil)).Elem()
}

type eventHubAuthorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName string `pulumi:"eventhubName"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	Location *string `pulumi:"location"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a EventHubAuthorizationRule resource.
type EventHubAuthorizationRuleArgs struct {
	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventhubName pulumi.StringInput
	// Does this Authorization Rule have permissions to Listen to the Event Hub? Defaults to `false`.
	Listen pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// Does this Authorization Rule have permissions to Manage to the Event Hub? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the EventHub Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Does this Authorization Rule have permissions to Send to the Event Hub? Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleArgs)(nil)).Elem()
}

