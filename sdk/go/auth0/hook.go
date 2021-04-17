// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth0

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hooks are secure, self-contained functions that allow you to customize the behavior of Auth0 when executed for selected extensibility points of the Auth0 platform. Auth0 invokes Hooks during runtime to execute your custom Node.js code.
//
// Depending on the extensibility point, you can use Hooks with Database Connections and/or Passwordless Connections.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-auth0/sdk/v2/go/auth0"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := auth0.NewHook(ctx, "myHook", &auth0.HookArgs{
// 			Dependencies: pulumi.StringMap{
// 				"auth0": pulumi.String("2.30.0"),
// 			},
// 			Enabled:   pulumi.Bool(true),
// 			Script:    pulumi.String(fmt.Sprintf("%v%v%v%v", "function (user, context, callback) {\n", "  callback(null, { user });\n", "}\n", "\n")),
// 			TriggerId: pulumi.String("pre-user-registration"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Hook struct {
	pulumi.CustomResourceState

	// Dependencies of this hook used by webtask server
	Dependencies pulumi.MapOutput `pulumi:"dependencies"`
	// Whether the hook is enabled, or disabled
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name of this hook
	Name pulumi.StringOutput `pulumi:"name"`
	// Code to be executed when this hook runs
	Script pulumi.StringOutput `pulumi:"script"`
	// The secrets associated with the hook
	Secrets pulumi.MapOutput `pulumi:"secrets"`
	// Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
	TriggerId pulumi.StringOutput `pulumi:"triggerId"`
}

// NewHook registers a new resource with the given unique name, arguments, and options.
func NewHook(ctx *pulumi.Context,
	name string, args *HookArgs, opts ...pulumi.ResourceOption) (*Hook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	if args.TriggerId == nil {
		return nil, errors.New("invalid value for required argument 'TriggerId'")
	}
	var resource Hook
	err := ctx.RegisterResource("auth0:index/hook:Hook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHook gets an existing Hook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HookState, opts ...pulumi.ResourceOption) (*Hook, error) {
	var resource Hook
	err := ctx.ReadResource("auth0:index/hook:Hook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hook resources.
type hookState struct {
	// Dependencies of this hook used by webtask server
	Dependencies map[string]interface{} `pulumi:"dependencies"`
	// Whether the hook is enabled, or disabled
	Enabled *bool `pulumi:"enabled"`
	// Name of this hook
	Name *string `pulumi:"name"`
	// Code to be executed when this hook runs
	Script *string `pulumi:"script"`
	// The secrets associated with the hook
	Secrets map[string]interface{} `pulumi:"secrets"`
	// Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
	TriggerId *string `pulumi:"triggerId"`
}

type HookState struct {
	// Dependencies of this hook used by webtask server
	Dependencies pulumi.MapInput
	// Whether the hook is enabled, or disabled
	Enabled pulumi.BoolPtrInput
	// Name of this hook
	Name pulumi.StringPtrInput
	// Code to be executed when this hook runs
	Script pulumi.StringPtrInput
	// The secrets associated with the hook
	Secrets pulumi.MapInput
	// Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
	TriggerId pulumi.StringPtrInput
}

func (HookState) ElementType() reflect.Type {
	return reflect.TypeOf((*hookState)(nil)).Elem()
}

type hookArgs struct {
	// Dependencies of this hook used by webtask server
	Dependencies map[string]interface{} `pulumi:"dependencies"`
	// Whether the hook is enabled, or disabled
	Enabled *bool `pulumi:"enabled"`
	// Name of this hook
	Name *string `pulumi:"name"`
	// Code to be executed when this hook runs
	Script string `pulumi:"script"`
	// The secrets associated with the hook
	Secrets map[string]interface{} `pulumi:"secrets"`
	// Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
	TriggerId string `pulumi:"triggerId"`
}

// The set of arguments for constructing a Hook resource.
type HookArgs struct {
	// Dependencies of this hook used by webtask server
	Dependencies pulumi.MapInput
	// Whether the hook is enabled, or disabled
	Enabled pulumi.BoolPtrInput
	// Name of this hook
	Name pulumi.StringPtrInput
	// Code to be executed when this hook runs
	Script pulumi.StringInput
	// The secrets associated with the hook
	Secrets pulumi.MapInput
	// Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
	TriggerId pulumi.StringInput
}

func (HookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hookArgs)(nil)).Elem()
}

type HookInput interface {
	pulumi.Input

	ToHookOutput() HookOutput
	ToHookOutputWithContext(ctx context.Context) HookOutput
}

func (*Hook) ElementType() reflect.Type {
	return reflect.TypeOf((*Hook)(nil))
}

func (i *Hook) ToHookOutput() HookOutput {
	return i.ToHookOutputWithContext(context.Background())
}

func (i *Hook) ToHookOutputWithContext(ctx context.Context) HookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookOutput)
}

