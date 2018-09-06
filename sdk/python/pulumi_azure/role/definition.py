# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Definition(pulumi.CustomResource):
    """
    Manages a custom Role Definition, used to assign Roles to Users/Principals. See ['Understand role definitions'](https://docs.microsoft.com/en-us/azure/role-based-access-control/role-definitions) in the Azure documentation for more details.
    """
    def __init__(__self__, __name__, __opts__=None, assignable_scopes=None, description=None, name=None, permissions=None, role_definition_id=None, scope=None):
        """Create a Definition resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not assignable_scopes:
            raise TypeError('Missing required property assignable_scopes')
        elif not isinstance(assignable_scopes, list):
            raise TypeError('Expected property assignable_scopes to be a list')
        __self__.assignable_scopes = assignable_scopes
        """
        One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
        """
        __props__['assignableScopes'] = assignable_scopes

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A description of the Role Definition.
        """
        __props__['description'] = description

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Role Definition. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not permissions:
            raise TypeError('Missing required property permissions')
        elif not isinstance(permissions, list):
            raise TypeError('Expected property permissions to be a list')
        __self__.permissions = permissions
        """
        A `permissions` block as defined below.
        """
        __props__['permissions'] = permissions

        if role_definition_id and not isinstance(role_definition_id, basestring):
            raise TypeError('Expected property role_definition_id to be a basestring')
        __self__.role_definition_id = role_definition_id
        """
        A unique UUID/GUID which identifies this role - one will be generated if not specified. Changing this forces a new resource to be created.
        """
        __props__['roleDefinitionId'] = role_definition_id

        if not scope:
            raise TypeError('Missing required property scope')
        elif not isinstance(scope, basestring):
            raise TypeError('Expected property scope to be a basestring')
        __self__.scope = scope
        """
        The scope at which the Role Definition applies too, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`. Changing this forces a new resource to be created.
        """
        __props__['scope'] = scope

        super(Definition, __self__).__init__(
            'azure:role/definition:Definition',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'assignableScopes' in outs:
            self.assignable_scopes = outs['assignableScopes']
        if 'description' in outs:
            self.description = outs['description']
        if 'name' in outs:
            self.name = outs['name']
        if 'permissions' in outs:
            self.permissions = outs['permissions']
        if 'roleDefinitionId' in outs:
            self.role_definition_id = outs['roleDefinitionId']
        if 'scope' in outs:
            self.scope = outs['scope']
