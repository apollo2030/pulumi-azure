# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VirtualMachineScaleSetExtension(pulumi.CustomResource):
    auto_upgrade_minor_version: pulumi.Output[bool]
    """
    Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
    """
    force_update_tag: pulumi.Output[str]
    """
    A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
    """
    name: pulumi.Output[str]
    """
    The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
    """
    protected_settings: pulumi.Output[str]
    """
    A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
    """
    provision_after_extensions: pulumi.Output[list]
    """
    An ordered list of Extension names which this should be provisioned after.
    """
    publisher: pulumi.Output[str]
    """
    Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
    """
    settings: pulumi.Output[str]
    """
    A JSON String which specifies Settings for the Extension.
    """
    type: pulumi.Output[str]
    """
    Specifies the Type of the Extension. Changing this forces a new resource to be created.
    """
    type_handler_version: pulumi.Output[str]
    """
    Specifies the version of the Script Handler which should be used.
    """
    virtual_machine_scale_set_id: pulumi.Output[str]
    """
    The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, auto_upgrade_minor_version=None, force_update_tag=None, name=None, protected_settings=None, provision_after_extensions=None, publisher=None, settings=None, type=None, type_handler_version=None, virtual_machine_scale_set_id=None, __props__=None, __name__=None, __opts__=None):
        """
        > **NOTE:** **This resource is in Beta** and as such the Schema can change in Minor versions of the Provider.
        
        Manages an Extension for a Virtual Machine Scale Set.
        
        > **NOTE:** This resource is not intended to be used with the `compute.ScaleSet` resource - instead it's intended for this to be used with the `compute.LinuxVirtualMachineScaleSet` and `compute.WindowsVirtualMachineScaleSet` resources.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        :param pulumi.Input[str] force_update_tag: A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        :param pulumi.Input[str] name: The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        :param pulumi.Input[list] provision_after_extensions: An ordered list of Extension names which this should be provisioned after.
        :param pulumi.Input[str] publisher: Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] settings: A JSON String which specifies Settings for the Extension.
        :param pulumi.Input[str] type: Specifies the Type of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the Script Handler which should be used.
        :param pulumi.Input[str] virtual_machine_scale_set_id: The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_scale_set_extension.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auto_upgrade_minor_version'] = auto_upgrade_minor_version
            __props__['force_update_tag'] = force_update_tag
            __props__['name'] = name
            __props__['protected_settings'] = protected_settings
            __props__['provision_after_extensions'] = provision_after_extensions
            if publisher is None:
                raise TypeError("Missing required property 'publisher'")
            __props__['publisher'] = publisher
            __props__['settings'] = settings
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            if type_handler_version is None:
                raise TypeError("Missing required property 'type_handler_version'")
            __props__['type_handler_version'] = type_handler_version
            if virtual_machine_scale_set_id is None:
                raise TypeError("Missing required property 'virtual_machine_scale_set_id'")
            __props__['virtual_machine_scale_set_id'] = virtual_machine_scale_set_id
        super(VirtualMachineScaleSetExtension, __self__).__init__(
            'azure:compute/virtualMachineScaleSetExtension:VirtualMachineScaleSetExtension',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_upgrade_minor_version=None, force_update_tag=None, name=None, protected_settings=None, provision_after_extensions=None, publisher=None, settings=None, type=None, type_handler_version=None, virtual_machine_scale_set_id=None):
        """
        Get an existing VirtualMachineScaleSetExtension resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_upgrade_minor_version: Should the latest version of the Extension be used at Deployment Time, if one is available? This won't auto-update the extension on existing installation. Defaults to `true`.
        :param pulumi.Input[str] force_update_tag: A value which, when different to the previous value can be used to force-run the Extension even if the Extension Configuration hasn't changed.
        :param pulumi.Input[str] name: The name for the Virtual Machine Scale Set Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protected_settings: A JSON String which specifies Sensitive Settings (such as Passwords) for the Extension.
        :param pulumi.Input[list] provision_after_extensions: An ordered list of Extension names which this should be provisioned after.
        :param pulumi.Input[str] publisher: Specifies the Publisher of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] settings: A JSON String which specifies Settings for the Extension.
        :param pulumi.Input[str] type: Specifies the Type of the Extension. Changing this forces a new resource to be created.
        :param pulumi.Input[str] type_handler_version: Specifies the version of the Script Handler which should be used.
        :param pulumi.Input[str] virtual_machine_scale_set_id: The ID of the Virtual Machine Scale Set. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_scale_set_extension.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["auto_upgrade_minor_version"] = auto_upgrade_minor_version
        __props__["force_update_tag"] = force_update_tag
        __props__["name"] = name
        __props__["protected_settings"] = protected_settings
        __props__["provision_after_extensions"] = provision_after_extensions
        __props__["publisher"] = publisher
        __props__["settings"] = settings
        __props__["type"] = type
        __props__["type_handler_version"] = type_handler_version
        __props__["virtual_machine_scale_set_id"] = virtual_machine_scale_set_id
        return VirtualMachineScaleSetExtension(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