func (i *Hook) ToHookPtrOutput() HookPtrOutput {
	return i.ToHookPtrOutputWithContext(context.Background())
}

func (i *Hook) ToHookPtrOutputWithContext(ctx context.Context) HookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookPtrOutput)
}

type HookPtrInput interface {
	pulumi.Input

	ToHookPtrOutput() HookPtrOutput
	ToHookPtrOutputWithContext(ctx context.Context) HookPtrOutput
}

type hookPtrType HookArgs

func (*hookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hook)(nil))
}

func (i *hookPtrType) ToHookPtrOutput() HookPtrOutput {
	return i.ToHookPtrOutputWithContext(context.Background())
}

func (i *hookPtrType) ToHookPtrOutputWithContext(ctx context.Context) HookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookPtrOutput)
}

// HookArrayInput is an input type that accepts HookArray and HookArrayOutput values.
// You can construct a concrete instance of `HookArrayInput` via:
//
//          HookArray{ HookArgs{...} }
type HookArrayInput interface {
	pulumi.Input

	ToHookArrayOutput() HookArrayOutput
	ToHookArrayOutputWithContext(context.Context) HookArrayOutput
}

type HookArray []HookInput

func (HookArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Hook)(nil))
}

func (i HookArray) ToHookArrayOutput() HookArrayOutput {
	return i.ToHookArrayOutputWithContext(context.Background())
}

func (i HookArray) ToHookArrayOutputWithContext(ctx context.Context) HookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookArrayOutput)
}

// HookMapInput is an input type that accepts HookMap and HookMapOutput values.
// You can construct a concrete instance of `HookMapInput` via:
//
//          HookMap{ "key": HookArgs{...} }
type HookMapInput interface {
	pulumi.Input

	ToHookMapOutput() HookMapOutput
	ToHookMapOutputWithContext(context.Context) HookMapOutput
}

type HookMap map[string]HookInput

func (HookMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Hook)(nil))
}

func (i HookMap) ToHookMapOutput() HookMapOutput {
	return i.ToHookMapOutputWithContext(context.Background())
}

func (i HookMap) ToHookMapOutputWithContext(ctx context.Context) HookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HookMapOutput)
}

type HookOutput struct {
	*pulumi.OutputState
}

func (HookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hook)(nil))
}

func (o HookOutput) ToHookOutput() HookOutput {
	return o
}

func (o HookOutput) ToHookOutputWithContext(ctx context.Context) HookOutput {
	return o
}

func (o HookOutput) ToHookPtrOutput() HookPtrOutput {
	return o.ToHookPtrOutputWithContext(context.Background())
}

func (o HookOutput) ToHookPtrOutputWithContext(ctx context.Context) HookPtrOutput {
	return o.ApplyT(func(v Hook) *Hook {
		return &v
	}).(HookPtrOutput)
}

type HookPtrOutput struct {
	*pulumi.OutputState
}

func (HookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hook)(nil))
}

func (o HookPtrOutput) ToHookPtrOutput() HookPtrOutput {
	return o
}

func (o HookPtrOutput) ToHookPtrOutputWithContext(ctx context.Context) HookPtrOutput {
	return o
}

type HookArrayOutput struct{ *pulumi.OutputState }

func (HookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Hook)(nil))
}

func (o HookArrayOutput) ToHookArrayOutput() HookArrayOutput {
	return o
}

func (o HookArrayOutput) ToHookArrayOutputWithContext(ctx context.Context) HookArrayOutput {
	return o
}

func (o HookArrayOutput) Index(i pulumi.IntInput) HookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Hook {
		return vs[0].([]Hook)[vs[1].(int)]
	}).(HookOutput)
}

type HookMapOutput struct{ *pulumi.OutputState }

func (HookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Hook)(nil))
}

func (o HookMapOutput) ToHookMapOutput() HookMapOutput {
	return o
}

func (o HookMapOutput) ToHookMapOutputWithContext(ctx context.Context) HookMapOutput {
	return o
}

func (o HookMapOutput) MapIndex(k pulumi.StringInput) HookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Hook {
		return vs[0].(map[string]Hook)[vs[1].(string)]
	}).(HookOutput)
}

func init() {
	pulumi.RegisterOutputType(HookOutput{})
	pulumi.RegisterOutputType(HookPtrOutput{})
	pulumi.RegisterOutputType(HookArrayOutput{})
	pulumi.RegisterOutputType(HookMapOutput{})
}
