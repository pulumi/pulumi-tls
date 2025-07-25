# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'Proxy',
]

@pulumi.output_type
class Proxy(dict):
    def __init__(__self__, *,
                 from_env: Optional[_builtins.bool] = None,
                 password: Optional[_builtins.str] = None,
                 url: Optional[_builtins.str] = None,
                 username: Optional[_builtins.str] = None):
        """
        :param _builtins.bool from_env: When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
        :param _builtins.str password: Password used for Basic authentication against the Proxy.
        :param _builtins.str url: URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
        :param _builtins.str username: Username (or Token) used for Basic authentication against the Proxy.
        """
        if from_env is not None:
            pulumi.set(__self__, "from_env", from_env)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @_builtins.property
    @pulumi.getter(name="fromEnv")
    def from_env(self) -> Optional[_builtins.bool]:
        """
        When `true` the provider will discover the proxy configuration from environment variables. This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
        """
        return pulumi.get(self, "from_env")

    @_builtins.property
    @pulumi.getter
    def password(self) -> Optional[_builtins.str]:
        """
        Password used for Basic authentication against the Proxy.
        """
        return pulumi.get(self, "password")

    @_builtins.property
    @pulumi.getter
    def url(self) -> Optional[_builtins.str]:
        """
        URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
        """
        return pulumi.get(self, "url")

    @_builtins.property
    @pulumi.getter
    def username(self) -> Optional[_builtins.str]:
        """
        Username (or Token) used for Basic authentication against the Proxy.
        """
        return pulumi.get(self, "username")


