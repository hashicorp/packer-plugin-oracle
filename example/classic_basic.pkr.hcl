
locals { timestamp = regex_replace(timestamp(), "[- TZ:]", "") }

source "oracle-classic" "basic" {
  api_endpoint      = "https://api-###.compute.###.oraclecloud.com/"
  attributes        = "{\"userdata\": {\"pre-bootstrap\": {\"script\": [\"...\"]}}}"
  dest_image_list   = "Packer_Builder_Test_List"
  identity_domain   = "#######"
  image_name        = "Packer_Builder_Test_${local.timestamp}"
  password          = "supersecretpasswordhere"
  shape             = "oc3"
  source_image_list = "/oracle/public/OL_7.2_UEKR4_x86_64"
  username          = "myuser@myaccount.com"
  ssh_username      = "opc"
}

build {
  sources = ["source.oracle-classic.basic"]

  provisioner "shell" {
    inline = ["echo hello"]
  }

}
