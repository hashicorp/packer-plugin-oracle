
source "oracle-oci" "base-image-filtered" {
  ssh_username        = "opc"
  availability_domain = "aaaa:PHX-AD-1"
  base_image_filter = {
    display_name_search      = "^Oracle-Linux-7\\.8-2020\\.\\d+"
    operating_system         = "Oracle Linux"
    operating_system_version = "7.8"
  }
}

build {
  sources = ["source.oracle-oci.base-image-filtered"]

}
