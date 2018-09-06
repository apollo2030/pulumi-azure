# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Database(pulumi.CustomResource):
    """
    Manages a MySQL Database within a MySQL Server
    """
    def __init__(__self__, __name__, __opts__=None, charset=None, collation=None, name=None, resource_group_name=None, server_name=None):
        """Create a Database resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not charset:
            raise TypeError('Missing required property charset')
        elif not isinstance(charset, basestring):
            raise TypeError('Expected property charset to be a basestring')
        __self__.charset = charset
        """
        Specifies the Charset for the MySQL Database, which needs [to be a valid MySQL Charset](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html). Changing this forces a new resource to be created.
        """
        __props__['charset'] = charset

        if not collation:
            raise TypeError('Missing required property collation')
        elif not isinstance(collation, basestring):
            raise TypeError('Expected property collation to be a basestring')
        __self__.collation = collation
        """
        Specifies the Collation for the MySQL Database, which needs [to be a valid MySQL Collation](https://dev.mysql.com/doc/refman/5.7/en/charset-mysql.html). Changing this forces a new resource to be created.
        """
        __props__['collation'] = collation

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the MySQL Database, which needs [to be a valid MySQL identifier](https://dev.mysql.com/doc/refman/5.7/en/identifiers.html). Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which the MySQL Server exists. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not server_name:
            raise TypeError('Missing required property server_name')
        elif not isinstance(server_name, basestring):
            raise TypeError('Expected property server_name to be a basestring')
        __self__.server_name = server_name
        """
        Specifies the name of the MySQL Server. Changing this forces a new resource to be created.
        """
        __props__['serverName'] = server_name

        super(Database, __self__).__init__(
            'azure:mysql/database:Database',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'charset' in outs:
            self.charset = outs['charset']
        if 'collation' in outs:
            self.collation = outs['collation']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'serverName' in outs:
            self.server_name = outs['serverName']
