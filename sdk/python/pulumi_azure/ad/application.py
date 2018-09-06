# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Application(pulumi.CustomResource):
    """
    Manages an Application within Azure Active Directory.
    
    -> **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
    """
    def __init__(__self__, __name__, __opts__=None, available_to_other_tenants=None, homepage=None, identifier_uris=None, name=None, oauth2_allow_implicit_flow=None, reply_urls=None):
        """Create a Application resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if available_to_other_tenants and not isinstance(available_to_other_tenants, bool):
            raise TypeError('Expected property available_to_other_tenants to be a bool')
        __self__.available_to_other_tenants = available_to_other_tenants
        """
        Is this Azure AD Application available to other tenants? Defaults to `false`.
        """
        __props__['availableToOtherTenants'] = available_to_other_tenants

        if homepage and not isinstance(homepage, basestring):
            raise TypeError('Expected property homepage to be a basestring')
        __self__.homepage = homepage
        """
        The URL to the application's home page. If no homepage is specified this defaults to `http://{name}`.
        """
        __props__['homepage'] = homepage

        if identifier_uris and not isinstance(identifier_uris, list):
            raise TypeError('Expected property identifier_uris to be a list')
        __self__.identifier_uris = identifier_uris
        """
        A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        """
        __props__['identifierUris'] = identifier_uris

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The display name for the application.
        """
        __props__['name'] = name

        if oauth2_allow_implicit_flow and not isinstance(oauth2_allow_implicit_flow, bool):
            raise TypeError('Expected property oauth2_allow_implicit_flow to be a bool')
        __self__.oauth2_allow_implicit_flow = oauth2_allow_implicit_flow
        """
        Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
        """
        __props__['oauth2AllowImplicitFlow'] = oauth2_allow_implicit_flow

        if reply_urls and not isinstance(reply_urls, list):
            raise TypeError('Expected property reply_urls to be a list')
        __self__.reply_urls = reply_urls
        """
        A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
        """
        __props__['replyUrls'] = reply_urls

        __self__.application_id = pulumi.runtime.UNKNOWN
        """
        The Application ID.
        """

        super(Application, __self__).__init__(
            'azure:ad/application:Application',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'applicationId' in outs:
            self.application_id = outs['applicationId']
        if 'availableToOtherTenants' in outs:
            self.available_to_other_tenants = outs['availableToOtherTenants']
        if 'homepage' in outs:
            self.homepage = outs['homepage']
        if 'identifierUris' in outs:
            self.identifier_uris = outs['identifierUris']
        if 'name' in outs:
            self.name = outs['name']
        if 'oauth2AllowImplicitFlow' in outs:
            self.oauth2_allow_implicit_flow = outs['oauth2AllowImplicitFlow']
        if 'replyUrls' in outs:
            self.reply_urls = outs['replyUrls']
