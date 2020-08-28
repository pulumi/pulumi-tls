# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CertRequest']


class CertRequest(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 subjects: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]]]] = None,
                 uris: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a CertRequest resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]]] subjects: The subject for which a certificate is being requested. This is
               a nested configuration block whose structure is described below.
        :param pulumi.Input[List[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['dns_names'] = dns_names
            __props__['ip_addresses'] = ip_addresses
            if key_algorithm is None:
                raise TypeError("Missing required property 'key_algorithm'")
            __props__['key_algorithm'] = key_algorithm
            if private_key_pem is None:
                raise TypeError("Missing required property 'private_key_pem'")
            __props__['private_key_pem'] = private_key_pem
            if subjects is None:
                raise TypeError("Missing required property 'subjects'")
            __props__['subjects'] = subjects
            __props__['uris'] = uris
            __props__['cert_request_pem'] = None
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
            dns_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            ip_addresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            key_algorithm: Optional[pulumi.Input[str]] = None,
            private_key_pem: Optional[pulumi.Input[str]] = None,
            subjects: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]]]] = None,
            uris: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'CertRequest':
        """
        Get an existing CertRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_request_pem: The certificate request data in PEM format.
        :param pulumi.Input[List[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[List[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CertRequestSubjectArgs']]]] subjects: The subject for which a certificate is being requested. This is
               a nested configuration block whose structure is described below.
        :param pulumi.Input[List[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cert_request_pem"] = cert_request_pem
        __props__["dns_names"] = dns_names
        __props__["ip_addresses"] = ip_addresses
        __props__["key_algorithm"] = key_algorithm
        __props__["private_key_pem"] = private_key_pem
        __props__["subjects"] = subjects
        __props__["uris"] = uris
        return CertRequest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> pulumi.Output[str]:
        """
        The certificate request data in PEM format.
        """
        return pulumi.get(self, "cert_request_pem")

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of DNS names for which a certificate is being requested.
        """
        return pulumi.get(self, "dns_names")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of IP addresses for which a certificate is being requested.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> pulumi.Output[str]:
        """
        The name of the algorithm for the key provided
        in `private_key_pem`.
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Output[str]:
        """
        PEM-encoded private key that the certificate will belong to
        """
        return pulumi.get(self, "private_key_pem")

    @property
    @pulumi.getter
    def subjects(self) -> pulumi.Output[List['outputs.CertRequestSubject']]:
        """
        The subject for which a certificate is being requested. This is
        a nested configuration block whose structure is described below.
        """
        return pulumi.get(self, "subjects")

    @property
    @pulumi.getter
    def uris(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of URIs for which a certificate is being requested.
        """
        return pulumi.get(self, "uris")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

