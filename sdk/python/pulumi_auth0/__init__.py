# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .client import *
from .client_grant import *
from .connection import *
from .custom_domain import *
from .email import *
from .email_template import *
from .global_client import *
from .hook import *
from .log_stream import *
from .prompt import *
from .provider import *
from .resource_server import *
from .role import *
from .rule import *
from .rule_config import *
from .tenant import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
from . import (
    config,
)

def _register_module():
    import pulumi
    from . import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "auth0:index/client:Client":
                return Client(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/clientGrant:ClientGrant":
                return ClientGrant(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/connection:Connection":
                return Connection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/customDomain:CustomDomain":
                return CustomDomain(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/email:Email":
                return Email(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/emailTemplate:EmailTemplate":
                return EmailTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/globalClient:GlobalClient":
                return GlobalClient(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/hook:Hook":
                return Hook(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/logStream:LogStream":
                return LogStream(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/prompt:Prompt":
                return Prompt(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/resourceServer:ResourceServer":
                return ResourceServer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/role:Role":
                return Role(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/rule:Rule":
                return Rule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/ruleConfig:RuleConfig":
                return RuleConfig(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/tenant:Tenant":
                return Tenant(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "auth0:index/user:User":
                return User(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("auth0", "index/client", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/clientGrant", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/connection", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/customDomain", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/email", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/emailTemplate", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/globalClient", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/hook", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/logStream", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/prompt", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/resourceServer", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/role", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/rule", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/ruleConfig", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/tenant", _module_instance)
    pulumi.runtime.register_resource_module("auth0", "index/user", _module_instance)


    class Package(pulumi.runtime.ResourcePackage):
        _version = _utilities.get_semver_version()

        def version(self):
            return Package._version

        def construct_provider(self, name: str, typ: str, urn: str) -> pulumi.ProviderResource:
            if typ != "pulumi:providers:auth0":
                raise Exception(f"unknown provider type {typ}")
            return Provider(name, pulumi.ResourceOptions(urn=urn))


    pulumi.runtime.register_resource_package("auth0", Package())

_register_module()
