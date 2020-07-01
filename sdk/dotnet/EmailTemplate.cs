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
    /// With Auth0, you can have standard welcome, password reset, and account verification email-based workflows built right into Auth0. This resource allows you to configure email templates to customize the look, feel, and sender identities of emails sent by Auth0. Used in conjunction with configured email providers.
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
    ///         var myEmailProvider = new Auth0.Email("myEmailProvider", new Auth0.EmailArgs
    ///         {
    ///             Enabled = true,
    ///             DefaultFromAddress = "accounts@example.com",
    ///             Credentials = new Auth0.Inputs.EmailCredentialsArgs
    ///             {
    ///                 AccessKeyId = "AKIAXXXXXXXXXXXXXXXX",
    ///                 SecretAccessKey = "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    ///                 Region = "us-east-1",
    ///             },
    ///         });
    ///         var myEmailTemplate = new Auth0.EmailTemplate("myEmailTemplate", new Auth0.EmailTemplateArgs
    ///         {
    ///             Template = "welcome_email",
    ///             Body = "&lt;html&gt;&lt;body&gt;&lt;h1&gt;Welcome!&lt;/h1&gt;&lt;/body&gt;&lt;/html&gt;",
    ///             From = "welcome@example.com",
    ///             ResultUrl = "https://example.com/welcome",
    ///             Subject = "Welcome",
    ///             Syntax = "liquid",
    ///             UrlLifetimeInSeconds = 3600,
    ///             Enabled = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 myEmailProvider,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class EmailTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the template is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Output("from")]
        public Output<string> From { get; private set; } = null!;

        /// <summary>
        /// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
        /// </summary>
        [Output("resultUrl")]
        public Output<string?> ResultUrl { get; private set; } = null!;

        /// <summary>
        /// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;

        /// <summary>
        /// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
        /// </summary>
        [Output("syntax")]
        public Output<string> Syntax { get; private set; } = null!;

        /// <summary>
        /// String. Template name. Options include `verify_email`, `reset_email`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `change_password` (legacy), and `password_reset` (legacy).
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;

        /// <summary>
        /// Integer. Number of seconds during which the link within the email will be valid.
        /// </summary>
        [Output("urlLifetimeInSeconds")]
        public Output<int?> UrlLifetimeInSeconds { get; private set; } = null!;


        /// <summary>
        /// Create a EmailTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmailTemplate(string name, EmailTemplateArgs args, CustomResourceOptions? options = null)
            : base("auth0:index/emailTemplate:EmailTemplate", name, args ?? new EmailTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmailTemplate(string name, Input<string> id, EmailTemplateState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/emailTemplate:EmailTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EmailTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmailTemplate Get(string name, Input<string> id, EmailTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new EmailTemplate(name, id, state, options);
        }
    }

    public sealed class EmailTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("body", required: true)]
        public Input<string> Body { get; set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the template is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("from", required: true)]
        public Input<string> From { get; set; } = null!;

        /// <summary>
        /// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
        /// </summary>
        [Input("resultUrl")]
        public Input<string>? ResultUrl { get; set; }

        /// <summary>
        /// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("subject", required: true)]
        public Input<string> Subject { get; set; } = null!;

        /// <summary>
        /// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
        /// </summary>
        [Input("syntax", required: true)]
        public Input<string> Syntax { get; set; } = null!;

        /// <summary>
        /// String. Template name. Options include `verify_email`, `reset_email`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `change_password` (legacy), and `password_reset` (legacy).
        /// </summary>
        [Input("template", required: true)]
        public Input<string> Template { get; set; } = null!;

        /// <summary>
        /// Integer. Number of seconds during which the link within the email will be valid.
        /// </summary>
        [Input("urlLifetimeInSeconds")]
        public Input<int>? UrlLifetimeInSeconds { get; set; }

        public EmailTemplateArgs()
        {
        }
    }

    public sealed class EmailTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String. Body of the email template. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the template is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// String. Email address to use as the sender. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("from")]
        public Input<string>? From { get; set; }

        /// <summary>
        /// String. URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/email/templates#configuring-the-redirect-to-url).
        /// </summary>
        [Input("resultUrl")]
        public Input<string>? ResultUrl { get; set; }

        /// <summary>
        /// String. Subject line of the email. You can include [common variables](https://auth0.com/docs/email/templates#common-variables).
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// String. Syntax of the template body. You can use either text or HTML + Liquid syntax.
        /// </summary>
        [Input("syntax")]
        public Input<string>? Syntax { get; set; }

        /// <summary>
        /// String. Template name. Options include `verify_email`, `reset_email`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `change_password` (legacy), and `password_reset` (legacy).
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Integer. Number of seconds during which the link within the email will be valid.
        /// </summary>
        [Input("urlLifetimeInSeconds")]
        public Input<int>? UrlLifetimeInSeconds { get; set; }

        public EmailTemplateState()
        {
        }
    }
}
