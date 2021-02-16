// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GlobalClient struct {
	pulumi.CustomResourceState

	Addons                         GlobalClientAddonsOutput           `pulumi:"addons"`
	AllowedLogoutUrls              pulumi.StringArrayOutput           `pulumi:"allowedLogoutUrls"`
	AllowedOrigins                 pulumi.StringArrayOutput           `pulumi:"allowedOrigins"`
	AppType                        pulumi.StringOutput                `pulumi:"appType"`
	Callbacks                      pulumi.StringArrayOutput           `pulumi:"callbacks"`
	ClientId                       pulumi.StringOutput                `pulumi:"clientId"`
	ClientMetadata                 pulumi.MapOutput                   `pulumi:"clientMetadata"`
	ClientSecret                   pulumi.StringOutput                `pulumi:"clientSecret"`
	ClientSecretRotationTrigger    pulumi.MapOutput                   `pulumi:"clientSecretRotationTrigger"`
	CrossOriginAuth                pulumi.BoolOutput                  `pulumi:"crossOriginAuth"`
	CrossOriginLoc                 pulumi.StringOutput                `pulumi:"crossOriginLoc"`
	CustomLoginPage                pulumi.StringOutput                `pulumi:"customLoginPage"`
	CustomLoginPageOn              pulumi.BoolOutput                  `pulumi:"customLoginPageOn"`
	CustomLoginPagePreview         pulumi.StringOutput                `pulumi:"customLoginPagePreview"`
	Description                    pulumi.StringOutput                `pulumi:"description"`
	EncryptionKey                  pulumi.StringMapOutput             `pulumi:"encryptionKey"`
	FormTemplate                   pulumi.StringOutput                `pulumi:"formTemplate"`
	GrantTypes                     pulumi.StringArrayOutput           `pulumi:"grantTypes"`
	InitiateLoginUri               pulumi.StringOutput                `pulumi:"initiateLoginUri"`
	IsFirstParty                   pulumi.BoolOutput                  `pulumi:"isFirstParty"`
	IsTokenEndpointIpHeaderTrusted pulumi.BoolOutput                  `pulumi:"isTokenEndpointIpHeaderTrusted"`
	JwtConfiguration               GlobalClientJwtConfigurationOutput `pulumi:"jwtConfiguration"`
	LogoUri                        pulumi.StringOutput                `pulumi:"logoUri"`
	Mobile                         GlobalClientMobileOutput           `pulumi:"mobile"`
	Name                           pulumi.StringOutput                `pulumi:"name"`
	OidcConformant                 pulumi.BoolOutput                  `pulumi:"oidcConformant"`
	RefreshToken                   GlobalClientRefreshTokenOutput     `pulumi:"refreshToken"`
	Sso                            pulumi.BoolOutput                  `pulumi:"sso"`
	SsoDisabled                    pulumi.BoolOutput                  `pulumi:"ssoDisabled"`
	TokenEndpointAuthMethod        pulumi.StringOutput                `pulumi:"tokenEndpointAuthMethod"`
	WebOrigins                     pulumi.StringArrayOutput           `pulumi:"webOrigins"`
}

