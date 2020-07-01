# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class Tenant(pulumi.CustomResource):
    allowed_logout_urls: pulumi.Output[list]
    """
    List(String). URLs that Auth0 may redirect to after logout.
    """
    change_password: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.

      * `enabled` (`bool`) - Boolean. Indicates whether or not to use the custom change password page.
      * `html` (`str`) - String, HTML format with supported Liquid syntax. Customized content of the change password page.
    """
    default_audience: pulumi.Output[str]
    """
    String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
    """
    default_directory: pulumi.Output[str]
    """
    String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
    """
    default_redirection_uri: pulumi.Output[str]
    """
    String. The default absolute redirection uri, must be https and cannot contain a fragment.
    """
    enabled_locales: pulumi.Output[list]
    error_page: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for error pages. For details, see Error Page.

      * `html` (`str`) - String, HTML format with supported Liquid syntax. Customized content of the error page.
      * `showLogLink` (`bool`) - Boolean. Indicates whether or not to show the link to logs as part of the default error page.
      * `url` (`str`) - String. URL to redirect to when an error occurs rather than showing the default error page.
    """
    flags: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for tenant flags. For details, see Flags.

      * `changePwdFlowV1` (`bool`) - Boolean. Indicates whether or not to use the older v1 change password flow. Not recommended except for backward compatibility.
      * `disableClickjackProtectionHeaders` (`bool`) - Boolean. Indicated whether or not classic Universal Login prompts include additional security headers to prevent clickjacking.
      * `enableApisSection` (`bool`) - Boolean. Indicates whether or not the APIs section is enabled for the tenant.
      * `enableClientConnections` (`bool`) - Boolean. Indicates whether or not all current connections should be enabled when a new client is created.
      * `enableCustomDomainInEmails` (`bool`) - Boolean. Indicates whether or not the tenant allows custom domains in emails.
      * `enableDynamicClientRegistration` (`bool`) - Boolean. Indicates whether or not the tenant allows dynamic client registration.
      * `enableLegacyLogsSearchV2` (`bool`) - Boolean. Indicates whether or not to use the older v2 legacy logs search.
      * `enablePipeline2` (`bool`) - Boolean. Indicates whether or not advanced API Authorization scenarios are enabled.
      * `enablePublicSignupUserExistsError` (`bool`) - Boolean. Indicates whether or not the public sign up process shows a user_exists error if the user already exists.
      * `universal_login` (`bool`) - Boolean. Indicates whether or not the tenant uses universal login.
      * `useScopeDescriptionsForConsent` (`bool`)
    """
    friendly_name: pulumi.Output[str]
    """
    String. Friendly name for the tenant.
    """
    guardian_mfa_page: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.

      * `enabled` (`bool`) - Boolean. Indicates whether or not to use the custom Guardian page.
      * `html` (`str`) - String, HTML format with supported Liquid syntax. Customized content of the Guardian page.
    """
    idle_session_lifetime: pulumi.Output[float]
    """
    Integer. Number of hours during which a session can be inactive before the user must log in again.
    """
    picture_url: pulumi.Output[str]
    """
    . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
    """
    sandbox_version: pulumi.Output[str]
    """
    String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
    """
    session_lifetime: pulumi.Output[float]
    """
    Integer. Number of hours during which a session will stay valid.
    """
    support_email: pulumi.Output[str]
    """
    String. Support email address for authenticating users.
    """
    support_url: pulumi.Output[str]
    """
    String. Support URL for authenticating users.
    """
    universal_login: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for Universal Login. For details, see Universal Login.

      * `colors` (`dict`) - List(Resource). Configuration settings for Universal Login colors. See Universal Login - Colors.
        * `pageBackground` (`str`) - String, Hexadecimal. Background color of login pages.
        * `primary` (`str`) - String, Hexadecimal. Primary button background color.
    """
    def __init__(__self__, resource_name, opts=None, allowed_logout_urls=None, change_password=None, default_audience=None, default_directory=None, default_redirection_uri=None, enabled_locales=None, error_page=None, flags=None, friendly_name=None, guardian_mfa_page=None, idle_session_lifetime=None, picture_url=None, sandbox_version=None, session_lifetime=None, support_email=None, support_url=None, universal_login=None, __props__=None, __name__=None, __opts__=None):
        """
        With this resource, you can manage Auth0 tenants, including setting logos and support contact information, setting error pages, and configuring default tenant behaviors.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        tenant = auth0.Tenant("tenant",
            allowed_logout_urls=["http://mysite/logout"],
            change_password={
                "enabled": True,
                "html": (lambda path: open(path).read())("./password_reset.html"),
            },
            default_audience="<client_id>",
            default_directory="Connection-Name",
            error_page={
                "html": (lambda path: open(path).read())("./error.html"),
                "showLogLink": True,
                "url": "http://mysite/errors",
            },
            friendly_name="Tenant Name",
            guardian_mfa_page={
                "enabled": True,
                "html": (lambda path: open(path).read())("./guardian_multifactor.html"),
            },
            picture_url="http://mysite/logo.png",
            sandbox_version="8",
            session_lifetime=46000,
            support_email="support@mysite",
            support_url="http://mysite/support")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_logout_urls: List(String). URLs that Auth0 may redirect to after logout.
        :param pulumi.Input[dict] change_password: List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
        :param pulumi.Input[str] default_audience: String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
        :param pulumi.Input[str] default_directory: String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
        :param pulumi.Input[str] default_redirection_uri: String. The default absolute redirection uri, must be https and cannot contain a fragment.
        :param pulumi.Input[dict] error_page: List(Resource). Configuration settings for error pages. For details, see Error Page.
        :param pulumi.Input[dict] flags: List(Resource). Configuration settings for tenant flags. For details, see Flags.
        :param pulumi.Input[str] friendly_name: String. Friendly name for the tenant.
        :param pulumi.Input[dict] guardian_mfa_page: List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
        :param pulumi.Input[float] idle_session_lifetime: Integer. Number of hours during which a session can be inactive before the user must log in again.
        :param pulumi.Input[str] picture_url: . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
        :param pulumi.Input[str] sandbox_version: String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
        :param pulumi.Input[float] session_lifetime: Integer. Number of hours during which a session will stay valid.
        :param pulumi.Input[str] support_email: String. Support email address for authenticating users.
        :param pulumi.Input[str] support_url: String. Support URL for authenticating users.
        :param pulumi.Input[dict] universal_login: List(Resource). Configuration settings for Universal Login. For details, see Universal Login.

        The **change_password** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the custom change password page.
          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the change password page.

        The **error_page** object supports the following:

          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the error page.
          * `showLogLink` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to show the link to logs as part of the default error page.
          * `url` (`pulumi.Input[str]`) - String. URL to redirect to when an error occurs rather than showing the default error page.

        The **flags** object supports the following:

          * `changePwdFlowV1` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the older v1 change password flow. Not recommended except for backward compatibility.
          * `disableClickjackProtectionHeaders` (`pulumi.Input[bool]`) - Boolean. Indicated whether or not classic Universal Login prompts include additional security headers to prevent clickjacking.
          * `enableApisSection` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the APIs section is enabled for the tenant.
          * `enableClientConnections` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not all current connections should be enabled when a new client is created.
          * `enableCustomDomainInEmails` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant allows custom domains in emails.
          * `enableDynamicClientRegistration` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant allows dynamic client registration.
          * `enableLegacyLogsSearchV2` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the older v2 legacy logs search.
          * `enablePipeline2` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not advanced API Authorization scenarios are enabled.
          * `enablePublicSignupUserExistsError` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the public sign up process shows a user_exists error if the user already exists.
          * `universal_login` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant uses universal login.
          * `useScopeDescriptionsForConsent` (`pulumi.Input[bool]`)

        The **guardian_mfa_page** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the custom Guardian page.
          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the Guardian page.

        The **universal_login** object supports the following:

          * `colors` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for Universal Login colors. See Universal Login - Colors.
            * `pageBackground` (`pulumi.Input[str]`) - String, Hexadecimal. Background color of login pages.
            * `primary` (`pulumi.Input[str]`) - String, Hexadecimal. Primary button background color.
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
    def get(resource_name, id, opts=None, allowed_logout_urls=None, change_password=None, default_audience=None, default_directory=None, default_redirection_uri=None, enabled_locales=None, error_page=None, flags=None, friendly_name=None, guardian_mfa_page=None, idle_session_lifetime=None, picture_url=None, sandbox_version=None, session_lifetime=None, support_email=None, support_url=None, universal_login=None):
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_logout_urls: List(String). URLs that Auth0 may redirect to after logout.
        :param pulumi.Input[dict] change_password: List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
        :param pulumi.Input[str] default_audience: String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
        :param pulumi.Input[str] default_directory: String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
        :param pulumi.Input[str] default_redirection_uri: String. The default absolute redirection uri, must be https and cannot contain a fragment.
        :param pulumi.Input[dict] error_page: List(Resource). Configuration settings for error pages. For details, see Error Page.
        :param pulumi.Input[dict] flags: List(Resource). Configuration settings for tenant flags. For details, see Flags.
        :param pulumi.Input[str] friendly_name: String. Friendly name for the tenant.
        :param pulumi.Input[dict] guardian_mfa_page: List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
        :param pulumi.Input[float] idle_session_lifetime: Integer. Number of hours during which a session can be inactive before the user must log in again.
        :param pulumi.Input[str] picture_url: . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
        :param pulumi.Input[str] sandbox_version: String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
        :param pulumi.Input[float] session_lifetime: Integer. Number of hours during which a session will stay valid.
        :param pulumi.Input[str] support_email: String. Support email address for authenticating users.
        :param pulumi.Input[str] support_url: String. Support URL for authenticating users.
        :param pulumi.Input[dict] universal_login: List(Resource). Configuration settings for Universal Login. For details, see Universal Login.

        The **change_password** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the custom change password page.
          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the change password page.

        The **error_page** object supports the following:

          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the error page.
          * `showLogLink` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to show the link to logs as part of the default error page.
          * `url` (`pulumi.Input[str]`) - String. URL to redirect to when an error occurs rather than showing the default error page.

        The **flags** object supports the following:

          * `changePwdFlowV1` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the older v1 change password flow. Not recommended except for backward compatibility.
          * `disableClickjackProtectionHeaders` (`pulumi.Input[bool]`) - Boolean. Indicated whether or not classic Universal Login prompts include additional security headers to prevent clickjacking.
          * `enableApisSection` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the APIs section is enabled for the tenant.
          * `enableClientConnections` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not all current connections should be enabled when a new client is created.
          * `enableCustomDomainInEmails` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant allows custom domains in emails.
          * `enableDynamicClientRegistration` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant allows dynamic client registration.
          * `enableLegacyLogsSearchV2` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the older v2 legacy logs search.
          * `enablePipeline2` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not advanced API Authorization scenarios are enabled.
          * `enablePublicSignupUserExistsError` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the public sign up process shows a user_exists error if the user already exists.
          * `universal_login` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the tenant uses universal login.
          * `useScopeDescriptionsForConsent` (`pulumi.Input[bool]`)

        The **guardian_mfa_page** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the custom Guardian page.
          * `html` (`pulumi.Input[str]`) - String, HTML format with supported Liquid syntax. Customized content of the Guardian page.

        The **universal_login** object supports the following:

          * `colors` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for Universal Login colors. See Universal Login - Colors.
            * `pageBackground` (`pulumi.Input[str]`) - String, Hexadecimal. Background color of login pages.
            * `primary` (`pulumi.Input[str]`) - String, Hexadecimal. Primary button background color.
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
