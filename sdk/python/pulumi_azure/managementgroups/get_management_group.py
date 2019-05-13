# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetManagementGroupResult:
    """
    A collection of values returned by getManagementGroup.
    """
    def __init__(__self__, display_name=None, group_id=None, parent_management_group_id=None, subscription_ids=None, id=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        __self__.display_name = display_name
        """
        A friendly name for the Management Group.
        """
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        __self__.group_id = group_id
        if parent_management_group_id and not isinstance(parent_management_group_id, str):
            raise TypeError("Expected argument 'parent_management_group_id' to be a str")
        __self__.parent_management_group_id = parent_management_group_id
        """
        The ID of any Parent Management Group.
        """
        if subscription_ids and not isinstance(subscription_ids, list):
            raise TypeError("Expected argument 'subscription_ids' to be a list")
        __self__.subscription_ids = subscription_ids
        """
        A list of Subscription ID's which are assigned to the Management Group.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_management_group(group_id=None,opts=None):
    """
    Use this data source to access information about an existing Management Group.
    """
    __args__ = dict()

    __args__['groupId'] = group_id
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('azure:managementgroups/getManagementGroup:getManagementGroup', __args__, opts=opts)

    return GetManagementGroupResult(
        display_name=__ret__.get('displayName'),
        group_id=__ret__.get('groupId'),
        parent_management_group_id=__ret__.get('parentManagementGroupId'),
        subscription_ids=__ret__.get('subscriptionIds'),
        id=__ret__.get('id'))
