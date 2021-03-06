// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package automation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Automation DSC Configuration.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/automation_dsc_configuration.html.markdown.
type DscConfiguration struct {
	pulumi.CustomResourceState

	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringOutput `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location pulumi.StringOutput `pulumi:"location"`
	// Verbose log option.
	LogVerbose pulumi.BoolPtrOutput `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	State pulumi.StringOutput `pulumi:"state"`
}

// NewDscConfiguration registers a new resource with the given unique name, arguments, and options.
func NewDscConfiguration(ctx *pulumi.Context,
	name string, args *DscConfigurationArgs, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	if args == nil || args.AutomationAccountName == nil {
		return nil, errors.New("missing required argument 'AutomationAccountName'")
	}
	if args == nil || args.ContentEmbedded == nil {
		return nil, errors.New("missing required argument 'ContentEmbedded'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DscConfigurationArgs{}
	}
	var resource DscConfiguration
	err := ctx.RegisterResource("azure:automation/dscConfiguration:DscConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDscConfiguration gets an existing DscConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDscConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscConfigurationState, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	var resource DscConfiguration
	err := ctx.ReadResource("azure:automation/dscConfiguration:DscConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DscConfiguration resources.
type dscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded *string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	State *string `pulumi:"state"`
}

type DscConfigurationState struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringPtrInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	State pulumi.StringPtrInput
}

func (DscConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationState)(nil)).Elem()
}

type dscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The PowerShell DSC Configuration script.
	ContentEmbedded string `pulumi:"contentEmbedded"`
	// Description to go with DSC Configuration.
	Description *string `pulumi:"description"`
	// Must be the same location as the Automation Account.
	Location *string `pulumi:"location"`
	// Verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DscConfiguration resource.
type DscConfigurationArgs struct {
	// The name of the automation account in which the DSC Configuration is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The PowerShell DSC Configuration script.
	ContentEmbedded pulumi.StringInput
	// Description to go with DSC Configuration.
	Description pulumi.StringPtrInput
	// Must be the same location as the Automation Account.
	Location pulumi.StringPtrInput
	// Verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Specifies the name of the DSC Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the DSC Configuration is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (DscConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationArgs)(nil)).Elem()
}

