// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// With this resource, you can manage Auth0 tenants, including setting logos and support contact information, setting error pages, and configuring default tenant behaviors.
//
// > Auth0 does not currently support creating tenants through the Management API. Therefore this resource can only manage an existing tenant created through the Auth0 dashboard.
//
// Auth0 does not currently support adding/removing extensions on tenants through their API. The Auth0 dashboard must be used to add/remove extensions.
type Tenant struct {
	pulumi.CustomResourceState

	// List(String). URLs that Auth0 may redirect to after logout.
	AllowedLogoutUrls pulumi.StringArrayOutput `pulumi:"allowedLogoutUrls"`
	// List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
	ChangePassword TenantChangePasswordOutput `pulumi:"changePassword"`
	// String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
	DefaultAudience pulumi.StringOutput `pulumi:"defaultAudience"`
	// String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
	DefaultDirectory pulumi.StringOutput `pulumi:"defaultDirectory"`
	// String. The default absolute redirection uri, must be https and cannot contain a fragment.
	DefaultRedirectionUri pulumi.StringOutput      `pulumi:"defaultRedirectionUri"`
	EnabledLocales        pulumi.StringArrayOutput `pulumi:"enabledLocales"`
	// List(Resource). Configuration settings for error pages. For details, see Error Page.
	ErrorPage TenantErrorPageOutput `pulumi:"errorPage"`
	// List(Resource). Configuration settings for tenant flags. For details, see Flags.
	Flags TenantFlagsOutput `pulumi:"flags"`
	// String. Friendly name for the tenant.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
	GuardianMfaPage TenantGuardianMfaPageOutput `pulumi:"guardianMfaPage"`
	// Integer. Number of hours during which a session can be inactive before the user must log in again.
	IdleSessionLifetime pulumi.Float64Output `pulumi:"idleSessionLifetime"`
	// . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
	PictureUrl pulumi.StringOutput `pulumi:"pictureUrl"`
	// String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
	SandboxVersion pulumi.StringOutput `pulumi:"sandboxVersion"`
	// Integer. Number of hours during which a session will stay valid.
	SessionLifetime pulumi.Float64Output `pulumi:"sessionLifetime"`
	// String. Support email address for authenticating users.
	SupportEmail pulumi.StringOutput `pulumi:"supportEmail"`
	// String. Support URL for authenticating users.
	SupportUrl pulumi.StringOutput `pulumi:"supportUrl"`
	// List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
	UniversalLogin TenantUniversalLoginOutput `pulumi:"universalLogin"`
}

