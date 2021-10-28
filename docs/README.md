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

⚠️  The usage represents, as of now, the desired state, not the current state. Below, you are able to see supported features in the feature matrix.

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
    key = "pricing:v1"
    attachment = "pricing:v1"
  }

  variant {
    key = "pricing:v2"
    attachment = "pricing:v2"
  }

  segments {
    description = "All traffic"
    rollout = 100

    distribution = {
        "pricing:v1" = 90
        "pricing:v2" = 10
    }
  }
}
```

The example above represents a release of a pricing algorithm, it will split the traffic in 2 parts, 90% gets the regular/existing pricing variant (v1) and 10% gets the new pricing (v2).

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


Features
---------

For more insights into what each entity means, please reffer to [openflagr-docs](https://openflagr.github.io/flagr/#/flagr_overview).

| Terraform     |        Entity | Feature           | Status   |
|--------------:|--------------:|------------------:|:--------:|
| Provider      |               | URL               |  ✅      |
|               |               | Path              |  ✅      |
|               |               | Authentication    |  ⚪️      |
| Data Source   | Flags         | ReadAll           |  ✅      |
|               |               | Search            |  ⛔️      |
|               | Flag          | Read (id)         |  ✅      |
|               |               | Search            |  ⛔️      |
|               | Variants      | ReadAll (flag-id) |  ⛔️      |
|               | Variant       | Read (flag-id/id) |  ⛔️      |
|               | Segments      | ReadAll (flag-id) |  ⛔️      |
|               | Segment       | Read (flag-id/id) |  ⛔️      |
| Resource      | Flag          | Read              |  ✅      |
|               |               | Import            |  ✅      |
|               |               | Create            |  ✅      |
|               |               | Update            |  ✅      |
|               |               | Delete            |  ✅      |
| Resource      | Variant       | Read              |  ⚪️      |
|               |               | Import            |  ⚪️      |
|               |               | Create            |  ⚪️      |
|               |               | Update            |  ⚪️      |
|               |               | Delete            |  ⚪️      |
| Resource      | Segment       | Read              |  ⚪️      |
|               |               | Import            |  ⚪️      |
|               |               | Create            |  ⚪️      |
|               |               | Update            |  ⚪️      |
|               |               | Delete            |  ⚪️      |

* ✅ - Supported
* ⚠️  - Partially supported / not stable
* ⚪️ - Not Supported but part of the roadmap
* ⛔️ - Not Supported & no plans for now

Maintainers
-----------

This provider plugin is maintained by the [OpenFlagr team](https://github.com/orgs/openflagr/people).

Contributing
------------

Read our [contributors](https://github.com/marceloboeira/terraform-provider-flagr/docs/CONTRIBUTING.md) guide for more info on contributing.
