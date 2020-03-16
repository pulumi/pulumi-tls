# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class CertRequest(pulumi.CustomResource):
    cert_request_pem: pulumi.Output[str]
    """
    The certificate request data in PEM format.
    """
    dns_names: pulumi.Output[list]
    """
    List of DNS names for which a certificate is being requested.
    """
    ip_addresses: pulumi.Output[list]
    """
    List of IP addresses for which a certificate is being requested.
    """
    key_algorithm: pulumi.Output[str]
    """
    The name of the algorithm for the key provided
    in `private_key_pem`.
    """
    private_key_pem: pulumi.Output[str]
    """
    PEM-encoded private key that the certificate will belong to
    """
    subjects: pulumi.Output[list]
    """
    The subject for which a certificate is being requested. This is
    a nested configuration block whose structure is described below.

      * `commonName` (`str`)
      * `country` (`str`)
      * `locality` (`str`)
      * `organization` (`str`)
      * `organizationalUnit` (`str`)
      * `postalCode` (`str`)
      * `province` (`str`)
      * `serialNumber` (`str`)
      * `streetAddresses` (`list`)
    """
    uris: pulumi.Output[list]
    """
    List of URIs for which a certificate is being requested.
    """
    def __init__(__self__, resource_name, opts=None, dns_names=None, ip_addresses=None, key_algorithm=None, private_key_pem=None, subjects=None, uris=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a CertRequest resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[list] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[list] subjects: The subject for which a certificate is being requested. This is
               a nested configuration block whose structure is described below.
        :param pulumi.Input[list] uris: List of URIs for which a certificate is being requested.

        The **subjects** object supports the following:

          * `commonName` (`pulumi.Input[str]`)
          * `country` (`pulumi.Input[str]`)
          * `locality` (`pulumi.Input[str]`)
          * `organization` (`pulumi.Input[str]`)
          * `organizationalUnit` (`pulumi.Input[str]`)
          * `postalCode` (`pulumi.Input[str]`)
          * `province` (`pulumi.Input[str]`)
          * `serialNumber` (`pulumi.Input[str]`)
          * `streetAddresses` (`pulumi.Input[list]`)
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, cert_request_pem=None, dns_names=None, ip_addresses=None, key_algorithm=None, private_key_pem=None, subjects=None, uris=None):
        """
        Get an existing CertRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_request_pem: The certificate request data in PEM format.
        :param pulumi.Input[list] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[list] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[list] subjects: The subject for which a certificate is being requested. This is
               a nested configuration block whose structure is described below.
        :param pulumi.Input[list] uris: List of URIs for which a certificate is being requested.

        The **subjects** object supports the following:

          * `commonName` (`pulumi.Input[str]`)
          * `country` (`pulumi.Input[str]`)
          * `locality` (`pulumi.Input[str]`)
          * `organization` (`pulumi.Input[str]`)
          * `organizationalUnit` (`pulumi.Input[str]`)
          * `postalCode` (`pulumi.Input[str]`)
          * `province` (`pulumi.Input[str]`)
          * `serialNumber` (`pulumi.Input[str]`)
          * `streetAddresses` (`pulumi.Input[list]`)
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

