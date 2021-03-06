# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Backend(pulumi.CustomResource):
    api_management_name: pulumi.Output[str]
    """
    The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
    """
    credentials: pulumi.Output[dict]
    """
    A `credentials` block as documented below.
    
      * `authorization` (`dict`)
    
        * `parameter` (`str`)
        * `scheme` (`str`)
    
      * `certificates` (`list`)
      * `header` (`dict`)
      * `query` (`dict`)
    """
    description: pulumi.Output[str]
    """
    The description of the backend.
    """
    name: pulumi.Output[str]
    """
    The name of the API Management backend. Changing this forces a new resource to be created.
    """
    protocol: pulumi.Output[str]
    """
    The protocol used by the backend host. Possible values are `http` or `soap`.
    """
    proxy: pulumi.Output[dict]
    """
    A `proxy` block as documented below.
    
      * `password` (`str`)
      * `url` (`str`) - The URL of the backend host.
      * `username` (`str`)
    """
    resource_group_name: pulumi.Output[str]
    """
    The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
    """
    resource_id: pulumi.Output[str]
    """
    The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
    """
    service_fabric_cluster: pulumi.Output[dict]
    """
    A `service_fabric_cluster` block as documented below.
    
      * `clientCertificateThumbprint` (`str`)
      * `managementEndpoints` (`list`)
      * `maxPartitionResolutionRetries` (`float`)
      * `serverCertificateThumbprints` (`list`)
      * `serverX509Names` (`list`)
    
        * `issuerCertificateThumbprint` (`str`)
        * `name` (`str`) - The name of the API Management backend. Changing this forces a new resource to be created.
    """
    title: pulumi.Output[str]
    """
    The title of the backend.
    """
    tls: pulumi.Output[dict]
    """
    A `tls` block as documented below.
    
      * `validateCertificateChain` (`bool`)
      * `validateCertificateName` (`bool`)
    """
    url: pulumi.Output[str]
    """
    The URL of the backend host.
    """
    def __init__(__self__, resource_name, opts=None, api_management_name=None, credentials=None, description=None, name=None, protocol=None, proxy=None, resource_group_name=None, resource_id=None, service_fabric_cluster=None, title=None, tls=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a backend within an API Management Service.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] credentials: A `credentials` block as documented below.
        :param pulumi.Input[str] description: The description of the backend.
        :param pulumi.Input[str] name: The name of the API Management backend. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The protocol used by the backend host. Possible values are `http` or `soap`.
        :param pulumi.Input[dict] proxy: A `proxy` block as documented below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_id: The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
        :param pulumi.Input[dict] service_fabric_cluster: A `service_fabric_cluster` block as documented below.
        :param pulumi.Input[str] title: The title of the backend.
        :param pulumi.Input[dict] tls: A `tls` block as documented below.
        :param pulumi.Input[str] url: The URL of the backend host.
        
        The **credentials** object supports the following:
        
          * `authorization` (`pulumi.Input[dict]`)
        
            * `parameter` (`pulumi.Input[str]`)
            * `scheme` (`pulumi.Input[str]`)
        
          * `certificates` (`pulumi.Input[list]`)
          * `header` (`pulumi.Input[dict]`)
          * `query` (`pulumi.Input[dict]`)
        
        The **proxy** object supports the following:
        
          * `password` (`pulumi.Input[str]`)
          * `url` (`pulumi.Input[str]`) - The URL of the backend host.
          * `username` (`pulumi.Input[str]`)
        
        The **service_fabric_cluster** object supports the following:
        
          * `clientCertificateThumbprint` (`pulumi.Input[str]`)
          * `managementEndpoints` (`pulumi.Input[list]`)
          * `maxPartitionResolutionRetries` (`pulumi.Input[float]`)
          * `serverCertificateThumbprints` (`pulumi.Input[list]`)
          * `serverX509Names` (`pulumi.Input[list]`)
        
            * `issuerCertificateThumbprint` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`) - The name of the API Management backend. Changing this forces a new resource to be created.
        
        The **tls** object supports the following:
        
          * `validateCertificateChain` (`pulumi.Input[bool]`)
          * `validateCertificateName` (`pulumi.Input[bool]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_backend.html.markdown.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            __props__['credentials'] = credentials
            __props__['description'] = description
            __props__['name'] = name
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            __props__['proxy'] = proxy
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_id'] = resource_id
            __props__['service_fabric_cluster'] = service_fabric_cluster
            __props__['title'] = title
            __props__['tls'] = tls
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
        super(Backend, __self__).__init__(
            'azure:apimanagement/backend:Backend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_management_name=None, credentials=None, description=None, name=None, protocol=None, proxy=None, resource_group_name=None, resource_id=None, service_fabric_cluster=None, title=None, tls=None, url=None):
        """
        Get an existing Backend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The Name of the API Management Service where this backend should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] credentials: A `credentials` block as documented below.
        :param pulumi.Input[str] description: The description of the backend.
        :param pulumi.Input[str] name: The name of the API Management backend. Changing this forces a new resource to be created.
        :param pulumi.Input[str] protocol: The protocol used by the backend host. Possible values are `http` or `soap`.
        :param pulumi.Input[dict] proxy: A `proxy` block as documented below.
        :param pulumi.Input[str] resource_group_name: The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_id: The management URI of the backend host in an external system. This URI can be the ARM Resource ID of Logic Apps, Function Apps or API Apps, or the management endpoint of a Service Fabric cluster.
        :param pulumi.Input[dict] service_fabric_cluster: A `service_fabric_cluster` block as documented below.
        :param pulumi.Input[str] title: The title of the backend.
        :param pulumi.Input[dict] tls: A `tls` block as documented below.
        :param pulumi.Input[str] url: The URL of the backend host.
        
        The **credentials** object supports the following:
        
          * `authorization` (`pulumi.Input[dict]`)
        
            * `parameter` (`pulumi.Input[str]`)
            * `scheme` (`pulumi.Input[str]`)
        
          * `certificates` (`pulumi.Input[list]`)
          * `header` (`pulumi.Input[dict]`)
          * `query` (`pulumi.Input[dict]`)
        
        The **proxy** object supports the following:
        
          * `password` (`pulumi.Input[str]`)
          * `url` (`pulumi.Input[str]`) - The URL of the backend host.
          * `username` (`pulumi.Input[str]`)
        
        The **service_fabric_cluster** object supports the following:
        
          * `clientCertificateThumbprint` (`pulumi.Input[str]`)
          * `managementEndpoints` (`pulumi.Input[list]`)
          * `maxPartitionResolutionRetries` (`pulumi.Input[float]`)
          * `serverCertificateThumbprints` (`pulumi.Input[list]`)
          * `serverX509Names` (`pulumi.Input[list]`)
        
            * `issuerCertificateThumbprint` (`pulumi.Input[str]`)
            * `name` (`pulumi.Input[str]`) - The name of the API Management backend. Changing this forces a new resource to be created.
        
        The **tls** object supports the following:
        
          * `validateCertificateChain` (`pulumi.Input[bool]`)
          * `validateCertificateName` (`pulumi.Input[bool]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_backend.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["api_management_name"] = api_management_name
        __props__["credentials"] = credentials
        __props__["description"] = description
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["proxy"] = proxy
        __props__["resource_group_name"] = resource_group_name
        __props__["resource_id"] = resource_id
        __props__["service_fabric_cluster"] = service_fabric_cluster
        __props__["title"] = title
        __props__["tls"] = tls
        __props__["url"] = url
        return Backend(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

