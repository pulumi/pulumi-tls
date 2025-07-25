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
    'CertRequestSubject',
    'SelfSignedCertSubject',
    'GetCertificateCertificateResult',
]

@pulumi.output_type
class CertRequestSubject(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "commonName":
            suggest = "common_name"
        elif key == "emailAddress":
            suggest = "email_address"
        elif key == "organizationalUnit":
            suggest = "organizational_unit"
        elif key == "postalCode":
            suggest = "postal_code"
        elif key == "serialNumber":
            suggest = "serial_number"
        elif key == "streetAddresses":
            suggest = "street_addresses"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CertRequestSubject. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CertRequestSubject.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CertRequestSubject.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 common_name: Optional[_builtins.str] = None,
                 country: Optional[_builtins.str] = None,
                 email_address: Optional[_builtins.str] = None,
                 locality: Optional[_builtins.str] = None,
                 organization: Optional[_builtins.str] = None,
                 organizational_unit: Optional[_builtins.str] = None,
                 postal_code: Optional[_builtins.str] = None,
                 province: Optional[_builtins.str] = None,
                 serial_number: Optional[_builtins.str] = None,
                 street_addresses: Optional[Sequence[_builtins.str]] = None):
        """
        :param _builtins.str common_name: Distinguished name: `CN`
        :param _builtins.str country: Distinguished name: `C`
        :param _builtins.str email_address: ASN.1 Object Identifier (OID): `1.2.840.113549.1.9.1`
        :param _builtins.str locality: Distinguished name: `L`
        :param _builtins.str organization: Distinguished name: `O`
        :param _builtins.str organizational_unit: Distinguished name: `OU`
        :param _builtins.str postal_code: Distinguished name: `PC`
        :param _builtins.str province: Distinguished name: `ST`
        :param _builtins.str serial_number: Distinguished name: `SERIALNUMBER`
        :param Sequence[_builtins.str] street_addresses: Distinguished name: `STREET`
        """
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if email_address is not None:
            pulumi.set(__self__, "email_address", email_address)
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

    @_builtins.property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `CN`
        """
        return pulumi.get(self, "common_name")

    @_builtins.property
    @pulumi.getter
    def country(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `C`
        """
        return pulumi.get(self, "country")

    @_builtins.property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> Optional[_builtins.str]:
        """
        ASN.1 Object Identifier (OID): `1.2.840.113549.1.9.1`
        """
        return pulumi.get(self, "email_address")

    @_builtins.property
    @pulumi.getter
    def locality(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `L`
        """
        return pulumi.get(self, "locality")

    @_builtins.property
    @pulumi.getter
    def organization(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `O`
        """
        return pulumi.get(self, "organization")

    @_builtins.property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `OU`
        """
        return pulumi.get(self, "organizational_unit")

    @_builtins.property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `PC`
        """
        return pulumi.get(self, "postal_code")

    @_builtins.property
    @pulumi.getter
    def province(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `ST`
        """
        return pulumi.get(self, "province")

    @_builtins.property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `SERIALNUMBER`
        """
        return pulumi.get(self, "serial_number")

    @_builtins.property
    @pulumi.getter(name="streetAddresses")
    def street_addresses(self) -> Optional[Sequence[_builtins.str]]:
        """
        Distinguished name: `STREET`
        """
        return pulumi.get(self, "street_addresses")


@pulumi.output_type
class SelfSignedCertSubject(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "commonName":
            suggest = "common_name"
        elif key == "emailAddress":
            suggest = "email_address"
        elif key == "organizationalUnit":
            suggest = "organizational_unit"
        elif key == "postalCode":
            suggest = "postal_code"
        elif key == "serialNumber":
            suggest = "serial_number"
        elif key == "streetAddresses":
            suggest = "street_addresses"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SelfSignedCertSubject. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SelfSignedCertSubject.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SelfSignedCertSubject.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 common_name: Optional[_builtins.str] = None,
                 country: Optional[_builtins.str] = None,
                 email_address: Optional[_builtins.str] = None,
                 locality: Optional[_builtins.str] = None,
                 organization: Optional[_builtins.str] = None,
                 organizational_unit: Optional[_builtins.str] = None,
                 postal_code: Optional[_builtins.str] = None,
                 province: Optional[_builtins.str] = None,
                 serial_number: Optional[_builtins.str] = None,
                 street_addresses: Optional[Sequence[_builtins.str]] = None):
        """
        :param _builtins.str common_name: Distinguished name: `CN`
        :param _builtins.str country: Distinguished name: `C`
        :param _builtins.str email_address: ASN.1 Object Identifier (OID): `1.2.840.113549.1.9.1`
        :param _builtins.str locality: Distinguished name: `L`
        :param _builtins.str organization: Distinguished name: `O`
        :param _builtins.str organizational_unit: Distinguished name: `OU`
        :param _builtins.str postal_code: Distinguished name: `PC`
        :param _builtins.str province: Distinguished name: `ST`
        :param _builtins.str serial_number: Distinguished name: `SERIALNUMBER`
        :param Sequence[_builtins.str] street_addresses: Distinguished name: `STREET`
        """
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if email_address is not None:
            pulumi.set(__self__, "email_address", email_address)
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

    @_builtins.property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `CN`
        """
        return pulumi.get(self, "common_name")

    @_builtins.property
    @pulumi.getter
    def country(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `C`
        """
        return pulumi.get(self, "country")

    @_builtins.property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> Optional[_builtins.str]:
        """
        ASN.1 Object Identifier (OID): `1.2.840.113549.1.9.1`
        """
        return pulumi.get(self, "email_address")

    @_builtins.property
    @pulumi.getter
    def locality(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `L`
        """
        return pulumi.get(self, "locality")

    @_builtins.property
    @pulumi.getter
    def organization(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `O`
        """
        return pulumi.get(self, "organization")

    @_builtins.property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `OU`
        """
        return pulumi.get(self, "organizational_unit")

    @_builtins.property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `PC`
        """
        return pulumi.get(self, "postal_code")

    @_builtins.property
    @pulumi.getter
    def province(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `ST`
        """
        return pulumi.get(self, "province")

    @_builtins.property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[_builtins.str]:
        """
        Distinguished name: `SERIALNUMBER`
        """
        return pulumi.get(self, "serial_number")

    @_builtins.property
    @pulumi.getter(name="streetAddresses")
    def street_addresses(self) -> Optional[Sequence[_builtins.str]]:
        """
        Distinguished name: `STREET`
        """
        return pulumi.get(self, "street_addresses")


@pulumi.output_type
class GetCertificateCertificateResult(dict):
    def __init__(__self__, *,
                 cert_pem: _builtins.str,
                 is_ca: _builtins.bool,
                 issuer: _builtins.str,
                 not_after: _builtins.str,
                 not_before: _builtins.str,
                 public_key_algorithm: _builtins.str,
                 serial_number: _builtins.str,
                 sha1_fingerprint: _builtins.str,
                 signature_algorithm: _builtins.str,
                 subject: _builtins.str,
                 version: _builtins.int):
        """
        :param _builtins.str cert_pem: Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        :param _builtins.bool is_ca: `true` if the certificate is of a CA (Certificate Authority).
        :param _builtins.str issuer: Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        :param _builtins.str not_after: The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param _builtins.str not_before: The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        :param _builtins.str public_key_algorithm: The key algorithm used to create the certificate.
        :param _builtins.str serial_number: Number that uniquely identifies the certificate with the CA's system.
               The `format` function can be used to convert this *base 10* number into other bases, such as hex.
        :param _builtins.str sha1_fingerprint: The SHA1 fingerprint of the public key of the certificate.
        :param _builtins.str signature_algorithm: The algorithm used to sign the certificate.
        :param _builtins.str subject: The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        :param _builtins.int version: The version the certificate is in.
        """
        pulumi.set(__self__, "cert_pem", cert_pem)
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

    @_builtins.property
    @pulumi.getter(name="certPem")
    def cert_pem(self) -> _builtins.str:
        """
        Certificate data in [PEM (RFC 1421)](https://datatracker.ietf.org/doc/html/rfc1421) format. **NOTE**: the [underlying](https://pkg.go.dev/encoding/pem#Encode) [libraries](https://pkg.go.dev/golang.org/x/crypto/ssh#MarshalAuthorizedKey) that generate this value append a `\\n` at the end of the PEM. In case this disrupts your use case, we recommend using `trimspace()`.
        """
        return pulumi.get(self, "cert_pem")

    @_builtins.property
    @pulumi.getter(name="isCa")
    def is_ca(self) -> _builtins.bool:
        """
        `true` if the certificate is of a CA (Certificate Authority).
        """
        return pulumi.get(self, "is_ca")

    @_builtins.property
    @pulumi.getter
    def issuer(self) -> _builtins.str:
        """
        Who verified and signed the certificate, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        """
        return pulumi.get(self, "issuer")

    @_builtins.property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> _builtins.str:
        """
        The time until which the certificate is invalid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "not_after")

    @_builtins.property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> _builtins.str:
        """
        The time after which the certificate is valid, as an [RFC3339](https://tools.ietf.org/html/rfc3339) timestamp.
        """
        return pulumi.get(self, "not_before")

    @_builtins.property
    @pulumi.getter(name="publicKeyAlgorithm")
    def public_key_algorithm(self) -> _builtins.str:
        """
        The key algorithm used to create the certificate.
        """
        return pulumi.get(self, "public_key_algorithm")

    @_builtins.property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> _builtins.str:
        """
        Number that uniquely identifies the certificate with the CA's system.
        The `format` function can be used to convert this *base 10* number into other bases, such as hex.
        """
        return pulumi.get(self, "serial_number")

    @_builtins.property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> _builtins.str:
        """
        The SHA1 fingerprint of the public key of the certificate.
        """
        return pulumi.get(self, "sha1_fingerprint")

    @_builtins.property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> _builtins.str:
        """
        The algorithm used to sign the certificate.
        """
        return pulumi.get(self, "signature_algorithm")

    @_builtins.property
    @pulumi.getter
    def subject(self) -> _builtins.str:
        """
        The entity the certificate belongs to, roughly following [RFC2253](https://tools.ietf.org/html/rfc2253).
        """
        return pulumi.get(self, "subject")

    @_builtins.property
    @pulumi.getter
    def version(self) -> _builtins.int:
        """
        The version the certificate is in.
        """
        return pulumi.get(self, "version")


