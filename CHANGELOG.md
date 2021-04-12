CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

## 3.4.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 3.3.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 3.3.0 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 3.2.0 (2021-02-22)
* Upgrade to v3.1.0 of the TLS Terraform Provider

## 3.1.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider

## 3.1.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.0

## 3.0.0 (2020-12-22)
* Upgrade to v3.0.0 of the TLS Terraform Provider

## 2.5.0 (2020-12-16)
* Upgrade to v2.16.0 of pulumi-terraform-bridge
    * Preserve unknowns during provider preview
* Upgrade NodeJS and Python versions to use Pulumi >= v2.15.0 

## 2.4.2 (2020-11-24)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.4.1 (2020-11-09)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.4.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.3.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.2.0 (2020-07-24)
* Upgrade to v2.2.0 of the TLS Terraform Provider

## 2.1.3 (2020-06-11)
* Switch to GitHub actions for build

## 2.1.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.1.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-28)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.6.0 (2020-03-31)
* Convert to go modules
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge 1.8.4

## 1.5.0 (2020-03-16)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.4.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.3.0 (2020-01-06)
* Upgrade to pulumi-terraform-bridge v1.5.2

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.1.0 (2019-11-12)
* Add a **preview** .NET SDK

## 1.0.0 (2019-10-09)
* Upgrade to v2.1.1 of the TLS Terraform Provider

## 0.18.5 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.18.4 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.18.3 (2019-08-20)
* Depend on latest pulumi package

## 0.18.2 (2019-08-09)
* Upgrade to pulumi-terraform@9db2fc93cd

## 0.18.1 (2019-08-08)
* Update to pulumi-terraform@013b95b1c8


## 0.18.0 (2019-07-10)
* Initial release based on v2.0.1 of the TLS Terraform Provider
