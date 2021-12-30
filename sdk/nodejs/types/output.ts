// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ActionDependency {
    /**
     * Secret name.
     */
    name: string;
    /**
     * Trigger version.
     */
    version: string;
}

export interface ActionSecret {
    /**
     * Secret name.
     */
    name: string;
    /**
     * Secret value.
     */
    value: string;
}

export interface ActionSupportedTriggers {
    /**
     * Trigger ID.
     */
    id: string;
    /**
     * Trigger version.
     */
    version: string;
}

export interface BrandingColors {
    /**
     * String, Hexadecimal. Background color of login pages.
     */
    pageBackground: string;
    /**
     * String, Hexadecimal. Primary button background color.
     */
    primary: string;
}

export interface BrandingFont {
    /**
     * String. URL for the custom font.
     */
    url: string;
}

export interface BrandingUniversalLogin {
    /**
     * String, body of login pages.
     */
    body: string;
}

export interface ClientAddons {
    /**
     * String
     */
    aws?: {[key: string]: any};
    /**
     * String
     */
    azureBlob?: {[key: string]: any};
    /**
     * String
     */
    azureSb?: {[key: string]: any};
    /**
     * String
     */
    box?: {[key: string]: any};
    /**
     * String
     */
    cloudbees?: {[key: string]: any};
    /**
     * String
     */
    concur?: {[key: string]: any};
    /**
     * String
     */
    dropbox?: {[key: string]: any};
    /**
     * String
     */
    echosign?: {[key: string]: any};
    /**
     * String
     */
    egnyte?: {[key: string]: any};
    /**
     * String
     */
    firebase?: {[key: string]: any};
    /**
     * String
     */
    layer?: {[key: string]: any};
    /**
     * String
     */
    mscrm?: {[key: string]: any};
    /**
     * String
     */
    newrelic?: {[key: string]: any};
    /**
     * String
     */
    office365?: {[key: string]: any};
    /**
     * String
     */
    rms?: {[key: string]: any};
    /**
     * String
     */
    salesforce?: {[key: string]: any};
    /**
     * String
     */
    salesforceApi?: {[key: string]: any};
    /**
     * String
     */
    salesforceSandboxApi?: {[key: string]: any};
    /**
     * List(Resource). Configuration settings for a SAML add-on. For details, see SAML.
     */
    samlp?: outputs.ClientAddonsSamlp;
    /**
     * String
     */
    sapApi?: {[key: string]: any};
    /**
     * String
     */
    sentry?: {[key: string]: any};
    /**
     * String
     */
    sharepoint?: {[key: string]: any};
    /**
     * String
     */
    slack?: {[key: string]: any};
    /**
     * String
     */
    springcm?: {[key: string]: any};
    /**
     * String
     */
    wams?: {[key: string]: any};
    /**
     * String
     */
    wsfed?: {[key: string]: any};
    /**
     * String
     */
    zendesk?: {[key: string]: any};
    /**
     * String
     */
    zoom?: {[key: string]: any};
}

