# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNetworkDdosProtectionPlanResult:
    """
    A collection of values returned by getNetworkDdosProtectionPlan.
    """
    def __init__(__self__, location=None, name=None, resource_group_name=None, tags=None, virtual_network_ids=None, id=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Specifies the supported Azure location where the resource exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if virtual_network_ids and not isinstance(virtual_network_ids, list):
            raise TypeError("Expected argument 'virtual_network_ids' to be a list")
        __self__.virtual_network_ids = virtual_network_ids
        """
        The Resource ID list of the Virtual Networks associated with DDoS Protection Plan.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetNetworkDdosProtectionPlanResult(GetNetworkDdosProtectionPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkDdosProtectionPlanResult(
            location=self.location,
            name=self.name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            virtual_network_ids=self.virtual_network_ids,
            id=self.id)

def get_network_ddos_protection_plan(name=None,resource_group_name=None,tags=None,opts=None):
    """
    Use this data source to access information about an existing Azure Network DDoS Protection Plan.
    
    :param str name: The name of the Network DDoS Protection Plan.
    :param str resource_group_name: The name of the resource group where the Network DDoS Protection Plan exists.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/network_ddos_protection_plan.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:network/getNetworkDdosProtectionPlan:getNetworkDdosProtectionPlan', __args__, opts=opts).value

    return AwaitableGetNetworkDdosProtectionPlanResult(
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        virtual_network_ids=__ret__.get('virtualNetworkIds'),
        id=__ret__.get('id'))
