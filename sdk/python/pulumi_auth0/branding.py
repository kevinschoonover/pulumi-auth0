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

__all__ = ['BrandingArgs', 'Branding']

@pulumi.input_type
class BrandingArgs:
    def __init__(__self__, *,
                 colors: Optional[pulumi.Input['BrandingColorsArgs']] = None,
                 favicon_url: Optional[pulumi.Input[str]] = None,
                 font: Optional[pulumi.Input['BrandingFontArgs']] = None,
                 logo_url: Optional[pulumi.Input[str]] = None,
                 universal_login: Optional[pulumi.Input['BrandingUniversalLoginArgs']] = None):
        """
        The set of arguments for constructing a Branding resource.
        :param pulumi.Input['BrandingColorsArgs'] colors: List(Resource). Configuration settings for colors for branding. See Colors.
        :param pulumi.Input[str] favicon_url: String. URL for the favicon.
        :param pulumi.Input['BrandingFontArgs'] font: List(Resource). Configuration settings to customize the font. See Font.
        :param pulumi.Input[str] logo_url: String. URL of logo for branding.
        :param pulumi.Input['BrandingUniversalLoginArgs'] universal_login: List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        if colors is not None:
            pulumi.set(__self__, "colors", colors)
        if favicon_url is not None:
            pulumi.set(__self__, "favicon_url", favicon_url)
        if font is not None:
            pulumi.set(__self__, "font", font)
        if logo_url is not None:
            pulumi.set(__self__, "logo_url", logo_url)
        if universal_login is not None:
            pulumi.set(__self__, "universal_login", universal_login)

    @property
    @pulumi.getter
    def colors(self) -> Optional[pulumi.Input['BrandingColorsArgs']]:
        """
        List(Resource). Configuration settings for colors for branding. See Colors.
        """
        return pulumi.get(self, "colors")

    @colors.setter
    def colors(self, value: Optional[pulumi.Input['BrandingColorsArgs']]):
        pulumi.set(self, "colors", value)

    @property
    @pulumi.getter(name="faviconUrl")
    def favicon_url(self) -> Optional[pulumi.Input[str]]:
        """
        String. URL for the favicon.
        """
        return pulumi.get(self, "favicon_url")

    @favicon_url.setter
    def favicon_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "favicon_url", value)

    @property
    @pulumi.getter
    def font(self) -> Optional[pulumi.Input['BrandingFontArgs']]:
        """
        List(Resource). Configuration settings to customize the font. See Font.
        """
        return pulumi.get(self, "font")

    @font.setter
    def font(self, value: Optional[pulumi.Input['BrandingFontArgs']]):
        pulumi.set(self, "font", value)

    @property
    @pulumi.getter(name="logoUrl")
    def logo_url(self) -> Optional[pulumi.Input[str]]:
        """
        String. URL of logo for branding.
        """
        return pulumi.get(self, "logo_url")

    @logo_url.setter
    def logo_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logo_url", value)

    @property
    @pulumi.getter(name="universalLogin")
    def universal_login(self) -> Optional[pulumi.Input['BrandingUniversalLoginArgs']]:
        """
        List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        return pulumi.get(self, "universal_login")

    @universal_login.setter
    def universal_login(self, value: Optional[pulumi.Input['BrandingUniversalLoginArgs']]):
        pulumi.set(self, "universal_login", value)


