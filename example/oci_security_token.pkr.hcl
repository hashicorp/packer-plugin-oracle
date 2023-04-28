# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

source "oracle-oci" "security_token_example" {
  access_cfg_file         = "<config-file-path>"
  access_cfg_file_account = "<profile-name>"
  availability_domain     = "aaaa:PHX-AD-1"
  base_image_ocid         = "ocid1.image.oc1.phx.aaaaaaaa5yu6pw3riqtuhxzov7fdngi4tsteganmao54nq3pyxu3hxcuzmoa"
  compartment_ocid        = "ocid1.compartment.oc1..aaa"
  image_name              = "SecurityTokenExampleImage"
  shape                   = "VM.Standard2.1"
  ssh_username            = "opc"
  subnet_ocid             = "ocid1.subnet.oc1..aaa"
}

build {
  sources = ["source.oracle-oci.security_token_example"]
}
