# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ActiveSlot(pulumi.CustomResource):
    app_service_name: pulumi.Output[str]
    """
    The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
    """
    app_service_slot_name: pulumi.Output[str]
    """
    The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, app_service_name=None, app_service_slot_name=None, resource_group_name=None, __name__=None, __opts__=None):
        """
        Promotes an App Service Slot to Production within an App Service.
        
        > **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azurerm_app_service` resource will be overwritten when promoting a Slot using the `azurerm_app_service_active_slot` resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service_name: The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] app_service_slot_name: The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if app_service_name is None:
            raise TypeError('Missing required property app_service_name')
        __props__['app_service_name'] = app_service_name

        if app_service_slot_name is None:
            raise TypeError('Missing required property app_service_slot_name')
        __props__['app_service_slot_name'] = app_service_slot_name

        if resource_group_name is None:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        super(ActiveSlot, __self__).__init__(
            'azure:appservice/activeSlot:ActiveSlot',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

