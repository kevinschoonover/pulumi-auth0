// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// With Auth0, you can have standard welcome, password reset, and account verification email-based workflows built right into Auth0. This resource allows you to configure email templates to customize the look, feel, and sender identities of emails sent by Auth0. Used in conjunction with configured email providers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-auth0/sdk/go/auth0"
// 	"github.com/pulumi/pulumi-auth0/sdk/go/auth0/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myEmailProvider, err := auth0.NewEmail(ctx, "myEmailProvider", &auth0.EmailArgs{
// 			Enabled:            pulumi.Bool(true),
// 			DefaultFromAddress: pulumi.String("accounts@example.com"),
// 			Credentials: &auth0.EmailCredentialsArgs{
// 				AccessKeyId:     pulumi.String("AKIAXXXXXXXXXXXXXXXX"),
// 				SecretAccessKey: pulumi.String("7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
// 				Region:          pulumi.String("us-east-1"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = auth0.NewEmailTemplate(ctx, "myEmailTemplate", &auth0.EmailTemplateArgs{
// 			Template:             pulumi.String("welcome_email"),
// 			Body:                 pulumi.String("<html><body><h1>Welcome!</h1></body></html>"),
// 			From:                 pulumi.String("welcome@example.com"),
// 			ResultUrl:            pulumi.String("https://example.com/welcome"),
// 			Subject:              pulumi.String("Welcome"),
// 			Syntax:               pulumi.String("liquid"),
// 			UrlLifetimeInSeconds: pulumi.Int(3600),
// 			Enabled:              pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			myEmailProvider,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EmailTemplate struct {
	pulumi.CustomResourceState

	// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Body pulumi.StringOutput `pulumi:"body"`
	// Boolean. Indicates whether or not the template is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	From pulumi.StringOutput `pulumi:"from"`
	// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
	ResultUrl pulumi.StringPtrOutput `pulumi:"resultUrl"`
	// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Subject pulumi.StringOutput `pulumi:"subject"`
	// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
	Syntax pulumi.StringOutput `pulumi:"syntax"`
	// String. Template name. Options include `verifyEmail`, `resetEmail`, `welcomeEmail`, `blockedAccount`, `stolenCredentials`, `enrollmentEmail`, `mfaOobCode`, `changePassword` (legacy), and `passwordReset` (legacy).
	Template pulumi.StringOutput `pulumi:"template"`
	// Integer. Number of seconds during which the link within the email will be valid.
	UrlLifetimeInSeconds pulumi.IntPtrOutput `pulumi:"urlLifetimeInSeconds"`
}

