// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package authorization

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/role_definition.html.markdown.
type RoleDefinition struct {
	pulumi.CustomResourceState

	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayOutput `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayOutput `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringOutput `pulumi:"roleDefinitionId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
	Scope pulumi.StringOutput `pulumi:"scope"`
}

// NewRoleDefinition registers a new resource with the given unique name, arguments, and options.
func NewRoleDefinition(ctx *pulumi.Context,
	name string, args *RoleDefinitionArgs, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	if args == nil || args.AssignableScopes == nil {
		return nil, errors.New("missing required argument 'AssignableScopes'")
	}
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &RoleDefinitionArgs{}
	}
	var resource RoleDefinition
	err := ctx.RegisterResource("azure:authorization/roleDefinition:RoleDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleDefinition gets an existing RoleDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleDefinitionState, opts ...pulumi.ResourceOption) (*RoleDefinition, error) {
	var resource RoleDefinition
	err := ctx.ReadResource("azure:authorization/roleDefinition:RoleDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleDefinition resources.
type roleDefinitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []RoleDefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
	Scope *string `pulumi:"scope"`
}

type RoleDefinitionState struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
	Scope pulumi.StringPtrInput
}

func (RoleDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionState)(nil)).Elem()
}

type roleDefinitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes []string `pulumi:"assignableScopes"`
	// A description of the Role Definition.
	Description *string `pulumi:"description"`
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `permissions` block as defined below.
	Permissions []RoleDefinitionPermission `pulumi:"permissions"`
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleDefinition resource.
type RoleDefinitionArgs struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes pulumi.StringArrayInput
	// A description of the Role Definition.
	Description pulumi.StringPtrInput
	// The name of the Role Definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `permissions` block as defined below.
	Permissions RoleDefinitionPermissionArrayInput
	// A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionId pulumi.StringPtrInput
	// The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
	Scope pulumi.StringInput
}

func (RoleDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleDefinitionArgs)(nil)).Elem()
}

