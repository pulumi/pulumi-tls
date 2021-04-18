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

__all__ = ['SelfSignedCertArgs', 'SelfSignedCert']

@pulumi.input_type
class SelfSignedCertArgs:
    def __init__(__self__, *,
                 allowed_uses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 key_algorithm: pulumi.Input[str],
                 private_key_pem: pulumi.Input[str],
                 subjects: pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]],
                 validity_period_hours: pulumi.Input[int],
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SelfSignedCert resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of keywords each describing a use that is permitted
               for the issued certificate. The valid keywords are listed below.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]] subjects: The subject for which a certificate is being requested.
               This is a nested configuration block whose structure matches the
               corresponding block for `CertRequest`.
        :param pulumi.Input[int] validity_period_hours: The number of hours after initial issuing that the
               certificate will become invalid.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: Number of hours before the certificates expiry when a new certificate will be generated
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[bool] is_ca_certificate: Boolean controlling whether the CA flag will be set in the
               generated certificate. Defaults to `false`, meaning that the certificate does not represent
               a certificate authority.
        :param pulumi.Input[bool] set_subject_key_id: If `true`, the certificate will include
               the subject key identifier. Defaults to `false`, in which case the subject
               key identifier is not set at all.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        """
        pulumi.set(__self__, "allowed_uses", allowed_uses)
        pulumi.set(__self__, "key_algorithm", key_algorithm)
        pulumi.set(__self__, "private_key_pem", private_key_pem)
        pulumi.set(__self__, "subjects", subjects)
        pulumi.set(__self__, "validity_period_hours", validity_period_hours)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if early_renewal_hours is not None:
            pulumi.set(__self__, "early_renewal_hours", early_renewal_hours)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if is_ca_certificate is not None:
            pulumi.set(__self__, "is_ca_certificate", is_ca_certificate)
        if set_subject_key_id is not None:
            pulumi.set(__self__, "set_subject_key_id", set_subject_key_id)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of keywords each describing a use that is permitted
        for the issued certificate. The valid keywords are listed below.
        """
        return pulumi.get(self, "allowed_uses")

    @allowed_uses.setter
    def allowed_uses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_uses", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> pulumi.Input[str]:
        """
        The name of the algorithm for the key provided
        in `private_key_pem`.
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> pulumi.Input[str]:
        """
        PEM-encoded private key that the certificate will belong to
        """
        return pulumi.get(self, "private_key_pem")

    @private_key_pem.setter
    def private_key_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key_pem", value)

    @property
    @pulumi.getter
    def subjects(self) -> pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]]:
        """
        The subject for which a certificate is being requested.
        This is a nested configuration block whose structure matches the
        corresponding block for `CertRequest`.
        """
        return pulumi.get(self, "subjects")

    @subjects.setter
    def subjects(self, value: pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]]):
        pulumi.set(self, "subjects", value)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Input[int]:
        """
        The number of hours after initial issuing that the
        certificate will become invalid.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: pulumi.Input[int]):
        pulumi.set(self, "validity_period_hours", value)

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DNS names for which a certificate is being requested.
        """
        return pulumi.get(self, "dns_names")

    @dns_names.setter
    def dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_names", value)

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours before the certificates expiry when a new certificate will be generated
        """
        return pulumi.get(self, "early_renewal_hours")

    @early_renewal_hours.setter
    def early_renewal_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "early_renewal_hours", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses for which a certificate is being requested.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean controlling whether the CA flag will be set in the
        generated certificate. Defaults to `false`, meaning that the certificate does not represent
        a certificate authority.
        """
        return pulumi.get(self, "is_ca_certificate")

    @is_ca_certificate.setter
    def is_ca_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ca_certificate", value)

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, the certificate will include
        the subject key identifier. Defaults to `false`, in which case the subject
        key identifier is not set at all.
        """
        return pulumi.get(self, "set_subject_key_id")

    @set_subject_key_id.setter
    def set_subject_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_subject_key_id", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URIs for which a certificate is being requested.
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)


