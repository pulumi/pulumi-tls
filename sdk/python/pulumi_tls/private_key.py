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
from . import _utilities

__all__ = ['PrivateKeyArgs', 'PrivateKey']

@pulumi.input_type
class PrivateKeyArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input[_builtins.str],
                 ecdsa_curve: Optional[pulumi.Input[_builtins.str]] = None,
                 rsa_bits: Optional[pulumi.Input[_builtins.int]] = None):
        """
        The set of arguments for constructing a PrivateKey resource.
        :param pulumi.Input[_builtins.str] algorithm: Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        :param pulumi.Input[_builtins.str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        :param pulumi.Input[_builtins.int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        pulumi.set(__self__, "algorithm", algorithm)
        if ecdsa_curve is not None:
            pulumi.set(__self__, "ecdsa_curve", ecdsa_curve)
        if rsa_bits is not None:
            pulumi.set(__self__, "rsa_bits", rsa_bits)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "algorithm", value)

    @_builtins.property
    @pulumi.getter(name="ecdsaCurve")
    def ecdsa_curve(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        """
        return pulumi.get(self, "ecdsa_curve")

    @ecdsa_curve.setter
    def ecdsa_curve(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "ecdsa_curve", value)

    @_builtins.property
    @pulumi.getter(name="rsaBits")
    def rsa_bits(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        return pulumi.get(self, "rsa_bits")

    @rsa_bits.setter
    def rsa_bits(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "rsa_bits", value)


@pulumi.input_type
class _PrivateKeyState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 ecdsa_curve: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key_openssh: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key_pem: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key_pem_pkcs8: Optional[pulumi.Input[_builtins.str]] = None,
                 public_key_fingerprint_md5: Optional[pulumi.Input[_builtins.str]] = None,
                 public_key_fingerprint_sha256: Optional[pulumi.Input[_builtins.str]] = None,
                 public_key_openssh: Optional[pulumi.Input[_builtins.str]] = None,
                 public_key_pem: Optional[pulumi.Input[_builtins.str]] = None,
                 rsa_bits: Optional[pulumi.Input[_builtins.int]] = None):
        """
        Input properties used for looking up and filtering PrivateKey resources.
        :param pulumi.Input[_builtins.str] algorithm: Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        :param pulumi.Input[_builtins.str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        :param pulumi.Input[_builtins.str] private_key_openssh: Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        :param pulumi.Input[_builtins.str] private_key_pem: Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[_builtins.str] private_key_pem_pkcs8: Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        :param pulumi.Input[_builtins.str] public_key_fingerprint_md5: The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        :param pulumi.Input[_builtins.str] public_key_fingerprint_sha256: The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        :param pulumi.Input[_builtins.str] public_key_openssh: The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        :param pulumi.Input[_builtins.str] public_key_pem: Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        :param pulumi.Input[_builtins.int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if ecdsa_curve is not None:
            pulumi.set(__self__, "ecdsa_curve", ecdsa_curve)
        if private_key_openssh is not None:
            pulumi.set(__self__, "private_key_openssh", private_key_openssh)
        if private_key_pem is not None:
            pulumi.set(__self__, "private_key_pem", private_key_pem)
        if private_key_pem_pkcs8 is not None:
            pulumi.set(__self__, "private_key_pem_pkcs8", private_key_pem_pkcs8)
        if public_key_fingerprint_md5 is not None:
            pulumi.set(__self__, "public_key_fingerprint_md5", public_key_fingerprint_md5)
        if public_key_fingerprint_sha256 is not None:
            pulumi.set(__self__, "public_key_fingerprint_sha256", public_key_fingerprint_sha256)
        if public_key_openssh is not None:
            pulumi.set(__self__, "public_key_openssh", public_key_openssh)
        if public_key_pem is not None:
            pulumi.set(__self__, "public_key_pem", public_key_pem)
        if rsa_bits is not None:
            pulumi.set(__self__, "rsa_bits", rsa_bits)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "algorithm", value)

    @_builtins.property
    @pulumi.getter(name="ecdsaCurve")
    def ecdsa_curve(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        """
        return pulumi.get(self, "ecdsa_curve")

    @ecdsa_curve.setter
    def ecdsa_curve(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "ecdsa_curve", value)

    @_builtins.property
    @pulumi.getter(name="privateKeyOpenssh")
    def private_key_openssh(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        """
        return pulumi.get(self, "private_key_openssh")

    @private_key_openssh.setter
    def private_key_openssh(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "private_key_openssh", value)

    @_builtins.property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "private_key_pem")

    @private_key_pem.setter
    def private_key_pem(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "private_key_pem", value)

    @_builtins.property
    @pulumi.getter(name="privateKeyPemPkcs8")
    def private_key_pem_pkcs8(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        """
        return pulumi.get(self, "private_key_pem_pkcs8")

    @private_key_pem_pkcs8.setter
    def private_key_pem_pkcs8(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "private_key_pem_pkcs8", value)

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintMd5")
    def public_key_fingerprint_md5(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_md5")

    @public_key_fingerprint_md5.setter
    def public_key_fingerprint_md5(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "public_key_fingerprint_md5", value)

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintSha256")
    def public_key_fingerprint_sha256(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_sha256")

    @public_key_fingerprint_sha256.setter
    def public_key_fingerprint_sha256(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "public_key_fingerprint_sha256", value)

    @_builtins.property
    @pulumi.getter(name="publicKeyOpenssh")
    def public_key_openssh(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_openssh")

    @public_key_openssh.setter
    def public_key_openssh(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "public_key_openssh", value)

    @_builtins.property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_pem")

    @public_key_pem.setter
    def public_key_pem(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "public_key_pem", value)

    @_builtins.property
    @pulumi.getter(name="rsaBits")
    def rsa_bits(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        return pulumi.get(self, "rsa_bits")

    @rsa_bits.setter
    def rsa_bits(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "rsa_bits", value)


@pulumi.type_token("tls:index/privateKey:PrivateKey")
class PrivateKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 ecdsa_curve: Optional[pulumi.Input[_builtins.str]] = None,
                 rsa_bits: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        """
        Create a PrivateKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] algorithm: Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        :param pulumi.Input[_builtins.str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        :param pulumi.Input[_builtins.int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PrivateKey resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PrivateKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 ecdsa_curve: Optional[pulumi.Input[_builtins.str]] = None,
                 rsa_bits: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateKeyArgs.__new__(PrivateKeyArgs)

            if algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'algorithm'")
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["ecdsa_curve"] = ecdsa_curve
            __props__.__dict__["rsa_bits"] = rsa_bits
            __props__.__dict__["private_key_openssh"] = None
            __props__.__dict__["private_key_pem"] = None
            __props__.__dict__["private_key_pem_pkcs8"] = None
            __props__.__dict__["public_key_fingerprint_md5"] = None
            __props__.__dict__["public_key_fingerprint_sha256"] = None
            __props__.__dict__["public_key_openssh"] = None
            __props__.__dict__["public_key_pem"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKeyOpenssh", "privateKeyPem", "privateKeyPemPkcs8"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PrivateKey, __self__).__init__(
            'tls:index/privateKey:PrivateKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[_builtins.str]] = None,
            ecdsa_curve: Optional[pulumi.Input[_builtins.str]] = None,
            private_key_openssh: Optional[pulumi.Input[_builtins.str]] = None,
            private_key_pem: Optional[pulumi.Input[_builtins.str]] = None,
            private_key_pem_pkcs8: Optional[pulumi.Input[_builtins.str]] = None,
            public_key_fingerprint_md5: Optional[pulumi.Input[_builtins.str]] = None,
            public_key_fingerprint_sha256: Optional[pulumi.Input[_builtins.str]] = None,
            public_key_openssh: Optional[pulumi.Input[_builtins.str]] = None,
            public_key_pem: Optional[pulumi.Input[_builtins.str]] = None,
            rsa_bits: Optional[pulumi.Input[_builtins.int]] = None) -> 'PrivateKey':
        """
        Get an existing PrivateKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] algorithm: Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        :param pulumi.Input[_builtins.str] ecdsa_curve: When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        :param pulumi.Input[_builtins.str] private_key_openssh: Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        :param pulumi.Input[_builtins.str] private_key_pem: Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[_builtins.str] private_key_pem_pkcs8: Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        :param pulumi.Input[_builtins.str] public_key_fingerprint_md5: The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        :param pulumi.Input[_builtins.str] public_key_fingerprint_sha256: The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        :param pulumi.Input[_builtins.str] public_key_openssh: The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        :param pulumi.Input[_builtins.str] public_key_pem: Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        :param pulumi.Input[_builtins.int] rsa_bits: When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateKeyState.__new__(_PrivateKeyState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["ecdsa_curve"] = ecdsa_curve
        __props__.__dict__["private_key_openssh"] = private_key_openssh
        __props__.__dict__["private_key_pem"] = private_key_pem
        __props__.__dict__["private_key_pem_pkcs8"] = private_key_pem_pkcs8
        __props__.__dict__["public_key_fingerprint_md5"] = public_key_fingerprint_md5
        __props__.__dict__["public_key_fingerprint_sha256"] = public_key_fingerprint_sha256
        __props__.__dict__["public_key_openssh"] = public_key_openssh
        __props__.__dict__["public_key_pem"] = public_key_pem
        __props__.__dict__["rsa_bits"] = rsa_bits
        return PrivateKey(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the algorithm to use when generating the private key. Currently-supported values are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "algorithm")

    @_builtins.property
    @pulumi.getter(name="ecdsaCurve")
    def ecdsa_curve(self) -> pulumi.Output[_builtins.str]:
        """
        When `algorithm` is `ECDSA`, the name of the elliptic curve to use. Currently-supported values are: `P224`, `P256`, `P384`, `P521`. (default: `P224`).
        """
        return pulumi.get(self, "ecdsa_curve")

    @_builtins.property
    @pulumi.getter(name="privateKeyOpenssh")
    def private_key_openssh(self) -> pulumi.Output[_builtins.str]:
        """
        Private key data in [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format.
        """
        return pulumi.get(self, "private_key_openssh")

    @_builtins.property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Output[_builtins.str]:
        """
        Private key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "private_key_pem")

    @_builtins.property
    @pulumi.getter(name="privateKeyPemPkcs8")
    def private_key_pem_pkcs8(self) -> pulumi.Output[_builtins.str]:
        """
        Private key data in [PKCS#8 PEM (RFC 5208)](https://datatracker.ietf.org/doc/html/rfc5208) format.
        """
        return pulumi.get(self, "private_key_pem_pkcs8")

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintMd5")
    def public_key_fingerprint_md5(self) -> pulumi.Output[_builtins.str]:
        """
        The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_md5")

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintSha256")
    def public_key_fingerprint_sha256(self) -> pulumi.Output[_builtins.str]:
        """
        The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, similarly to `public_key_openssh` and the ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_sha256")

    @_builtins.property
    @pulumi.getter(name="publicKeyOpenssh")
    def public_key_openssh(self) -> pulumi.Output[_builtins.str]:
        """
        The public key data in ["Authorized Keys"](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_openssh")

    @_builtins.property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> pulumi.Output[_builtins.str]:
        """
        Public key data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_pem")

    @_builtins.property
    @pulumi.getter(name="rsaBits")
    def rsa_bits(self) -> pulumi.Output[_builtins.int]:
        """
        When `algorithm` is `RSA`, the size of the generated RSA key, in bits (default: `2048`).
        """
        return pulumi.get(self, "rsa_bits")

