// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package frontdoor

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Front Door Web Application Firewall Policy instance.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/frontdoor_firewall_policy.html.markdown.
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrOutput `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrOutput `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayOutput `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds pulumi.StringArrayOutput `pulumi:"frontendEndpointIds"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayOutput `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrOutput `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &FirewallPolicyArgs{}
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("azure:frontdoor/firewallPolicy:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azure:frontdoor/firewallPolicy:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode *int `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules []FirewallPolicyCustomRule `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds []string `pulumi:"frontendEndpointIds"`
	// Resource location.
	Location *string `pulumi:"location"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules []FirewallPolicyManagedRule `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode *string `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

type FirewallPolicyState struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrInput
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrInput
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayInput
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// the Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds pulumi.StringArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayInput
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `pulumi:"customBlockResponseBody"`
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode *int `pulumi:"customBlockResponseStatusCode"`
	// One or more `customRule` blocks as defined below.
	CustomRules []FirewallPolicyCustomRule `pulumi:"customRules"`
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// One or more `managedRule` blocks as defined below.
	ManagedRules []FirewallPolicyManagedRule `pulumi:"managedRules"`
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode *string `pulumi:"mode"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl *string `pulumi:"redirectUrl"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// If a `customRule` block's action type is `block`, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody pulumi.StringPtrInput
	// If a `customRule` block's action type is `block`, this is the response status code. Possible values are `200`, `403`, `405`, `406`, or `429`.
	CustomBlockResponseStatusCode pulumi.IntPtrInput
	// One or more `customRule` blocks as defined below.
	CustomRules FirewallPolicyCustomRuleArrayInput
	// Is the policy a enabled state or disabled state. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// One or more `managedRule` blocks as defined below.
	ManagedRules FirewallPolicyManagedRuleArrayInput
	// The firewall policy mode. Possible values are `Detection`, `Prevention` and defaults to `Prevention`.
	Mode pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// If action type is redirect, this field represents redirect URL for the client.
	RedirectUrl pulumi.StringPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