@pulumi.input_type
class _BrandingState:
    def __init__(__self__, *,
                 colors: Optional[pulumi.Input['BrandingColorsArgs']] = None,
                 favicon_url: Optional[pulumi.Input[str]] = None,
                 font: Optional[pulumi.Input['BrandingFontArgs']] = None,
                 logo_url: Optional[pulumi.Input[str]] = None,
                 universal_login: Optional[pulumi.Input['BrandingUniversalLoginArgs']] = None):
        """
        Input properties used for looking up and filtering Branding resources.
        :param pulumi.Input['BrandingColorsArgs'] colors: List(Resource). Configuration settings for colors for branding. See Colors.
        :param pulumi.Input[str] favicon_url: String. URL for the favicon.
        :param pulumi.Input['BrandingFontArgs'] font: List(Resource). Configuration settings to customize the font. See Font.
        :param pulumi.Input[str] logo_url: String. URL of logo for branding.
        :param pulumi.Input['BrandingUniversalLoginArgs'] universal_login: List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        if colors is not None:
            pulumi.set(__self__, "colors", colors)
        if favicon_url is not None:
            pulumi.set(__self__, "favicon_url", favicon_url)
        if font is not None:
            pulumi.set(__self__, "font", font)
        if logo_url is not None:
            pulumi.set(__self__, "logo_url", logo_url)
        if universal_login is not None:
            pulumi.set(__self__, "universal_login", universal_login)

    @property
    @pulumi.getter
    def colors(self) -> Optional[pulumi.Input['BrandingColorsArgs']]:
        """
        List(Resource). Configuration settings for colors for branding. See Colors.
        """
        return pulumi.get(self, "colors")

    @colors.setter
    def colors(self, value: Optional[pulumi.Input['BrandingColorsArgs']]):
        pulumi.set(self, "colors", value)

    @property
    @pulumi.getter(name="faviconUrl")
    def favicon_url(self) -> Optional[pulumi.Input[str]]:
        """
        String. URL for the favicon.
        """
        return pulumi.get(self, "favicon_url")

    @favicon_url.setter
    def favicon_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "favicon_url", value)

    @property
    @pulumi.getter
    def font(self) -> Optional[pulumi.Input['BrandingFontArgs']]:
        """
        List(Resource). Configuration settings to customize the font. See Font.
        """
        return pulumi.get(self, "font")

    @font.setter
    def font(self, value: Optional[pulumi.Input['BrandingFontArgs']]):
        pulumi.set(self, "font", value)

    @property
    @pulumi.getter(name="logoUrl")
    def logo_url(self) -> Optional[pulumi.Input[str]]:
        """
        String. URL of logo for branding.
        """
        return pulumi.get(self, "logo_url")

    @logo_url.setter
    def logo_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logo_url", value)

    @property
    @pulumi.getter(name="universalLogin")
    def universal_login(self) -> Optional[pulumi.Input['BrandingUniversalLoginArgs']]:
        """
        List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        return pulumi.get(self, "universal_login")

    @universal_login.setter
    def universal_login(self, value: Optional[pulumi.Input['BrandingUniversalLoginArgs']]):
        pulumi.set(self, "universal_login", value)


