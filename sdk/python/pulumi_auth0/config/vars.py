# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'client_id',
    'client_secret',
    'debug',
    'domain',
]

__config__ = pulumi.Config('auth0')

client_id = __config__.get('clientId')

client_secret = __config__.get('clientSecret')

debug = __config__.get('debug') or _utilities.get_env_bool('AUTH0_DEBUG')

domain = __config__.get('domain')

