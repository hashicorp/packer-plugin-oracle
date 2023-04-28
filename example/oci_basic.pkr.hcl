# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


source "oracle-oci" "example" {
  availability_domain = "aaaa:PHX-AD-1"
  base_image_ocid     = "ocid1.image.oc1.phx.aaaaaaaa5yu6pw3riqtuhxzov7fdngi4tsteganmao54nq3pyxu3hxcuzmoa"
  compartment_ocid    = "ocid1.compartment.oc1..aaa"
  image_name          = "ExampleImage"
  shape               = "VM.Standard1.1"
  ssh_username        = "opc"
  subnet_ocid         = "ocid1.subnet.oc1..aaa"
}

build {
  sources = ["source.oracle-oci.example"]

}
