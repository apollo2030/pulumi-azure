# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    data_ingestion_uri: pulumi.Output[str]
    """
    The Kusto Cluster URI to be used for data ingestion.
    """
    location: pulumi.Output[str]
    """
    The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
    """
    sku: pulumi.Output[dict]
    """
    A `sku` block as defined below.
    
      * `capacity` (`float`)
      * `name` (`str`) - The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    uri: pulumi.Output[str]
    """
    The FQDN of the Azure Kusto Cluster.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Kusto (also known as Azure Data Explorer) Cluster
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] sku: A `sku` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        
        The **sku** object supports the following:
        
          * `capacity` (`pulumi.Input[float]`)
          * `name` (`pulumi.Input[str]`) - The name of the Kusto Cluster to create. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kusto_cluster.html.markdown.
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

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['data_ingestion_uri'] = None
            __props__['uri'] = None
        super(Cluster, __self__).__init__(
            'azure:kusto/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data_ingestion_uri=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, uri=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_ingestion_uri: The Kusto Cluster URI to be used for data ingestion.
        :param pulumi.Input[str] location: The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] sku: A `sku` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] uri: The FQDN of the Azure Kusto Cluster.
        
        The **sku** object supports the following:
        
          * `capacity` (`pulumi.Input[float]`)
          * `name` (`pulumi.Input[str]`) - The name of the Kusto Cluster to create. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kusto_cluster.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["data_ingestion_uri"] = data_ingestion_uri
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku"] = sku
        __props__["tags"] = tags
        __props__["uri"] = uri
        return Cluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