export interface ClientAddonsSamlp {
    /**
     * String. Audience of the SAML Assertion. Default will be the Issuer on SAMLRequest.
     */
    audience?: string;
    /**
     * String. Class reference of the authentication context.
     */
    authnContextClassRef?: string;
    /**
     * String. Protocol binding used for SAML logout responses.
     */
    binding?: string;
    /**
     * Boolean, (Default=true) Indicates whether or not a UPN claim should be created.
     */
    createUpnClaim?: boolean;
    /**
     * String. Destination of the SAML Response. If not specified, it will be AssertionConsumerUrlof SAMLRequest or Callback URL if there was no SAMLRequest.
     */
    destination?: string;
    /**
     * String, (Default=`sha1`). Algorithm used to calculate the digest of the SAML Assertion or response. Options include `defaultsha1` and `sha256`.
     */
    digestAlgorithm?: string;
    /**
     * Boolean,(Default=true). Indicates whether or not we should infer the NameFormat based on the attribute name. If set to false, the attribute NameFormat is not set in the assertion.
     */
    includeAttributeNameFormat?: boolean;
    /**
     * Integer, (Default=3600). Number of seconds during which the token is valid.
     */
    lifetimeInSeconds?: number;
    /**
     * Map(Resource). Configuration settings for logout. For details, see Logout.
     */
    logout?: outputs.ClientAddonsSamlpLogout;
    /**
     * Boolean, (Default=true). Indicates whether or not to add additional identity information in the token, such as the provider used and the access_token, if available.
     */
    mapIdentities?: boolean;
    /**
     * Boolean, (Default=false). Indicates whether or not to add a prefix of `http://schema.auth0.com` to any claims that are not mapped to the common profile when passed through in the output assertion.
     */
    mapUnknownClaimsAsIs?: boolean;
    /**
     * Map(String). Mappings between the Auth0 user profile property name (`name`) and the output attributes on the SAML attribute in the assertion (`value`).
     */
    mappings?: {[key: string]: any};
    /**
     * String, (Default=`urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified`). Format of the name identifier.
     */
    nameIdentifierFormat?: string;
    /**
     * List(String). Attributes that can be used for Subject/NameID. Auth0 will try each of the attributes of this array in order and use the first value it finds.
     */
    nameIdentifierProbes?: string[];
    /**
     * Boolean, (Default=true). Indicates whether or not to passthrough claims that are not mapped to the common profile in the output assertion.
     */
    passthroughClaimsWithNoMapping?: boolean;
    /**
     * String. Recipient of the SAML Assertion (SubjectConfirmationData). Default is AssertionConsumerUrl on SAMLRequest or Callback URL if no SAMLRequest was sent.
     */
    recipient?: string;
    /**
     * Boolean. Indicates whether or not the SAML Response should be signed instead of the SAML Assertion.
     */
    signResponse?: boolean;
    /**
     * String, (Default=`rsa-sha1`). Algorithm used to sign the SAML Assertion or response. Options include `rsa-sha1` and `rsa-sha256`.
     */
    signatureAlgorithm?: string;
    /**
     * Boolean, (Default=true). Indicates whether or not we should infer the `xs:type` of the element. Types include `xs:string`, `xs:boolean`, `xs:double`, and `xs:anyType`. When set to false, all `xs:type` are `xs:anyType`.
     */
    typedAttributes?: boolean;
}

export interface ClientAddonsSamlpLogout {
    /**
     * String. Service provider's Single Logout Service URL, to which Auth0 will send logout requests and responses.
     */
    callback?: string;
    /**
     * Boolean. Indicates whether or not Auth0 should notify service providers of session termination.
     */
    sloEnabled?: boolean;
}

export interface ClientJwtConfiguration {
    /**
     * String. Algorithm used to sign JWTs.
     */
    alg?: string;
    /**
     * Integer. Number of seconds during which the JWT will be valid.
     */
    lifetimeInSeconds: number;
    /**
     * Map(String). Permissions (scopes) included in JWTs.
     */
    scopes?: {[key: string]: string};
    /**
     * Boolean. Indicates whether or not the client secret is base64 encoded.
     */
    secretEncoded: boolean;
}

export interface ClientMobile {
    /**
     * List(Resource). Configuration settings for Android native apps. For details, see Android.
     */
    android?: outputs.ClientMobileAndroid;
    /**
     * List(Resource). Configuration settings for i0S native apps. For details, see iOS.
     */
    ios?: outputs.ClientMobileIos;
}

export interface ClientMobileAndroid {
    /**
     * String
     */
    appPackageName?: string;
    /**
     * List(String)
     */
    sha256CertFingerprints?: string[];
}

export interface ClientMobileIos {
    /**
     * String
     */
    appBundleIdentifier?: string;
    /**
     * String
     */
    teamId?: string;
}

