# The "legacy_isotime" function has been provided for backwards compatability,
# but we recommend switching to the timestamp and formatdate functions.

locals {
  timestamp = "${legacy_isotime("20060102030405")}"
}

source "oracle-oci" "base-image-example" {
  availability_domain = "aaaa:PHX-AD-1"
  base_image_ocid     = "ocid1.image.oc1.iad.aaa"
  compartment_ocid    = "ocid1.compartment.oc1..aaa"
  create_vnic_details {
    assign_public_ip = "false"
    display_name     = "testing-123"
    nsg_ids          = ["ocid1.networksecuritygroup.oc1.iad.aaa"]
  }
  image_name = "my-image-${local.timestamp}"
  instance_defined_tags {
    Operations = {
      Environment = "prod"
      Team        = "CostCenter"
    }
  }
  instance_name = "packer-build-${local.timestamp}"
  instance_tags = {
    testing = "yes"
  }
  shape       = "VM.Standard.E2.1"
  subnet_ocid = "ocid1.subnet.oc1.iad.aaa"
  tags = {
    CreationDate = "${legacy_isotime("20060102 03:04:05 MST")}"
  }
  use_private_ip = "true"
}

build {
  sources = ["source.oracle-oci.base-image-example"]

}
