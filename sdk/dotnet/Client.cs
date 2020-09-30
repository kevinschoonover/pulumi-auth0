// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Auth0
{
    /// <summary>
    /// With this resource, you can set up applications that use Auth0 for authentication and configure allowed callback URLs and secrets for these applications. Depending on your plan, you may also configure add-ons to allow your application to call another application's API (such as Firebase and AWS) on behalf of an authenticated user.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Auth0 = Pulumi.Auth0;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myClient = new Auth0.Client("myClient", new Auth0.ClientArgs
    ///         {
    ///             Addons = new Auth0.Inputs.ClientAddonsArgs
    ///             {
    ///                 Firebase = 
    ///                 {
    ///                     { "clientEmail", "john.doe@example.com" },
    ///                     { "lifetimeInSeconds", 1 },
    ///                     { "privateKey", "wer" },
    ///                     { "privateKeyId", "qwreerwerwe" },
    ///                 },
    ///                 Samlp = new Auth0.Inputs.ClientAddonsSamlpArgs
    ///                 {
    ///                     Audience = "https://example.com/saml",
    ///                     CreateUpnClaim = false,
    ///                     MapIdentities = false,
    ///                     MapUnknownClaimsAsIs = false,
    ///                     Mappings = 
    ///                     {
    ///                         { "email", "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress" },
    ///                         { "name", "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name" },
    ///                     },
    ///                     NameIdentifierFormat = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
    ///                     NameIdentifierProbes = 
    ///                     {
    ///                         "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
    ///                     },
    ///                     PassthroughClaimsWithNoMapping = false,
    ///                 },
    ///             },
    ///             AllowedLogoutUrls = 
    ///             {
    ///                 "https://example.com",
    ///             },
    ///             AllowedOrigins = 
    ///             {
    ///                 "https://example.com",
    ///             },
    ///             AppType = "non_interactive",
    ///             Callbacks = 
    ///             {
    ///                 "https://example.com/callback",
    ///             },
    ///             ClientMetadata = 
    ///             {
    ///                 { "foo", "zoo" },
    ///             },
    ///             CustomLoginPageOn = true,
    ///             Description = "Test Applications Long Description",
    ///             GrantTypes = 
    ///             {
    ///                 "authorization_code",
    ///                 "http://auth0.com/oauth/grant-type/password-realm",
    ///                 "implicit",
    ///                 "password",
    ///                 "refresh_token",
    ///             },
    ///             IsFirstParty = true,
    ///             IsTokenEndpointIpHeaderTrusted = true,
    ///             JwtConfiguration = new Auth0.Inputs.ClientJwtConfigurationArgs
    ///             {
    ///                 Alg = "RS256",
    ///                 LifetimeInSeconds = 300,
    ///                 Scopes = 
    ///                 {
    ///                     { "foo", "bar" },
    ///                 },
    ///                 SecretEncoded = true,
    ///             },
    ///             Mobile = new Auth0.Inputs.ClientMobileArgs
    ///             {
    ///                 Ios = new Auth0.Inputs.ClientMobileIosArgs
    ///                 {
    ///                     AppBundleIdentifier = "com.my.bundle.id",
    ///                     TeamId = "9JA89QQLNQ",
    ///                 },
    ///             },
    ///             OidcConformant = false,
    ///             TokenEndpointAuthMethod = "client_secret_post",
    ///             WebOrigins = 
    ///             {
    ///                 "https://example.com",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Client : Pulumi.CustomResource
    {
        /// <summary>
        /// List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
        /// </summary>
        [Output("addons")]
        public Output<Outputs.ClientAddons?> Addons { get; private set; } = null!;

        /// <summary>
        /// List(String). URLs that Auth0 may redirect to after logout.
        /// </summary>
        [Output("allowedLogoutUrls")]
        public Output<ImmutableArray<string>> AllowedLogoutUrls { get; private set; } = null!;

        /// <summary>
        /// List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
        /// </summary>
        [Output("allowedOrigins")]
        public Output<ImmutableArray<string>> AllowedOrigins { get; private set; } = null!;

        /// <summary>
        /// String. Type of application the client represents. Options include `native`, `spa`, `regular_web`, `non_interactive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
        /// </summary>
        [Output("appType")]
        public Output<string?> AppType { get; private set; } = null!;

        /// <summary>
        /// List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
        /// </summary>
        [Output("callbacks")]
        public Output<ImmutableArray<string>> Callbacks { get; private set; } = null!;

        /// <summary>
        /// String. ID of the client.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Map(String)
        /// </summary>
        [Output("clientMetadata")]
        public Output<ImmutableDictionary<string, object>?> ClientMetadata { get; private set; } = null!;

        /// <summary>
        /// String. Secret for the client; keep this private.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Map.
        /// </summary>
        [Output("clientSecretRotationTrigger")]
        public Output<ImmutableDictionary<string, object>?> ClientSecretRotationTrigger { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests.
        /// </summary>
        [Output("crossOriginAuth")]
        public Output<bool?> CrossOriginAuth { get; private set; } = null!;

        /// <summary>
        /// String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
        /// </summary>
        [Output("crossOriginLoc")]
        public Output<string?> CrossOriginLoc { get; private set; } = null!;

        /// <summary>
        /// String. Content of the custom login page.
        /// </summary>
        [Output("customLoginPage")]
        public Output<string?> CustomLoginPage { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not a custom login page is to be used.
        /// </summary>
        [Output("customLoginPageOn")]
        public Output<bool> CustomLoginPageOn { get; private set; } = null!;

        /// <summary>
        /// String.
        /// </summary>
        [Output("customLoginPagePreview")]
        public Output<string?> CustomLoginPagePreview { get; private set; } = null!;

        /// <summary>
        /// String, (Max length = 140 characters). Description of the purpose of the client.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Map(String).
        /// </summary>
        [Output("encryptionKey")]
        public Output<ImmutableDictionary<string, string>?> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// String. Form template for WS-Federation protocol.
        /// </summary>
        [Output("formTemplate")]
        public Output<string?> FormTemplate { get; private set; } = null!;

        /// <summary>
        /// List(String). Types of grants that this client is authorized to use.
        /// </summary>
        [Output("grantTypes")]
        public Output<ImmutableArray<string>> GrantTypes { get; private set; } = null!;

        [Output("initiateLoginUri")]
        public Output<string?> InitiateLoginUri { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not this client is a first-party client.
        /// </summary>
        [Output("isFirstParty")]
        public Output<bool> IsFirstParty { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the token endpoint IP header is trusted.
        /// </summary>
        [Output("isTokenEndpointIpHeaderTrusted")]
        public Output<bool> IsTokenEndpointIpHeaderTrusted { get; private set; } = null!;

        /// <summary>
        /// List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
        /// </summary>
        [Output("jwtConfiguration")]
        public Output<Outputs.ClientJwtConfiguration> JwtConfiguration { get; private set; } = null!;

        /// <summary>
        /// String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
        /// </summary>
        [Output("logoUri")]
        public Output<string?> LogoUri { get; private set; } = null!;

        /// <summary>
        /// List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
        /// </summary>
        [Output("mobile")]
        public Output<Outputs.ClientMobile?> Mobile { get; private set; } = null!;

        /// <summary>
        /// String. Name of the client.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
        /// </summary>
        [Output("oidcConformant")]
        public Output<bool> OidcConformant { get; private set; } = null!;

        [Output("refreshToken")]
        public Output<Outputs.ClientRefreshToken> RefreshToken { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
        /// </summary>
        [Output("sso")]
        public Output<bool?> Sso { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not SSO is disabled.
        /// </summary>
        [Output("ssoDisabled")]
        public Output<bool?> SsoDisabled { get; private set; } = null!;

        /// <summary>
        /// String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters), `client_secret_basic` (client uses HTTP Basic).
        /// </summary>
        [Output("tokenEndpointAuthMethod")]
        public Output<string> TokenEndpointAuthMethod { get; private set; } = null!;

        /// <summary>
        /// List(String). URLs that represent valid web origins for use with web message response mode.
        /// </summary>
        [Output("webOrigins")]
        public Output<ImmutableArray<string>> WebOrigins { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs? args = null, CustomResourceOptions? options = null)
            : base("auth0:index/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/client:Client", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
        /// </summary>
        [Input("addons")]
        public Input<Inputs.ClientAddonsArgs>? Addons { get; set; }

        [Input("allowedLogoutUrls")]
        private InputList<string>? _allowedLogoutUrls;

        /// <summary>
        /// List(String). URLs that Auth0 may redirect to after logout.
        /// </summary>
        public InputList<string> AllowedLogoutUrls
        {
            get => _allowedLogoutUrls ?? (_allowedLogoutUrls = new InputList<string>());
            set => _allowedLogoutUrls = value;
        }

        [Input("allowedOrigins")]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        /// <summary>
        /// String. Type of application the client represents. Options include `native`, `spa`, `regular_web`, `non_interactive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
        /// </summary>
        [Input("appType")]
        public Input<string>? AppType { get; set; }

        [Input("callbacks")]
        private InputList<string>? _callbacks;

        /// <summary>
        /// List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
        /// </summary>
        public InputList<string> Callbacks
        {
            get => _callbacks ?? (_callbacks = new InputList<string>());
            set => _callbacks = value;
        }

        [Input("clientMetadata")]
        private InputMap<object>? _clientMetadata;

        /// <summary>
        /// Map(String)
        /// </summary>
        public InputMap<object> ClientMetadata
        {
            get => _clientMetadata ?? (_clientMetadata = new InputMap<object>());
            set => _clientMetadata = value;
        }

        [Input("clientSecretRotationTrigger")]
        private InputMap<object>? _clientSecretRotationTrigger;

        /// <summary>
        /// Map.
        /// </summary>
        public InputMap<object> ClientSecretRotationTrigger
        {
            get => _clientSecretRotationTrigger ?? (_clientSecretRotationTrigger = new InputMap<object>());
            set => _clientSecretRotationTrigger = value;
        }

        /// <summary>
        /// Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests.
        /// </summary>
        [Input("crossOriginAuth")]
        public Input<bool>? CrossOriginAuth { get; set; }

        /// <summary>
        /// String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
        /// </summary>
        [Input("crossOriginLoc")]
        public Input<string>? CrossOriginLoc { get; set; }

        /// <summary>
        /// String. Content of the custom login page.
        /// </summary>
        [Input("customLoginPage")]
        public Input<string>? CustomLoginPage { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not a custom login page is to be used.
        /// </summary>
        [Input("customLoginPageOn")]
        public Input<bool>? CustomLoginPageOn { get; set; }

        /// <summary>
        /// String.
        /// </summary>
        [Input("customLoginPagePreview")]
        public Input<string>? CustomLoginPagePreview { get; set; }

        /// <summary>
        /// String, (Max length = 140 characters). Description of the purpose of the client.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("encryptionKey")]
        private InputMap<string>? _encryptionKey;

        /// <summary>
        /// Map(String).
        /// </summary>
        public InputMap<string> EncryptionKey
        {
            get => _encryptionKey ?? (_encryptionKey = new InputMap<string>());
            set => _encryptionKey = value;
        }

        /// <summary>
        /// String. Form template for WS-Federation protocol.
        /// </summary>
        [Input("formTemplate")]
        public Input<string>? FormTemplate { get; set; }

        [Input("grantTypes")]
        private InputList<string>? _grantTypes;

        /// <summary>
        /// List(String). Types of grants that this client is authorized to use.
        /// </summary>
        public InputList<string> GrantTypes
        {
            get => _grantTypes ?? (_grantTypes = new InputList<string>());
            set => _grantTypes = value;
        }

        [Input("initiateLoginUri")]
        public Input<string>? InitiateLoginUri { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not this client is a first-party client.
        /// </summary>
        [Input("isFirstParty")]
        public Input<bool>? IsFirstParty { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the token endpoint IP header is trusted.
        /// </summary>
        [Input("isTokenEndpointIpHeaderTrusted")]
        public Input<bool>? IsTokenEndpointIpHeaderTrusted { get; set; }

        /// <summary>
        /// List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
        /// </summary>
        [Input("jwtConfiguration")]
        public Input<Inputs.ClientJwtConfigurationArgs>? JwtConfiguration { get; set; }

        /// <summary>
        /// String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
        /// </summary>
        [Input("logoUri")]
        public Input<string>? LogoUri { get; set; }

        /// <summary>
        /// List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
        /// </summary>
        [Input("mobile")]
        public Input<Inputs.ClientMobileArgs>? Mobile { get; set; }

        /// <summary>
        /// String. Name of the client.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
        /// </summary>
        [Input("oidcConformant")]
        public Input<bool>? OidcConformant { get; set; }

        [Input("refreshToken")]
        public Input<Inputs.ClientRefreshTokenArgs>? RefreshToken { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
        /// </summary>
        [Input("sso")]
        public Input<bool>? Sso { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not SSO is disabled.
        /// </summary>
        [Input("ssoDisabled")]
        public Input<bool>? SsoDisabled { get; set; }

        /// <summary>
        /// String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters), `client_secret_basic` (client uses HTTP Basic).
        /// </summary>
        [Input("tokenEndpointAuthMethod")]
        public Input<string>? TokenEndpointAuthMethod { get; set; }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;

        /// <summary>
        /// List(String). URLs that represent valid web origins for use with web message response mode.
        /// </summary>
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ClientState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
        /// </summary>
        [Input("addons")]
        public Input<Inputs.ClientAddonsGetArgs>? Addons { get; set; }

        [Input("allowedLogoutUrls")]
        private InputList<string>? _allowedLogoutUrls;

        /// <summary>
        /// List(String). URLs that Auth0 may redirect to after logout.
        /// </summary>
        public InputList<string> AllowedLogoutUrls
        {
            get => _allowedLogoutUrls ?? (_allowedLogoutUrls = new InputList<string>());
            set => _allowedLogoutUrls = value;
        }

        [Input("allowedOrigins")]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        /// <summary>
        /// String. Type of application the client represents. Options include `native`, `spa`, `regular_web`, `non_interactive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
        /// </summary>
        [Input("appType")]
        public Input<string>? AppType { get; set; }

        [Input("callbacks")]
        private InputList<string>? _callbacks;

        /// <summary>
        /// List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
        /// </summary>
        public InputList<string> Callbacks
        {
            get => _callbacks ?? (_callbacks = new InputList<string>());
            set => _callbacks = value;
        }

        /// <summary>
        /// String. ID of the client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientMetadata")]
        private InputMap<object>? _clientMetadata;

        /// <summary>
        /// Map(String)
        /// </summary>
        public InputMap<object> ClientMetadata
        {
            get => _clientMetadata ?? (_clientMetadata = new InputMap<object>());
            set => _clientMetadata = value;
        }

        /// <summary>
        /// String. Secret for the client; keep this private.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        [Input("clientSecretRotationTrigger")]
        private InputMap<object>? _clientSecretRotationTrigger;

        /// <summary>
        /// Map.
        /// </summary>
        public InputMap<object> ClientSecretRotationTrigger
        {
            get => _clientSecretRotationTrigger ?? (_clientSecretRotationTrigger = new InputMap<object>());
            set => _clientSecretRotationTrigger = value;
        }

        /// <summary>
        /// Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests.
        /// </summary>
        [Input("crossOriginAuth")]
        public Input<bool>? CrossOriginAuth { get; set; }

        /// <summary>
        /// String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
        /// </summary>
        [Input("crossOriginLoc")]
        public Input<string>? CrossOriginLoc { get; set; }

        /// <summary>
        /// String. Content of the custom login page.
        /// </summary>
        [Input("customLoginPage")]
        public Input<string>? CustomLoginPage { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not a custom login page is to be used.
        /// </summary>
        [Input("customLoginPageOn")]
        public Input<bool>? CustomLoginPageOn { get; set; }

        /// <summary>
        /// String.
        /// </summary>
        [Input("customLoginPagePreview")]
        public Input<string>? CustomLoginPagePreview { get; set; }

        /// <summary>
        /// String, (Max length = 140 characters). Description of the purpose of the client.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("encryptionKey")]
        private InputMap<string>? _encryptionKey;

        /// <summary>
        /// Map(String).
        /// </summary>
        public InputMap<string> EncryptionKey
        {
            get => _encryptionKey ?? (_encryptionKey = new InputMap<string>());
            set => _encryptionKey = value;
        }

        /// <summary>
        /// String. Form template for WS-Federation protocol.
        /// </summary>
        [Input("formTemplate")]
        public Input<string>? FormTemplate { get; set; }

        [Input("grantTypes")]
        private InputList<string>? _grantTypes;

        /// <summary>
        /// List(String). Types of grants that this client is authorized to use.
        /// </summary>
        public InputList<string> GrantTypes
        {
            get => _grantTypes ?? (_grantTypes = new InputList<string>());
            set => _grantTypes = value;
        }

        [Input("initiateLoginUri")]
        public Input<string>? InitiateLoginUri { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not this client is a first-party client.
        /// </summary>
        [Input("isFirstParty")]
        public Input<bool>? IsFirstParty { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the token endpoint IP header is trusted.
        /// </summary>
        [Input("isTokenEndpointIpHeaderTrusted")]
        public Input<bool>? IsTokenEndpointIpHeaderTrusted { get; set; }

        /// <summary>
        /// List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
        /// </summary>
        [Input("jwtConfiguration")]
        public Input<Inputs.ClientJwtConfigurationGetArgs>? JwtConfiguration { get; set; }

        /// <summary>
        /// String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
        /// </summary>
        [Input("logoUri")]
        public Input<string>? LogoUri { get; set; }

        /// <summary>
        /// List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
        /// </summary>
        [Input("mobile")]
        public Input<Inputs.ClientMobileGetArgs>? Mobile { get; set; }

        /// <summary>
        /// String. Name of the client.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
        /// </summary>
        [Input("oidcConformant")]
        public Input<bool>? OidcConformant { get; set; }

        [Input("refreshToken")]
        public Input<Inputs.ClientRefreshTokenGetArgs>? RefreshToken { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
        /// </summary>
        [Input("sso")]
        public Input<bool>? Sso { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not SSO is disabled.
        /// </summary>
        [Input("ssoDisabled")]
        public Input<bool>? SsoDisabled { get; set; }

        /// <summary>
        /// String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters), `client_secret_basic` (client uses HTTP Basic).
        /// </summary>
        [Input("tokenEndpointAuthMethod")]
        public Input<string>? TokenEndpointAuthMethod { get; set; }

        [Input("webOrigins")]
        private InputList<string>? _webOrigins;

        /// <summary>
        /// List(String). URLs that represent valid web origins for use with web message response mode.
        /// </summary>
        public InputList<string> WebOrigins
        {
            get => _webOrigins ?? (_webOrigins = new InputList<string>());
            set => _webOrigins = value;
        }

        public ClientState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
