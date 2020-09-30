// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// With this resource, you can manage user identities, including resetting passwords, and creating, provisioning, blocking, and deleting users.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-auth0/sdk/go/auth0"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := auth0.NewRole(ctx, "admin", &auth0.RoleArgs{
// 			Description: pulumi.String("Administrator"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = auth0.NewUser(ctx, "user", &auth0.UserArgs{
// 			ConnectionName: pulumi.String("Username-Password-Authentication"),
// 			UserId:         pulumi.String("12345"),
// 			Username:       pulumi.String("unique_username"),
// 			GivenName:      pulumi.String("Firstname"),
// 			FamilyName:     pulumi.String("Lastname"),
// 			Nickname:       pulumi.String("some.nickname"),
// 			Email:          pulumi.String("test@test.com"),
// 			EmailVerified:  pulumi.Bool(true),
// 			Password:       pulumi.String(fmt.Sprintf("%v%v%v%v%v", "passpass", "$", "12", "$", "12")),
// 			Roles: pulumi.StringArray{
// 				admin.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type User struct {
	pulumi.CustomResourceState

	// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
	AppMetadata pulumi.StringPtrOutput `pulumi:"appMetadata"`
	Blocked     pulumi.BoolPtrOutput   `pulumi:"blocked"`
	// String. Name of the connection from which the user information was sourced.
	ConnectionName pulumi.StringOutput `pulumi:"connectionName"`
	// String. Email address of the user.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// Boolean. Indicates whether or not the email address has been verified.
	EmailVerified pulumi.BoolPtrOutput   `pulumi:"emailVerified"`
	FamilyName    pulumi.StringPtrOutput `pulumi:"familyName"`
	GivenName     pulumi.StringPtrOutput `pulumi:"givenName"`
	Name          pulumi.StringOutput    `pulumi:"name"`
	// String. Preferred nickname or alias of the user.
	Nickname pulumi.StringOutput `pulumi:"nickname"`
	// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
	PhoneNumber pulumi.StringPtrOutput `pulumi:"phoneNumber"`
	// Boolean. Indicates whether or not the phone number has been verified.
	PhoneVerified pulumi.BoolPtrOutput `pulumi:"phoneVerified"`
	Picture       pulumi.StringOutput  `pulumi:"picture"`
	// Set(String). Set of IDs of roles assigned to the user.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// String. ID of the user.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
	// String. Username of the user. Only valid if the connection requires a username.
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
	VerifyEmail pulumi.BoolPtrOutput `pulumi:"verifyEmail"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("auth0:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("auth0:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
	AppMetadata *string `pulumi:"appMetadata"`
	Blocked     *bool   `pulumi:"blocked"`
	// String. Name of the connection from which the user information was sourced.
	ConnectionName *string `pulumi:"connectionName"`
	// String. Email address of the user.
	Email *string `pulumi:"email"`
	// Boolean. Indicates whether or not the email address has been verified.
	EmailVerified *bool   `pulumi:"emailVerified"`
	FamilyName    *string `pulumi:"familyName"`
	GivenName     *string `pulumi:"givenName"`
	Name          *string `pulumi:"name"`
	// String. Preferred nickname or alias of the user.
	Nickname *string `pulumi:"nickname"`
	// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
	Password *string `pulumi:"password"`
	// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
	PhoneNumber *string `pulumi:"phoneNumber"`
	// Boolean. Indicates whether or not the phone number has been verified.
	PhoneVerified *bool   `pulumi:"phoneVerified"`
	Picture       *string `pulumi:"picture"`
	// Set(String). Set of IDs of roles assigned to the user.
	Roles []string `pulumi:"roles"`
	// String. ID of the user.
	UserId *string `pulumi:"userId"`
	// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
	UserMetadata *string `pulumi:"userMetadata"`
	// String. Username of the user. Only valid if the connection requires a username.
	Username *string `pulumi:"username"`
	// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
	VerifyEmail *bool `pulumi:"verifyEmail"`
}

type UserState struct {
	// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
	AppMetadata pulumi.StringPtrInput
	Blocked     pulumi.BoolPtrInput
	// String. Name of the connection from which the user information was sourced.
	ConnectionName pulumi.StringPtrInput
	// String. Email address of the user.
	Email pulumi.StringPtrInput
	// Boolean. Indicates whether or not the email address has been verified.
	EmailVerified pulumi.BoolPtrInput
	FamilyName    pulumi.StringPtrInput
	GivenName     pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	// String. Preferred nickname or alias of the user.
	Nickname pulumi.StringPtrInput
	// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
	Password pulumi.StringPtrInput
	// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
	PhoneNumber pulumi.StringPtrInput
	// Boolean. Indicates whether or not the phone number has been verified.
	PhoneVerified pulumi.BoolPtrInput
	Picture       pulumi.StringPtrInput
	// Set(String). Set of IDs of roles assigned to the user.
	Roles pulumi.StringArrayInput
	// String. ID of the user.
	UserId pulumi.StringPtrInput
	// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
	UserMetadata pulumi.StringPtrInput
	// String. Username of the user. Only valid if the connection requires a username.
	Username pulumi.StringPtrInput
	// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
	VerifyEmail pulumi.BoolPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
	AppMetadata *string `pulumi:"appMetadata"`
	Blocked     *bool   `pulumi:"blocked"`
	// String. Name of the connection from which the user information was sourced.
	ConnectionName string `pulumi:"connectionName"`
	// String. Email address of the user.
	Email *string `pulumi:"email"`
	// Boolean. Indicates whether or not the email address has been verified.
	EmailVerified *bool   `pulumi:"emailVerified"`
	FamilyName    *string `pulumi:"familyName"`
	GivenName     *string `pulumi:"givenName"`
	Name          *string `pulumi:"name"`
	// String. Preferred nickname or alias of the user.
	Nickname *string `pulumi:"nickname"`
	// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
	Password *string `pulumi:"password"`
	// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
	PhoneNumber *string `pulumi:"phoneNumber"`
	// Boolean. Indicates whether or not the phone number has been verified.
	PhoneVerified *bool   `pulumi:"phoneVerified"`
	Picture       *string `pulumi:"picture"`
	// Set(String). Set of IDs of roles assigned to the user.
	Roles []string `pulumi:"roles"`
	// String. ID of the user.
	UserId *string `pulumi:"userId"`
	// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
	UserMetadata *string `pulumi:"userMetadata"`
	// String. Username of the user. Only valid if the connection requires a username.
	Username *string `pulumi:"username"`
	// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
	VerifyEmail *bool `pulumi:"verifyEmail"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
	AppMetadata pulumi.StringPtrInput
	Blocked     pulumi.BoolPtrInput
	// String. Name of the connection from which the user information was sourced.
	ConnectionName pulumi.StringInput
	// String. Email address of the user.
	Email pulumi.StringPtrInput
	// Boolean. Indicates whether or not the email address has been verified.
	EmailVerified pulumi.BoolPtrInput
	FamilyName    pulumi.StringPtrInput
	GivenName     pulumi.StringPtrInput
	Name          pulumi.StringPtrInput
	// String. Preferred nickname or alias of the user.
	Nickname pulumi.StringPtrInput
	// String, Case-sensitive. Initial password for this user. Used for non-SMS connections.
	Password pulumi.StringPtrInput
	// String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
	PhoneNumber pulumi.StringPtrInput
	// Boolean. Indicates whether or not the phone number has been verified.
	PhoneVerified pulumi.BoolPtrInput
	Picture       pulumi.StringPtrInput
	// Set(String). Set of IDs of roles assigned to the user.
	Roles pulumi.StringArrayInput
	// String. ID of the user.
	UserId pulumi.StringPtrInput
	// String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
	UserMetadata pulumi.StringPtrInput
	// String. Username of the user. Only valid if the connection requires a username.
	Username pulumi.StringPtrInput
	// Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `emailVerified` parameter.
	VerifyEmail pulumi.BoolPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
