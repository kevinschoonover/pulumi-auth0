// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// With Auth0, you can use a custom domain to maintain a consistent user experience. This resource allows you to create and manage a custom domain within your Auth0 tenant.
type CustomDomain struct {
	pulumi.CustomResourceState

	// String. Name of the custom domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Boolean. Indicates whether or not this is a primary domain.
	Primary pulumi.BoolOutput `pulumi:"primary"`
	// String. Configuration status for the custom domain. Options include `disabled`, `pending`, `pendingVerification`, and `ready`.
	Status pulumi.StringOutput `pulumi:"status"`
	// String. Provisioning type for the custom domain. Options include `auth0ManagedCerts` and `selfManagedCerts`.
	Type pulumi.StringOutput `pulumi:"type"`
	// List(Resource). Configuration settings for verification. For details, see Verification.
	Verification CustomDomainVerificationOutput `pulumi:"verification"`
	// String. Domain verification method. Options include `txt`.
	VerificationMethod pulumi.StringOutput `pulumi:"verificationMethod"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil || args.Domain == nil {
		return nil, errors.New("missing required argument 'Domain'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.VerificationMethod == nil {
		return nil, errors.New("missing required argument 'VerificationMethod'")
	}
	if args == nil {
		args = &CustomDomainArgs{}
	}
	var resource CustomDomain
	err := ctx.RegisterResource("auth0:index/customDomain:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("auth0:index/customDomain:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
	// String. Name of the custom domain.
	Domain *string `pulumi:"domain"`
	// Boolean. Indicates whether or not this is a primary domain.
	Primary *bool `pulumi:"primary"`
	// String. Configuration status for the custom domain. Options include `disabled`, `pending`, `pendingVerification`, and `ready`.
	Status *string `pulumi:"status"`
	// String. Provisioning type for the custom domain. Options include `auth0ManagedCerts` and `selfManagedCerts`.
	Type *string `pulumi:"type"`
	// List(Resource). Configuration settings for verification. For details, see Verification.
	Verification *CustomDomainVerification `pulumi:"verification"`
	// String. Domain verification method. Options include `txt`.
	VerificationMethod *string `pulumi:"verificationMethod"`
}

type CustomDomainState struct {
	// String. Name of the custom domain.
	Domain pulumi.StringPtrInput
	// Boolean. Indicates whether or not this is a primary domain.
	Primary pulumi.BoolPtrInput
	// String. Configuration status for the custom domain. Options include `disabled`, `pending`, `pendingVerification`, and `ready`.
	Status pulumi.StringPtrInput
	// String. Provisioning type for the custom domain. Options include `auth0ManagedCerts` and `selfManagedCerts`.
	Type pulumi.StringPtrInput
	// List(Resource). Configuration settings for verification. For details, see Verification.
	Verification CustomDomainVerificationPtrInput
	// String. Domain verification method. Options include `txt`.
	VerificationMethod pulumi.StringPtrInput
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// String. Name of the custom domain.
	Domain string `pulumi:"domain"`
	// String. Provisioning type for the custom domain. Options include `auth0ManagedCerts` and `selfManagedCerts`.
	Type string `pulumi:"type"`
	// String. Domain verification method. Options include `txt`.
	VerificationMethod string `pulumi:"verificationMethod"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// String. Name of the custom domain.
	Domain pulumi.StringInput
	// String. Provisioning type for the custom domain. Options include `auth0ManagedCerts` and `selfManagedCerts`.
	Type pulumi.StringInput
	// String. Domain verification method. Options include `txt`.
	VerificationMethod pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}