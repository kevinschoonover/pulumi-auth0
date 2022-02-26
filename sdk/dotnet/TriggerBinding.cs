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
    /// With this resource, you can bind an action to a trigger. Once an action is
    /// created and deployed, it can be attached (i.e. bound) to a trigger so that it
    /// will be executed as part of a flow.
    /// 
    /// The list of actions reflects the order in which they will be executed during the
    /// appropriate flow.
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
    ///         var actionFoo = new Auth0.Action("actionFoo", new Auth0.ActionArgs
    ///         {
    ///             SupportedTriggers = new Auth0.Inputs.ActionSupportedTriggersArgs
    ///             {
    ///                 Id = "post-login",
    ///                 Version = "v2",
    ///             },
    ///             Code = @"exports.onContinuePostLogin = async (event, api) =&gt; { 
    /// 	console.log(""foo"") 
    /// };""
    /// ",
    ///             Deploy = true,
    ///         });
    ///         var actionBar = new Auth0.Action("actionBar", new Auth0.ActionArgs
    ///         {
    ///             SupportedTriggers = new Auth0.Inputs.ActionSupportedTriggersArgs
    ///             {
    ///                 Id = "post-login",
    ///                 Version = "v2",
    ///             },
    ///             Code = @"exports.onContinuePostLogin = async (event, api) =&gt; { 
    /// 	console.log(""bar"") 
    /// };""
    /// ",
    ///             Deploy = true,
    ///         });
    ///         var loginFlow = new Auth0.TriggerBinding("loginFlow", new Auth0.TriggerBindingArgs
    ///         {
    ///             Trigger = "post-login",
    ///             Actions = 
    ///             {
    ///                 new Auth0.Inputs.TriggerBindingActionArgs
    ///                 {
    ///                     Id = actionFoo.Id,
    ///                     DisplayName = actionFoo.Name,
    ///                 },
    ///                 new Auth0.Inputs.TriggerBindingActionArgs
    ///                 {
    ///                     Id = actionBar.Id,
    ///                     DisplayName = actionBar.Name,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// auth0_trigger_binding can be imported using the bindings trigger ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import auth0:index/triggerBinding:TriggerBinding example "post-login"
    /// ```
    /// </summary>
    [Auth0ResourceType("auth0:index/triggerBinding:TriggerBinding")]
    public partial class TriggerBinding : Pulumi.CustomResource
    {
        /// <summary>
        /// The actions bound to this trigger. For details, see
        /// Actions.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.TriggerBindingAction>> Actions { get; private set; } = null!;

        /// <summary>
        /// The id of the trigger to bind with
        /// </summary>
        [Output("trigger")]
        public Output<string> Trigger { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerBinding(string name, TriggerBindingArgs args, CustomResourceOptions? options = null)
            : base("auth0:index/triggerBinding:TriggerBinding", name, args ?? new TriggerBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerBinding(string name, Input<string> id, TriggerBindingState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/triggerBinding:TriggerBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TriggerBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerBinding Get(string name, Input<string> id, TriggerBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerBinding(name, id, state, options);
        }
    }

    public sealed class TriggerBindingArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.TriggerBindingActionArgs>? _actions;

        /// <summary>
        /// The actions bound to this trigger. For details, see
        /// Actions.
        /// </summary>
        public InputList<Inputs.TriggerBindingActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.TriggerBindingActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The id of the trigger to bind with
        /// </summary>
        [Input("trigger", required: true)]
        public Input<string> Trigger { get; set; } = null!;

        public TriggerBindingArgs()
        {
        }
    }

    public sealed class TriggerBindingState : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.TriggerBindingActionGetArgs>? _actions;

        /// <summary>
        /// The actions bound to this trigger. For details, see
        /// Actions.
        /// </summary>
        public InputList<Inputs.TriggerBindingActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.TriggerBindingActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// The id of the trigger to bind with
        /// </summary>
        [Input("trigger")]
        public Input<string>? Trigger { get; set; }

        public TriggerBindingState()
        {
        }
    }
}