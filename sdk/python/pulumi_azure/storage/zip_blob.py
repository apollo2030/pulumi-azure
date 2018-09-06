# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ZipBlob(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, attempts=None, content_type=None, name=None, parallelism=None, resource_group_name=None, size=None, content=None, source_uri=None, storage_account_name=None, storage_container_name=None, type=None):
        """Create a ZipBlob resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if attempts and not isinstance(attempts, int):
            raise TypeError('Expected property attempts to be a int')
        __self__.attempts = attempts
        __props__['attempts'] = attempts

        if content_type and not isinstance(content_type, basestring):
            raise TypeError('Expected property content_type to be a basestring')
        __self__.content_type = content_type
        __props__['contentType'] = content_type

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        __props__['name'] = name

        if parallelism and not isinstance(parallelism, int):
            raise TypeError('Expected property parallelism to be a int')
        __self__.parallelism = parallelism
        __props__['parallelism'] = parallelism

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        __props__['resourceGroupName'] = resource_group_name

        if size and not isinstance(size, int):
            raise TypeError('Expected property size to be a int')
        __self__.size = size
        __props__['size'] = size

        if content and not isinstance(content, pulumi.Archive):
            raise TypeError('Expected property content to be a pulumi.Archive')
        __self__.content = content
        __props__['content'] = content

        if source_uri and not isinstance(source_uri, basestring):
            raise TypeError('Expected property source_uri to be a basestring')
        __self__.source_uri = source_uri
        __props__['sourceUri'] = source_uri

        if not storage_account_name:
            raise TypeError('Missing required property storage_account_name')
        elif not isinstance(storage_account_name, basestring):
            raise TypeError('Expected property storage_account_name to be a basestring')
        __self__.storage_account_name = storage_account_name
        __props__['storageAccountName'] = storage_account_name

        if not storage_container_name:
            raise TypeError('Missing required property storage_container_name')
        elif not isinstance(storage_container_name, basestring):
            raise TypeError('Expected property storage_container_name to be a basestring')
        __self__.storage_container_name = storage_container_name
        __props__['storageContainerName'] = storage_container_name

        if type and not isinstance(type, basestring):
            raise TypeError('Expected property type to be a basestring')
        __self__.type = type
        __props__['type'] = type

        __self__.url = pulumi.runtime.UNKNOWN

        super(ZipBlob, __self__).__init__(
            'azure:storage/zipBlob:ZipBlob',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'attempts' in outs:
            self.attempts = outs['attempts']
        if 'contentType' in outs:
            self.content_type = outs['contentType']
        if 'name' in outs:
            self.name = outs['name']
        if 'parallelism' in outs:
            self.parallelism = outs['parallelism']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'size' in outs:
            self.size = outs['size']
        if 'content' in outs:
            self.content = outs['content']
        if 'sourceUri' in outs:
            self.source_uri = outs['sourceUri']
        if 'storageAccountName' in outs:
            self.storage_account_name = outs['storageAccountName']
        if 'storageContainerName' in outs:
            self.storage_container_name = outs['storageContainerName']
        if 'type' in outs:
            self.type = outs['type']
        if 'url' in outs:
            self.url = outs['url']
