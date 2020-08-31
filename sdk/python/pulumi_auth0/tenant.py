# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Tenant']


class Tenant(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_logout_urls: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 change_password: Optional[pulumi.Input[pulumi.InputType['TenantChangePasswordArgs']]] = None,
                 default_audience: Optional[pulumi.Input[str]] = None,
                 default_directory: Optional[pulumi.Input[str]] = None,
                 default_redirection_uri: Optional[pulumi.Input[str]] = None,
                 enabled_locales: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 error_page: Optional[pulumi.Input[pulumi.InputType['TenantErrorPageArgs']]] = None,
                 flags: Optional[pulumi.Input[pulumi.InputType['TenantFlagsArgs']]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 guardian_mfa_page: Optional[pulumi.Input[pulumi.InputType['TenantGuardianMfaPageArgs']]] = None,
                 idle_session_lifetime: Optional[pulumi.Input[float]] = None,
                 picture_url: Optional[pulumi.Input[str]] = None,
                 sandbox_version: Optional[pulumi.Input[str]] = None,
                 session_lifetime: Optional[pulumi.Input[float]] = None,
                 support_email: Optional[pulumi.Input[str]] = None,
                 support_url: Optional[pulumi.Input[str]] = None,
                 universal_login: Optional[pulumi.Input[pulumi.InputType['TenantUniversalLoginArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        With this resource, you can manage Auth0 tenants, including setting logos and support contact information, setting error pages, and configuring default tenant behaviors.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        tenant = auth0.Tenant("tenant",
            allowed_logout_urls=["http://mysite/logout"],
            change_password=auth0.TenantChangePasswordArgs(
                enabled=True,
                html=(lambda path: open(path).read())("./password_reset.html"),
            ),
            default_audience="<client_id>",
            default_directory="Connection-Name",
            error_page=auth0.TenantErrorPageArgs(
                html=(lambda path: open(path).read())("./error.html"),
                show_log_link=True,
                url="http://mysite/errors",
            ),
            friendly_name="Tenant Name",
            guardian_mfa_page=auth0.TenantGuardianMfaPageArgs(
                enabled=True,
                html=(lambda path: open(path).read())("./guardian_multifactor.html"),
            ),
            picture_url="http://mysite/logo.png",
            sandbox_version="8",
            session_lifetime=46000,
            support_email="support@mysite",
            support_url="http://mysite/support")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_logout_urls: List(String). URLs that Auth0 may redirect to after logout.
        :param pulumi.Input[pulumi.InputType['TenantChangePasswordArgs']] change_password: List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
        :param pulumi.Input[str] default_audience: String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
        :param pulumi.Input[str] default_directory: String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
        :param pulumi.Input[str] default_redirection_uri: String. The default absolute redirection uri, must be https and cannot contain a fragment.
        :param pulumi.Input[pulumi.InputType['TenantErrorPageArgs']] error_page: List(Resource). Configuration settings for error pages. For details, see Error Page.
        :param pulumi.Input[pulumi.InputType['TenantFlagsArgs']] flags: List(Resource). Configuration settings for tenant flags. For details, see Flags.
        :param pulumi.Input[str] friendly_name: String. Friendly name for the tenant.
        :param pulumi.Input[pulumi.InputType['TenantGuardianMfaPageArgs']] guardian_mfa_page: List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
        :param pulumi.Input[float] idle_session_lifetime: Integer. Number of hours during which a session can be inactive before the user must log in again.
        :param pulumi.Input[str] picture_url: . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
        :param pulumi.Input[str] sandbox_version: String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
        :param pulumi.Input[float] session_lifetime: Integer. Number of hours during which a session will stay valid.
        :param pulumi.Input[str] support_email: String. Support email address for authenticating users.
        :param pulumi.Input[str] support_url: String. Support URL for authenticating users.
        :param pulumi.Input[pulumi.InputType['TenantUniversalLoginArgs']] universal_login: List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['allowed_logout_urls'] = allowed_logout_urls
            __props__['change_password'] = change_password
            __props__['default_audience'] = default_audience
            __props__['default_directory'] = default_directory
            __props__['default_redirection_uri'] = default_redirection_uri
            __props__['enabled_locales'] = enabled_locales
            __props__['error_page'] = error_page
            __props__['flags'] = flags
            __props__['friendly_name'] = friendly_name
            __props__['guardian_mfa_page'] = guardian_mfa_page
            __props__['idle_session_lifetime'] = idle_session_lifetime
            __props__['picture_url'] = picture_url
            __props__['sandbox_version'] = sandbox_version
            __props__['session_lifetime'] = session_lifetime
            __props__['support_email'] = support_email
            __props__['support_url'] = support_url
            __props__['universal_login'] = universal_login
        super(Tenant, __self__).__init__(
            'auth0:index/tenant:Tenant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_logout_urls: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            change_password: Optional[pulumi.Input[pulumi.InputType['TenantChangePasswordArgs']]] = None,
            default_audience: Optional[pulumi.Input[str]] = None,
            default_directory: Optional[pulumi.Input[str]] = None,
            default_redirection_uri: Optional[pulumi.Input[str]] = None,
            enabled_locales: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            error_page: Optional[pulumi.Input[pulumi.InputType['TenantErrorPageArgs']]] = None,
            flags: Optional[pulumi.Input[pulumi.InputType['TenantFlagsArgs']]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            guardian_mfa_page: Optional[pulumi.Input[pulumi.InputType['TenantGuardianMfaPageArgs']]] = None,
            idle_session_lifetime: Optional[pulumi.Input[float]] = None,
            picture_url: Optional[pulumi.Input[str]] = None,
            sandbox_version: Optional[pulumi.Input[str]] = None,
            session_lifetime: Optional[pulumi.Input[float]] = None,
            support_email: Optional[pulumi.Input[str]] = None,
            support_url: Optional[pulumi.Input[str]] = None,
            universal_login: Optional[pulumi.Input[pulumi.InputType['TenantUniversalLoginArgs']]] = None) -> 'Tenant':
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_logout_urls: List(String). URLs that Auth0 may redirect to after logout.
        :param pulumi.Input[pulumi.InputType['TenantChangePasswordArgs']] change_password: List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
        :param pulumi.Input[str] default_audience: String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
        :param pulumi.Input[str] default_directory: String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
        :param pulumi.Input[str] default_redirection_uri: String. The default absolute redirection uri, must be https and cannot contain a fragment.
        :param pulumi.Input[pulumi.InputType['TenantErrorPageArgs']] error_page: List(Resource). Configuration settings for error pages. For details, see Error Page.
        :param pulumi.Input[pulumi.InputType['TenantFlagsArgs']] flags: List(Resource). Configuration settings for tenant flags. For details, see Flags.
        :param pulumi.Input[str] friendly_name: String. Friendly name for the tenant.
        :param pulumi.Input[pulumi.InputType['TenantGuardianMfaPageArgs']] guardian_mfa_page: List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
        :param pulumi.Input[float] idle_session_lifetime: Integer. Number of hours during which a session can be inactive before the user must log in again.
        :param pulumi.Input[str] picture_url: . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
        :param pulumi.Input[str] sandbox_version: String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
        :param pulumi.Input[float] session_lifetime: Integer. Number of hours during which a session will stay valid.
        :param pulumi.Input[str] support_email: String. Support email address for authenticating users.
        :param pulumi.Input[str] support_url: String. Support URL for authenticating users.
        :param pulumi.Input[pulumi.InputType['TenantUniversalLoginArgs']] universal_login: List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_logout_urls"] = allowed_logout_urls
        __props__["change_password"] = change_password
        __props__["default_audience"] = default_audience
        __props__["default_directory"] = default_directory
        __props__["default_redirection_uri"] = default_redirection_uri
        __props__["enabled_locales"] = enabled_locales
        __props__["error_page"] = error_page
        __props__["flags"] = flags
        __props__["friendly_name"] = friendly_name
        __props__["guardian_mfa_page"] = guardian_mfa_page
        __props__["idle_session_lifetime"] = idle_session_lifetime
        __props__["picture_url"] = picture_url
        __props__["sandbox_version"] = sandbox_version
        __props__["session_lifetime"] = session_lifetime
        __props__["support_email"] = support_email
        __props__["support_url"] = support_url
        __props__["universal_login"] = universal_login
        return Tenant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedLogoutUrls")
    def allowed_logout_urls(self) -> pulumi.Output[List[str]]:
        """
        List(String). URLs that Auth0 may redirect to after logout.
        """
        return pulumi.get(self, "allowed_logout_urls")

    @property
    @pulumi.getter(name="changePassword")
    def change_password(self) -> pulumi.Output['outputs.TenantChangePassword']:
        """
        List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
        """
        return pulumi.get(self, "change_password")

    @property
    @pulumi.getter(name="defaultAudience")
    def default_audience(self) -> pulumi.Output[str]:
        """
        String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
        """
        return pulumi.get(self, "default_audience")

    @property
    @pulumi.getter(name="defaultDirectory")
    def default_directory(self) -> pulumi.Output[str]:
        """
        String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
        """
        return pulumi.get(self, "default_directory")

    @property
    @pulumi.getter(name="defaultRedirectionUri")
    def default_redirection_uri(self) -> pulumi.Output[str]:
        """
        String. The default absolute redirection uri, must be https and cannot contain a fragment.
        """
        return pulumi.get(self, "default_redirection_uri")

    @property
    @pulumi.getter(name="enabledLocales")
    def enabled_locales(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "enabled_locales")

    @property
    @pulumi.getter(name="errorPage")
    def error_page(self) -> pulumi.Output['outputs.TenantErrorPage']:
        """
        List(Resource). Configuration settings for error pages. For details, see Error Page.
        """
        return pulumi.get(self, "error_page")

    @property
    @pulumi.getter
    def flags(self) -> pulumi.Output['outputs.TenantFlags']:
        """
        List(Resource). Configuration settings for tenant flags. For details, see Flags.
        """
        return pulumi.get(self, "flags")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[str]:
        """
        String. Friendly name for the tenant.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="guardianMfaPage")
    def guardian_mfa_page(self) -> pulumi.Output['outputs.TenantGuardianMfaPage']:
        """
        List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
        """
        return pulumi.get(self, "guardian_mfa_page")

    @property
    @pulumi.getter(name="idleSessionLifetime")
    def idle_session_lifetime(self) -> pulumi.Output[float]:
        """
        Integer. Number of hours during which a session can be inactive before the user must log in again.
        """
        return pulumi.get(self, "idle_session_lifetime")

    @property
    @pulumi.getter(name="pictureUrl")
    def picture_url(self) -> pulumi.Output[str]:
        """
        . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
        """
        return pulumi.get(self, "picture_url")

    @property
    @pulumi.getter(name="sandboxVersion")
    def sandbox_version(self) -> pulumi.Output[str]:
        """
        String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
        """
        return pulumi.get(self, "sandbox_version")

    @property
    @pulumi.getter(name="sessionLifetime")
    def session_lifetime(self) -> pulumi.Output[float]:
        """
        Integer. Number of hours during which a session will stay valid.
        """
        return pulumi.get(self, "session_lifetime")

    @property
    @pulumi.getter(name="supportEmail")
    def support_email(self) -> pulumi.Output[str]:
        """
        String. Support email address for authenticating users.
        """
        return pulumi.get(self, "support_email")

    @property
    @pulumi.getter(name="supportUrl")
    def support_url(self) -> pulumi.Output[str]:
        """
        String. Support URL for authenticating users.
        """
        return pulumi.get(self, "support_url")

    @property
    @pulumi.getter(name="universalLogin")
    def universal_login(self) -> pulumi.Output['outputs.TenantUniversalLogin']:
        """
        List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
        """
        return pulumi.get(self, "universal_login")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

