# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'CertRequestSubject',
    'SelfSignedCertSubject',
    'GetCertificateCertificateResult',
]

@pulumi.output_type
class CertRequestSubject(dict):
    def __init__(__self__, *,
                 common_name: Optional[str] = None,
                 country: Optional[str] = None,
                 locality: Optional[str] = None,
                 organization: Optional[str] = None,
                 organizational_unit: Optional[str] = None,
                 postal_code: Optional[str] = None,
                 province: Optional[str] = None,
                 serial_number: Optional[str] = None,
                 street_addresses: Optional[Sequence[str]] = None):
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if locality is not None:
            pulumi.set(__self__, "locality", locality)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if organizational_unit is not None:
            pulumi.set(__self__, "organizational_unit", organizational_unit)
        if postal_code is not None:
            pulumi.set(__self__, "postal_code", postal_code)
        if province is not None:
            pulumi.set(__self__, "province", province)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if street_addresses is not None:
            pulumi.set(__self__, "street_addresses", street_addresses)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[str]:
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def country(self) -> Optional[str]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def locality(self) -> Optional[str]:
        return pulumi.get(self, "locality")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[str]:
        return pulumi.get(self, "organizational_unit")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[str]:
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter
    def province(self) -> Optional[str]:
        return pulumi.get(self, "province")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[str]:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="streetAddresses")
    def street_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "street_addresses")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SelfSignedCertSubject(dict):
    def __init__(__self__, *,
                 common_name: Optional[str] = None,
                 country: Optional[str] = None,
                 locality: Optional[str] = None,
                 organization: Optional[str] = None,
                 organizational_unit: Optional[str] = None,
                 postal_code: Optional[str] = None,
                 province: Optional[str] = None,
                 serial_number: Optional[str] = None,
                 street_addresses: Optional[Sequence[str]] = None):
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if locality is not None:
            pulumi.set(__self__, "locality", locality)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if organizational_unit is not None:
            pulumi.set(__self__, "organizational_unit", organizational_unit)
        if postal_code is not None:
            pulumi.set(__self__, "postal_code", postal_code)
        if province is not None:
            pulumi.set(__self__, "province", province)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if street_addresses is not None:
            pulumi.set(__self__, "street_addresses", street_addresses)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[str]:
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def country(self) -> Optional[str]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def locality(self) -> Optional[str]:
        return pulumi.get(self, "locality")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[str]:
        return pulumi.get(self, "organizational_unit")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[str]:
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter
    def province(self) -> Optional[str]:
        return pulumi.get(self, "province")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[str]:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="streetAddresses")
    def street_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "street_addresses")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetCertificateCertificateResult(dict):
    def __init__(__self__, *,
                 is_ca: bool,
                 issuer: str,
                 not_after: str,
                 not_before: str,
                 public_key_algorithm: str,
                 serial_number: str,
                 sha1_fingerprint: str,
                 signature_algorithm: str,
                 subject: str,
                 version: int):
        pulumi.set(__self__, "is_ca", is_ca)
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "not_after", not_after)
        pulumi.set(__self__, "not_before", not_before)
        pulumi.set(__self__, "public_key_algorithm", public_key_algorithm)
        pulumi.set(__self__, "serial_number", serial_number)
        pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)
        pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        pulumi.set(__self__, "subject", subject)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="isCa")
    def is_ca(self) -> bool:
        return pulumi.get(self, "is_ca")

    @property
    @pulumi.getter
    def issuer(self) -> str:
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> str:
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> str:
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="publicKeyAlgorithm")
    def public_key_algorithm(self) -> str:
        return pulumi.get(self, "public_key_algorithm")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> str:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> str:
        return pulumi.get(self, "sha1_fingerprint")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> str:
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter
    def subject(self) -> str:
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def version(self) -> int:
        return pulumi.get(self, "version")


