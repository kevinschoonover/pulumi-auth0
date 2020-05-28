# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Connection(pulumi.CustomResource):
    display_name: pulumi.Output[str]
    """
    Name used in login screen
    """
    enabled_clients: pulumi.Output[list]
    """
    Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
    """
    is_domain_connection: pulumi.Output[bool]
    """
    Boolean. Indicates whether or not the connection is domain level.
    """
    name: pulumi.Output[str]
    """
    String. Name of the connection.
    """
    options: pulumi.Output[dict]
    """
    List(Resource). Configuration settings for connection options. For details, see Options.

      * `adfsServer` (`str`) - String. ADFS Metadata source.
      * `allowedAudiences` (`list`)
      * `apiEnableUsers` (`bool`) - Boolean.
      * `appDomain` (`str`) - String. Azure AD domain name.
      * `appId` (`str`) - String
      * `authorizationEndpoint` (`str`) - String.
      * `bruteForceProtection` (`bool`) - Boolean. Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.
      * `client_id` (`str`) - String. Client ID given by your OIDC provider.
      * `client_secret` (`str`) - String, Case-sensitive. Client secret given by your OIDC provider.
      * `communityBaseUrl` (`str`) - String.
      * `configuration` (`dict`) - Map(String), Case-sensitive.
      * `customScripts` (`dict`) - Map(String). 
      * `disableCache` (`bool`)
      * `disableSignup` (`bool`) - Boolean. Indicates whether or not to allow user sign-ups to your application.
      * `discoveryUrl` (`str`) - String. Usually an URL ending with `/.well-known/openid-configuration`
      * `domain` (`str`)
      * `domainAliases` (`list`) - List(String). List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
      * `enabledDatabaseCustomization` (`bool`) - Boolean.
      * `from_` (`str`) - String. SMS number for the sender. Used when SMS Source is From.
      * `iconUrl` (`str`)
      * `identityApi` (`str`)
      * `importMode` (`bool`) - Boolean. Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).
      * `ips` (`list`)
      * `issuer` (`str`) - String. URL of the issuer.
      * `jwksUri` (`str`) - String.
      * `keyId` (`str`)
      * `maxGroupsToRetrieve` (`str`) - String. Maximum number of groups to retrieve.
      * `messagingServiceSid` (`str`) - String. SID for Copilot. Used when SMS Source is Copilot.
      * `name` (`str`) - String. 
      * `passwordComplexityOptions` (`dict`) - List(Resource). Configuration settings for password complexity. For details, see Password Complexity Options.
        * `minLength` (`float`) - Integer. Minimum number of characters allowed in passwords.

      * `passwordDictionary` (`dict`) - List(Resource). Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see Password Dictionary.
        * `dictionaries` (`list`) - Set(String), (Maximum=2000 characters). Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive.
        * `enable` (`bool`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.

      * `passwordHistories` (`list`) - List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see Password History.
        * `enable` (`bool`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
        * `size` (`float`) - Integer, (Maximum=24). Indicates the number of passwords to keep in history. 

      * `passwordNoPersonalInfo` (`dict`) - List(Resource). Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see Password No Personal Info.
        * `enable` (`bool`) - Boolean. Indicates whether the password personal info check is enabled for this connection.

      * `passwordPolicy` (`str`) - String. Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`.
      * `requiresUsername` (`bool`) - Boolean. Indicates whether or not the user is required to provide a username in addition to an email address.
      * `scopes` (`list`) - List(String). Value must be a list of scopes. For example `["openid", "profile", "email"]`
      * `strategy_version` (`float`) - Int. Version 1 is deprecated, use version 2.
      * `subject` (`str`)
      * `syntax` (`str`) - String. Syntax of the SMS. Options include `markdown` and `liquid`.
      * `teamId` (`str`)
      * `template` (`str`) - String. Template for the SMS. You can use `@@password@@` as a placeholder for the password value.
      * `tenantDomain` (`str`) - String
      * `tokenEndpoint` (`str`) - String.
      * `totp` (`dict`) - Map(Resource). Configuration options for one-time passwords. For details, see TOTP.
        * `length` (`float`) - Integer. Length of the one-time password.
        * `timeStep` (`float`) - Integer. Seconds between allowed generation of new passwords.

      * `twilioSid` (`str`) - String. SID for your Twilio account.
      * `twilioToken` (`str`) - String, Case-sensitive. AuthToken for your Twilio account.
      * `type` (`str`) - String. Value must be `back_channel` or `front_channel`
      * `useCertAuth` (`bool`)
      * `useKerberos` (`bool`)
      * `useWsfed` (`bool`) - Bool
      * `userinfoEndpoint` (`str`) - String.
      * `validation` (`dict`) - String.
      * `waadCommonEndpoint` (`bool`) - Boolean. Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD.
      * `waadProtocol` (`str`) - String
    """
    realms: pulumi.Output[list]
    """
    List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
    """
    strategy: pulumi.Output[str]
    """
    String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
    """
    strategy_version: pulumi.Output[str]
    """
    Int. Version 1 is deprecated, use version 2.
    """
    def __init__(__self__, resource_name, opts=None, display_name=None, enabled_clients=None, is_domain_connection=None, name=None, options=None, realms=None, strategy=None, strategy_version=None, __props__=None, __name__=None, __opts__=None):
        """
        With Auth0, you can define sources of users, otherwise known as connections, which may include identity providers (such as Google or LinkedIn), databases, or passwordless authentication methods. This resource allows you to configure and manage connections to be used with your clients and users.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_auth0 as auth0

        my_connection = auth0.Connection("myConnection",
            options={
                "bruteForceProtection": "true",
                "configuration": {
                    "bar": "baz",
                    "foo": "bar",
                },
                "customScripts": {
                    "getUser": \"\"\"function getByEmail (email, callback) {
          return callback(new Error("Whoops!"))
        }

        \"\"\",
                },
                "enabledDatabaseCustomization": "true",
                "passwordHistory": [{
                    "enable": True,
                    "size": 3,
                }],
                "passwordPolicy": "excellent",
            },
            strategy="auth0")
        my_waad_connection = auth0.Connection("myWaadConnection",
            options={
                "apiEnableUsers": True,
                "appDomain": "my-auth0-app.eu.auth0.com",
                "basicProfile": True,
                "client_id": "1234",
                "client_secret": "1234",
                "domainAliases": ["example.io"],
                "extGroups": True,
                "extProfile": True,
                "tenantDomain": "exmaple.onmicrosoft.com",
                "useWsfed": False,
                "waadCommonEndpoint": False,
                "waadProtocol": "openid-connect",
            },
            strategy="waad")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Name used in login screen
        :param pulumi.Input[list] enabled_clients: Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
        :param pulumi.Input[bool] is_domain_connection: Boolean. Indicates whether or not the connection is domain level.
        :param pulumi.Input[str] name: String. Name of the connection.
        :param pulumi.Input[dict] options: List(Resource). Configuration settings for connection options. For details, see Options.
        :param pulumi.Input[list] realms: List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
        :param pulumi.Input[str] strategy: String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
        :param pulumi.Input[str] strategy_version: Int. Version 1 is deprecated, use version 2.

        The **options** object supports the following:

          * `adfsServer` (`pulumi.Input[str]`) - String. ADFS Metadata source.
          * `allowedAudiences` (`pulumi.Input[list]`)
          * `apiEnableUsers` (`pulumi.Input[bool]`) - Boolean.
          * `appDomain` (`pulumi.Input[str]`) - String. Azure AD domain name.
          * `appId` (`pulumi.Input[str]`) - String
          * `authorizationEndpoint` (`pulumi.Input[str]`) - String.
          * `bruteForceProtection` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.
          * `client_id` (`pulumi.Input[str]`) - String. Client ID given by your OIDC provider.
          * `client_secret` (`pulumi.Input[str]`) - String, Case-sensitive. Client secret given by your OIDC provider.
          * `communityBaseUrl` (`pulumi.Input[str]`) - String.
          * `configuration` (`pulumi.Input[dict]`) - Map(String), Case-sensitive.
          * `customScripts` (`pulumi.Input[dict]`) - Map(String). 
          * `disableCache` (`pulumi.Input[bool]`)
          * `disableSignup` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to allow user sign-ups to your application.
          * `discoveryUrl` (`pulumi.Input[str]`) - String. Usually an URL ending with `/.well-known/openid-configuration`
          * `domain` (`pulumi.Input[str]`)
          * `domainAliases` (`pulumi.Input[list]`) - List(String). List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
          * `enabledDatabaseCustomization` (`pulumi.Input[bool]`) - Boolean.
          * `from_` (`pulumi.Input[str]`) - String. SMS number for the sender. Used when SMS Source is From.
          * `iconUrl` (`pulumi.Input[str]`)
          * `identityApi` (`pulumi.Input[str]`)
          * `importMode` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).
          * `ips` (`pulumi.Input[list]`)
          * `issuer` (`pulumi.Input[str]`) - String. URL of the issuer.
          * `jwksUri` (`pulumi.Input[str]`) - String.
          * `keyId` (`pulumi.Input[str]`)
          * `maxGroupsToRetrieve` (`pulumi.Input[str]`) - String. Maximum number of groups to retrieve.
          * `messagingServiceSid` (`pulumi.Input[str]`) - String. SID for Copilot. Used when SMS Source is Copilot.
          * `name` (`pulumi.Input[str]`) - String. 
          * `passwordComplexityOptions` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for password complexity. For details, see Password Complexity Options.
            * `minLength` (`pulumi.Input[float]`) - Integer. Minimum number of characters allowed in passwords.

          * `passwordDictionary` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see Password Dictionary.
            * `dictionaries` (`pulumi.Input[list]`) - Set(String), (Maximum=2000 characters). Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.

          * `passwordHistories` (`pulumi.Input[list]`) - List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see Password History.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
            * `size` (`pulumi.Input[float]`) - Integer, (Maximum=24). Indicates the number of passwords to keep in history. 

          * `passwordNoPersonalInfo` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see Password No Personal Info.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether the password personal info check is enabled for this connection.

          * `passwordPolicy` (`pulumi.Input[str]`) - String. Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`.
          * `requiresUsername` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the user is required to provide a username in addition to an email address.
          * `scopes` (`pulumi.Input[list]`) - List(String). Value must be a list of scopes. For example `["openid", "profile", "email"]`
          * `strategy_version` (`pulumi.Input[float]`) - Int. Version 1 is deprecated, use version 2.
          * `subject` (`pulumi.Input[str]`)
          * `syntax` (`pulumi.Input[str]`) - String. Syntax of the SMS. Options include `markdown` and `liquid`.
          * `teamId` (`pulumi.Input[str]`)
          * `template` (`pulumi.Input[str]`) - String. Template for the SMS. You can use `@@password@@` as a placeholder for the password value.
          * `tenantDomain` (`pulumi.Input[str]`) - String
          * `tokenEndpoint` (`pulumi.Input[str]`) - String.
          * `totp` (`pulumi.Input[dict]`) - Map(Resource). Configuration options for one-time passwords. For details, see TOTP.
            * `length` (`pulumi.Input[float]`) - Integer. Length of the one-time password.
            * `timeStep` (`pulumi.Input[float]`) - Integer. Seconds between allowed generation of new passwords.

          * `twilioSid` (`pulumi.Input[str]`) - String. SID for your Twilio account.
          * `twilioToken` (`pulumi.Input[str]`) - String, Case-sensitive. AuthToken for your Twilio account.
          * `type` (`pulumi.Input[str]`) - String. Value must be `back_channel` or `front_channel`
          * `useCertAuth` (`pulumi.Input[bool]`)
          * `useKerberos` (`pulumi.Input[bool]`)
          * `useWsfed` (`pulumi.Input[bool]`) - Bool
          * `userinfoEndpoint` (`pulumi.Input[str]`) - String.
          * `validation` (`pulumi.Input[dict]`) - String.
          * `waadCommonEndpoint` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD.
          * `waadProtocol` (`pulumi.Input[str]`) - String
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

            __props__['display_name'] = display_name
            __props__['enabled_clients'] = enabled_clients
            __props__['is_domain_connection'] = is_domain_connection
            __props__['name'] = name
            __props__['options'] = options
            __props__['realms'] = realms
            if strategy is None:
                raise TypeError("Missing required property 'strategy'")
            __props__['strategy'] = strategy
            __props__['strategy_version'] = strategy_version
        super(Connection, __self__).__init__(
            'auth0:index/connection:Connection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, display_name=None, enabled_clients=None, is_domain_connection=None, name=None, options=None, realms=None, strategy=None, strategy_version=None):
        """
        Get an existing Connection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Name used in login screen
        :param pulumi.Input[list] enabled_clients: Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
        :param pulumi.Input[bool] is_domain_connection: Boolean. Indicates whether or not the connection is domain level.
        :param pulumi.Input[str] name: String. Name of the connection.
        :param pulumi.Input[dict] options: List(Resource). Configuration settings for connection options. For details, see Options.
        :param pulumi.Input[list] realms: List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
        :param pulumi.Input[str] strategy: String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
        :param pulumi.Input[str] strategy_version: Int. Version 1 is deprecated, use version 2.

        The **options** object supports the following:

          * `adfsServer` (`pulumi.Input[str]`) - String. ADFS Metadata source.
          * `allowedAudiences` (`pulumi.Input[list]`)
          * `apiEnableUsers` (`pulumi.Input[bool]`) - Boolean.
          * `appDomain` (`pulumi.Input[str]`) - String. Azure AD domain name.
          * `appId` (`pulumi.Input[str]`) - String
          * `authorizationEndpoint` (`pulumi.Input[str]`) - String.
          * `bruteForceProtection` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.
          * `client_id` (`pulumi.Input[str]`) - String. Client ID given by your OIDC provider.
          * `client_secret` (`pulumi.Input[str]`) - String, Case-sensitive. Client secret given by your OIDC provider.
          * `communityBaseUrl` (`pulumi.Input[str]`) - String.
          * `configuration` (`pulumi.Input[dict]`) - Map(String), Case-sensitive.
          * `customScripts` (`pulumi.Input[dict]`) - Map(String). 
          * `disableCache` (`pulumi.Input[bool]`)
          * `disableSignup` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to allow user sign-ups to your application.
          * `discoveryUrl` (`pulumi.Input[str]`) - String. Usually an URL ending with `/.well-known/openid-configuration`
          * `domain` (`pulumi.Input[str]`)
          * `domainAliases` (`pulumi.Input[list]`) - List(String). List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
          * `enabledDatabaseCustomization` (`pulumi.Input[bool]`) - Boolean.
          * `from_` (`pulumi.Input[str]`) - String. SMS number for the sender. Used when SMS Source is From.
          * `iconUrl` (`pulumi.Input[str]`)
          * `identityApi` (`pulumi.Input[str]`)
          * `importMode` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).
          * `ips` (`pulumi.Input[list]`)
          * `issuer` (`pulumi.Input[str]`) - String. URL of the issuer.
          * `jwksUri` (`pulumi.Input[str]`) - String.
          * `keyId` (`pulumi.Input[str]`)
          * `maxGroupsToRetrieve` (`pulumi.Input[str]`) - String. Maximum number of groups to retrieve.
          * `messagingServiceSid` (`pulumi.Input[str]`) - String. SID for Copilot. Used when SMS Source is Copilot.
          * `name` (`pulumi.Input[str]`) - String. 
          * `passwordComplexityOptions` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for password complexity. For details, see Password Complexity Options.
            * `minLength` (`pulumi.Input[float]`) - Integer. Minimum number of characters allowed in passwords.

          * `passwordDictionary` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see Password Dictionary.
            * `dictionaries` (`pulumi.Input[list]`) - Set(String), (Maximum=2000 characters). Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.

          * `passwordHistories` (`pulumi.Input[list]`) - List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see Password History.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
            * `size` (`pulumi.Input[float]`) - Integer, (Maximum=24). Indicates the number of passwords to keep in history. 

          * `passwordNoPersonalInfo` (`pulumi.Input[dict]`) - List(Resource). Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see Password No Personal Info.
            * `enable` (`pulumi.Input[bool]`) - Boolean. Indicates whether the password personal info check is enabled for this connection.

          * `passwordPolicy` (`pulumi.Input[str]`) - String. Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`.
          * `requiresUsername` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not the user is required to provide a username in addition to an email address.
          * `scopes` (`pulumi.Input[list]`) - List(String). Value must be a list of scopes. For example `["openid", "profile", "email"]`
          * `strategy_version` (`pulumi.Input[float]`) - Int. Version 1 is deprecated, use version 2.
          * `subject` (`pulumi.Input[str]`)
          * `syntax` (`pulumi.Input[str]`) - String. Syntax of the SMS. Options include `markdown` and `liquid`.
          * `teamId` (`pulumi.Input[str]`)
          * `template` (`pulumi.Input[str]`) - String. Template for the SMS. You can use `@@password@@` as a placeholder for the password value.
          * `tenantDomain` (`pulumi.Input[str]`) - String
          * `tokenEndpoint` (`pulumi.Input[str]`) - String.
          * `totp` (`pulumi.Input[dict]`) - Map(Resource). Configuration options for one-time passwords. For details, see TOTP.
            * `length` (`pulumi.Input[float]`) - Integer. Length of the one-time password.
            * `timeStep` (`pulumi.Input[float]`) - Integer. Seconds between allowed generation of new passwords.

          * `twilioSid` (`pulumi.Input[str]`) - String. SID for your Twilio account.
          * `twilioToken` (`pulumi.Input[str]`) - String, Case-sensitive. AuthToken for your Twilio account.
          * `type` (`pulumi.Input[str]`) - String. Value must be `back_channel` or `front_channel`
          * `useCertAuth` (`pulumi.Input[bool]`)
          * `useKerberos` (`pulumi.Input[bool]`)
          * `useWsfed` (`pulumi.Input[bool]`) - Bool
          * `userinfoEndpoint` (`pulumi.Input[str]`) - String.
          * `validation` (`pulumi.Input[dict]`) - String.
          * `waadCommonEndpoint` (`pulumi.Input[bool]`) - Boolean. Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD.
          * `waadProtocol` (`pulumi.Input[str]`) - String
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["display_name"] = display_name
        __props__["enabled_clients"] = enabled_clients
        __props__["is_domain_connection"] = is_domain_connection
        __props__["name"] = name
        __props__["options"] = options
        __props__["realms"] = realms
        __props__["strategy"] = strategy
        __props__["strategy_version"] = strategy_version
        return Connection(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

