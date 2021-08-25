# CHANGELOG

## 1.0.1 (August 25, 2021)

* Added support for [token-based authentication](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/clitoken.htm)

## 1.0.0 (June 14, 2021)

The code base for this plugin has been stable since the Packer core split.
We are marking this plugin as v1.0.0 to indicate that it is stable and ready for consumption via `packer init`.

* Update packer-plugin-sdk to v0.2.3
* Add `instance_defined_tags_json` and `defined_tags_json` as HCL2-only equivalent options
  to `instance_defined_tags` and `defined_tags` to properly support JSON tag-mappings in HCL templates. [GH-15]


## 0.0.3 (April 21, 2021)

* Oracle plugin break out from Packer core. Changes prior to break out can be found in [Packer's CHANGELOG](https://github.com/hashicorp/packer/blob/master/CHANGELOG.md)
