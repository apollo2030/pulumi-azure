# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DiagnosticSetting(pulumi.CustomResource):
    eventhub_authorization_rule_id: pulumi.Output[str]
    """
    Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
    """
    eventhub_name: pulumi.Output[str]
    """
    Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
    """
    logs: pulumi.Output[list]
    """
    One or more `log` blocks as defined below.
    """
    log_analytics_workspace_id: pulumi.Output[str]
    """
    Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
    """
    metrics: pulumi.Output[list]
    """
    One or more `metric` blocks as defined below.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
    """
    storage_account_id: pulumi.Output[str]
    """
    With this parameter you can specify a storage account which should be used to send the logs to. Parameter must be a valid Azure Resource ID. Changing this forces a new resource to be created.
    """
    target_resource_id: pulumi.Output[str]
    """
    The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, eventhub_authorization_rule_id=None, eventhub_name=None, logs=None, log_analytics_workspace_id=None, metrics=None, name=None, storage_account_id=None, target_resource_id=None, __name__=None, __opts__=None):
        """
        Manages a Diagnostic Setting for an existing Resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_authorization_rule_id: Specifies the ID of an Event Hub Namespace Authorization Rule used to send Diagnostics Data. Changing this forces a new resource to be created.
        :param pulumi.Input[str] eventhub_name: Specifies the name of the Event Hub where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[list] logs: One or more `log` blocks as defined below.
        :param pulumi.Input[str] log_analytics_workspace_id: Specifies the ID of a Log Analytics Workspace where Diagnostics Data should be sent. Changing this forces a new resource to be created.
        :param pulumi.Input[list] metrics: One or more `metric` blocks as defined below.
        :param pulumi.Input[str] name: Specifies the name of the Diagnostic Setting. Changing this forces a new resource to be created.
        :param pulumi.Input[str] storage_account_id: With this parameter you can specify a storage account which should be used to send the logs to. Parameter must be a valid Azure Resource ID. Changing this forces a new resource to be created.
        :param pulumi.Input[str] target_resource_id: The ID of an existing Resource on which to configure Diagnostic Settings. Changing this forces a new resource to be created.
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

        __props__['eventhub_authorization_rule_id'] = eventhub_authorization_rule_id

        __props__['eventhub_name'] = eventhub_name

        __props__['logs'] = logs

        __props__['log_analytics_workspace_id'] = log_analytics_workspace_id

        __props__['metrics'] = metrics

        __props__['name'] = name

        __props__['storage_account_id'] = storage_account_id

        if target_resource_id is None:
            raise TypeError('Missing required property target_resource_id')
        __props__['target_resource_id'] = target_resource_id

        super(DiagnosticSetting, __self__).__init__(
            'azure:monitoring/diagnosticSetting:DiagnosticSetting',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

