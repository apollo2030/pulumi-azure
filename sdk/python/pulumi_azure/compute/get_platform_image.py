# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetPlatformImageResult(object):
    """
    A collection of values returned by getPlatformImage.
    """
    def __init__(__self__, version=None, id=None):
        if version and not isinstance(version, basestring):
            raise TypeError('Expected argument version to be a basestring')
        __self__.version = version
        """
        The latest version of the Platform Image.
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_platform_image(location=None, offer=None, publisher=None, sku=None):
    """
    Use this data source to access the properties of an Azure Platform Image.
    """
    __args__ = dict()

    __args__['location'] = location
    __args__['offer'] = offer
    __args__['publisher'] = publisher
    __args__['sku'] = sku
    __ret__ = pulumi.runtime.invoke('azure:compute/getPlatformImage:getPlatformImage', __args__)

    return GetPlatformImageResult(
        version=__ret__.get('version'),
        id=__ret__.get('id'))