export interface ClientRefreshToken {
    /**
     * String. Options include `expiring`, `non-expiring`. Whether a refresh token will expire based on an absolute lifetime, after which the token can no longer be used. If rotation is `rotating`, this must be set to `expiring`.
     */
    expirationType: string;
    /**
     * Integer. The time in seconds after which inactive refresh tokens will expire.
     */
    idleTokenLifetime?: number;
    /**
     * Boolean, (Default=false) Whether or not inactive refresh tokens should be remain valid indefinitely.
     */
    infiniteIdleTokenLifetime?: boolean;
    /**
     * Boolean, (Default=false) Whether or not refresh tokens should remain valid indefinitely. If false, `tokenLifetime` should also be set
     */
    infiniteTokenLifetime?: boolean;
    /**
     * Integer. The amount of time in seconds in which a refresh token may be reused without trigging reuse detection.
     */
    leeway?: number;
    /**
     * String. Options include `rotating`, `non-rotating`. When `rotating`, exchanging a refresh token will cause a new refresh token to be issued and the existing token will be invalidated. This allows for automatic detection of token reuse if the token is leaked.
     */
    rotationType: string;
    /**
     * Integer. The absolute lifetime of a refresh token in seconds.
     */
    tokenLifetime?: number;
}

export interface ConnectionOptions {
    /**
     * ADFS Metadata source.
     */
    adfsServer?: string;
    /**
     * List of allowed audiences.
     */
    allowedAudiences?: string[];
    apiEnableUsers?: boolean;
    /**
     * Azure AD domain name.
     *
     * @deprecated Use domain instead
     */
    appDomain?: string;
    /**
     * Azure AD app ID.
     */
    appId?: string;
    authorizationEndpoint?: string;
    /**
     * Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.
     */
    bruteForceProtection?: boolean;
    /**
     * OIDC provider client ID.
     */
    clientId?: string;
    /**
     * OIDC provider client secret.
     */
    clientSecret?: string;
    /**
     * String.
     */
    communityBaseUrl?: string;
    /**
     * A case-sensitive map of key value pairs used as configuration variables for the `customScript`.
     */
    configuration?: {[key: string]: string};
    /**
     * Custom database action scripts. For more information, read [Custom Database Action Script Templates](https://auth0.com/docs/connections/database/custom-db/templates).
     */
    customScripts?: {[key: string]: string};
    /**
     * (Boolean) When enabled additional debugging information will be generated.
     */
    debug?: boolean;
    /**
     * Sign Request Algorithm Digest
     */
    digestAlgorithm?: string;
    disableCache?: boolean;
    /**
     * Boolean. Indicates whether or not to allow user sign-ups to your application.
     */
    disableSignup?: boolean;
    /**
     * OpenID discovery URL. E.g. `https://auth.example.com/.well-known/openid-configuration`.
     */
    discoveryUrl?: string;
    domain?: string;
    /**
     * List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
     */
    domainAliases?: string[];
    enabledDatabaseCustomization?: boolean;
    /**
     * Custom Entity ID for the connection.
     */
    entityId?: string;
    /**
     * SAML Attributes mapping. If you're configuring a SAML enterprise connection for a non-standard PingFederate Server, you must update the attribute mappings.
     */
    fieldsMap?: {[key: string]: string};
    /**
     * SMS number for the sender. Used when SMS Source is From.
     */
    from?: string;
    iconUrl?: string;
    identityApi?: string;
    /**
     * Configuration Options for IDP Initiated Authentication.  This is an object with the properties: `clientId`, `clientProtocol`, and `clientAuthorizeQuery`
     */
    idpInitiated?: outputs.ConnectionOptionsIdpInitiated;
    /**
     * Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).
     */
    importMode?: boolean;
    ips?: string[];
    /**
     * Issuer URL. E.g. `https://auth.example.com`
     */
    issuer?: string;
    jwksUri?: string;
    /**
     * Key ID.
     */
    keyId?: string;
    /**
     * Maximum number of groups to retrieve.
     */
    maxGroupsToRetrieve?: string;
    /**
     * SID for Copilot. Used when SMS Source is Copilot.
     */
    messagingServiceSid?: string;
    /**
     * Configuration settings Options for multifactor authentication. For details, see MFA Options.
     */
    mfa?: outputs.ConnectionOptionsMfa;
    /**
     * Name of the connection.
     */
    name?: string;
    /**
     * If there are user fields that should not be stored in Auth0 databases due to privacy reasons, you can add them to the denylist. See [here](https://auth0.com/docs/security/denylist-user-attributes) for more info.
     */
    nonPersistentAttrs: string[];
    /**
     * Configuration settings for password complexity. For details, see Password Complexity Options.
     */
    passwordComplexityOptions?: outputs.ConnectionOptionsPasswordComplexityOptions;
    /**
     * Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see Password Dictionary.
     */
    passwordDictionary?: outputs.ConnectionOptionsPasswordDictionary;
    /**
     * Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see Password History.
     */
    passwordHistories: outputs.ConnectionOptionsPasswordHistory[];
    /**
     * Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see Password No Personal Info.
     */
    passwordNoPersonalInfo?: outputs.ConnectionOptionsPasswordNoPersonalInfo;
    /**
     * Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`.
     */
    passwordPolicy: string;
    /**
     * The SAML Response Binding - how the SAML token is received by Auth0 from IdP. Two possible values are `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` (default) and `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST`
     */
    protocolBinding?: string;
    /**
     * Template that formats the SAML request
     */
    requestTemplate?: string;
    /**
     * Indicates whether or not the user is required to provide a username in addition to an email address.
     */
    requiresUsername?: boolean;
    /**
     * Scopes required by the connection. The value must be a list, for example `["openid", "profile", "email"]`.
     */
    scopes?: string[];
    scripts?: {[key: string]: string};
    /**
     * Determines whether the 'name', 'given_name', 'family_name', 'nickname', and 'picture' attributes can be independently updated when using the external IdP. Default is `onEachLogin` and can be set to `onFirstLogin`.
     */
    setUserRootAttributes: string;
    /**
     * Determines how Auth0 sets the emailVerified field in the user profile. Can either be set to `neverSetEmailsAsVerified` or `alwaysSetEmailsAsVerified`.
     */
    shouldTrustEmailVerifiedConnection?: string;
    /**
     * SAML single login URL for the connection.
     */
    signInEndpoint?: string;
    /**
     * SAML single logout URL for the connection.
     */
    signOutEndpoint?: string;
    /**
     * (Boolean) When enabled, the SAML authentication request will be signed.
     */
    signSamlRequest?: boolean;
    /**
     * Sign Request Algorithm
     */
    signatureAlgorithm?: string;
    /**
     * The X.509 signing certificate (encoded in PEM or CER) you retrieved from the IdP, Base64-encoded
     */
    signingCert?: string;
    /**
     * Version 1 is deprecated, use version 2.
     */
    strategyVersion: number;
    subject?: string;
    /**
     * Syntax of the SMS. Options include `markdown` and `liquid`.
     */
    syntax?: string;
    /**
     * Team ID.
     */
    teamId?: string;
    /**
     * Template for the SMS. You can use `@@password@@` as a placeholder for the password value.
     */
    template?: string;
    tenantDomain?: string;
    tokenEndpoint?: string;
    /**
     * Configuration options for one-time passwords. For details, see TOTP.
     */
    totp?: outputs.ConnectionOptionsTotp;
    /**
     * SID for your Twilio account.
     */
    twilioSid?: string;
    /**
     * AuthToken for your Twilio account.
     */
    twilioToken?: string;
    /**
     * Value can be `backChannel` or `frontChannel`.
     */
    type?: string;
    useCertAuth?: boolean;
    useKerberos?: boolean;
    useWsfed?: boolean;
    /**
     * Attribute in the SAML token that will be mapped to the userId property in Auth0.
     */
    userIdAttribute?: string;
    userinfoEndpoint?: string;
    /**
     * Validation of the minimum and maximum values allowed for a user to have as username. For details, see Validation.
     */
    validation?: outputs.ConnectionOptionsValidation;
    /**
     * Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD.
     */
    waadCommonEndpoint?: boolean;
    waadProtocol?: string;
}

