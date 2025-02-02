# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GuardianArgs', 'Guardian']

@pulumi.input_type
class GuardianArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input[str],
                 email: Optional[pulumi.Input[bool]] = None,
                 phone: Optional[pulumi.Input['GuardianPhoneArgs']] = None):
        """
        The set of arguments for constructing a Guardian resource.
        :param pulumi.Input[str] policy: String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        :param pulumi.Input[bool] email: Boolean. Indicates whether or not email MFA is enabled.
        :param pulumi.Input['GuardianPhoneArgs'] phone: List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        """
        pulumi.set(__self__, "policy", policy)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean. Indicates whether or not email MFA is enabled.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input['GuardianPhoneArgs']]:
        """
        List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input['GuardianPhoneArgs']]):
        pulumi.set(self, "phone", value)


@pulumi.input_type
class _GuardianState:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[bool]] = None,
                 phone: Optional[pulumi.Input['GuardianPhoneArgs']] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Guardian resources.
        :param pulumi.Input[bool] email: Boolean. Indicates whether or not email MFA is enabled.
        :param pulumi.Input['GuardianPhoneArgs'] phone: List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        :param pulumi.Input[str] policy: String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean. Indicates whether or not email MFA is enabled.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input['GuardianPhoneArgs']]:
        """
        List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input['GuardianPhoneArgs']]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class Guardian(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[bool]] = None,
                 phone: Optional[pulumi.Input[pulumi.InputType['GuardianPhoneArgs']]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Multi-factor Authentication works by requiring additional factors during the login process to prevent unauthorized access. With this resource you can configure some of
        the options available for MFA.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        default = auth0.Guardian("default",
            email=True,
            phone=auth0.GuardianPhoneArgs(
                message_types=["sms"],
                options=auth0.GuardianPhoneOptionsArgs(
                    enrollment_message="{{code}} is your verification code for {{tenant.friendly_name}}. Please enter this code to verify your enrollment.",
                    verification_message="{{code}} is your verification code for {{tenant.friendly_name}}.",
                ),
                provider="auth0",
            ),
            policy="all-applications")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] email: Boolean. Indicates whether or not email MFA is enabled.
        :param pulumi.Input[pulumi.InputType['GuardianPhoneArgs']] phone: List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        :param pulumi.Input[str] policy: String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GuardianArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Multi-factor Authentication works by requiring additional factors during the login process to prevent unauthorized access. With this resource you can configure some of
        the options available for MFA.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        default = auth0.Guardian("default",
            email=True,
            phone=auth0.GuardianPhoneArgs(
                message_types=["sms"],
                options=auth0.GuardianPhoneOptionsArgs(
                    enrollment_message="{{code}} is your verification code for {{tenant.friendly_name}}. Please enter this code to verify your enrollment.",
                    verification_message="{{code}} is your verification code for {{tenant.friendly_name}}.",
                ),
                provider="auth0",
            ),
            policy="all-applications")
        ```

        :param str resource_name: The name of the resource.
        :param GuardianArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GuardianArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[bool]] = None,
                 phone: Optional[pulumi.Input[pulumi.InputType['GuardianPhoneArgs']]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GuardianArgs.__new__(GuardianArgs)

            __props__.__dict__["email"] = email
            __props__.__dict__["phone"] = phone
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
        super(Guardian, __self__).__init__(
            'auth0:index/guardian:Guardian',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email: Optional[pulumi.Input[bool]] = None,
            phone: Optional[pulumi.Input[pulumi.InputType['GuardianPhoneArgs']]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'Guardian':
        """
        Get an existing Guardian resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] email: Boolean. Indicates whether or not email MFA is enabled.
        :param pulumi.Input[pulumi.InputType['GuardianPhoneArgs']] phone: List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        :param pulumi.Input[str] policy: String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GuardianState.__new__(_GuardianState)

        __props__.__dict__["email"] = email
        __props__.__dict__["phone"] = phone
        __props__.__dict__["policy"] = policy
        return Guardian(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean. Indicates whether or not email MFA is enabled.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def phone(self) -> pulumi.Output[Optional['outputs.GuardianPhone']]:
        """
        List(Resource). Configuration settings for the phone MFA. For details, see Phone.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        String. Policy to use. Available options are `never`, `all-applications` and `confidence-score`. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
        """
        return pulumi.get(self, "policy")