// NewGlobalClient registers a new resource with the given unique name, arguments, and options.
func NewGlobalClient(ctx *pulumi.Context,
	name string, args *GlobalClientArgs, opts ...pulumi.ResourceOption) (*GlobalClient, error) {
	if args == nil {
		args = &GlobalClientArgs{}
	}

	var resource GlobalClient
	err := ctx.RegisterResource("auth0:index/globalClient:GlobalClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalClient gets an existing GlobalClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalClientState, opts ...pulumi.ResourceOption) (*GlobalClient, error) {
	var resource GlobalClient
	err := ctx.ReadResource("auth0:index/globalClient:GlobalClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalClient resources.
type globalClientState struct {
	Addons                         *GlobalClientAddons           `pulumi:"addons"`
	AllowedLogoutUrls              []string                      `pulumi:"allowedLogoutUrls"`
	AllowedOrigins                 []string                      `pulumi:"allowedOrigins"`
	AppType                        *string                       `pulumi:"appType"`
	Callbacks                      []string                      `pulumi:"callbacks"`
	ClientId                       *string                       `pulumi:"clientId"`
	ClientMetadata                 map[string]interface{}        `pulumi:"clientMetadata"`
	ClientSecret                   *string                       `pulumi:"clientSecret"`
	ClientSecretRotationTrigger    map[string]interface{}        `pulumi:"clientSecretRotationTrigger"`
	CrossOriginAuth                *bool                         `pulumi:"crossOriginAuth"`
	CrossOriginLoc                 *string                       `pulumi:"crossOriginLoc"`
	CustomLoginPage                *string                       `pulumi:"customLoginPage"`
	CustomLoginPageOn              *bool                         `pulumi:"customLoginPageOn"`
	CustomLoginPagePreview         *string                       `pulumi:"customLoginPagePreview"`
	Description                    *string                       `pulumi:"description"`
	EncryptionKey                  map[string]string             `pulumi:"encryptionKey"`
	FormTemplate                   *string                       `pulumi:"formTemplate"`
	GrantTypes                     []string                      `pulumi:"grantTypes"`
	InitiateLoginUri               *string                       `pulumi:"initiateLoginUri"`
	IsFirstParty                   *bool                         `pulumi:"isFirstParty"`
	IsTokenEndpointIpHeaderTrusted *bool                         `pulumi:"isTokenEndpointIpHeaderTrusted"`
	JwtConfiguration               *GlobalClientJwtConfiguration `pulumi:"jwtConfiguration"`
	LogoUri                        *string                       `pulumi:"logoUri"`
	Mobile                         *GlobalClientMobile           `pulumi:"mobile"`
	Name                           *string                       `pulumi:"name"`
	OidcConformant                 *bool                         `pulumi:"oidcConformant"`
	RefreshToken                   *GlobalClientRefreshToken     `pulumi:"refreshToken"`
	Sso                            *bool                         `pulumi:"sso"`
	SsoDisabled                    *bool                         `pulumi:"ssoDisabled"`
	TokenEndpointAuthMethod        *string                       `pulumi:"tokenEndpointAuthMethod"`
	WebOrigins                     []string                      `pulumi:"webOrigins"`
}

type GlobalClientState struct {
	Addons                         GlobalClientAddonsPtrInput
	AllowedLogoutUrls              pulumi.StringArrayInput
	AllowedOrigins                 pulumi.StringArrayInput
	AppType                        pulumi.StringPtrInput
	Callbacks                      pulumi.StringArrayInput
	ClientId                       pulumi.StringPtrInput
	ClientMetadata                 pulumi.MapInput
	ClientSecret                   pulumi.StringPtrInput
	ClientSecretRotationTrigger    pulumi.MapInput
	CrossOriginAuth                pulumi.BoolPtrInput
	CrossOriginLoc                 pulumi.StringPtrInput
	CustomLoginPage                pulumi.StringPtrInput
	CustomLoginPageOn              pulumi.BoolPtrInput
	CustomLoginPagePreview         pulumi.StringPtrInput
	Description                    pulumi.StringPtrInput
	EncryptionKey                  pulumi.StringMapInput
	FormTemplate                   pulumi.StringPtrInput
	GrantTypes                     pulumi.StringArrayInput
	InitiateLoginUri               pulumi.StringPtrInput
	IsFirstParty                   pulumi.BoolPtrInput
	IsTokenEndpointIpHeaderTrusted pulumi.BoolPtrInput
	JwtConfiguration               GlobalClientJwtConfigurationPtrInput
	LogoUri                        pulumi.StringPtrInput
	Mobile                         GlobalClientMobilePtrInput
	Name                           pulumi.StringPtrInput
	OidcConformant                 pulumi.BoolPtrInput
	RefreshToken                   GlobalClientRefreshTokenPtrInput
	Sso                            pulumi.BoolPtrInput
	SsoDisabled                    pulumi.BoolPtrInput
	TokenEndpointAuthMethod        pulumi.StringPtrInput
	WebOrigins                     pulumi.StringArrayInput
}

func (GlobalClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClientState)(nil)).Elem()
}

type globalClientArgs struct {
	Addons                         *GlobalClientAddons           `pulumi:"addons"`
	AllowedLogoutUrls              []string                      `pulumi:"allowedLogoutUrls"`
	AllowedOrigins                 []string                      `pulumi:"allowedOrigins"`
	AppType                        *string                       `pulumi:"appType"`
	Callbacks                      []string                      `pulumi:"callbacks"`
	ClientId                       *string                       `pulumi:"clientId"`
	ClientMetadata                 map[string]interface{}        `pulumi:"clientMetadata"`
	ClientSecret                   *string                       `pulumi:"clientSecret"`
	ClientSecretRotationTrigger    map[string]interface{}        `pulumi:"clientSecretRotationTrigger"`
	CrossOriginAuth                *bool                         `pulumi:"crossOriginAuth"`
	CrossOriginLoc                 *string                       `pulumi:"crossOriginLoc"`
	CustomLoginPage                *string                       `pulumi:"customLoginPage"`
	CustomLoginPageOn              *bool                         `pulumi:"customLoginPageOn"`
	CustomLoginPagePreview         *string                       `pulumi:"customLoginPagePreview"`
	Description                    *string                       `pulumi:"description"`
	EncryptionKey                  map[string]string             `pulumi:"encryptionKey"`
	FormTemplate                   *string                       `pulumi:"formTemplate"`
	GrantTypes                     []string                      `pulumi:"grantTypes"`
	InitiateLoginUri               *string                       `pulumi:"initiateLoginUri"`
	IsFirstParty                   *bool                         `pulumi:"isFirstParty"`
	IsTokenEndpointIpHeaderTrusted *bool                         `pulumi:"isTokenEndpointIpHeaderTrusted"`
	JwtConfiguration               *GlobalClientJwtConfiguration `pulumi:"jwtConfiguration"`
	LogoUri                        *string                       `pulumi:"logoUri"`
	Mobile                         *GlobalClientMobile           `pulumi:"mobile"`
	Name                           *string                       `pulumi:"name"`
	OidcConformant                 *bool                         `pulumi:"oidcConformant"`
	RefreshToken                   *GlobalClientRefreshToken     `pulumi:"refreshToken"`
	Sso                            *bool                         `pulumi:"sso"`
	SsoDisabled                    *bool                         `pulumi:"ssoDisabled"`
	TokenEndpointAuthMethod        *string                       `pulumi:"tokenEndpointAuthMethod"`
	WebOrigins                     []string                      `pulumi:"webOrigins"`
}

// The set of arguments for constructing a GlobalClient resource.
type GlobalClientArgs struct {
	Addons                         GlobalClientAddonsPtrInput
	AllowedLogoutUrls              pulumi.StringArrayInput
	AllowedOrigins                 pulumi.StringArrayInput
	AppType                        pulumi.StringPtrInput
	Callbacks                      pulumi.StringArrayInput
	ClientId                       pulumi.StringPtrInput
	ClientMetadata                 pulumi.MapInput
	ClientSecret                   pulumi.StringPtrInput
	ClientSecretRotationTrigger    pulumi.MapInput
	CrossOriginAuth                pulumi.BoolPtrInput
	CrossOriginLoc                 pulumi.StringPtrInput
	CustomLoginPage                pulumi.StringPtrInput
	CustomLoginPageOn              pulumi.BoolPtrInput
	CustomLoginPagePreview         pulumi.StringPtrInput
	Description                    pulumi.StringPtrInput
	EncryptionKey                  pulumi.StringMapInput
	FormTemplate                   pulumi.StringPtrInput
	GrantTypes                     pulumi.StringArrayInput
	InitiateLoginUri               pulumi.StringPtrInput
	IsFirstParty                   pulumi.BoolPtrInput
	IsTokenEndpointIpHeaderTrusted pulumi.BoolPtrInput
	JwtConfiguration               GlobalClientJwtConfigurationPtrInput
	LogoUri                        pulumi.StringPtrInput
	Mobile                         GlobalClientMobilePtrInput
	Name                           pulumi.StringPtrInput
	OidcConformant                 pulumi.BoolPtrInput
	RefreshToken                   GlobalClientRefreshTokenPtrInput
	Sso                            pulumi.BoolPtrInput
	SsoDisabled                    pulumi.BoolPtrInput
	TokenEndpointAuthMethod        pulumi.StringPtrInput
	WebOrigins                     pulumi.StringArrayInput
}

func (GlobalClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClientArgs)(nil)).Elem()
}