export interface ConnectionOptionsIdpInitiated {
    clientAuthorizeQuery?: string;
    /**
     * Google client ID.
     */
    clientId?: string;
    clientProtocol?: string;
}

export interface ConnectionOptionsMfa {
    /**
     * Indicates whether multifactor authentication is enabled for this connection.
     */
    active?: boolean;
    /**
     * Indicates whether multifactor authentication enrollment settings will be returned.
     */
    returnEnrollSettings?: boolean;
}

export interface ConnectionOptionsPasswordComplexityOptions {
    /**
     * Minimum number of characters allowed in passwords.
     */
    minLength?: number;
}

export interface ConnectionOptionsPasswordDictionary {
    /**
     * Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive.
     */
    dictionaries?: string[];
    /**
     * Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
     */
    enable?: boolean;
}

export interface ConnectionOptionsPasswordHistory {
    /**
     * Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
     */
    enable?: boolean;
    /**
     * Indicates the number of passwords to keep in history with a maximum of 24.
     */
    size?: number;
}

export interface ConnectionOptionsPasswordNoPersonalInfo {
    /**
     * Indicates whether the password personal info check is enabled for this connection.
     */
    enable?: boolean;
}

export interface ConnectionOptionsTotp {
    /**
     * Integer. Length of the one-time password.
     */
    length?: number;
    /**
     * Integer. Seconds between allowed generation of new passwords.
     */
    timeStep?: number;
}

