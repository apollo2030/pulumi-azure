# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProtectionPolicyVM(pulumi.CustomResource):
    backup: pulumi.Output[dict]
    """
    Configures the Policy backup frequecent, times & days as documented in the `backup` block below. 
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Recovery Services Vault Policy. Changing this forces a new resource to be created.
    """
    recovery_vault_name: pulumi.Output[str]
    """
    Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
    """
    retention_daily: pulumi.Output[dict]
    """
    Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
    """
    retention_monthly: pulumi.Output[dict]
    """
    Configures the policy monthly retention as documented in the `retention_monthly` block below.
    """
    retention_weekly: pulumi.Output[dict]
    """
    Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
    """
    retention_yearly: pulumi.Output[dict]
    """
    Configures the policy yearly retention as documented in the `retention_yearly` block below.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    timezone: pulumi.Output[str]
    """
    Specifies the timezone. Defaults to `UTC`
    """
    def __init__(__self__, __name__, __opts__=None, backup=None, name=None, recovery_vault_name=None, resource_group_name=None, retention_daily=None, retention_monthly=None, retention_weekly=None, retention_yearly=None, tags=None, timezone=None):
        """
        Manages an Recovery Services VM Protection Policy.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] backup: Configures the Policy backup frequecent, times & days as documented in the `backup` block below. 
        :param pulumi.Input[str] name: Specifies the name of the Recovery Services Vault Policy. Changing this forces a new resource to be created.
        :param pulumi.Input[str] recovery_vault_name: Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] retention_daily: Configures the policy daily retention as documented in the `retention_daily` block below. Required when backup frequency is `Daily`.
        :param pulumi.Input[dict] retention_monthly: Configures the policy monthly retention as documented in the `retention_monthly` block below.
        :param pulumi.Input[dict] retention_weekly: Configures the policy weekly retention as documented in the `retention_weekly` block below. Required when backup frequency is `Weekly`.
        :param pulumi.Input[dict] retention_yearly: Configures the policy yearly retention as documented in the `retention_yearly` block below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] timezone: Specifies the timezone. Defaults to `UTC`
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not backup:
            raise TypeError('Missing required property backup')
        __props__['backup'] = backup

        __props__['name'] = name

        if not recovery_vault_name:
            raise TypeError('Missing required property recovery_vault_name')
        __props__['recovery_vault_name'] = recovery_vault_name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['retention_daily'] = retention_daily

        __props__['retention_monthly'] = retention_monthly

        __props__['retention_weekly'] = retention_weekly

        __props__['retention_yearly'] = retention_yearly

        __props__['tags'] = tags

        __props__['timezone'] = timezone

        super(ProtectionPolicyVM, __self__).__init__(
            'azure:recoveryservices/protectionPolicyVM:ProtectionPolicyVM',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