type GlobalClientInput interface {
	pulumi.Input

	ToGlobalClientOutput() GlobalClientOutput
	ToGlobalClientOutputWithContext(ctx context.Context) GlobalClientOutput
}

func (*GlobalClient) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalClient)(nil))
}

func (i *GlobalClient) ToGlobalClientOutput() GlobalClientOutput {
	return i.ToGlobalClientOutputWithContext(context.Background())
}

func (i *GlobalClient) ToGlobalClientOutputWithContext(ctx context.Context) GlobalClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClientOutput)
}

func (i *GlobalClient) ToGlobalClientPtrOutput() GlobalClientPtrOutput {
	return i.ToGlobalClientPtrOutputWithContext(context.Background())
}

func (i *GlobalClient) ToGlobalClientPtrOutputWithContext(ctx context.Context) GlobalClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClientPtrOutput)
}

type GlobalClientPtrInput interface {
	pulumi.Input

	ToGlobalClientPtrOutput() GlobalClientPtrOutput
	ToGlobalClientPtrOutputWithContext(ctx context.Context) GlobalClientPtrOutput
}

type globalClientPtrType GlobalClientArgs

func (*globalClientPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalClient)(nil))
}

func (i *globalClientPtrType) ToGlobalClientPtrOutput() GlobalClientPtrOutput {
	return i.ToGlobalClientPtrOutputWithContext(context.Background())
}

