// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// With Auth0, you can have standard welcome, password reset, and account verification email-based workflows built right into Auth0. This resource allows you to configure email providers so you can route all emails that are part of Auth0's authentication workflows through the supported high-volume email service of your choice.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-auth0/sdk/v2/go/auth0"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := auth0.NewEmail(ctx, "myEmailProvider", &auth0.EmailArgs{
// 			Credentials: &EmailCredentialsArgs{
// 				AccessKeyId:     pulumi.String("AKIAXXXXXXXXXXXXXXXX"),
// 				Region:          pulumi.String("us-east-1"),
// 				SecretAccessKey: pulumi.String("7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
// 			},
// 			DefaultFromAddress: pulumi.String("accounts@example.com"),
// 			Enabled:            pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Email struct {
	pulumi.CustomResourceState

	// List(Resource). Configuration settings for the credentials for the email provider. For details, see Credentials.
	Credentials EmailCredentialsOutput `pulumi:"credentials"`
	// String. Email address to use as the sender when no other "from" address is specified.
	DefaultFromAddress pulumi.StringOutput `pulumi:"defaultFromAddress"`
	// Boolean. Indicates whether or not the email provider is enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// String. Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEmail registers a new resource with the given unique name, arguments, and options.
func NewEmail(ctx *pulumi.Context,
	name string, args *EmailArgs, opts ...pulumi.ResourceOption) (*Email, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Credentials == nil {
		return nil, errors.New("invalid value for required argument 'Credentials'")
	}
	if args.DefaultFromAddress == nil {
		return nil, errors.New("invalid value for required argument 'DefaultFromAddress'")
	}
	var resource Email
	err := ctx.RegisterResource("auth0:index/email:Email", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmail gets an existing Email resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailState, opts ...pulumi.ResourceOption) (*Email, error) {
	var resource Email
	err := ctx.ReadResource("auth0:index/email:Email", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Email resources.
type emailState struct {
	// List(Resource). Configuration settings for the credentials for the email provider. For details, see Credentials.
	Credentials *EmailCredentials `pulumi:"credentials"`
	// String. Email address to use as the sender when no other "from" address is specified.
	DefaultFromAddress *string `pulumi:"defaultFromAddress"`
	// Boolean. Indicates whether or not the email provider is enabled.
	Enabled *bool `pulumi:"enabled"`
	// String. Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`.
	Name *string `pulumi:"name"`
}

type EmailState struct {
	// List(Resource). Configuration settings for the credentials for the email provider. For details, see Credentials.
	Credentials EmailCredentialsPtrInput
	// String. Email address to use as the sender when no other "from" address is specified.
	DefaultFromAddress pulumi.StringPtrInput
	// Boolean. Indicates whether or not the email provider is enabled.
	Enabled pulumi.BoolPtrInput
	// String. Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`.
	Name pulumi.StringPtrInput
}

func (EmailState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailState)(nil)).Elem()
}

type emailArgs struct {
	// List(Resource). Configuration settings for the credentials for the email provider. For details, see Credentials.
	Credentials EmailCredentials `pulumi:"credentials"`
	// String. Email address to use as the sender when no other "from" address is specified.
	DefaultFromAddress string `pulumi:"defaultFromAddress"`
	// Boolean. Indicates whether or not the email provider is enabled.
	Enabled *bool `pulumi:"enabled"`
	// String. Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Email resource.
type EmailArgs struct {
	// List(Resource). Configuration settings for the credentials for the email provider. For details, see Credentials.
	Credentials EmailCredentialsInput
	// String. Email address to use as the sender when no other "from" address is specified.
	DefaultFromAddress pulumi.StringInput
	// Boolean. Indicates whether or not the email provider is enabled.
	Enabled pulumi.BoolPtrInput
	// String. Name of the email provider. Options include `mailgun`, `mandrill`, `sendgrid`, `ses`, `smtp`, and `sparkpost`.
	Name pulumi.StringPtrInput
}

func (EmailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailArgs)(nil)).Elem()
}

