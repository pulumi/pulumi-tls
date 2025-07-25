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

__all__ = [
    'GetPublicKeyResult',
    'AwaitableGetPublicKeyResult',
    'get_public_key',
    'get_public_key_output',
]

@pulumi.output_type
class GetPublicKeyResult:
    """
    A collection of values returned by getPublicKey.
    """
    def __init__(__self__, algorithm=None, id=None, private_key_openssh=None, private_key_pem=None, public_key_fingerprint_md5=None, public_key_fingerprint_sha256=None, public_key_openssh=None, public_key_pem=None):
        if algorithm and not isinstance(algorithm, str):
            raise TypeError("Expected argument 'algorithm' to be a str")
        pulumi.set(__self__, "algorithm", algorithm)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if private_key_openssh and not isinstance(private_key_openssh, str):
            raise TypeError("Expected argument 'private_key_openssh' to be a str")
        pulumi.set(__self__, "private_key_openssh", private_key_openssh)
        if private_key_pem and not isinstance(private_key_pem, str):
            raise TypeError("Expected argument 'private_key_pem' to be a str")
        pulumi.set(__self__, "private_key_pem", private_key_pem)
        if public_key_fingerprint_md5 and not isinstance(public_key_fingerprint_md5, str):
            raise TypeError("Expected argument 'public_key_fingerprint_md5' to be a str")
        pulumi.set(__self__, "public_key_fingerprint_md5", public_key_fingerprint_md5)
        if public_key_fingerprint_sha256 and not isinstance(public_key_fingerprint_sha256, str):
            raise TypeError("Expected argument 'public_key_fingerprint_sha256' to be a str")
        pulumi.set(__self__, "public_key_fingerprint_sha256", public_key_fingerprint_sha256)
        if public_key_openssh and not isinstance(public_key_openssh, str):
            raise TypeError("Expected argument 'public_key_openssh' to be a str")
        pulumi.set(__self__, "public_key_openssh", public_key_openssh)
        if public_key_pem and not isinstance(public_key_pem, str):
            raise TypeError("Expected argument 'public_key_pem' to be a str")
        pulumi.set(__self__, "public_key_pem", public_key_pem)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> _builtins.str:
        """
        The name of the algorithm used by the given private key. Possible values are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "algorithm")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        Unique identifier for this data source: hexadecimal representation of the SHA1 checksum of the data source.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="privateKeyOpenssh")
    def private_key_openssh(self) -> Optional[_builtins.str]:
        """
        The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "private_key_openssh")

    @_builtins.property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> Optional[_builtins.str]:
        """
        The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
        """
        return pulumi.get(self, "private_key_pem")

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintMd5")
    def public_key_fingerprint_md5(self) -> _builtins.str:
        """
        The fingerprint of the public key data in OpenSSH MD5 hash format, e.g. `aa:bb:cc:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_md5")

    @_builtins.property
    @pulumi.getter(name="publicKeyFingerprintSha256")
    def public_key_fingerprint_sha256(self) -> _builtins.str:
        """
        The fingerprint of the public key data in OpenSSH SHA256 hash format, e.g. `SHA256:...`. Only available if the selected private key format is compatible, as per the rules for `public_key_openssh` and ECDSA P224 limitations.
        """
        return pulumi.get(self, "public_key_fingerprint_sha256")

    @_builtins.property
    @pulumi.getter(name="publicKeyOpenssh")
    def public_key_openssh(self) -> _builtins.str:
        """
        The public key, in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format. This is also known as ['Authorized Keys'](https://www.ssh.com/academy/ssh/authorized_keys/openssh#format-of-the-authorized-keys-file) format. This is not populated for `ECDSA` with curve `P224`, as it is not supported. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_openssh")

    @_builtins.property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> _builtins.str:
        """
        The public key, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "public_key_pem")


class AwaitableGetPublicKeyResult(GetPublicKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPublicKeyResult(
            algorithm=self.algorithm,
            id=self.id,
            private_key_openssh=self.private_key_openssh,
            private_key_pem=self.private_key_pem,
            public_key_fingerprint_md5=self.public_key_fingerprint_md5,
            public_key_fingerprint_sha256=self.public_key_fingerprint_sha256,
            public_key_openssh=self.public_key_openssh,
            public_key_pem=self.public_key_pem)


def get_public_key(private_key_openssh: Optional[_builtins.str] = None,
                   private_key_pem: Optional[_builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPublicKeyResult:
    """
    Get a public key from a PEM-encoded private key.

    Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_std as std
    import pulumi_tls as tls

    ed25519_example = tls.PrivateKey("ed25519-example", algorithm="ED25519")
    # Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
    private_key_pem_example = tls.get_public_key_output(private_key_pem=ed25519_example.private_key_pem)
    # Public key loaded from filesystem, using the Open SSH (RFC 4716) format
    private_key_openssh_example = tls.get_public_key(private_key_openssh=std.file(input="~/.ssh/id_rsa_rfc4716").result)
    ```


    :param _builtins.str private_key_openssh: The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
    :param _builtins.str private_key_pem: The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
    """
    __args__ = dict()
    __args__['privateKeyOpenssh'] = private_key_openssh
    __args__['privateKeyPem'] = private_key_pem
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tls:index/getPublicKey:getPublicKey', __args__, opts=opts, typ=GetPublicKeyResult).value

    return AwaitableGetPublicKeyResult(
        algorithm=pulumi.get(__ret__, 'algorithm'),
        id=pulumi.get(__ret__, 'id'),
        private_key_openssh=pulumi.get(__ret__, 'private_key_openssh'),
        private_key_pem=pulumi.get(__ret__, 'private_key_pem'),
        public_key_fingerprint_md5=pulumi.get(__ret__, 'public_key_fingerprint_md5'),
        public_key_fingerprint_sha256=pulumi.get(__ret__, 'public_key_fingerprint_sha256'),
        public_key_openssh=pulumi.get(__ret__, 'public_key_openssh'),
        public_key_pem=pulumi.get(__ret__, 'public_key_pem'))
def get_public_key_output(private_key_openssh: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          private_key_pem: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPublicKeyResult]:
    """
    Get a public key from a PEM-encoded private key.

    Use this data source to get the public key from a [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) or [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) formatted private key, for use in other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_std as std
    import pulumi_tls as tls

    ed25519_example = tls.PrivateKey("ed25519-example", algorithm="ED25519")
    # Public key loaded from a terraform-generated private key, using the PEM (RFC 1421) format
    private_key_pem_example = tls.get_public_key_output(private_key_pem=ed25519_example.private_key_pem)
    # Public key loaded from filesystem, using the Open SSH (RFC 4716) format
    private_key_openssh_example = tls.get_public_key(private_key_openssh=std.file(input="~/.ssh/id_rsa_rfc4716").result)
    ```


    :param _builtins.str private_key_openssh: The private key (in  [OpenSSH PEM (RFC 4716)](https://datatracker.ietf.org/doc/html/rfc4716) format) to extract the public key from. This is *mutually exclusive* with `private_key_pem`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
    :param _builtins.str private_key_pem: The private key (in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format) to extract the public key from. This is *mutually exclusive* with `private_key_openssh`. Currently-supported algorithms for keys are: `RSA`, `ECDSA`, `ED25519`.
    """
    __args__ = dict()
    __args__['privateKeyOpenssh'] = private_key_openssh
    __args__['privateKeyPem'] = private_key_pem
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('tls:index/getPublicKey:getPublicKey', __args__, opts=opts, typ=GetPublicKeyResult)
    return __ret__.apply(lambda __response__: GetPublicKeyResult(
        algorithm=pulumi.get(__response__, 'algorithm'),
        id=pulumi.get(__response__, 'id'),
        private_key_openssh=pulumi.get(__response__, 'private_key_openssh'),
        private_key_pem=pulumi.get(__response__, 'private_key_pem'),
        public_key_fingerprint_md5=pulumi.get(__response__, 'public_key_fingerprint_md5'),
        public_key_fingerprint_sha256=pulumi.get(__response__, 'public_key_fingerprint_sha256'),
        public_key_openssh=pulumi.get(__response__, 'public_key_openssh'),
        public_key_pem=pulumi.get(__response__, 'public_key_pem')))
