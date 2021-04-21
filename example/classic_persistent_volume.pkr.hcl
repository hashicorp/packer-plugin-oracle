
variable "opc_api_endpoint" {
  type    = string
  default = "${env("OPC_ENDPOINT")}"
}

variable "opc_identity_domain" {
  type    = string
  default = "${env("OPC_IDENTITY_DOMAIN")}"
}

variable "opc_password" {
  type    = string
  default = "${env("OPC_PASSWORD")}"
}

variable "opc_username" {
  type    = string
  default = "${env("OPC_USERNAME")}"
}

locals { timestamp = regex_replace(timestamp(), "[- TZ:]", "") }

source "oracle-classic" "persistent" {
  api_endpoint           = "${var.opc_api_endpoint}"
  dest_image_list        = "Packer_Builder_Test_List"
  identity_domain        = "${var.opc_identity_domain}"
  image_name             = "Packer_Builder_Test_${local.timestamp}"
  password               = "${var.opc_password}"
  persistent_volume_size = 15
  shape                  = "oc3"
  source_image_list      = "/oracle/public/OL_7.2_UEKR4_x86_64"
  ssh_username           = "opc"
  username               = "${var.opc_username}"
}

build {
  sources = ["source.oracle-classic.persistent"]

  provisioner "shell" {
    inline = ["echo hello"]
  }

}
