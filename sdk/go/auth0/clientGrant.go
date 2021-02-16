// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Auth0 uses various grant types, or methods by which you grant limited access to your resources to another entity without exposing credentials. The OAuth 2.0 protocol supports several types of grants, which allow different types of access. This resource allows you to create and manage client grants used with configured Auth0 clients.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-auth0/sdk/go/auth0"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myClient, err := auth0.NewClient(ctx, "myClient", nil)
// 		if err != nil {
// 			return err
// 		}
// 		myResourceServer, err := auth0.NewResourceServer(ctx, "myResourceServer", &auth0.ResourceServerArgs{
// 			Identifier: pulumi.String("https://api.example.com/client-grant"),
// 			Scopes: auth0.ResourceServerScopeArray{
// 				&auth0.ResourceServerScopeArgs{
// 					Description: pulumi.String("Create foos"),
// 					Value:       pulumi.String("create:foo"),
// 				},
// 				&auth0.ResourceServerScopeArgs{
// 					Description: pulumi.String("Create bars"),
// 					Value:       pulumi.String("create:bar"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = auth0.NewClientGrant(ctx, "myClientGrant", &auth0.ClientGrantArgs{
// 			Audience: myResourceServer.Identifier,
// 			ClientId: myClient.ID(),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("create:foo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ClientGrant struct {
	pulumi.CustomResourceState

	// String. Audience or API Identifier for this grant.
	Audience pulumi.StringOutput `pulumi:"audience"`
	// String. ID of the client for this grant.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// List(String). Permissions (scopes) included in this grant.
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewClientGrant registers a new resource with the given unique name, arguments, and options.
func NewClientGrant(ctx *pulumi.Context,
	name string, args *ClientGrantArgs, opts ...pulumi.ResourceOption) (*ClientGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Audience == nil {
		return nil, errors.New("invalid value for required argument 'Audience'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	var resource ClientGrant
	err := ctx.RegisterResource("auth0:index/clientGrant:ClientGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientGrant gets an existing ClientGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientGrantState, opts ...pulumi.ResourceOption) (*ClientGrant, error) {
	var resource ClientGrant
	err := ctx.ReadResource("auth0:index/clientGrant:ClientGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientGrant resources.
type clientGrantState struct {
	// String. Audience or API Identifier for this grant.
	Audience *string `pulumi:"audience"`
	// String. ID of the client for this grant.
	ClientId *string `pulumi:"clientId"`
	// List(String). Permissions (scopes) included in this grant.
	Scopes []string `pulumi:"scopes"`
}

type ClientGrantState struct {
	// String. Audience or API Identifier for this grant.
	Audience pulumi.StringPtrInput
	// String. ID of the client for this grant.
	ClientId pulumi.StringPtrInput
	// List(String). Permissions (scopes) included in this grant.
	Scopes pulumi.StringArrayInput
}

func (ClientGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGrantState)(nil)).Elem()
}

type clientGrantArgs struct {
	// String. Audience or API Identifier for this grant.
	Audience string `pulumi:"audience"`
	// String. ID of the client for this grant.
	ClientId string `pulumi:"clientId"`
	// List(String). Permissions (scopes) included in this grant.
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a ClientGrant resource.
type ClientGrantArgs struct {
	// String. Audience or API Identifier for this grant.
	Audience pulumi.StringInput
	// String. ID of the client for this grant.
	ClientId pulumi.StringInput
	// List(String). Permissions (scopes) included in this grant.
	Scopes pulumi.StringArrayInput
}

func (ClientGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGrantArgs)(nil)).Elem()
}

type ClientGrantInput interface {
	pulumi.Input

	ToClientGrantOutput() ClientGrantOutput
	ToClientGrantOutputWithContext(ctx context.Context) ClientGrantOutput
}

func (*ClientGrant) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGrant)(nil))
}

func (i *ClientGrant) ToClientGrantOutput() ClientGrantOutput {
	return i.ToClientGrantOutputWithContext(context.Background())
}

func (i *ClientGrant) ToClientGrantOutputWithContext(ctx context.Context) ClientGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGrantOutput)
}

func (i *ClientGrant) ToClientGrantPtrOutput() ClientGrantPtrOutput {
	return i.ToClientGrantPtrOutputWithContext(context.Background())
}

func (i *ClientGrant) ToClientGrantPtrOutputWithContext(ctx context.Context) ClientGrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGrantPtrOutput)
}

type ClientGrantPtrInput interface {
	pulumi.Input

	ToClientGrantPtrOutput() ClientGrantPtrOutput
	ToClientGrantPtrOutputWithContext(ctx context.Context) ClientGrantPtrOutput
}

type clientGrantPtrType ClientGrantArgs

func (*clientGrantPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGrant)(nil))
}

