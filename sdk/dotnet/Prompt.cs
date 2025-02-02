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
    /// With this resource, you can manage your Auth0 prompts, including choosing the login experience version.
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
    ///         var example = new Auth0.Prompt("example", new Auth0.PromptArgs
    ///         {
    ///             IdentifierFirst = false,
    ///             UniversalLoginExperience = "classic",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [Auth0ResourceType("auth0:index/prompt:Prompt")]
    public partial class Prompt : Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean. Indicates whether or not identifier first is used when using the new universal login experience.
        /// </summary>
        [Output("identifierFirst")]
        public Output<bool?> IdentifierFirst { get; private set; } = null!;

        /// <summary>
        /// Which login experience to use. Options include `classic` and `new`.
        /// </summary>
        [Output("universalLoginExperience")]
        public Output<string?> UniversalLoginExperience { get; private set; } = null!;


        /// <summary>
        /// Create a Prompt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Prompt(string name, PromptArgs? args = null, CustomResourceOptions? options = null)
            : base("auth0:index/prompt:Prompt", name, args ?? new PromptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Prompt(string name, Input<string> id, PromptState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/prompt:Prompt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Prompt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Prompt Get(string name, Input<string> id, PromptState? state = null, CustomResourceOptions? options = null)
        {
            return new Prompt(name, id, state, options);
        }
    }

    public sealed class PromptArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean. Indicates whether or not identifier first is used when using the new universal login experience.
        /// </summary>
        [Input("identifierFirst")]
        public Input<bool>? IdentifierFirst { get; set; }

        /// <summary>
        /// Which login experience to use. Options include `classic` and `new`.
        /// </summary>
        [Input("universalLoginExperience")]
        public Input<string>? UniversalLoginExperience { get; set; }

        public PromptArgs()
        {
        }
    }

    public sealed class PromptState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean. Indicates whether or not identifier first is used when using the new universal login experience.
        /// </summary>
        [Input("identifierFirst")]
        public Input<bool>? IdentifierFirst { get; set; }

        /// <summary>
        /// Which login experience to use. Options include `classic` and `new`.
        /// </summary>
        [Input("universalLoginExperience")]
        public Input<string>? UniversalLoginExperience { get; set; }

        public PromptState()
        {
        }
    }
}