class Branding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 colors: Optional[pulumi.Input[pulumi.InputType['BrandingColorsArgs']]] = None,
                 favicon_url: Optional[pulumi.Input[str]] = None,
                 font: Optional[pulumi.Input[pulumi.InputType['BrandingFontArgs']]] = None,
                 logo_url: Optional[pulumi.Input[str]] = None,
                 universal_login: Optional[pulumi.Input[pulumi.InputType['BrandingUniversalLoginArgs']]] = None,
                 __props__=None):
        """
        With Auth0, you can setting logo, color to maintain a consistent service brand. This resource allows you to manage a branding within your Auth0 tenant.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        my_brand = auth0.Branding("myBrand",
            colors=auth0.BrandingColorsArgs(
                page_background="#000000",
                primary="#0059d6",
            ),
            logo_url="https://mycompany.org/logo.png",
            universal_login=auth0.BrandingUniversalLoginArgs(
                body="<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BrandingColorsArgs']] colors: List(Resource). Configuration settings for colors for branding. See Colors.
        :param pulumi.Input[str] favicon_url: String. URL for the favicon.
        :param pulumi.Input[pulumi.InputType['BrandingFontArgs']] font: List(Resource). Configuration settings to customize the font. See Font.
        :param pulumi.Input[str] logo_url: String. URL of logo for branding.
        :param pulumi.Input[pulumi.InputType['BrandingUniversalLoginArgs']] universal_login: List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BrandingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        With Auth0, you can setting logo, color to maintain a consistent service brand. This resource allows you to manage a branding within your Auth0 tenant.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_auth0 as auth0

        my_brand = auth0.Branding("myBrand",
            colors=auth0.BrandingColorsArgs(
                page_background="#000000",
                primary="#0059d6",
            ),
            logo_url="https://mycompany.org/logo.png",
            universal_login=auth0.BrandingUniversalLoginArgs(
                body="<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param BrandingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BrandingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 colors: Optional[pulumi.Input[pulumi.InputType['BrandingColorsArgs']]] = None,
                 favicon_url: Optional[pulumi.Input[str]] = None,
                 font: Optional[pulumi.Input[pulumi.InputType['BrandingFontArgs']]] = None,
                 logo_url: Optional[pulumi.Input[str]] = None,
                 universal_login: Optional[pulumi.Input[pulumi.InputType['BrandingUniversalLoginArgs']]] = None,
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
            __props__ = BrandingArgs.__new__(BrandingArgs)

            __props__.__dict__["colors"] = colors
            __props__.__dict__["favicon_url"] = favicon_url
            __props__.__dict__["font"] = font
            __props__.__dict__["logo_url"] = logo_url
            __props__.__dict__["universal_login"] = universal_login
        super(Branding, __self__).__init__(
            'auth0:index/branding:Branding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            colors: Optional[pulumi.Input[pulumi.InputType['BrandingColorsArgs']]] = None,
            favicon_url: Optional[pulumi.Input[str]] = None,
            font: Optional[pulumi.Input[pulumi.InputType['BrandingFontArgs']]] = None,
            logo_url: Optional[pulumi.Input[str]] = None,
            universal_login: Optional[pulumi.Input[pulumi.InputType['BrandingUniversalLoginArgs']]] = None) -> 'Branding':
        """
        Get an existing Branding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BrandingColorsArgs']] colors: List(Resource). Configuration settings for colors for branding. See Colors.
        :param pulumi.Input[str] favicon_url: String. URL for the favicon.
        :param pulumi.Input[pulumi.InputType['BrandingFontArgs']] font: List(Resource). Configuration settings to customize the font. See Font.
        :param pulumi.Input[str] logo_url: String. URL of logo for branding.
        :param pulumi.Input[pulumi.InputType['BrandingUniversalLoginArgs']] universal_login: List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BrandingState.__new__(_BrandingState)

        __props__.__dict__["colors"] = colors
        __props__.__dict__["favicon_url"] = favicon_url
        __props__.__dict__["font"] = font
        __props__.__dict__["logo_url"] = logo_url
        __props__.__dict__["universal_login"] = universal_login
        return Branding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def colors(self) -> pulumi.Output[Optional['outputs.BrandingColors']]:
        """
        List(Resource). Configuration settings for colors for branding. See Colors.
        """
        return pulumi.get(self, "colors")

    @property
    @pulumi.getter(name="faviconUrl")
    def favicon_url(self) -> pulumi.Output[str]:
        """
        String. URL for the favicon.
        """
        return pulumi.get(self, "favicon_url")

    @property
    @pulumi.getter
    def font(self) -> pulumi.Output[Optional['outputs.BrandingFont']]:
        """
        List(Resource). Configuration settings to customize the font. See Font.
        """
        return pulumi.get(self, "font")

    @property
    @pulumi.getter(name="logoUrl")
    def logo_url(self) -> pulumi.Output[str]:
        """
        String. URL of logo for branding.
        """
        return pulumi.get(self, "logo_url")

    @property
    @pulumi.getter(name="universalLogin")
    def universal_login(self) -> pulumi.Output[Optional['outputs.BrandingUniversalLogin']]:
        """
        List(Resource). Configuration settings for Universal Login. See Universal Login. This capability can only be used if the tenant has [Custom Domains](https://auth0.com/docs/custom-domains) enabled.
        """
        return pulumi.get(self, "universal_login")
