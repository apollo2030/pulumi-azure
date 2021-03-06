// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package backup

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Backup VM Backup Policy.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_policy_vm.html.markdown.
type PolicyVM struct {
	pulumi.CustomResourceState

	// Configures the Policy backup frequency, times & days as documented in the `backup` block below.
	Backup PolicyVMBackupOutput `pulumi:"backup"`
	// Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringOutput `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
	RetentionDaily PolicyVMRetentionDailyPtrOutput `pulumi:"retentionDaily"`
	// Configures the policy monthly retention as documented in the `retentionMonthly` block below.
	RetentionMonthly PolicyVMRetentionMonthlyPtrOutput `pulumi:"retentionMonthly"`
	// Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
	RetentionWeekly PolicyVMRetentionWeeklyPtrOutput `pulumi:"retentionWeekly"`
	// Configures the policy yearly retention as documented in the `retentionYearly` block below.
	RetentionYearly PolicyVMRetentionYearlyPtrOutput `pulumi:"retentionYearly"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewPolicyVM registers a new resource with the given unique name, arguments, and options.
func NewPolicyVM(ctx *pulumi.Context,
	name string, args *PolicyVMArgs, opts ...pulumi.ResourceOption) (*PolicyVM, error) {
	if args == nil || args.Backup == nil {
		return nil, errors.New("missing required argument 'Backup'")
	}
	if args == nil || args.RecoveryVaultName == nil {
		return nil, errors.New("missing required argument 'RecoveryVaultName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PolicyVMArgs{}
	}
	var resource PolicyVM
	err := ctx.RegisterResource("azure:backup/policyVM:PolicyVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyVM gets an existing PolicyVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyVMState, opts ...pulumi.ResourceOption) (*PolicyVM, error) {
	var resource PolicyVM
	err := ctx.ReadResource("azure:backup/policyVM:PolicyVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyVM resources.
type policyVMState struct {
	// Configures the Policy backup frequency, times & days as documented in the `backup` block below.
	Backup *PolicyVMBackup `pulumi:"backup"`
	// Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
	RetentionDaily *PolicyVMRetentionDaily `pulumi:"retentionDaily"`
	// Configures the policy monthly retention as documented in the `retentionMonthly` block below.
	RetentionMonthly *PolicyVMRetentionMonthly `pulumi:"retentionMonthly"`
	// Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
	RetentionWeekly *PolicyVMRetentionWeekly `pulumi:"retentionWeekly"`
	// Configures the policy yearly retention as documented in the `retentionYearly` block below.
	RetentionYearly *PolicyVMRetentionYearly `pulumi:"retentionYearly"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

type PolicyVMState struct {
	// Configures the Policy backup frequency, times & days as documented in the `backup` block below.
	Backup PolicyVMBackupPtrInput
	// Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringPtrInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
	RetentionDaily PolicyVMRetentionDailyPtrInput
	// Configures the policy monthly retention as documented in the `retentionMonthly` block below.
	RetentionMonthly PolicyVMRetentionMonthlyPtrInput
	// Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
	RetentionWeekly PolicyVMRetentionWeeklyPtrInput
	// Configures the policy yearly retention as documented in the `retentionYearly` block below.
	RetentionYearly PolicyVMRetentionYearlyPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVMState)(nil)).Elem()
}

type policyVMArgs struct {
	// Configures the Policy backup frequency, times & days as documented in the `backup` block below.
	Backup PolicyVMBackup `pulumi:"backup"`
	// Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName string `pulumi:"recoveryVaultName"`
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
	RetentionDaily *PolicyVMRetentionDaily `pulumi:"retentionDaily"`
	// Configures the policy monthly retention as documented in the `retentionMonthly` block below.
	RetentionMonthly *PolicyVMRetentionMonthly `pulumi:"retentionMonthly"`
	// Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
	RetentionWeekly *PolicyVMRetentionWeekly `pulumi:"retentionWeekly"`
	// Configures the policy yearly retention as documented in the `retentionYearly` block below.
	RetentionYearly *PolicyVMRetentionYearly `pulumi:"retentionYearly"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the timezone. Defaults to `UTC`
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a PolicyVM resource.
type PolicyVMArgs struct {
	// Configures the Policy backup frequency, times & days as documented in the `backup` block below.
	Backup PolicyVMBackupInput
	// Specifies the name of the Backup Policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName pulumi.StringInput
	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Configures the policy daily retention as documented in the `retentionDaily` block below. Required when backup frequency is `Daily`.
	RetentionDaily PolicyVMRetentionDailyPtrInput
	// Configures the policy monthly retention as documented in the `retentionMonthly` block below.
	RetentionMonthly PolicyVMRetentionMonthlyPtrInput
	// Configures the policy weekly retention as documented in the `retentionWeekly` block below. Required when backup frequency is `Weekly`.
	RetentionWeekly PolicyVMRetentionWeeklyPtrInput
	// Configures the policy yearly retention as documented in the `retentionYearly` block below.
	RetentionYearly PolicyVMRetentionYearlyPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the timezone. Defaults to `UTC`
	Timezone pulumi.StringPtrInput
}

func (PolicyVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyVMArgs)(nil)).Elem()
}

