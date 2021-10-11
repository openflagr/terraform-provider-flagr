Flagr Terraform Provider
==================
<p align="center" style="display: flex;justify-content: center; align-items: center; height: 200px;">
    <img src="https://avatars.githubusercontent.com/u/49816112?s=400&v=4" height="100px">        <>
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" height="100px">
</p>

Welcome to the Flagr Terraform provider! With this provider you can create and manage your flags on [flagr](https://github.com/openflagr/flagr)!

To view the full documentation of this provider, we recommend checking the [Terraform Registry](https://registry.terraform.io/providers/marceloboeira/flagr/latest) - Coming soon!

Usage
-----

```hcl
terraform {
  required_providers {
    flagr = {
      source  = "openflagr/flagr"
      version = "1.0.0"
    }
  }
}

provider "flagr" {
  ## Flagr Host
  host = "http://flagr.yourdomain.com:18000"
  ## Flagr Path - Optional, in case your flagr API runs on a custom path
  # path = "/api/v1"
}

# Example key for a release of a new pricing algorithm
resource "flagr_flag" "pricing-algorithm-v2" {
  description = "Price Algorithm V2"

  enabled              = true
  data_records_enabled = true

  variant {
    key = "pricing-v1"
    attachment = "pricing:v1"
  }

  variant {
    key = "pricing-v2"
    attachment = "pricing:v2"
  }
}
```

Releases
---------

Coming soon!


### Compatibility Matrix

|   Target                | Provider Version   |
|------------------------:|:------------------:|
|                         |       0.0.0        |
|  openflagr/flagr:1.1.13 | :white_check_mark: |
|  openflagr/flagr:1.1.12 | :white_check_mark: |
|     checkr/flagr:1.1.12 | :white_check_mark: |

Maintainers
-----------

This provider plugin is maintained by the [OpenFlagr team](https://github.com/orgs/openflagr/people).

Contributing
------------

Read our [contributors](https://github.com/marceloboeira/terraform-provider-flagr/docs/CONTRIBUTING.md) guide for more info on contributing.