export interface ConnectionOptionsValidation {
    /**
     * Specifies the `min` and `max` values of username length. `min` and `max` are integers.
     */
    username?: outputs.ConnectionOptionsValidationUsername;
}

export interface ConnectionOptionsValidationUsername {
    max?: number;
    min?: number;
}

export interface CustomDomainVerification {
    /**
     * List(Map). Verification methods for the domain.
     */
    methods: any[];
}

export interface EmailCredentials {
    /**
     * String, Case-sensitive. AWS Access Key ID. Used only for AWS.
     */
    accessKeyId?: string;
    /**
     * String, Case-sensitive. API Key for your email service. Will always be encrypted in our database.
     */
    apiKey?: string;
    /**
     * String. API User for your email service.
     */
    apiUser?: string;
    domain?: string;
    /**
     * String. Default region. Used only for AWS, Mailgun, and SparkPost.
     */
    region?: string;
    /**
     * String, Case-sensitive. AWS Secret Key. Will always be encrypted in our database. Used only for AWS.
     */
    secretAccessKey?: string;
    /**
     * String. Hostname or IP address of your SMTP server. Used only for SMTP.
     */
    smtpHost?: string;
    /**
     * String, Case-sensitive. SMTP password. Used only for SMTP.
     */
    smtpPass?: string;
    /**
     * Integer. Port used by your SMTP server. Please avoid using port 25 if possible because many providers have limitations on this port. Used only for SMTP.
     */
    smtpPort?: number;
    /**
     * String. SMTP username. Used only for SMTP.
     */
    smtpUser?: string;
}

export interface GlobalClientAddons {
    aws?: {[key: string]: any};
    azureBlob?: {[key: string]: any};
    azureSb?: {[key: string]: any};
    box?: {[key: string]: any};
    cloudbees?: {[key: string]: any};
    concur?: {[key: string]: any};
    dropbox?: {[key: string]: any};
    echosign?: {[key: string]: any};
    egnyte?: {[key: string]: any};
    firebase?: {[key: string]: any};
    layer?: {[key: string]: any};
    mscrm?: {[key: string]: any};
    newrelic?: {[key: string]: any};
    office365?: {[key: string]: any};
    rms?: {[key: string]: any};
    salesforce?: {[key: string]: any};
    salesforceApi?: {[key: string]: any};
    salesforceSandboxApi?: {[key: string]: any};
    samlp?: outputs.GlobalClientAddonsSamlp;
    sapApi?: {[key: string]: any};
    sentry?: {[key: string]: any};
    sharepoint?: {[key: string]: any};
    slack?: {[key: string]: any};
    springcm?: {[key: string]: any};
    wams?: {[key: string]: any};
    wsfed?: {[key: string]: any};
    zendesk?: {[key: string]: any};
    zoom?: {[key: string]: any};
}

export interface GlobalClientAddonsSamlp {
    audience?: string;
    authnContextClassRef?: string;
    binding?: string;
    createUpnClaim?: boolean;
    destination?: string;
    digestAlgorithm?: string;
    includeAttributeNameFormat?: boolean;
    lifetimeInSeconds?: number;
    logout?: outputs.GlobalClientAddonsSamlpLogout;
    mapIdentities?: boolean;
    mapUnknownClaimsAsIs?: boolean;
    mappings?: {[key: string]: any};
    nameIdentifierFormat?: string;
    nameIdentifierProbes?: string[];
    passthroughClaimsWithNoMapping?: boolean;
    recipient?: string;
    signResponse?: boolean;
    signatureAlgorithm?: string;
    typedAttributes?: boolean;
}

