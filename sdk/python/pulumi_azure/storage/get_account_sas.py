# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetAccountSASResult(object):
    """
    A collection of values returned by getAccountSAS.
    """
    def __init__(__self__, sas=None, id=None):
        if sas and not isinstance(sas, basestring):
            raise TypeError('Expected argument sas to be a basestring')
        __self__.sas = sas
        """
        The computed Account Shared Access Signature (SAS). 
        """
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_account_sas(connection_string=None, expiry=None, https_only=None, permissions=None, resource_types=None, services=None, start=None):
    """
    Use this data source to create a Shared Access Signature (SAS) for an Azure Storage Account.
    
    Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account.
    
    Note that this is an [Account SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-an-account-sas)
    and *not* a [Service SAS](https://docs.microsoft.com/en-us/rest/api/storageservices/constructing-a-service-sas).
    """
    __args__ = dict()

    __args__['connectionString'] = connection_string
    __args__['expiry'] = expiry
    __args__['httpsOnly'] = https_only
    __args__['permissions'] = permissions
    __args__['resourceTypes'] = resource_types
    __args__['services'] = services
    __args__['start'] = start
    __ret__ = pulumi.runtime.invoke('azure:storage/getAccountSAS:getAccountSAS', __args__)

    return GetAccountSASResult(
        sas=__ret__.get('sas'),
        id=__ret__.get('id'))
