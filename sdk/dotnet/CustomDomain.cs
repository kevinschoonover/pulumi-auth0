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
    /// With Auth0, you can use a custom domain to maintain a consistent user experience. This resource allows you to create and manage a custom domain within your Auth0 tenant.
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
    ///         var myCustomDomain = new Auth0.CustomDomain("myCustomDomain", new Auth0.CustomDomainArgs
    ///         {
    ///             Domain = "auth.example.com",
    ///             Type = "auth0_managed_certs",
    ///             VerificationMethod = "txt",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [Auth0ResourceType("auth0:index/customDomain:CustomDomain")]
    public partial class CustomDomain : Pulumi.CustomResource
    {
        /// <summary>
        /// String. Name of the custom domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not this is a primary domain.
        /// </summary>
        [Output("primary")]
        public Output<bool> Primary { get; private set; } = null!;

        /// <summary>
        /// String. Configuration status for the custom domain. Options include `disabled`, `pending`, `pending_verification`, and `ready`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// String. Provisioning type for the custom domain. Options include `auth0_managed_certs` and `self_managed_certs`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// List(Resource). Configuration settings for verification. For details, see Verification.
        /// </summary>
        [Output("verification")]
        public Output<Outputs.CustomDomainVerification> Verification { get; private set; } = null!;

        /// <summary>
        /// String. Domain verification method. Options include `txt`.
        /// </summary>
        [Output("verificationMethod")]
        public Output<string?> VerificationMethod { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs args, CustomResourceOptions? options = null)
            : base("auth0:index/customDomain:CustomDomain", name, args ?? new CustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDomain(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/customDomain:CustomDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDomain Get(string name, Input<string> id, CustomDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomDomain(name, id, state, options);
        }
    }

    public sealed class CustomDomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String. Name of the custom domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// String. Provisioning type for the custom domain. Options include `auth0_managed_certs` and `self_managed_certs`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// String. Domain verification method. Options include `txt`.
        /// </summary>
        [Input("verificationMethod")]
        public Input<string>? VerificationMethod { get; set; }

        public CustomDomainArgs()
        {
        }
    }

    public sealed class CustomDomainState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String. Name of the custom domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not this is a primary domain.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// String. Configuration status for the custom domain. Options include `disabled`, `pending`, `pending_verification`, and `ready`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// String. Provisioning type for the custom domain. Options include `auth0_managed_certs` and `self_managed_certs`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// List(Resource). Configuration settings for verification. For details, see Verification.
        /// </summary>
        [Input("verification")]
        public Input<Inputs.CustomDomainVerificationGetArgs>? Verification { get; set; }

        /// <summary>
        /// String. Domain verification method. Options include `txt`.
        /// </summary>
        [Input("verificationMethod")]
        public Input<string>? VerificationMethod { get; set; }

        public CustomDomainState()
        {
        }
    }
}
