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
    /// With this resource, you can manage user identities, including resetting passwords, and creating, provisioning, blocking, and deleting users.
    /// </summary>
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
        /// </summary>
        [Output("appMetadata")]
        public Output<string?> AppMetadata { get; private set; } = null!;

        [Output("blocked")]
        public Output<bool?> Blocked { get; private set; } = null!;

        /// <summary>
        /// String. Name of the connection from which the user information was sourced.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// String. Email address of the user.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the email address has been verified.
        /// </summary>
        [Output("emailVerified")]
        public Output<bool?> EmailVerified { get; private set; } = null!;

        [Output("familyName")]
        public Output<string?> FamilyName { get; private set; } = null!;

        [Output("givenName")]
        public Output<string?> GivenName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// String. Preferred nickname or alias of the user.
        /// </summary>
        [Output("nickname")]
        public Output<string?> Nickname { get; private set; } = null!;

        /// <summary>
        /// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
        /// </summary>
        [Output("phoneNumber")]
        public Output<string?> PhoneNumber { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the phone number has been verified.
        /// </summary>
        [Output("phoneVerified")]
        public Output<bool?> PhoneVerified { get; private set; } = null!;

        [Output("picture")]
        public Output<string?> Picture { get; private set; } = null!;

        /// <summary>
        /// Set(String). Set of IDs of roles assigned to the user.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// String. ID of the user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
        /// </summary>
        [Output("userMetadata")]
        public Output<string?> UserMetadata { get; private set; } = null!;

        /// <summary>
        /// String. Username of the user. Only valid if the connection requires a username.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `email_verified` parameter.
        /// </summary>
        [Output("verifyEmail")]
        public Output<bool?> VerifyEmail { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("auth0:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("auth0:index/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
        /// </summary>
        [Input("appMetadata")]
        public Input<string>? AppMetadata { get; set; }

        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// String. Name of the connection from which the user information was sourced.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// String. Email address of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the email address has been verified.
        /// </summary>
        [Input("emailVerified")]
        public Input<bool>? EmailVerified { get; set; }

        [Input("familyName")]
        public Input<string>? FamilyName { get; set; }

        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// String. Preferred nickname or alias of the user.
        /// </summary>
        [Input("nickname")]
        public Input<string>? Nickname { get; set; }

        /// <summary>
        /// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
        /// </summary>
        [Input("phoneNumber")]
        public Input<string>? PhoneNumber { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the phone number has been verified.
        /// </summary>
        [Input("phoneVerified")]
        public Input<bool>? PhoneVerified { get; set; }

        [Input("picture")]
        public Input<string>? Picture { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Set(String). Set of IDs of roles assigned to the user.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// String. ID of the user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        /// <summary>
        /// String. Username of the user. Only valid if the connection requires a username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `email_verified` parameter.
        /// </summary>
        [Input("verifyEmail")]
        public Input<bool>? VerifyEmail { get; set; }

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
        /// </summary>
        [Input("appMetadata")]
        public Input<string>? AppMetadata { get; set; }

        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        /// <summary>
        /// String. Name of the connection from which the user information was sourced.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// String. Email address of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the email address has been verified.
        /// </summary>
        [Input("emailVerified")]
        public Input<bool>? EmailVerified { get; set; }

        [Input("familyName")]
        public Input<string>? FamilyName { get; set; }

        [Input("givenName")]
        public Input<string>? GivenName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// String. Preferred nickname or alias of the user.
        /// </summary>
        [Input("nickname")]
        public Input<string>? Nickname { get; set; }

        /// <summary>
        /// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections. 
        /// </summary>
        [Input("phoneNumber")]
        public Input<string>? PhoneNumber { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the phone number has been verified.
        /// </summary>
        [Input("phoneVerified")]
        public Input<bool>? PhoneVerified { get; set; }

        [Input("picture")]
        public Input<string>? Picture { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Set(String). Set of IDs of roles assigned to the user.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// String. ID of the user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
        /// </summary>
        [Input("userMetadata")]
        public Input<string>? UserMetadata { get; set; }

        /// <summary>
        /// String. Username of the user. Only valid if the connection requires a username.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `email_verified` parameter.
        /// </summary>
        [Input("verifyEmail")]
        public Input<bool>? VerifyEmail { get; set; }

        public UserState()
        {
        }
    }
}