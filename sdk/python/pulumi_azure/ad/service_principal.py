# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServicePrincipal(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The ID of the Azure AD Application for which to create a Service Principal.
    """
    display_name: pulumi.Output[str]
    """
    The Display Name of the Azure Active Directory Application associated with this Service Principal.
    """
    def __init__(__self__, resource_name, opts=None, application_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ServicePrincipal resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The ID of the Azure AD Application for which to create a Service Principal.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/azuread_service_principal.html.markdown.
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

            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['display_name'] = None
        super(ServicePrincipal, __self__).__init__(
            'azure:ad/servicePrincipal:ServicePrincipal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, application_id=None, display_name=None):
        """
        Get an existing ServicePrincipal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The ID of the Azure AD Application for which to create a Service Principal.
        :param pulumi.Input[str] display_name: The Display Name of the Azure Active Directory Application associated with this Service Principal.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/azuread_service_principal.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["application_id"] = application_id
        __props__["display_name"] = display_name
        return ServicePrincipal(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

