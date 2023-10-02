# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Oracle"
  description = "The Oracle multi-component plugin can be used with HashiCorp Packer to create custom images."
  identifier = "packer/hashicorp/oracle"
  flags = ["hcp-ready"]
  component {
    type = "builder"
    name = "Oracle Cloud Infrastructure"
    slug = "oci"
  }
  component {
    type = "builder"
    name = "Oracle Cloud Infrastructure Classic Compute"
    slug = "classic"
  }
}
