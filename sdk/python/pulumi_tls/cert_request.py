# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CertRequestArgs', 'CertRequest']

@pulumi.input_type
class CertRequestArgs:
    def __init__(__self__, *,
                 private_key_pem: pulumi.Input[str],
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subject: Optional[pulumi.Input['CertRequestSubjectArgs']] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CertRequest resource.
        :param pulumi.Input[str] private_key_pem: Private key in PEM (RFC 1421) interpolation function.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input['CertRequestSubjectArgs'] subject: The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        pulumi.set(__self__, "private_key_pem", private_key_pem)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Input[str]:
        """
        Private key in PEM (RFC 1421) interpolation function.
        """
        return pulumi.get(self, "private_key_pem")

    @private_key_pem.setter
    def private_key_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key_pem", value)

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "dns_names")

    @dns_names.setter
    def dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_names", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input['CertRequestSubjectArgs']]:
        """
        The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input['CertRequestSubjectArgs']]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)


@pulumi.input_type
class _CertRequestState:
    def __init__(__self__, *,
                 cert_request_pem: Optional[pulumi.Input[str]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input['CertRequestSubjectArgs']] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering CertRequest resources.
        :param pulumi.Input[str] cert_request_pem: The certificate request data in PEM (RFC 1421).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[str] key_algorithm: Name of the algorithm used when generating the private key provided in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: Private key in PEM (RFC 1421) interpolation function.
        :param pulumi.Input['CertRequestSubjectArgs'] subject: The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        if cert_request_pem is not None:
            pulumi.set(__self__, "cert_request_pem", cert_request_pem)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if private_key_pem is not None:
            pulumi.set(__self__, "private_key_pem", private_key_pem)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate request data in PEM (RFC 1421).
        """
        return pulumi.get(self, "cert_request_pem")

    @cert_request_pem.setter
    def cert_request_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_request_pem", value)

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "dns_names")

    @dns_names.setter
    def dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_names", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the algorithm used when generating the private key provided in `private_key_pem`.
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> Optional[pulumi.Input[str]]:
        """
        Private key in PEM (RFC 1421) interpolation function.
        """
        return pulumi.get(self, "private_key_pem")

    @private_key_pem.setter
    def private_key_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_pem", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input['CertRequestSubjectArgs']]:
        """
        The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input['CertRequestSubjectArgs']]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)


class CertRequest(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_tls as tls

        example = tls.CertRequest("example",
            private_key_pem=(lambda path: open(path).read())("private_key.pem"),
            subject=tls.CertRequestSubjectArgs(
                common_name="example.com",
                organization="ACME Examples, Inc",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[str] private_key_pem: Private key in PEM (RFC 1421) interpolation function.
        :param pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']] subject: The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertRequestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_tls as tls

        example = tls.CertRequest("example",
            private_key_pem=(lambda path: open(path).read())("private_key.pem"),
            subject=tls.CertRequestSubjectArgs(
                common_name="example.com",
                organization="ACME Examples, Inc",
            ))
        ```

        :param str resource_name: The name of the resource.
        :param CertRequestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertRequestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CertRequestArgs.__new__(CertRequestArgs)

            __props__.__dict__["dns_names"] = dns_names
            __props__.__dict__["ip_addresses"] = ip_addresses
            if private_key_pem is None and not opts.urn:
                raise TypeError("Missing required property 'private_key_pem'")
            __props__.__dict__["private_key_pem"] = None if private_key_pem is None else pulumi.Output.secret(private_key_pem)
            __props__.__dict__["subject"] = subject
            __props__.__dict__["uris"] = uris
            __props__.__dict__["cert_request_pem"] = None
            __props__.__dict__["key_algorithm"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKeyPem"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CertRequest, __self__).__init__(
            'tls:index/certRequest:CertRequest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert_request_pem: Optional[pulumi.Input[str]] = None,
            dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key_algorithm: Optional[pulumi.Input[str]] = None,
            private_key_pem: Optional[pulumi.Input[str]] = None,
            subject: Optional[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]] = None,
            uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'CertRequest':
        """
        Get an existing CertRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_request_pem: The certificate request data in PEM (RFC 1421).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        :param pulumi.Input[str] key_algorithm: Name of the algorithm used when generating the private key provided in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: Private key in PEM (RFC 1421) interpolation function.
        :param pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']] subject: The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertRequestState.__new__(_CertRequestState)

        __props__.__dict__["cert_request_pem"] = cert_request_pem
        __props__.__dict__["dns_names"] = dns_names
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["key_algorithm"] = key_algorithm
        __props__.__dict__["private_key_pem"] = private_key_pem
        __props__.__dict__["subject"] = subject
        __props__.__dict__["uris"] = uris
        return CertRequest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> pulumi.Output[str]:
        """
        The certificate request data in PEM (RFC 1421).
        """
        return pulumi.get(self, "cert_request_pem")

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DNS names for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "dns_names")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP addresses for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> pulumi.Output[str]:
        """
        Name of the algorithm used when generating the private key provided in `private_key_pem`.
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Output[str]:
        """
        Private key in PEM (RFC 1421) interpolation function.
        """
        return pulumi.get(self, "private_key_pem")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[Optional['outputs.CertRequestSubject']]:
        """
        The subject for which a certificate is being requested. The acceptable arguments are all optional and their naming is based upon [Issuer Distinguished Names (RFC5280)](https://tools.ietf.org/html/rfc5280#section-4.1.2.4) section.
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def uris(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of URIs for which a certificate is being requested (i.e. certificate subjects).
        """
        return pulumi.get(self, "uris")

