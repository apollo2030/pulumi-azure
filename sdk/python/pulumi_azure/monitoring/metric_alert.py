# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class MetricAlert(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    One or more `action` blocks as defined below.
    
      * `actionGroupId` (`str`)
      * `webhookProperties` (`dict`)
    """
    auto_mitigate: pulumi.Output[bool]
    """
    Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
    """
    criterias: pulumi.Output[list]
    """
    One or more `criteria` blocks as defined below.
    
      * `aggregation` (`str`)
      * `dimensions` (`list`)
    
        * `name` (`str`) - The name of the Metric Alert. Changing this forces a new resource to be created.
        * `operator` (`str`)
        * `values` (`list`)
    
      * `metric_name` (`str`)
      * `metricNamespace` (`str`)
      * `operator` (`str`)
      * `threshold` (`float`)
    """
    description: pulumi.Output[str]
    """
    The description of this Metric Alert.
    """
    enabled: pulumi.Output[bool]
    """
    Should this Metric Alert be enabled? Defaults to `true`.
    """
    frequency: pulumi.Output[str]
    """
    The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
    """
    name: pulumi.Output[str]
    """
    The name of the Metric Alert. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Metric Alert instance.
    """
    scopes: pulumi.Output[str]
    """
    A set of strings of resource IDs at which the metric criteria should be applied.
    """
    severity: pulumi.Output[float]
    """
    The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    window_size: pulumi.Output[str]
    """
    The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, auto_mitigate=None, criterias=None, description=None, enabled=None, frequency=None, name=None, resource_group_name=None, scopes=None, severity=None, tags=None, window_size=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Metric Alert within Azure Monitor.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[bool] auto_mitigate: Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
        :param pulumi.Input[list] criterias: One or more `criteria` blocks as defined below.
        :param pulumi.Input[str] description: The description of this Metric Alert.
        :param pulumi.Input[bool] enabled: Should this Metric Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] frequency: The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
        :param pulumi.Input[str] name: The name of the Metric Alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Metric Alert instance.
        :param pulumi.Input[str] scopes: A set of strings of resource IDs at which the metric criteria should be applied.
        :param pulumi.Input[float] severity: The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] window_size: The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
        
        The **actions** object supports the following:
        
          * `actionGroupId` (`pulumi.Input[str]`)
          * `webhookProperties` (`pulumi.Input[dict]`)
        
        The **criterias** object supports the following:
        
          * `aggregation` (`pulumi.Input[str]`)
          * `dimensions` (`pulumi.Input[list]`)
        
            * `name` (`pulumi.Input[str]`) - The name of the Metric Alert. Changing this forces a new resource to be created.
            * `operator` (`pulumi.Input[str]`)
            * `values` (`pulumi.Input[list]`)
        
          * `metric_name` (`pulumi.Input[str]`)
          * `metricNamespace` (`pulumi.Input[str]`)
          * `operator` (`pulumi.Input[str]`)
          * `threshold` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_metric_alert.html.markdown.
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

            __props__['actions'] = actions
            __props__['auto_mitigate'] = auto_mitigate
            if criterias is None:
                raise TypeError("Missing required property 'criterias'")
            __props__['criterias'] = criterias
            __props__['description'] = description
            __props__['enabled'] = enabled
            __props__['frequency'] = frequency
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if scopes is None:
                raise TypeError("Missing required property 'scopes'")
            __props__['scopes'] = scopes
            __props__['severity'] = severity
            __props__['tags'] = tags
            __props__['window_size'] = window_size
        super(MetricAlert, __self__).__init__(
            'azure:monitoring/metricAlert:MetricAlert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, auto_mitigate=None, criterias=None, description=None, enabled=None, frequency=None, name=None, resource_group_name=None, scopes=None, severity=None, tags=None, window_size=None):
        """
        Get an existing MetricAlert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: One or more `action` blocks as defined below.
        :param pulumi.Input[bool] auto_mitigate: Should the alerts in this Metric Alert be auto resolved? Defaults to `false`.
        :param pulumi.Input[list] criterias: One or more `criteria` blocks as defined below.
        :param pulumi.Input[str] description: The description of this Metric Alert.
        :param pulumi.Input[bool] enabled: Should this Metric Alert be enabled? Defaults to `true`.
        :param pulumi.Input[str] frequency: The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
        :param pulumi.Input[str] name: The name of the Metric Alert. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Metric Alert instance.
        :param pulumi.Input[str] scopes: A set of strings of resource IDs at which the metric criteria should be applied.
        :param pulumi.Input[float] severity: The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] window_size: The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
        
        The **actions** object supports the following:
        
          * `actionGroupId` (`pulumi.Input[str]`)
          * `webhookProperties` (`pulumi.Input[dict]`)
        
        The **criterias** object supports the following:
        
          * `aggregation` (`pulumi.Input[str]`)
          * `dimensions` (`pulumi.Input[list]`)
        
            * `name` (`pulumi.Input[str]`) - The name of the Metric Alert. Changing this forces a new resource to be created.
            * `operator` (`pulumi.Input[str]`)
            * `values` (`pulumi.Input[list]`)
        
          * `metric_name` (`pulumi.Input[str]`)
          * `metricNamespace` (`pulumi.Input[str]`)
          * `operator` (`pulumi.Input[str]`)
          * `threshold` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_metric_alert.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["actions"] = actions
        __props__["auto_mitigate"] = auto_mitigate
        __props__["criterias"] = criterias
        __props__["description"] = description
        __props__["enabled"] = enabled
        __props__["frequency"] = frequency
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["scopes"] = scopes
        __props__["severity"] = severity
        __props__["tags"] = tags
        __props__["window_size"] = window_size
        return MetricAlert(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

