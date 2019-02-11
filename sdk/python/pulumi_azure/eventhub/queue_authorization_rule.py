# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class QueueAuthorizationRule(pulumi.CustomResource):
    listen: pulumi.Output[bool]
    """
    Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
    """
    manage: pulumi.Output[bool]
    """
    Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
    """
    primary_connection_string: pulumi.Output[str]
    """
    The Primary Connection String for the Authorization Rule.
    """
    primary_key: pulumi.Output[str]
    """
    The Primary Key for the Authorization Rule.
    """
    queue_name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
    """
    secondary_connection_string: pulumi.Output[str]
    """
    The Secondary Connection String for the Authorization Rule.
    """
    secondary_key: pulumi.Output[str]
    """
    The Secondary Key for the Authorization Rule.
    """
    send: pulumi.Output[bool]
    """
    Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
    """
    def __init__(__self__, resource_name, opts=None, listen=None, manage=None, name=None, namespace_name=None, queue_name=None, resource_group_name=None, send=None, __name__=None, __opts__=None):
        """
        Manages an Authorization Rule for a ServiceBus Queue.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] listen: Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to `false`.
        :param pulumi.Input[bool] manage: Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
        :param pulumi.Input[str] name: Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: Specifies the name of the ServiceBus Namespace in which the Queue exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] queue_name: Specifies the name of the ServiceBus Queue. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] send: Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to `false`.
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

        __props__['listen'] = listen

        __props__['manage'] = manage

        __props__['name'] = name

        if namespace_name is None:
            raise TypeError('Missing required property namespace_name')
        __props__['namespace_name'] = namespace_name

        if queue_name is None:
            raise TypeError('Missing required property queue_name')
        __props__['queue_name'] = queue_name

        if resource_group_name is None:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        __props__['send'] = send

        __props__['primary_connection_string'] = None
        __props__['primary_key'] = None
        __props__['secondary_connection_string'] = None
        __props__['secondary_key'] = None

        super(QueueAuthorizationRule, __self__).__init__(
            'azure:eventhub/queueAuthorizationRule:QueueAuthorizationRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

