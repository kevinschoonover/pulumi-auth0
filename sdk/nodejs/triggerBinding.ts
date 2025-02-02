// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * With this resource, you can bind an action to a trigger. Once an action is
 * created and deployed, it can be attached (i.e. bound) to a trigger so that it
 * will be executed as part of a flow.
 *
 * The list of actions reflects the order in which they will be executed during the
 * appropriate flow.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as auth0 from "@pulumi/auth0";
 *
 * const actionFoo = new auth0.Action("actionFoo", {
 *     supportedTriggers: {
 *         id: "post-login",
 *         version: "v2",
 *     },
 *     code: `exports.onContinuePostLogin = async (event, api) => { 
 * 	console.log("foo") 
 * };"
 * `,
 *     deploy: true,
 * });
 * const actionBar = new auth0.Action("actionBar", {
 *     supportedTriggers: {
 *         id: "post-login",
 *         version: "v2",
 *     },
 *     code: `exports.onContinuePostLogin = async (event, api) => { 
 * 	console.log("bar") 
 * };"
 * `,
 *     deploy: true,
 * });
 * const loginFlow = new auth0.TriggerBinding("loginFlow", {
 *     trigger: "post-login",
 *     actions: [
 *         {
 *             id: actionFoo.id,
 *             displayName: actionFoo.name,
 *         },
 *         {
 *             id: actionBar.id,
 *             displayName: actionBar.name,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * auth0_trigger_binding can be imported using the bindings trigger ID, e.g.
 *
 * ```sh
 *  $ pulumi import auth0:index/triggerBinding:TriggerBinding example "post-login"
 * ```
 */
export class TriggerBinding extends pulumi.CustomResource {
    /**
     * Get an existing TriggerBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerBindingState, opts?: pulumi.CustomResourceOptions): TriggerBinding {
        return new TriggerBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'auth0:index/triggerBinding:TriggerBinding';

    /**
     * Returns true if the given object is an instance of TriggerBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TriggerBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TriggerBinding.__pulumiType;
    }

    /**
     * The actions bound to this trigger. For details, see
     * Actions.
     */
    public readonly actions!: pulumi.Output<outputs.TriggerBindingAction[]>;
    /**
     * The id of the trigger to bind with
     */
    public readonly trigger!: pulumi.Output<string>;

    /**
     * Create a TriggerBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerBindingArgs | TriggerBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerBindingState | undefined;
            resourceInputs["actions"] = state ? state.actions : undefined;
            resourceInputs["trigger"] = state ? state.trigger : undefined;
        } else {
            const args = argsOrState as TriggerBindingArgs | undefined;
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.trigger === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trigger'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["trigger"] = args ? args.trigger : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TriggerBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TriggerBinding resources.
 */
export interface TriggerBindingState {
    /**
     * The actions bound to this trigger. For details, see
     * Actions.
     */
    actions?: pulumi.Input<pulumi.Input<inputs.TriggerBindingAction>[]>;
    /**
     * The id of the trigger to bind with
     */
    trigger?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TriggerBinding resource.
 */
export interface TriggerBindingArgs {
    /**
     * The actions bound to this trigger. For details, see
     * Actions.
     */
    actions: pulumi.Input<pulumi.Input<inputs.TriggerBindingAction>[]>;
    /**
     * The id of the trigger to bind with
     */
    trigger: pulumi.Input<string>;
}