@pulumi.input_type
class _SelfSignedCertState:
    def __init__(__self__, *,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cert_pem: Optional[pulumi.Input[str]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 ready_for_renewal: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subjects: Optional[pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validity_end_time: Optional[pulumi.Input[str]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 validity_start_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SelfSignedCert resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of keywords each describing a use that is permitted
               for the issued certificate. The valid keywords are listed below.
        :param pulumi.Input[str] cert_pem: The certificate data in PEM format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: Number of hours before the certificates expiry when a new certificate will be generated
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[bool] is_ca_certificate: Boolean controlling whether the CA flag will be set in the
               generated certificate. Defaults to `false`, meaning that the certificate does not represent
               a certificate authority.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[bool] set_subject_key_id: If `true`, the certificate will include
               the subject key identifier. Defaults to `false`, in which case the subject
               key identifier is not set at all.
        :param pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]] subjects: The subject for which a certificate is being requested.
               This is a nested configuration block whose structure matches the
               corresponding block for `CertRequest`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        :param pulumi.Input[str] validity_end_time: The time until which the certificate is invalid, as an
               [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param pulumi.Input[int] validity_period_hours: The number of hours after initial issuing that the
               certificate will become invalid.
        :param pulumi.Input[str] validity_start_time: The time after which the certificate is valid, as an
               [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        if allowed_uses is not None:
            pulumi.set(__self__, "allowed_uses", allowed_uses)
        if cert_pem is not None:
            pulumi.set(__self__, "cert_pem", cert_pem)
        if dns_names is not None:
            pulumi.set(__self__, "dns_names", dns_names)
        if early_renewal_hours is not None:
            pulumi.set(__self__, "early_renewal_hours", early_renewal_hours)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if is_ca_certificate is not None:
            pulumi.set(__self__, "is_ca_certificate", is_ca_certificate)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if private_key_pem is not None:
            pulumi.set(__self__, "private_key_pem", private_key_pem)
        if ready_for_renewal is not None:
            pulumi.set(__self__, "ready_for_renewal", ready_for_renewal)
        if set_subject_key_id is not None:
            pulumi.set(__self__, "set_subject_key_id", set_subject_key_id)
        if subjects is not None:
            pulumi.set(__self__, "subjects", subjects)
        if uris is not None:
            pulumi.set(__self__, "uris", uris)
        if validity_end_time is not None:
            pulumi.set(__self__, "validity_end_time", validity_end_time)
        if validity_period_hours is not None:
            pulumi.set(__self__, "validity_period_hours", validity_period_hours)
        if validity_start_time is not None:
            pulumi.set(__self__, "validity_start_time", validity_start_time)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of keywords each describing a use that is permitted
        for the issued certificate. The valid keywords are listed below.
        """
        return pulumi.get(self, "allowed_uses")

    @allowed_uses.setter
    def allowed_uses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_uses", value)

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate data in PEM format.
        """
        return pulumi.get(self, "cert_pem")

    @cert_pem.setter
    def cert_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_pem", value)

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of DNS names for which a certificate is being requested.
        """
        return pulumi.get(self, "dns_names")

    @dns_names.setter
    def dns_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "dns_names", value)

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours before the certificates expiry when a new certificate will be generated
        """
        return pulumi.get(self, "early_renewal_hours")

    @early_renewal_hours.setter
    def early_renewal_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "early_renewal_hours", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP addresses for which a certificate is being requested.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean controlling whether the CA flag will be set in the
        generated certificate. Defaults to `false`, meaning that the certificate does not represent
        a certificate authority.
        """
        return pulumi.get(self, "is_ca_certificate")

    @is_ca_certificate.setter
    def is_ca_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ca_certificate", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the algorithm for the key provided
        in `private_key_pem`.
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="privateKeyPem")
    def private_key_pem(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded private key that the certificate will belong to
        """
        return pulumi.get(self, "private_key_pem")

    @private_key_pem.setter
    def private_key_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_pem", value)

    @property
    @pulumi.getter(name="readyForRenewal")
    def ready_for_renewal(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ready_for_renewal")

    @ready_for_renewal.setter
    def ready_for_renewal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ready_for_renewal", value)

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, the certificate will include
        the subject key identifier. Defaults to `false`, in which case the subject
        key identifier is not set at all.
        """
        return pulumi.get(self, "set_subject_key_id")

    @set_subject_key_id.setter
    def set_subject_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_subject_key_id", value)

    @property
    @pulumi.getter
    def subjects(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]]]:
        """
        The subject for which a certificate is being requested.
        This is a nested configuration block whose structure matches the
        corresponding block for `CertRequest`.
        """
        return pulumi.get(self, "subjects")

    @subjects.setter
    def subjects(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SelfSignedCertSubjectArgs']]]]):
        pulumi.set(self, "subjects", value)

    @property
    @pulumi.getter
    def uris(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of URIs for which a certificate is being requested.
        """
        return pulumi.get(self, "uris")

    @uris.setter
    def uris(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uris", value)

    @property
    @pulumi.getter(name="validityEndTime")
    def validity_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time until which the certificate is invalid, as an
        [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_end_time")

    @validity_end_time.setter
    def validity_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validity_end_time", value)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        The number of hours after initial issuing that the
        certificate will become invalid.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "validity_period_hours", value)

    @property
    @pulumi.getter(name="validityStartTime")
    def validity_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time after which the certificate is valid, as an
        [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_start_time")

    @validity_start_time.setter
    def validity_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validity_start_time", value)


class SelfSignedCert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subjects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SelfSignedCertSubjectArgs']]]]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SelfSignedCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of keywords each describing a use that is permitted
               for the issued certificate. The valid keywords are listed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: Number of hours before the certificates expiry when a new certificate will be generated
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[bool] is_ca_certificate: Boolean controlling whether the CA flag will be set in the
               generated certificate. Defaults to `false`, meaning that the certificate does not represent
               a certificate authority.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[bool] set_subject_key_id: If `true`, the certificate will include
               the subject key identifier. Defaults to `false`, in which case the subject
               key identifier is not set at all.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SelfSignedCertSubjectArgs']]]] subjects: The subject for which a certificate is being requested.
               This is a nested configuration block whose structure matches the
               corresponding block for `CertRequest`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        :param pulumi.Input[int] validity_period_hours: The number of hours after initial issuing that the
               certificate will become invalid.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SelfSignedCertArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SelfSignedCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SelfSignedCertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SelfSignedCertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_pem: Optional[pulumi.Input[str]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 subjects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SelfSignedCertSubjectArgs']]]]] = None,
                 uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = SelfSignedCertArgs.__new__(SelfSignedCertArgs)

            if allowed_uses is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_uses'")
            __props__.__dict__["allowed_uses"] = allowed_uses
            __props__.__dict__["dns_names"] = dns_names
            __props__.__dict__["early_renewal_hours"] = early_renewal_hours
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["is_ca_certificate"] = is_ca_certificate
            if key_algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'key_algorithm'")
            __props__.__dict__["key_algorithm"] = key_algorithm
            if private_key_pem is None and not opts.urn:
                raise TypeError("Missing required property 'private_key_pem'")
            __props__.__dict__["private_key_pem"] = private_key_pem
            __props__.__dict__["set_subject_key_id"] = set_subject_key_id
            if subjects is None and not opts.urn:
                raise TypeError("Missing required property 'subjects'")
            __props__.__dict__["subjects"] = subjects
            __props__.__dict__["uris"] = uris
            if validity_period_hours is None and not opts.urn:
                raise TypeError("Missing required property 'validity_period_hours'")
            __props__.__dict__["validity_period_hours"] = validity_period_hours
            __props__.__dict__["cert_pem"] = None
            __props__.__dict__["ready_for_renewal"] = None
            __props__.__dict__["validity_end_time"] = None
            __props__.__dict__["validity_start_time"] = None
        super(SelfSignedCert, __self__).__init__(
            'tls:index/selfSignedCert:SelfSignedCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cert_pem: Optional[pulumi.Input[str]] = None,
            dns_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            early_renewal_hours: Optional[pulumi.Input[int]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            is_ca_certificate: Optional[pulumi.Input[bool]] = None,
            key_algorithm: Optional[pulumi.Input[str]] = None,
            private_key_pem: Optional[pulumi.Input[str]] = None,
            ready_for_renewal: Optional[pulumi.Input[bool]] = None,
            set_subject_key_id: Optional[pulumi.Input[bool]] = None,
            subjects: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SelfSignedCertSubjectArgs']]]]] = None,
            uris: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            validity_end_time: Optional[pulumi.Input[str]] = None,
            validity_period_hours: Optional[pulumi.Input[int]] = None,
            validity_start_time: Optional[pulumi.Input[str]] = None) -> 'SelfSignedCert':
        """
        Get an existing SelfSignedCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of keywords each describing a use that is permitted
               for the issued certificate. The valid keywords are listed below.
        :param pulumi.Input[str] cert_pem: The certificate data in PEM format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_names: List of DNS names for which a certificate is being requested.
        :param pulumi.Input[int] early_renewal_hours: Number of hours before the certificates expiry when a new certificate will be generated
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: List of IP addresses for which a certificate is being requested.
        :param pulumi.Input[bool] is_ca_certificate: Boolean controlling whether the CA flag will be set in the
               generated certificate. Defaults to `false`, meaning that the certificate does not represent
               a certificate authority.
        :param pulumi.Input[str] key_algorithm: The name of the algorithm for the key provided
               in `private_key_pem`.
        :param pulumi.Input[str] private_key_pem: PEM-encoded private key that the certificate will belong to
        :param pulumi.Input[bool] set_subject_key_id: If `true`, the certificate will include
               the subject key identifier. Defaults to `false`, in which case the subject
               key identifier is not set at all.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SelfSignedCertSubjectArgs']]]] subjects: The subject for which a certificate is being requested.
               This is a nested configuration block whose structure matches the
               corresponding block for `CertRequest`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uris: List of URIs for which a certificate is being requested.
        :param pulumi.Input[str] validity_end_time: The time until which the certificate is invalid, as an
               [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param pulumi.Input[int] validity_period_hours: The number of hours after initial issuing that the
               certificate will become invalid.
        :param pulumi.Input[str] validity_start_time: The time after which the certificate is valid, as an
               [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SelfSignedCertState.__new__(_SelfSignedCertState)

        __props__.__dict__["allowed_uses"] = allowed_uses
        __props__.__dict__["cert_pem"] = cert_pem
        __props__.__dict__["dns_names"] = dns_names
        __props__.__dict__["early_renewal_hours"] = early_renewal_hours
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["is_ca_certificate"] = is_ca_certificate
        __props__.__dict__["key_algorithm"] = key_algorithm
        __props__.__dict__["private_key_pem"] = private_key_pem
        __props__.__dict__["ready_for_renewal"] = ready_for_renewal
        __props__.__dict__["set_subject_key_id"] = set_subject_key_id
        __props__.__dict__["subjects"] = subjects
        __props__.__dict__["uris"] = uris
        __props__.__dict__["validity_end_time"] = validity_end_time
        __props__.__dict__["validity_period_hours"] = validity_period_hours
        __props__.__dict__["validity_start_time"] = validity_start_time
        return SelfSignedCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Output[Sequence[str]]:
        """
        List of keywords each describing a use that is permitted
        for the issued certificate. The valid keywords are listed below.
        """
        return pulumi.get(self, "allowed_uses")

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Output[str]:
        """
        The certificate data in PEM format.
        """
        return pulumi.get(self, "cert_pem")

    @property
    @pulumi.getter(name="dnsNames")
    def dns_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of DNS names for which a certificate is being requested.
        """
        return pulumi.get(self, "dns_names")

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> pulumi.Output[Optional[int]]:
        """
        Number of hours before the certificates expiry when a new certificate will be generated
        """
        return pulumi.get(self, "early_renewal_hours")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP addresses for which a certificate is being requested.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean controlling whether the CA flag will be set in the
        generated certificate. Defaults to `false`, meaning that the certificate does not represent
        a certificate authority.
        """
        return pulumi.get(self, "is_ca_certificate")

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
    @pulumi.getter(name="readyForRenewal")
    def ready_for_renewal(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "ready_for_renewal")

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, the certificate will include
        the subject key identifier. Defaults to `false`, in which case the subject
        key identifier is not set at all.
        """
        return pulumi.get(self, "set_subject_key_id")

    @property
    @pulumi.getter
    def subjects(self) -> pulumi.Output[Sequence['outputs.SelfSignedCertSubject']]:
        """
        The subject for which a certificate is being requested.
        This is a nested configuration block whose structure matches the
        corresponding block for `CertRequest`.
        """
        return pulumi.get(self, "subjects")

    @property
    @pulumi.getter
    def uris(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of URIs for which a certificate is being requested.
        """
        return pulumi.get(self, "uris")

    @property
    @pulumi.getter(name="validityEndTime")
    def validity_end_time(self) -> pulumi.Output[str]:
        """
        The time until which the certificate is invalid, as an
        [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_end_time")

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Output[int]:
        """
        The number of hours after initial issuing that the
        certificate will become invalid.
        """
        return pulumi.get(self, "validity_period_hours")

    @property
    @pulumi.getter(name="validityStartTime")
    def validity_start_time(self) -> pulumi.Output[str]:
        """
        The time after which the certificate is valid, as an
        [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_start_time")