type EmailInput interface {
	pulumi.Input

	ToEmailOutput() EmailOutput
	ToEmailOutputWithContext(ctx context.Context) EmailOutput
}

func (*Email) ElementType() reflect.Type {
	return reflect.TypeOf((**Email)(nil)).Elem()
}

func (i *Email) ToEmailOutput() EmailOutput {
	return i.ToEmailOutputWithContext(context.Background())
}

func (i *Email) ToEmailOutputWithContext(ctx context.Context) EmailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailOutput)
}

// EmailArrayInput is an input type that accepts EmailArray and EmailArrayOutput values.
// You can construct a concrete instance of `EmailArrayInput` via:
//
//          EmailArray{ EmailArgs{...} }
type EmailArrayInput interface {
	pulumi.Input

	ToEmailArrayOutput() EmailArrayOutput
	ToEmailArrayOutputWithContext(context.Context) EmailArrayOutput
}

type EmailArray []EmailInput

func (EmailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Email)(nil)).Elem()
}

func (i EmailArray) ToEmailArrayOutput() EmailArrayOutput {
	return i.ToEmailArrayOutputWithContext(context.Background())
}

func (i EmailArray) ToEmailArrayOutputWithContext(ctx context.Context) EmailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailArrayOutput)
}

// EmailMapInput is an input type that accepts EmailMap and EmailMapOutput values.
// You can construct a concrete instance of `EmailMapInput` via:
//
//          EmailMap{ "key": EmailArgs{...} }
type EmailMapInput interface {
	pulumi.Input

	ToEmailMapOutput() EmailMapOutput
	ToEmailMapOutputWithContext(context.Context) EmailMapOutput
}

type EmailMap map[string]EmailInput

func (EmailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Email)(nil)).Elem()
}

func (i EmailMap) ToEmailMapOutput() EmailMapOutput {
	return i.ToEmailMapOutputWithContext(context.Background())
}

func (i EmailMap) ToEmailMapOutputWithContext(ctx context.Context) EmailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailMapOutput)
}

type EmailOutput struct{ *pulumi.OutputState }

func (EmailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Email)(nil)).Elem()
}

func (o EmailOutput) ToEmailOutput() EmailOutput {
	return o
}

func (o EmailOutput) ToEmailOutputWithContext(ctx context.Context) EmailOutput {
	return o
}

type EmailArrayOutput struct{ *pulumi.OutputState }

func (EmailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Email)(nil)).Elem()
}

func (o EmailArrayOutput) ToEmailArrayOutput() EmailArrayOutput {
	return o
}

func (o EmailArrayOutput) ToEmailArrayOutputWithContext(ctx context.Context) EmailArrayOutput {
	return o
}

func (o EmailArrayOutput) Index(i pulumi.IntInput) EmailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Email {
		return vs[0].([]*Email)[vs[1].(int)]
	}).(EmailOutput)
}

type EmailMapOutput struct{ *pulumi.OutputState }

func (EmailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Email)(nil)).Elem()
}

func (o EmailMapOutput) ToEmailMapOutput() EmailMapOutput {
	return o
}

func (o EmailMapOutput) ToEmailMapOutputWithContext(ctx context.Context) EmailMapOutput {
	return o
}

func (o EmailMapOutput) MapIndex(k pulumi.StringInput) EmailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Email {
		return vs[0].(map[string]*Email)[vs[1].(string)]
	}).(EmailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailInput)(nil)).Elem(), &Email{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailArrayInput)(nil)).Elem(), EmailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailMapInput)(nil)).Elem(), EmailMap{})
	pulumi.RegisterOutputType(EmailOutput{})
	pulumi.RegisterOutputType(EmailArrayOutput{})
	pulumi.RegisterOutputType(EmailMapOutput{})
}
