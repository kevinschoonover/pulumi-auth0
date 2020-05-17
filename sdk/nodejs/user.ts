// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * With this resource, you can manage user identities, including resetting passwords, and creating, provisioning, blocking, and deleting users.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-auth0/blob/master/website/docs/r/user.html.md.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'auth0:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
     */
    public readonly appMetadata!: pulumi.Output<string | undefined>;
    public readonly blocked!: pulumi.Output<boolean | undefined>;
    /**
     * String. Name of the connection from which the user information was sourced.
     */
    public readonly connectionName!: pulumi.Output<string>;
    /**
     * String. Email address of the user.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * Boolean. Indicates whether or not the email address has been verified.
     */
    public readonly emailVerified!: pulumi.Output<boolean | undefined>;
    public readonly familyName!: pulumi.Output<string | undefined>;
    public readonly givenName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    /**
     * String. Preferred nickname or alias of the user.
     */
    public readonly nickname!: pulumi.Output<string | undefined>;
    /**
     * String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
     */
    public readonly phoneNumber!: pulumi.Output<string | undefined>;
    /**
     * Boolean. Indicates whether or not the phone number has been verified.
     */
    public readonly phoneVerified!: pulumi.Output<boolean | undefined>;
    public readonly picture!: pulumi.Output<string | undefined>;
    /**
     * Set(String). Set of IDs of roles assigned to the user.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * String. ID of the user.
     */
    public readonly userId!: pulumi.Output<string>;
    /**
     * String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
     */
    public readonly userMetadata!: pulumi.Output<string | undefined>;
    /**
     * String. Username of the user. Only valid if the connection requires a username.
     */
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
     */
    public readonly verifyEmail!: pulumi.Output<boolean | undefined>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["appMetadata"] = state ? state.appMetadata : undefined;
            inputs["blocked"] = state ? state.blocked : undefined;
            inputs["connectionName"] = state ? state.connectionName : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["emailVerified"] = state ? state.emailVerified : undefined;
            inputs["familyName"] = state ? state.familyName : undefined;
            inputs["givenName"] = state ? state.givenName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nickname"] = state ? state.nickname : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["phoneNumber"] = state ? state.phoneNumber : undefined;
            inputs["phoneVerified"] = state ? state.phoneVerified : undefined;
            inputs["picture"] = state ? state.picture : undefined;
            inputs["roles"] = state ? state.roles : undefined;
            inputs["userId"] = state ? state.userId : undefined;
            inputs["userMetadata"] = state ? state.userMetadata : undefined;
            inputs["username"] = state ? state.username : undefined;
            inputs["verifyEmail"] = state ? state.verifyEmail : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if (!args || args.connectionName === undefined) {
                throw new Error("Missing required property 'connectionName'");
            }
            inputs["appMetadata"] = args ? args.appMetadata : undefined;
            inputs["blocked"] = args ? args.blocked : undefined;
            inputs["connectionName"] = args ? args.connectionName : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["emailVerified"] = args ? args.emailVerified : undefined;
            inputs["familyName"] = args ? args.familyName : undefined;
            inputs["givenName"] = args ? args.givenName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nickname"] = args ? args.nickname : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["phoneNumber"] = args ? args.phoneNumber : undefined;
            inputs["phoneVerified"] = args ? args.phoneVerified : undefined;
            inputs["picture"] = args ? args.picture : undefined;
            inputs["roles"] = args ? args.roles : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["userMetadata"] = args ? args.userMetadata : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["verifyEmail"] = args ? args.verifyEmail : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
     */
    readonly appMetadata?: pulumi.Input<string>;
    readonly blocked?: pulumi.Input<boolean>;
    /**
     * String. Name of the connection from which the user information was sourced.
     */
    readonly connectionName?: pulumi.Input<string>;
    /**
     * String. Email address of the user.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the email address has been verified.
     */
    readonly emailVerified?: pulumi.Input<boolean>;
    readonly familyName?: pulumi.Input<string>;
    readonly givenName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * String. Preferred nickname or alias of the user.
     */
    readonly nickname?: pulumi.Input<string>;
    /**
     * String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
     */
    readonly phoneNumber?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the phone number has been verified.
     */
    readonly phoneVerified?: pulumi.Input<boolean>;
    readonly picture?: pulumi.Input<string>;
    /**
     * Set(String). Set of IDs of roles assigned to the user.
     */
    readonly roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. ID of the user.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
     */
    readonly userMetadata?: pulumi.Input<string>;
    /**
     * String. Username of the user. Only valid if the connection requires a username.
     */
    readonly username?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
     */
    readonly verifyEmail?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
     */
    readonly appMetadata?: pulumi.Input<string>;
    readonly blocked?: pulumi.Input<boolean>;
    /**
     * String. Name of the connection from which the user information was sourced.
     */
    readonly connectionName: pulumi.Input<string>;
    /**
     * String. Email address of the user.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the email address has been verified.
     */
    readonly emailVerified?: pulumi.Input<boolean>;
    readonly familyName?: pulumi.Input<string>;
    readonly givenName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * String. Preferred nickname or alias of the user.
     */
    readonly nickname?: pulumi.Input<string>;
    /**
     * String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
     */
    readonly phoneNumber?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the phone number has been verified.
     */
    readonly phoneVerified?: pulumi.Input<boolean>;
    readonly picture?: pulumi.Input<string>;
    /**
     * Set(String). Set of IDs of roles assigned to the user.
     */
    readonly roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. ID of the user.
     */
    readonly userId?: pulumi.Input<string>;
    /**
     * String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
     */
    readonly userMetadata?: pulumi.Input<string>;
    /**
     * String. Username of the user. Only valid if the connection requires a username.
     */
    readonly username?: pulumi.Input<string>;
    /**
     * Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
     */
    readonly verifyEmail?: pulumi.Input<boolean>;
}