func (i *clientGrantPtrType) ToClientGrantPtrOutput() ClientGrantPtrOutput {
	return i.ToClientGrantPtrOutputWithContext(context.Background())
}

func (i *clientGrantPtrType) ToClientGrantPtrOutputWithContext(ctx context.Context) ClientGrantPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGrantPtrOutput)
}

// ClientGrantArrayInput is an input type that accepts ClientGrantArray and ClientGrantArrayOutput values.
// You can construct a concrete instance of `ClientGrantArrayInput` via:
//
//          ClientGrantArray{ ClientGrantArgs{...} }
type ClientGrantArrayInput interface {
	pulumi.Input

	ToClientGrantArrayOutput() ClientGrantArrayOutput
	ToClientGrantArrayOutputWithContext(context.Context) ClientGrantArrayOutput
}

type ClientGrantArray []ClientGrantInput

func (ClientGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ClientGrant)(nil))
}

func (i ClientGrantArray) ToClientGrantArrayOutput() ClientGrantArrayOutput {
	return i.ToClientGrantArrayOutputWithContext(context.Background())
}

func (i ClientGrantArray) ToClientGrantArrayOutputWithContext(ctx context.Context) ClientGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGrantArrayOutput)
}

// ClientGrantMapInput is an input type that accepts ClientGrantMap and ClientGrantMapOutput values.
// You can construct a concrete instance of `ClientGrantMapInput` via:
//
//          ClientGrantMap{ "key": ClientGrantArgs{...} }
type ClientGrantMapInput interface {
	pulumi.Input

	ToClientGrantMapOutput() ClientGrantMapOutput
	ToClientGrantMapOutputWithContext(context.Context) ClientGrantMapOutput
}

type ClientGrantMap map[string]ClientGrantInput

func (ClientGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ClientGrant)(nil))
}

func (i ClientGrantMap) ToClientGrantMapOutput() ClientGrantMapOutput {
	return i.ToClientGrantMapOutputWithContext(context.Background())
}

func (i ClientGrantMap) ToClientGrantMapOutputWithContext(ctx context.Context) ClientGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGrantMapOutput)
}

type ClientGrantOutput struct {
	*pulumi.OutputState
}

func (ClientGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGrant)(nil))
}

func (o ClientGrantOutput) ToClientGrantOutput() ClientGrantOutput {
	return o
}

func (o ClientGrantOutput) ToClientGrantOutputWithContext(ctx context.Context) ClientGrantOutput {
	return o
}

func (o ClientGrantOutput) ToClientGrantPtrOutput() ClientGrantPtrOutput {
	return o.ToClientGrantPtrOutputWithContext(context.Background())
}

func (o ClientGrantOutput) ToClientGrantPtrOutputWithContext(ctx context.Context) ClientGrantPtrOutput {
	return o.ApplyT(func(v ClientGrant) *ClientGrant {
		return &v
	}).(ClientGrantPtrOutput)
}

type ClientGrantPtrOutput struct {
	*pulumi.OutputState
}

func (ClientGrantPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientGrant)(nil))
}

func (o ClientGrantPtrOutput) ToClientGrantPtrOutput() ClientGrantPtrOutput {
	return o
}

func (o ClientGrantPtrOutput) ToClientGrantPtrOutputWithContext(ctx context.Context) ClientGrantPtrOutput {
	return o
}

type ClientGrantArrayOutput struct{ *pulumi.OutputState }

func (ClientGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientGrant)(nil))
}

func (o ClientGrantArrayOutput) ToClientGrantArrayOutput() ClientGrantArrayOutput {
	return o
}

func (o ClientGrantArrayOutput) ToClientGrantArrayOutputWithContext(ctx context.Context) ClientGrantArrayOutput {
	return o
}

func (o ClientGrantArrayOutput) Index(i pulumi.IntInput) ClientGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientGrant {
		return vs[0].([]ClientGrant)[vs[1].(int)]
	}).(ClientGrantOutput)
}

type ClientGrantMapOutput struct{ *pulumi.OutputState }

func (ClientGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ClientGrant)(nil))
}

func (o ClientGrantMapOutput) ToClientGrantMapOutput() ClientGrantMapOutput {
	return o
}

func (o ClientGrantMapOutput) ToClientGrantMapOutputWithContext(ctx context.Context) ClientGrantMapOutput {
	return o
}

func (o ClientGrantMapOutput) MapIndex(k pulumi.StringInput) ClientGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ClientGrant {
		return vs[0].(map[string]ClientGrant)[vs[1].(string)]
	}).(ClientGrantOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientGrantOutput{})
	pulumi.RegisterOutputType(ClientGrantPtrOutput{})
	pulumi.RegisterOutputType(ClientGrantArrayOutput{})
	pulumi.RegisterOutputType(ClientGrantMapOutput{})
}
