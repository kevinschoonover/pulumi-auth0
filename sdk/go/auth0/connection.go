// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// With Auth0, you can define sources of users, otherwise known as connections, which may include identity providers (such as Google or LinkedIn), databases, or passwordless authentication methods. This resource allows you to configure and manage connections to be used with your clients and users.
type Connection struct {
	pulumi.CustomResourceState

	// Name used in login screen
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
	EnabledClients pulumi.StringArrayOutput `pulumi:"enabledClients"`
	// Boolean. Indicates whether or not the connection is domain level.
	IsDomainConnection pulumi.BoolOutput `pulumi:"isDomainConnection"`
	// String. Name of the connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// List(Resource). Configuration settings for connection options. For details, see Options.
	Options ConnectionOptionsPtrOutput `pulumi:"options"`
	// List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
	Realms pulumi.StringArrayOutput `pulumi:"realms"`
	// String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil || args.Strategy == nil {
		return nil, errors.New("missing required argument 'Strategy'")
	}
	if args == nil {
		args = &ConnectionArgs{}
	}
	var resource Connection
	err := ctx.RegisterResource("auth0:index/connection:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("auth0:index/connection:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Name used in login screen
	DisplayName *string `pulumi:"displayName"`
	// Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
	EnabledClients []string `pulumi:"enabledClients"`
	// Boolean. Indicates whether or not the connection is domain level.
	IsDomainConnection *bool `pulumi:"isDomainConnection"`
	// String. Name of the connection.
	Name *string `pulumi:"name"`
	// List(Resource). Configuration settings for connection options. For details, see Options.
	Options *ConnectionOptions `pulumi:"options"`
	// List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
	Realms []string `pulumi:"realms"`
	// String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
	Strategy *string `pulumi:"strategy"`
}

type ConnectionState struct {
	// Name used in login screen
	DisplayName pulumi.StringPtrInput
	// Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
	EnabledClients pulumi.StringArrayInput
	// Boolean. Indicates whether or not the connection is domain level.
	IsDomainConnection pulumi.BoolPtrInput
	// String. Name of the connection.
	Name pulumi.StringPtrInput
	// List(Resource). Configuration settings for connection options. For details, see Options.
	Options ConnectionOptionsPtrInput
	// List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
	Realms pulumi.StringArrayInput
	// String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
	Strategy pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Name used in login screen
	DisplayName *string `pulumi:"displayName"`
	// Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
	EnabledClients []string `pulumi:"enabledClients"`
	// Boolean. Indicates whether or not the connection is domain level.
	IsDomainConnection *bool `pulumi:"isDomainConnection"`
	// String. Name of the connection.
	Name *string `pulumi:"name"`
	// List(Resource). Configuration settings for connection options. For details, see Options.
	Options *ConnectionOptions `pulumi:"options"`
	// List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
	Realms []string `pulumi:"realms"`
	// String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
	Strategy string `pulumi:"strategy"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Name used in login screen
	DisplayName pulumi.StringPtrInput
	// Set(String). IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
	EnabledClients pulumi.StringArrayInput
	// Boolean. Indicates whether or not the connection is domain level.
	IsDomainConnection pulumi.BoolPtrInput
	// String. Name of the connection.
	Name pulumi.StringPtrInput
	// List(Resource). Configuration settings for connection options. For details, see Options.
	Options ConnectionOptionsPtrInput
	// List(String). Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.
	Realms pulumi.StringArrayInput
	// String. Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
	Strategy pulumi.StringInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}