func (i *globalClientPtrType) ToGlobalClientPtrOutputWithContext(ctx context.Context) GlobalClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClientPtrOutput)
}

// GlobalClientArrayInput is an input type that accepts GlobalClientArray and GlobalClientArrayOutput values.
// You can construct a concrete instance of `GlobalClientArrayInput` via:
//
//          GlobalClientArray{ GlobalClientArgs{...} }
type GlobalClientArrayInput interface {
	pulumi.Input

	ToGlobalClientArrayOutput() GlobalClientArrayOutput
	ToGlobalClientArrayOutputWithContext(context.Context) GlobalClientArrayOutput
}

type GlobalClientArray []GlobalClientInput

func (GlobalClientArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GlobalClient)(nil))
}

func (i GlobalClientArray) ToGlobalClientArrayOutput() GlobalClientArrayOutput {
	return i.ToGlobalClientArrayOutputWithContext(context.Background())
}

func (i GlobalClientArray) ToGlobalClientArrayOutputWithContext(ctx context.Context) GlobalClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClientArrayOutput)
}

// GlobalClientMapInput is an input type that accepts GlobalClientMap and GlobalClientMapOutput values.
// You can construct a concrete instance of `GlobalClientMapInput` via:
//
//          GlobalClientMap{ "key": GlobalClientArgs{...} }
type GlobalClientMapInput interface {
	pulumi.Input

	ToGlobalClientMapOutput() GlobalClientMapOutput
	ToGlobalClientMapOutputWithContext(context.Context) GlobalClientMapOutput
}

type GlobalClientMap map[string]GlobalClientInput

func (GlobalClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GlobalClient)(nil))
}

func (i GlobalClientMap) ToGlobalClientMapOutput() GlobalClientMapOutput {
	return i.ToGlobalClientMapOutputWithContext(context.Background())
}

func (i GlobalClientMap) ToGlobalClientMapOutputWithContext(ctx context.Context) GlobalClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClientMapOutput)
}

type GlobalClientOutput struct {
	*pulumi.OutputState
}

func (GlobalClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalClient)(nil))
}

func (o GlobalClientOutput) ToGlobalClientOutput() GlobalClientOutput {
	return o
}

func (o GlobalClientOutput) ToGlobalClientOutputWithContext(ctx context.Context) GlobalClientOutput {
	return o
}

func (o GlobalClientOutput) ToGlobalClientPtrOutput() GlobalClientPtrOutput {
	return o.ToGlobalClientPtrOutputWithContext(context.Background())
}

func (o GlobalClientOutput) ToGlobalClientPtrOutputWithContext(ctx context.Context) GlobalClientPtrOutput {
	return o.ApplyT(func(v GlobalClient) *GlobalClient {
		return &v
	}).(GlobalClientPtrOutput)
}

type GlobalClientPtrOutput struct {
	*pulumi.OutputState
}

func (GlobalClientPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalClient)(nil))
}

func (o GlobalClientPtrOutput) ToGlobalClientPtrOutput() GlobalClientPtrOutput {
	return o
}

func (o GlobalClientPtrOutput) ToGlobalClientPtrOutputWithContext(ctx context.Context) GlobalClientPtrOutput {
	return o
}

type GlobalClientArrayOutput struct{ *pulumi.OutputState }

func (GlobalClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalClient)(nil))
}

func (o GlobalClientArrayOutput) ToGlobalClientArrayOutput() GlobalClientArrayOutput {
	return o
}

func (o GlobalClientArrayOutput) ToGlobalClientArrayOutputWithContext(ctx context.Context) GlobalClientArrayOutput {
	return o
}

func (o GlobalClientArrayOutput) Index(i pulumi.IntInput) GlobalClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalClient {
		return vs[0].([]GlobalClient)[vs[1].(int)]
	}).(GlobalClientOutput)
}

type GlobalClientMapOutput struct{ *pulumi.OutputState }

func (GlobalClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GlobalClient)(nil))
}

func (o GlobalClientMapOutput) ToGlobalClientMapOutput() GlobalClientMapOutput {
	return o
}

func (o GlobalClientMapOutput) ToGlobalClientMapOutputWithContext(ctx context.Context) GlobalClientMapOutput {
	return o
}

func (o GlobalClientMapOutput) MapIndex(k pulumi.StringInput) GlobalClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GlobalClient {
		return vs[0].(map[string]GlobalClient)[vs[1].(string)]
	}).(GlobalClientOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalClientOutput{})
	pulumi.RegisterOutputType(GlobalClientPtrOutput{})
	pulumi.RegisterOutputType(GlobalClientArrayOutput{})
	pulumi.RegisterOutputType(GlobalClientMapOutput{})
}
