# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LocallySignedCertArgs', 'LocallySignedCert']

@pulumi.input_type
class LocallySignedCertArgs:
    def __init__(__self__, *,
                 allowed_uses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 ca_cert_pem: pulumi.Input[str],
                 ca_private_key_pem: pulumi.Input[str],
                 cert_request_pem: pulumi.Input[str],
                 validity_period_hours: pulumi.Input[int],
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a LocallySignedCert resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        :param pulumi.Input[str] ca_cert_pem: Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] ca_private_key_pem: Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] cert_request_pem: Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid for.
        :param pulumi.Input[bool] is_ca_certificate: Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        """
        pulumi.set(__self__, "allowed_uses", allowed_uses)
        pulumi.set(__self__, "ca_cert_pem", ca_cert_pem)
        pulumi.set(__self__, "ca_private_key_pem", ca_private_key_pem)
        pulumi.set(__self__, "cert_request_pem", cert_request_pem)
        pulumi.set(__self__, "validity_period_hours", validity_period_hours)
        if early_renewal_hours is not None:
            pulumi.set(__self__, "early_renewal_hours", early_renewal_hours)
        if is_ca_certificate is not None:
            pulumi.set(__self__, "is_ca_certificate", is_ca_certificate)
        if set_subject_key_id is not None:
            pulumi.set(__self__, "set_subject_key_id", set_subject_key_id)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        """
        return pulumi.get(self, "allowed_uses")

    @allowed_uses.setter
    def allowed_uses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_uses", value)

    @property
    @pulumi.getter(name="caCertPem")
    def ca_cert_pem(self) -> pulumi.Input[str]:
        """
        Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_cert_pem")

    @ca_cert_pem.setter
    def ca_cert_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_cert_pem", value)

    @property
    @pulumi.getter(name="caPrivateKeyPem")
    def ca_private_key_pem(self) -> pulumi.Input[str]:
        """
        Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_private_key_pem")

    @ca_private_key_pem.setter
    def ca_private_key_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_private_key_pem", value)

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> pulumi.Input[str]:
        """
        Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "cert_request_pem")

    @cert_request_pem.setter
    def cert_request_pem(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_request_pem", value)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Input[int]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid for.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: pulumi.Input[int]):
        pulumi.set(self, "validity_period_hours", value)

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "early_renewal_hours")

    @early_renewal_hours.setter
    def early_renewal_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "early_renewal_hours", value)

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        """
        return pulumi.get(self, "is_ca_certificate")

    @is_ca_certificate.setter
    def is_ca_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ca_certificate", value)

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        """
        return pulumi.get(self, "set_subject_key_id")

    @set_subject_key_id.setter
    def set_subject_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_subject_key_id", value)


@pulumi.input_type
class _LocallySignedCertState:
    def __init__(__self__, *,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert_pem: Optional[pulumi.Input[str]] = None,
                 ca_key_algorithm: Optional[pulumi.Input[str]] = None,
                 ca_private_key_pem: Optional[pulumi.Input[str]] = None,
                 cert_pem: Optional[pulumi.Input[str]] = None,
                 cert_request_pem: Optional[pulumi.Input[str]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 ready_for_renewal: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 validity_end_time: Optional[pulumi.Input[str]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 validity_start_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LocallySignedCert resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        :param pulumi.Input[str] ca_cert_pem: Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] ca_key_algorithm: Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
        :param pulumi.Input[str] ca_private_key_pem: Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] cert_pem: Certificate data in PEM (RFC 1421).
        :param pulumi.Input[str] cert_request_pem: Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[bool] is_ca_certificate: Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        :param pulumi.Input[bool] ready_for_renewal: Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        :param pulumi.Input[str] validity_end_time: The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid for.
        :param pulumi.Input[str] validity_start_time: The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        if allowed_uses is not None:
            pulumi.set(__self__, "allowed_uses", allowed_uses)
        if ca_cert_pem is not None:
            pulumi.set(__self__, "ca_cert_pem", ca_cert_pem)
        if ca_key_algorithm is not None:
            pulumi.set(__self__, "ca_key_algorithm", ca_key_algorithm)
        if ca_private_key_pem is not None:
            pulumi.set(__self__, "ca_private_key_pem", ca_private_key_pem)
        if cert_pem is not None:
            pulumi.set(__self__, "cert_pem", cert_pem)
        if cert_request_pem is not None:
            pulumi.set(__self__, "cert_request_pem", cert_request_pem)
        if early_renewal_hours is not None:
            pulumi.set(__self__, "early_renewal_hours", early_renewal_hours)
        if is_ca_certificate is not None:
            pulumi.set(__self__, "is_ca_certificate", is_ca_certificate)
        if ready_for_renewal is not None:
            pulumi.set(__self__, "ready_for_renewal", ready_for_renewal)
        if set_subject_key_id is not None:
            pulumi.set(__self__, "set_subject_key_id", set_subject_key_id)
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
        List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        """
        return pulumi.get(self, "allowed_uses")

    @allowed_uses.setter
    def allowed_uses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_uses", value)

    @property
    @pulumi.getter(name="caCertPem")
    def ca_cert_pem(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_cert_pem")

    @ca_cert_pem.setter
    def ca_cert_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert_pem", value)

    @property
    @pulumi.getter(name="caKeyAlgorithm")
    def ca_key_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
        """
        return pulumi.get(self, "ca_key_algorithm")

    @ca_key_algorithm.setter
    def ca_key_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_key_algorithm", value)

    @property
    @pulumi.getter(name="caPrivateKeyPem")
    def ca_private_key_pem(self) -> Optional[pulumi.Input[str]]:
        """
        Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_private_key_pem")

    @ca_private_key_pem.setter
    def ca_private_key_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_private_key_pem", value)

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate data in PEM (RFC 1421).
        """
        return pulumi.get(self, "cert_pem")

    @cert_pem.setter
    def cert_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_pem", value)

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "cert_request_pem")

    @cert_request_pem.setter
    def cert_request_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_request_pem", value)

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "early_renewal_hours")

    @early_renewal_hours.setter
    def early_renewal_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "early_renewal_hours", value)

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        """
        return pulumi.get(self, "is_ca_certificate")

    @is_ca_certificate.setter
    def is_ca_certificate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ca_certificate", value)

    @property
    @pulumi.getter(name="readyForRenewal")
    def ready_for_renewal(self) -> Optional[pulumi.Input[bool]]:
        """
        Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        """
        return pulumi.get(self, "ready_for_renewal")

    @ready_for_renewal.setter
    def ready_for_renewal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ready_for_renewal", value)

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        """
        return pulumi.get(self, "set_subject_key_id")

    @set_subject_key_id.setter
    def set_subject_key_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "set_subject_key_id", value)

    @property
    @pulumi.getter(name="validityEndTime")
    def validity_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_end_time")

    @validity_end_time.setter
    def validity_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validity_end_time", value)

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> Optional[pulumi.Input[int]]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid for.
        """
        return pulumi.get(self, "validity_period_hours")

    @validity_period_hours.setter
    def validity_period_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "validity_period_hours", value)

    @property
    @pulumi.getter(name="validityStartTime")
    def validity_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_start_time")

    @validity_start_time.setter
    def validity_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validity_start_time", value)


