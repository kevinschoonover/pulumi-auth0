// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * With this resource, you can set up applications that use Auth0 for authentication and configure allowed callback URLs and secrets for these applications. Depending on your plan, you may also configure add-ons to allow your application to call another application's API (such as Firebase and AWS) on behalf of an authenticated user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-auth0/blob/master/website/docs/r/client.html.md.
 */
export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'auth0:index/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    /**
     * List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
     */
    public readonly addons!: pulumi.Output<outputs.ClientAddons | undefined>;
    /**
     * List(String). URLs that Auth0 may redirect to after logout.
     */
    public readonly allowedLogoutUrls!: pulumi.Output<string[] | undefined>;
    /**
     * List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
     */
    public readonly allowedOrigins!: pulumi.Output<string[] | undefined>;
    /**
     * String. Type of application the client represents. Options include `native`, `spa`, `regularWeb`, `nonInteractive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
     */
    public readonly appType!: pulumi.Output<string | undefined>;
    /**
     * List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
     */
    public readonly callbacks!: pulumi.Output<string[] | undefined>;
    /**
     * String. ID of the client.
     */
    public /*out*/ readonly clientId!: pulumi.Output<string>;
    /**
     * Map(String)
     */
    public readonly clientMetadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * String. Secret for the client; keep this private.
     */
    public /*out*/ readonly clientSecret!: pulumi.Output<string>;
    /**
     * Map.
     */
    public readonly clientSecretRotationTrigger!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests. 
     */
    public readonly crossOriginAuth!: pulumi.Output<boolean | undefined>;
    /**
     * String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
     */
    public readonly crossOriginLoc!: pulumi.Output<string | undefined>;
    /**
     * String. Content of the custom login page.
     */
    public readonly customLoginPage!: pulumi.Output<string | undefined>;
    /**
     * Boolean. Indicates whether or not a custom login page is to be used.
     */
    public readonly customLoginPageOn!: pulumi.Output<boolean>;
    /**
     * String.
     */
    public readonly customLoginPagePreview!: pulumi.Output<string | undefined>;
    /**
     * String, (Max length = 140 characters). Description of the purpose of the client.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Map(String).
     */
    public readonly encryptionKey!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * String. Form template for WS-Federation protocol.
     */
    public readonly formTemplate!: pulumi.Output<string | undefined>;
    /**
     * List(String). Types of grants that this client is authorized to use.
     */
    public readonly grantTypes!: pulumi.Output<string[]>;
    public readonly initiateLoginUri!: pulumi.Output<string | undefined>;
    /**
     * Boolean. Indicates whether or not this client is a first-party client.
     */
    public readonly isFirstParty!: pulumi.Output<boolean>;
    /**
     * Boolean. Indicates whether or not the token endpoint IP header is trusted.
     */
    public readonly isTokenEndpointIpHeaderTrusted!: pulumi.Output<boolean>;
    /**
     * List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
     */
    public readonly jwtConfiguration!: pulumi.Output<outputs.ClientJwtConfiguration>;
    /**
     * String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
     */
    public readonly logoUri!: pulumi.Output<string | undefined>;
    /**
     * List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
     */
    public readonly mobile!: pulumi.Output<outputs.ClientMobile | undefined>;
    /**
     * String. Name of the client.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
     */
    public readonly oidcConformant!: pulumi.Output<boolean>;
    /**
     * Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
     */
    public readonly sso!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean. Indicates whether or not SSO is disabled.
     */
    public readonly ssoDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `clientSecretPost` (client uses HTTP POST parameters), `clientSecretBasic` (client uses HTTP Basic).
     */
    public readonly tokenEndpointAuthMethod!: pulumi.Output<string>;
    /**
     * List(String). URLs that represent valid web origins for use with web message response mode.
     */
    public readonly webOrigins!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientState | undefined;
            inputs["addons"] = state ? state.addons : undefined;
            inputs["allowedLogoutUrls"] = state ? state.allowedLogoutUrls : undefined;
            inputs["allowedOrigins"] = state ? state.allowedOrigins : undefined;
            inputs["appType"] = state ? state.appType : undefined;
            inputs["callbacks"] = state ? state.callbacks : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientMetadata"] = state ? state.clientMetadata : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["clientSecretRotationTrigger"] = state ? state.clientSecretRotationTrigger : undefined;
            inputs["crossOriginAuth"] = state ? state.crossOriginAuth : undefined;
            inputs["crossOriginLoc"] = state ? state.crossOriginLoc : undefined;
            inputs["customLoginPage"] = state ? state.customLoginPage : undefined;
            inputs["customLoginPageOn"] = state ? state.customLoginPageOn : undefined;
            inputs["customLoginPagePreview"] = state ? state.customLoginPagePreview : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptionKey"] = state ? state.encryptionKey : undefined;
            inputs["formTemplate"] = state ? state.formTemplate : undefined;
            inputs["grantTypes"] = state ? state.grantTypes : undefined;
            inputs["initiateLoginUri"] = state ? state.initiateLoginUri : undefined;
            inputs["isFirstParty"] = state ? state.isFirstParty : undefined;
            inputs["isTokenEndpointIpHeaderTrusted"] = state ? state.isTokenEndpointIpHeaderTrusted : undefined;
            inputs["jwtConfiguration"] = state ? state.jwtConfiguration : undefined;
            inputs["logoUri"] = state ? state.logoUri : undefined;
            inputs["mobile"] = state ? state.mobile : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["oidcConformant"] = state ? state.oidcConformant : undefined;
            inputs["sso"] = state ? state.sso : undefined;
            inputs["ssoDisabled"] = state ? state.ssoDisabled : undefined;
            inputs["tokenEndpointAuthMethod"] = state ? state.tokenEndpointAuthMethod : undefined;
            inputs["webOrigins"] = state ? state.webOrigins : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            inputs["addons"] = args ? args.addons : undefined;
            inputs["allowedLogoutUrls"] = args ? args.allowedLogoutUrls : undefined;
            inputs["allowedOrigins"] = args ? args.allowedOrigins : undefined;
            inputs["appType"] = args ? args.appType : undefined;
            inputs["callbacks"] = args ? args.callbacks : undefined;
            inputs["clientMetadata"] = args ? args.clientMetadata : undefined;
            inputs["clientSecretRotationTrigger"] = args ? args.clientSecretRotationTrigger : undefined;
            inputs["crossOriginAuth"] = args ? args.crossOriginAuth : undefined;
            inputs["crossOriginLoc"] = args ? args.crossOriginLoc : undefined;
            inputs["customLoginPage"] = args ? args.customLoginPage : undefined;
            inputs["customLoginPageOn"] = args ? args.customLoginPageOn : undefined;
            inputs["customLoginPagePreview"] = args ? args.customLoginPagePreview : undefined;
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            inputs["formTemplate"] = args ? args.formTemplate : undefined;
            inputs["grantTypes"] = args ? args.grantTypes : undefined;
            inputs["initiateLoginUri"] = args ? args.initiateLoginUri : undefined;
            inputs["isFirstParty"] = args ? args.isFirstParty : undefined;
            inputs["isTokenEndpointIpHeaderTrusted"] = args ? args.isTokenEndpointIpHeaderTrusted : undefined;
            inputs["jwtConfiguration"] = args ? args.jwtConfiguration : undefined;
            inputs["logoUri"] = args ? args.logoUri : undefined;
            inputs["mobile"] = args ? args.mobile : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oidcConformant"] = args ? args.oidcConformant : undefined;
            inputs["sso"] = args ? args.sso : undefined;
            inputs["ssoDisabled"] = args ? args.ssoDisabled : undefined;
            inputs["tokenEndpointAuthMethod"] = args ? args.tokenEndpointAuthMethod : undefined;
            inputs["webOrigins"] = args ? args.webOrigins : undefined;
            inputs["clientId"] = undefined /*out*/;
            inputs["clientSecret"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Client.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    /**
     * List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
     */
    readonly addons?: pulumi.Input<inputs.ClientAddons>;
    /**
     * List(String). URLs that Auth0 may redirect to after logout.
     */
    readonly allowedLogoutUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
     */
    readonly allowedOrigins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. Type of application the client represents. Options include `native`, `spa`, `regularWeb`, `nonInteractive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
     */
    readonly appType?: pulumi.Input<string>;
    /**
     * List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
     */
    readonly callbacks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. ID of the client.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Map(String)
     */
    readonly clientMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * String. Secret for the client; keep this private.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * Map.
     */
    readonly clientSecretRotationTrigger?: pulumi.Input<{[key: string]: any}>;
    /**
     * Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests. 
     */
    readonly crossOriginAuth?: pulumi.Input<boolean>;
    /**
     * String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
     */
    readonly crossOriginLoc?: pulumi.Input<string>;
    /**
     * String. Content of the custom login page.
     */
    readonly customLoginPage?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not a custom login page is to be used.
     */
    readonly customLoginPageOn?: pulumi.Input<boolean>;
    /**
     * String.
     */
    readonly customLoginPagePreview?: pulumi.Input<string>;
    /**
     * String, (Max length = 140 characters). Description of the purpose of the client.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Map(String).
     */
    readonly encryptionKey?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * String. Form template for WS-Federation protocol.
     */
    readonly formTemplate?: pulumi.Input<string>;
    /**
     * List(String). Types of grants that this client is authorized to use.
     */
    readonly grantTypes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly initiateLoginUri?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not this client is a first-party client.
     */
    readonly isFirstParty?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not the token endpoint IP header is trusted.
     */
    readonly isTokenEndpointIpHeaderTrusted?: pulumi.Input<boolean>;
    /**
     * List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
     */
    readonly jwtConfiguration?: pulumi.Input<inputs.ClientJwtConfiguration>;
    /**
     * String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
     */
    readonly logoUri?: pulumi.Input<string>;
    /**
     * List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
     */
    readonly mobile?: pulumi.Input<inputs.ClientMobile>;
    /**
     * String. Name of the client.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
     */
    readonly oidcConformant?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
     */
    readonly sso?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not SSO is disabled.
     */
    readonly ssoDisabled?: pulumi.Input<boolean>;
    /**
     * String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `clientSecretPost` (client uses HTTP POST parameters), `clientSecretBasic` (client uses HTTP Basic).
     */
    readonly tokenEndpointAuthMethod?: pulumi.Input<string>;
    /**
     * List(String). URLs that represent valid web origins for use with web message response mode.
     */
    readonly webOrigins?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    /**
     * List(Resource). Configuration settings for add-ons for this client. For details, see Add-ons.
     */
    readonly addons?: pulumi.Input<inputs.ClientAddons>;
    /**
     * List(String). URLs that Auth0 may redirect to after logout.
     */
    readonly allowedLogoutUrls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
     */
    readonly allowedOrigins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. Type of application the client represents. Options include `native`, `spa`, `regularWeb`, `nonInteractive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
     */
    readonly appType?: pulumi.Input<string>;
    /**
     * List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
     */
    readonly callbacks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map(String)
     */
    readonly clientMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Map.
     */
    readonly clientSecretRotationTrigger?: pulumi.Input<{[key: string]: any}>;
    /**
     * Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests. 
     */
    readonly crossOriginAuth?: pulumi.Input<boolean>;
    /**
     * String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
     */
    readonly crossOriginLoc?: pulumi.Input<string>;
    /**
     * String. Content of the custom login page.
     */
    readonly customLoginPage?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not a custom login page is to be used.
     */
    readonly customLoginPageOn?: pulumi.Input<boolean>;
    /**
     * String.
     */
    readonly customLoginPagePreview?: pulumi.Input<string>;
    /**
     * String, (Max length = 140 characters). Description of the purpose of the client.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Map(String).
     */
    readonly encryptionKey?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * String. Form template for WS-Federation protocol.
     */
    readonly formTemplate?: pulumi.Input<string>;
    /**
     * List(String). Types of grants that this client is authorized to use.
     */
    readonly grantTypes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly initiateLoginUri?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not this client is a first-party client.
     */
    readonly isFirstParty?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not the token endpoint IP header is trusted.
     */
    readonly isTokenEndpointIpHeaderTrusted?: pulumi.Input<boolean>;
    /**
     * List(Resource). Configuration settings for the JWTs issued for this client. For details, see JWT Configuration.
     */
    readonly jwtConfiguration?: pulumi.Input<inputs.ClientJwtConfiguration>;
    /**
     * String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
     */
    readonly logoUri?: pulumi.Input<string>;
    /**
     * List(Resource). Configuration settings for mobile native applications. For details, see Mobile.
     */
    readonly mobile?: pulumi.Input<inputs.ClientMobile>;
    /**
     * String. Name of the client.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
     */
    readonly oidcConformant?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
     */
    readonly sso?: pulumi.Input<boolean>;
    /**
     * Boolean. Indicates whether or not SSO is disabled.
     */
    readonly ssoDisabled?: pulumi.Input<boolean>;
    /**
     * String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `clientSecretPost` (client uses HTTP POST parameters), `clientSecretBasic` (client uses HTTP Basic).
     */
    readonly tokenEndpointAuthMethod?: pulumi.Input<string>;
    /**
     * List(String). URLs that represent valid web origins for use with web message response mode.
     */
    readonly webOrigins?: pulumi.Input<pulumi.Input<string>[]>;
}