// NewEmailTemplate registers a new resource with the given unique name, arguments, and options.
func NewEmailTemplate(ctx *pulumi.Context,
	name string, args *EmailTemplateArgs, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Body == nil {
		return nil, errors.New("invalid value for required argument 'Body'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.From == nil {
		return nil, errors.New("invalid value for required argument 'From'")
	}
	if args.Subject == nil {
		return nil, errors.New("invalid value for required argument 'Subject'")
	}
	if args.Syntax == nil {
		return nil, errors.New("invalid value for required argument 'Syntax'")
	}
	if args.Template == nil {
		return nil, errors.New("invalid value for required argument 'Template'")
	}
	var resource EmailTemplate
	err := ctx.RegisterResource("auth0:index/emailTemplate:EmailTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailTemplate gets an existing EmailTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailTemplateState, opts ...pulumi.ResourceOption) (*EmailTemplate, error) {
	var resource EmailTemplate
	err := ctx.ReadResource("auth0:index/emailTemplate:EmailTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailTemplate resources.
type emailTemplateState struct {
	// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Body *string `pulumi:"body"`
	// Boolean. Indicates whether or not the template is enabled.
	Enabled *bool `pulumi:"enabled"`
	// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	From *string `pulumi:"from"`
	// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
	ResultUrl *string `pulumi:"resultUrl"`
	// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Subject *string `pulumi:"subject"`
	// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
	Syntax *string `pulumi:"syntax"`
	// String. Template name. Options include `verifyEmail`, `resetEmail`, `welcomeEmail`, `blockedAccount`, `stolenCredentials`, `enrollmentEmail`, `mfaOobCode`, `changePassword` (legacy), and `passwordReset` (legacy).
	Template *string `pulumi:"template"`
	// Integer. Number of seconds during which the link within the email will be valid.
	UrlLifetimeInSeconds *int `pulumi:"urlLifetimeInSeconds"`
}

type EmailTemplateState struct {
	// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Body pulumi.StringPtrInput
	// Boolean. Indicates whether or not the template is enabled.
	Enabled pulumi.BoolPtrInput
	// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	From pulumi.StringPtrInput
	// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
	ResultUrl pulumi.StringPtrInput
	// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Subject pulumi.StringPtrInput
	// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
	Syntax pulumi.StringPtrInput
	// String. Template name. Options include `verifyEmail`, `resetEmail`, `welcomeEmail`, `blockedAccount`, `stolenCredentials`, `enrollmentEmail`, `mfaOobCode`, `changePassword` (legacy), and `passwordReset` (legacy).
	Template pulumi.StringPtrInput
	// Integer. Number of seconds during which the link within the email will be valid.
	UrlLifetimeInSeconds pulumi.IntPtrInput
}

func (EmailTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateState)(nil)).Elem()
}

type emailTemplateArgs struct {
	// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Body string `pulumi:"body"`
	// Boolean. Indicates whether or not the template is enabled.
	Enabled bool `pulumi:"enabled"`
	// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	From string `pulumi:"from"`
	// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
	ResultUrl *string `pulumi:"resultUrl"`
	// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Subject string `pulumi:"subject"`
	// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
	Syntax string `pulumi:"syntax"`
	// String. Template name. Options include `verifyEmail`, `resetEmail`, `welcomeEmail`, `blockedAccount`, `stolenCredentials`, `enrollmentEmail`, `mfaOobCode`, `changePassword` (legacy), and `passwordReset` (legacy).
	Template string `pulumi:"template"`
	// Integer. Number of seconds during which the link within the email will be valid.
	UrlLifetimeInSeconds *int `pulumi:"urlLifetimeInSeconds"`
}

// The set of arguments for constructing a EmailTemplate resource.
type EmailTemplateArgs struct {
	// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Body pulumi.StringInput
	// Boolean. Indicates whether or not the template is enabled.
	Enabled pulumi.BoolInput
	// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	From pulumi.StringInput
	// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
	ResultUrl pulumi.StringPtrInput
	// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
	Subject pulumi.StringInput
	// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
	Syntax pulumi.StringInput
	// String. Template name. Options include `verifyEmail`, `resetEmail`, `welcomeEmail`, `blockedAccount`, `stolenCredentials`, `enrollmentEmail`, `mfaOobCode`, `changePassword` (legacy), and `passwordReset` (legacy).
	Template pulumi.StringInput
	// Integer. Number of seconds during which the link within the email will be valid.
	UrlLifetimeInSeconds pulumi.IntPtrInput
}

func (EmailTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailTemplateArgs)(nil)).Elem()
}

type EmailTemplateInput interface {
	pulumi.Input

	ToEmailTemplateOutput() EmailTemplateOutput
	ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput
}

func (*EmailTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplate)(nil))
}

func (i *EmailTemplate) ToEmailTemplateOutput() EmailTemplateOutput {
	return i.ToEmailTemplateOutputWithContext(context.Background())
}

func (i *EmailTemplate) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailTemplateOutput)
}

type EmailTemplateOutput struct {
	*pulumi.OutputState
}

func (EmailTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailTemplate)(nil))
}

func (o EmailTemplateOutput) ToEmailTemplateOutput() EmailTemplateOutput {
	return o
}

func (o EmailTemplateOutput) ToEmailTemplateOutputWithContext(ctx context.Context) EmailTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EmailTemplateOutput{})
}
