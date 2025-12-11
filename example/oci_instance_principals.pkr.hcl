# Copyright IBM Corp. 2013, 2025
# SPDX-License-Identifier: MPL-2.0

source "oracle-oci" "instance_principal_example" {
  availability_domain     = "aaaa:PHX-AD-1"
  base_image_ocid         = "ocid1.image.oc1.phx.aaaaaaaa5yu6pw3riqtuhxzov7fdngi4tsteganmao54nq3pyxu3hxcuzmoa"
  compartment_ocid        = "ocid1.compartment.oc1..aaa"
  image_name              = "InstancePrincipalExampleImage"
  shape                   = "VM.Standard2.1"
  ssh_username            = "opc"
  subnet_ocid             = "ocid1.subnet.oc1..aaa"
  use_instance_principals = "true"
  ssh_username            = "opc"
}

build {
  sources = ["source.oracle-oci.instance_principal_example"]
}
