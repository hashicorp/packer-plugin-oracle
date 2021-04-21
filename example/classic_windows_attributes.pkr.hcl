
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

source "oracle-classic" "windows_attrs" {
  api_endpoint      = "${var.opc_api_endpoint}"
  attributes_file   = "./windows_attributes.json"
  communicator      = "winrm"
  dest_image_list   = "Packer_Windows_Demo"
  identity_domain   = "${var.opc_identity_domain}"
  image_name        = "Packer_Windows_Demo_${local.timestamp}"
  password          = "${var.opc_password}"
  shape             = "oc3"
  source_image_list = "/Compute-${var.opc_identity_domain}/${var.opc_username}/Microsoft_Windows_Server_2012_R2-17.3.6-20170930-124649"
  username          = "${var.opc_username}"
  winrm_password    = "password"
  winrm_username    = "Administrator"
}

build {
  sources = ["source.oracle-classic.windows_attrs"]

  provisioner "powershell" {
    inline = "Write-Output(\"HELLO WORLD\")"
  }

}