export interface GlobalClientAddonsSamlpLogout {
    callback?: string;
    sloEnabled?: boolean;
}

export interface GlobalClientJwtConfiguration {
    alg?: string;
    lifetimeInSeconds: number;
    scopes?: {[key: string]: string};
    secretEncoded: boolean;
}

export interface GlobalClientMobile {
    android?: outputs.GlobalClientMobileAndroid;
    ios?: outputs.GlobalClientMobileIos;
}

export interface GlobalClientMobileAndroid {
    appPackageName?: string;
    sha256CertFingerprints?: string[];
}

export interface GlobalClientMobileIos {
    appBundleIdentifier?: string;
    teamId?: string;
}

export interface GlobalClientRefreshToken {
    expirationType: string;
    idleTokenLifetime?: number;
    infiniteIdleTokenLifetime?: boolean;
    infiniteTokenLifetime?: boolean;
    leeway?: number;
    rotationType: string;
    tokenLifetime?: number;
}

export interface GuardianPhone {
    /**
     * List(String). Message types to use, array of `sms` and or `voice`. Adding both to array should enable the user to choose.
     */
    messageTypes: string[];
    /**
     * List(Resource). Options for the various providers. See Options.
     */
    options?: outputs.GuardianPhoneOptions;
    /**
     * String, Case-sensitive. Provider to use, one of `auth0`, `twilio` or `phone-message-hook`.
     */
    provider: string;
}

export interface GuardianPhoneOptions {
    /**
     * String.
     */
    authToken?: string;
    /**
     * String. This message will be sent whenever a user enrolls a new device for the first time using MFA. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).
     */
    enrollmentMessage?: string;
    /**
     * String.
     */
    from?: string;
    /**
     * String.
     */
    messagingServiceSid?: string;
    /**
     * String.
     */
    sid?: string;
    /**
     * String. This message will be sent whenever a user logs in after the enrollment. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).
     */
    verificationMessage?: string;
}

export interface LogStreamSink {
    /**
     * The AWS Account ID
     */
    awsAccountId?: string;
    /**
     * Name of the Partner Event Source to be used with AWS. Generally generated by Auth0 and passed to AWS so this should generally be an output attribute.
     */
    awsPartnerEventSource: string;
    /**
     * The AWS Region (i.e "us-east-2")
     */
    awsRegion?: string;
    /**
     * Name of the Partner Topic to be used with Azure.  Generally should not be specified.
     */
    azurePartnerTopic: string;
    /**
     * The Azure region code (i.e. "ne")
     */
    azureRegion?: string;
    /**
     * The Azure EventGrid resource group which allows you to manage all Azure assets within one subscription
     */
    azureResourceGroup?: string;
    /**
     * The unique alphanumeric string that identifies your Azure subscription
     */
    azureSubscriptionId?: string;
    /**
     * The Datadog API key
     */
    datadogApiKey?: string;
    /**
     * The Datadog region
     */
    datadogRegion?: string;
    /**
     * Sent in the HTTP "Authorization" header with each request
     */
    httpAuthorization?: string;
    /**
     * The format of data sent over HTTP. Options are "JSONLINES" or "JSONARRAY"
     */
    httpContentFormat?: string;
    /**
     * The ContentType header to send over HTTP.  Common value is "application/json"
     */
    httpContentType?: string;
    /**
     * Additional HTTP headers to be included as part of the HTTP request
     */
    httpCustomHeaders?: string[];
    /**
     * The HTTP endpoint to send streaming logs
     */
    httpEndpoint?: string;
    /**
     * The Splunk domain name
     */
    splunkDomain?: string;
    splunkPort?: string;
    /**
     * This toggle should be turned off when using self-signed certificates
     */
    splunkSecure?: boolean;
    /**
     * The Splunk access token
     */
    splunkToken?: string;
    /**
     * Generated URL for your defined HTTP source in Sumo Logic for collecting streaming data from Auth0
     */
    sumoSourceAddress?: string;
}