class LocallySignedCert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert_pem: Optional[pulumi.Input[str]] = None,
                 ca_private_key_pem: Optional[pulumi.Input[str]] = None,
                 cert_request_pem: Optional[pulumi.Input[str]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a LocallySignedCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        :param pulumi.Input[str] ca_cert_pem: Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] ca_private_key_pem: Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] cert_request_pem: Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[bool] is_ca_certificate: Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid for.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LocallySignedCertArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LocallySignedCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LocallySignedCertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LocallySignedCertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_cert_pem: Optional[pulumi.Input[str]] = None,
                 ca_private_key_pem: Optional[pulumi.Input[str]] = None,
                 cert_request_pem: Optional[pulumi.Input[str]] = None,
                 early_renewal_hours: Optional[pulumi.Input[int]] = None,
                 is_ca_certificate: Optional[pulumi.Input[bool]] = None,
                 set_subject_key_id: Optional[pulumi.Input[bool]] = None,
                 validity_period_hours: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LocallySignedCertArgs.__new__(LocallySignedCertArgs)

            if allowed_uses is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_uses'")
            __props__.__dict__["allowed_uses"] = allowed_uses
            if ca_cert_pem is None and not opts.urn:
                raise TypeError("Missing required property 'ca_cert_pem'")
            __props__.__dict__["ca_cert_pem"] = ca_cert_pem
            if ca_private_key_pem is None and not opts.urn:
                raise TypeError("Missing required property 'ca_private_key_pem'")
            __props__.__dict__["ca_private_key_pem"] = None if ca_private_key_pem is None else pulumi.Output.secret(ca_private_key_pem)
            if cert_request_pem is None and not opts.urn:
                raise TypeError("Missing required property 'cert_request_pem'")
            __props__.__dict__["cert_request_pem"] = cert_request_pem
            __props__.__dict__["early_renewal_hours"] = early_renewal_hours
            __props__.__dict__["is_ca_certificate"] = is_ca_certificate
            __props__.__dict__["set_subject_key_id"] = set_subject_key_id
            if validity_period_hours is None and not opts.urn:
                raise TypeError("Missing required property 'validity_period_hours'")
            __props__.__dict__["validity_period_hours"] = validity_period_hours
            __props__.__dict__["ca_key_algorithm"] = None
            __props__.__dict__["cert_pem"] = None
            __props__.__dict__["ready_for_renewal"] = None
            __props__.__dict__["validity_end_time"] = None
            __props__.__dict__["validity_start_time"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["caPrivateKeyPem"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LocallySignedCert, __self__).__init__(
            'tls:index/locallySignedCert:LocallySignedCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_uses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ca_cert_pem: Optional[pulumi.Input[str]] = None,
            ca_key_algorithm: Optional[pulumi.Input[str]] = None,
            ca_private_key_pem: Optional[pulumi.Input[str]] = None,
            cert_pem: Optional[pulumi.Input[str]] = None,
            cert_request_pem: Optional[pulumi.Input[str]] = None,
            early_renewal_hours: Optional[pulumi.Input[int]] = None,
            is_ca_certificate: Optional[pulumi.Input[bool]] = None,
            ready_for_renewal: Optional[pulumi.Input[bool]] = None,
            set_subject_key_id: Optional[pulumi.Input[bool]] = None,
            validity_end_time: Optional[pulumi.Input[str]] = None,
            validity_period_hours: Optional[pulumi.Input[int]] = None,
            validity_start_time: Optional[pulumi.Input[str]] = None) -> 'LocallySignedCert':
        """
        Get an existing LocallySignedCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uses: List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        :param pulumi.Input[str] ca_cert_pem: Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] ca_key_algorithm: Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
        :param pulumi.Input[str] ca_private_key_pem: Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[str] cert_pem: Certificate data in PEM (RFC 1421).
        :param pulumi.Input[str] cert_request_pem: Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        :param pulumi.Input[bool] is_ca_certificate: Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        :param pulumi.Input[bool] ready_for_renewal: Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        :param pulumi.Input[bool] set_subject_key_id: Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        :param pulumi.Input[str] validity_end_time: The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param pulumi.Input[int] validity_period_hours: Number of hours, after initial issuing, that the certificate will remain valid for.
        :param pulumi.Input[str] validity_start_time: The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LocallySignedCertState.__new__(_LocallySignedCertState)

        __props__.__dict__["allowed_uses"] = allowed_uses
        __props__.__dict__["ca_cert_pem"] = ca_cert_pem
        __props__.__dict__["ca_key_algorithm"] = ca_key_algorithm
        __props__.__dict__["ca_private_key_pem"] = ca_private_key_pem
        __props__.__dict__["cert_pem"] = cert_pem
        __props__.__dict__["cert_request_pem"] = cert_request_pem
        __props__.__dict__["early_renewal_hours"] = early_renewal_hours
        __props__.__dict__["is_ca_certificate"] = is_ca_certificate
        __props__.__dict__["ready_for_renewal"] = ready_for_renewal
        __props__.__dict__["set_subject_key_id"] = set_subject_key_id
        __props__.__dict__["validity_end_time"] = validity_end_time
        __props__.__dict__["validity_period_hours"] = validity_period_hours
        __props__.__dict__["validity_start_time"] = validity_start_time
        return LocallySignedCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedUses")
    def allowed_uses(self) -> pulumi.Output[Sequence[str]]:
        """
        List of key usages allowed for the issued certificate. Values are defined in [RFC 5280](https://datatracker.ietf.org/doc/html/rfc5280) and combine flags defined by both [Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3) and [Extended Key Usages](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12). Accepted values: `any_extended`, `cert_signing`, `client_auth`, `code_signing`, `content_commitment`, `crl_signing`, `data_encipherment`, `decipher_only`, `digital_signature`, `email_protection`, `encipher_only`, `ipsec_end_system`, `ipsec_tunnel`, `ipsec_user`, `key_agreement`, `key_encipherment`, `microsoft_commercial_code_signing`, `microsoft_kernel_code_signing`, `microsoft_server_gated_crypto`, `netscape_server_gated_crypto`, `ocsp_signing`, `server_auth`, `timestamping`.
        """
        return pulumi.get(self, "allowed_uses")

    @property
    @pulumi.getter(name="caCertPem")
    def ca_cert_pem(self) -> pulumi.Output[str]:
        """
        Certificate data of the Certificate Authority (CA) in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_cert_pem")

    @property
    @pulumi.getter(name="caKeyAlgorithm")
    def ca_key_algorithm(self) -> pulumi.Output[str]:
        """
        Name of the algorithm used when generating the private key provided in `ca_private_key_pem`.
        """
        return pulumi.get(self, "ca_key_algorithm")

    @property
    @pulumi.getter(name="caPrivateKeyPem")
    def ca_private_key_pem(self) -> pulumi.Output[str]:
        """
        Private key of the Certificate Authority (CA) used to sign the certificate, in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "ca_private_key_pem")

    @property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> pulumi.Output[str]:
        """
        Certificate data in PEM (RFC 1421).
        """
        return pulumi.get(self, "cert_pem")

    @property
    @pulumi.getter(name="certRequestPem")
    def cert_request_pem(self) -> pulumi.Output[str]:
        """
        Certificate request data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format.
        """
        return pulumi.get(self, "cert_request_pem")

    @property
    @pulumi.getter(name="earlyRenewalHours")
    def early_renewal_hours(self) -> pulumi.Output[int]:
        return pulumi.get(self, "early_renewal_hours")

    @property
    @pulumi.getter(name="isCaCertificate")
    def is_ca_certificate(self) -> pulumi.Output[bool]:
        """
        Is the generated certificate representing a Certificate Authority (CA) (default: `false`).
        """
        return pulumi.get(self, "is_ca_certificate")

    @property
    @pulumi.getter(name="readyForRenewal")
    def ready_for_renewal(self) -> pulumi.Output[bool]:
        """
        Is the certificate either expired (i.e. beyond the `validity_period_hours`) or ready for an early renewal (i.e. within the `early_renewal_hours`)?
        """
        return pulumi.get(self, "ready_for_renewal")

    @property
    @pulumi.getter(name="setSubjectKeyId")
    def set_subject_key_id(self) -> pulumi.Output[bool]:
        """
        Should the generated certificate include a [subject key identifier](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.2) (default: `false`).
        """
        return pulumi.get(self, "set_subject_key_id")

    @property
    @pulumi.getter(name="validityEndTime")
    def validity_end_time(self) -> pulumi.Output[str]:
        """
        The time until which the certificate is invalid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_end_time")

    @property
    @pulumi.getter(name="validityPeriodHours")
    def validity_period_hours(self) -> pulumi.Output[int]:
        """
        Number of hours, after initial issuing, that the certificate will remain valid for.
        """
        return pulumi.get(self, "validity_period_hours")

    @property
    @pulumi.getter(name="validityStartTime")
    def validity_start_time(self) -> pulumi.Output[str]:
        """
        The time after which the certificate is valid, expressed as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "validity_start_time")