// NewTenant registers a new resource with the given unique name, arguments, and options.
func NewTenant(ctx *pulumi.Context,
	name string, args *TenantArgs, opts ...pulumi.ResourceOption) (*Tenant, error) {
	if args == nil {
		args = &TenantArgs{}
	}

	var resource Tenant
	err := ctx.RegisterResource("auth0:index/tenant:Tenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTenant gets an existing Tenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TenantState, opts ...pulumi.ResourceOption) (*Tenant, error) {
	var resource Tenant
	err := ctx.ReadResource("auth0:index/tenant:Tenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tenant resources.
type tenantState struct {
	// List(String). URLs that Auth0 may redirect to after logout.
	AllowedLogoutUrls []string `pulumi:"allowedLogoutUrls"`
	// List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
	ChangePassword *TenantChangePassword `pulumi:"changePassword"`
	// String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
	DefaultAudience *string `pulumi:"defaultAudience"`
	// String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
	DefaultDirectory *string `pulumi:"defaultDirectory"`
	// String. The default absolute redirection uri, must be https and cannot contain a fragment.
	DefaultRedirectionUri *string  `pulumi:"defaultRedirectionUri"`
	EnabledLocales        []string `pulumi:"enabledLocales"`
	// List(Resource). Configuration settings for error pages. For details, see Error Page.
	ErrorPage *TenantErrorPage `pulumi:"errorPage"`
	// List(Resource). Configuration settings for tenant flags. For details, see Flags.
	Flags *TenantFlags `pulumi:"flags"`
	// String. Friendly name for the tenant.
	FriendlyName *string `pulumi:"friendlyName"`
	// List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
	GuardianMfaPage *TenantGuardianMfaPage `pulumi:"guardianMfaPage"`
	// Integer. Number of hours during which a session can be inactive before the user must log in again.
	IdleSessionLifetime *float64 `pulumi:"idleSessionLifetime"`
	// . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
	PictureUrl *string `pulumi:"pictureUrl"`
	// String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
	SandboxVersion *string `pulumi:"sandboxVersion"`
	// Integer. Number of hours during which a session will stay valid.
	SessionLifetime *float64 `pulumi:"sessionLifetime"`
	// String. Support email address for authenticating users.
	SupportEmail *string `pulumi:"supportEmail"`
	// String. Support URL for authenticating users.
	SupportUrl *string `pulumi:"supportUrl"`
	// List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
	UniversalLogin *TenantUniversalLogin `pulumi:"universalLogin"`
}

type TenantState struct {
	// List(String). URLs that Auth0 may redirect to after logout.
	AllowedLogoutUrls pulumi.StringArrayInput
	// List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
	ChangePassword TenantChangePasswordPtrInput
	// String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
	DefaultAudience pulumi.StringPtrInput
	// String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
	DefaultDirectory pulumi.StringPtrInput
	// String. The default absolute redirection uri, must be https and cannot contain a fragment.
	DefaultRedirectionUri pulumi.StringPtrInput
	EnabledLocales        pulumi.StringArrayInput
	// List(Resource). Configuration settings for error pages. For details, see Error Page.
	ErrorPage TenantErrorPagePtrInput
	// List(Resource). Configuration settings for tenant flags. For details, see Flags.
	Flags TenantFlagsPtrInput
	// String. Friendly name for the tenant.
	FriendlyName pulumi.StringPtrInput
	// List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
	GuardianMfaPage TenantGuardianMfaPagePtrInput
	// Integer. Number of hours during which a session can be inactive before the user must log in again.
	IdleSessionLifetime pulumi.Float64PtrInput
	// . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
	PictureUrl pulumi.StringPtrInput
	// String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
	SandboxVersion pulumi.StringPtrInput
	// Integer. Number of hours during which a session will stay valid.
	SessionLifetime pulumi.Float64PtrInput
	// String. Support email address for authenticating users.
	SupportEmail pulumi.StringPtrInput
	// String. Support URL for authenticating users.
	SupportUrl pulumi.StringPtrInput
	// List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
	UniversalLogin TenantUniversalLoginPtrInput
}

func (TenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantState)(nil)).Elem()
}

type tenantArgs struct {
	// List(String). URLs that Auth0 may redirect to after logout.
	AllowedLogoutUrls []string `pulumi:"allowedLogoutUrls"`
	// List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
	ChangePassword *TenantChangePassword `pulumi:"changePassword"`
	// String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
	DefaultAudience *string `pulumi:"defaultAudience"`
	// String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
	DefaultDirectory *string `pulumi:"defaultDirectory"`
	// String. The default absolute redirection uri, must be https and cannot contain a fragment.
	DefaultRedirectionUri *string  `pulumi:"defaultRedirectionUri"`
	EnabledLocales        []string `pulumi:"enabledLocales"`
	// List(Resource). Configuration settings for error pages. For details, see Error Page.
	ErrorPage *TenantErrorPage `pulumi:"errorPage"`
	// List(Resource). Configuration settings for tenant flags. For details, see Flags.
	Flags *TenantFlags `pulumi:"flags"`
	// String. Friendly name for the tenant.
	FriendlyName *string `pulumi:"friendlyName"`
	// List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
	GuardianMfaPage *TenantGuardianMfaPage `pulumi:"guardianMfaPage"`
	// Integer. Number of hours during which a session can be inactive before the user must log in again.
	IdleSessionLifetime *float64 `pulumi:"idleSessionLifetime"`
	// . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
	PictureUrl *string `pulumi:"pictureUrl"`
	// String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
	SandboxVersion *string `pulumi:"sandboxVersion"`
	// Integer. Number of hours during which a session will stay valid.
	SessionLifetime *float64 `pulumi:"sessionLifetime"`
	// String. Support email address for authenticating users.
	SupportEmail *string `pulumi:"supportEmail"`
	// String. Support URL for authenticating users.
	SupportUrl *string `pulumi:"supportUrl"`
	// List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
	UniversalLogin *TenantUniversalLogin `pulumi:"universalLogin"`
}

// The set of arguments for constructing a Tenant resource.
type TenantArgs struct {
	// List(String). URLs that Auth0 may redirect to after logout.
	AllowedLogoutUrls pulumi.StringArrayInput
	// List(Resource). Configuration settings for change passsword page. For details, see Change Password Page.
	ChangePassword TenantChangePasswordPtrInput
	// String. API Audience to use by default for API Authorization flows. This setting is equivalent to appending the audience to every authorization request made to the tenant for every application.
	DefaultAudience pulumi.StringPtrInput
	// String. Name of the connection to be used for Password Grant exchanges. Options include `auth0-adldap`, `ad`, `auth0`, `email`, `sms`, `waad`, and `adfs`.
	DefaultDirectory pulumi.StringPtrInput
	// String. The default absolute redirection uri, must be https and cannot contain a fragment.
	DefaultRedirectionUri pulumi.StringPtrInput
	EnabledLocales        pulumi.StringArrayInput
	// List(Resource). Configuration settings for error pages. For details, see Error Page.
	ErrorPage TenantErrorPagePtrInput
	// List(Resource). Configuration settings for tenant flags. For details, see Flags.
	Flags TenantFlagsPtrInput
	// String. Friendly name for the tenant.
	FriendlyName pulumi.StringPtrInput
	// List(Resource). Configuration settings for the Guardian MFA page. For details, see Guardian MFA Page.
	GuardianMfaPage TenantGuardianMfaPagePtrInput
	// Integer. Number of hours during which a session can be inactive before the user must log in again.
	IdleSessionLifetime pulumi.Float64PtrInput
	// . String URL of logo to be shown for the tenant. Recommended size is 150px x 150px. If no URL is provided, the Auth0 logo will be used.
	PictureUrl pulumi.StringPtrInput
	// String. Selected sandbox version for the extensibility environment, which allows you to use custom scripts to extend parts of Auth0's functionality.
	SandboxVersion pulumi.StringPtrInput
	// Integer. Number of hours during which a session will stay valid.
	SessionLifetime pulumi.Float64PtrInput
	// String. Support email address for authenticating users.
	SupportEmail pulumi.StringPtrInput
	// String. Support URL for authenticating users.
	SupportUrl pulumi.StringPtrInput
	// List(Resource). Configuration settings for Universal Login. For details, see Universal Login.
	UniversalLogin TenantUniversalLoginPtrInput
}

func (TenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tenantArgs)(nil)).Elem()
}

type TenantInput interface {
	pulumi.Input

	ToTenantOutput() TenantOutput
	ToTenantOutputWithContext(ctx context.Context) TenantOutput
}

func (Tenant) ElementType() reflect.Type {
	return reflect.TypeOf((*Tenant)(nil)).Elem()
}

func (i Tenant) ToTenantOutput() TenantOutput {
	return i.ToTenantOutputWithContext(context.Background())
}

func (i Tenant) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TenantOutput)
}

type TenantOutput struct {
	*pulumi.OutputState
}

func (TenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TenantOutput)(nil)).Elem()
}

func (o TenantOutput) ToTenantOutput() TenantOutput {
	return o
}

func (o TenantOutput) ToTenantOutputWithContext(ctx context.Context) TenantOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TenantOutput{})
}