export interface OrganizationBranding {
    /**
     * Color scheme used to customize the login pages
     */
    colors?: {[key: string]: string};
    /**
     * URL of logo to display on login page
     */
    logoUrl?: string;
}

export interface OrganizationConnection {
    /**
     * When true, all users that log in
     * with this connection will be automatically granted membership in the
     * organization. When false, users must be granted membership in the organization
     * before logging in with this connection.
     */
    assignMembershipOnLogin?: boolean;
    /**
     * The connection ID of the connection to add to the
     * organization
     */
    connectionId: string;
}

export interface ResourceServerScope {
    /**
     * String. Description of the permission (scope).
     */
    description?: string;
    /**
     * String. Name of the permission (scope). Examples include `read:appointments` or `delete:appointments`.
     */
    value: string;
}

export interface RolePermission {
    /**
     * String. Name of the permission (scope).
     */
    name: string;
    /**
     * String. Unique identifier for the resource server.
     */
    resourceServerIdentifier: string;
}

export interface TenantChangePassword {
    /**
     * Boolean. Indicates whether or not to use the custom change password page.
     */
    enabled: boolean;
    /**
     * String, HTML format with supported Liquid syntax. Customized content of the change password page.
     */
    html: string;
}

export interface TenantErrorPage {
    /**
     * String, HTML format with supported Liquid syntax. Customized content of the error page.
     */
    html: string;
    /**
     * Boolean. Indicates whether or not to show the link to logs as part of the default error page.
     */
    showLogLink: boolean;
    /**
     * String. URL to redirect to when an error occurs rather than showing the default error page.
     */
    url: string;
}

export interface TenantFlags {
    /**
     * Boolean. Indicates whether or not to use the older v1 change password flow. Not recommended except for backward compatibility.
     */
    changePwdFlowV1: boolean;
    /**
     * Boolean. Indicated whether or not classic Universal Login prompts include additional security headers to prevent clickjacking.
     */
    disableClickjackProtectionHeaders: boolean;
    /**
     * Boolean. Indicates whether or not the APIs section is enabled for the tenant.
     */
    enableApisSection: boolean;
    /**
     * Boolean. Indicates whether or not all current connections should be enabled when a new client is created.
     */
    enableClientConnections: boolean;
    /**
     * Boolean. Indicates whether or not the tenant allows custom domains in emails.
     */
    enableCustomDomainInEmails: boolean;
    /**
     * Boolean. Indicates whether or not the tenant allows dynamic client registration.
     */
    enableDynamicClientRegistration: boolean;
    /**
     * Boolean. Indicates whether or not to use the older v2 legacy logs search.
     */
    enableLegacyLogsSearchV2: boolean;
    /**
     * Boolean. Indicates whether or not advanced API Authorization scenarios are enabled.
     */
    enablePipeline2: boolean;
    /**
     * Boolean. Indicates whether or not the public sign up process shows a userExists error if the user already exists.
     */
    enablePublicSignupUserExistsError: boolean;
    /**
     * Boolean. Indicates whether or not the tenant uses universal login.
     */
    universalLogin: boolean;
    useScopeDescriptionsForConsent: boolean;
}

export interface TenantGuardianMfaPage {
    /**
     * Boolean. Indicates whether or not to use the custom Guardian page.
     */
    enabled: boolean;
    /**
     * String, HTML format with supported Liquid syntax. Customized content of the Guardian page.
     */
    html: string;
}

export interface TenantUniversalLogin {
    /**
     * List(Resource). Configuration settings for Universal Login colors. See Universal Login - Colors.
     */
    colors?: outputs.TenantUniversalLoginColors;
}

export interface TenantUniversalLoginColors {
    /**
     * String, Hexadecimal. Background color of login pages.
     */
    pageBackground: string;
    /**
     * String, Hexadecimal. Primary button background color.
     */
    primary: string;
}

export interface TriggerBindingAction {
    /**
     * The name of an action.
     */
    displayName: string;
    /**
     * Trigger ID.
     */
    id: string;
}

