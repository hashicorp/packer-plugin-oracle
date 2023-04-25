# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Oracle"
  description = "The Oracle multi-component plugin can be used with HashiCorp Packer to create custom images."
  identifier = "packer/BrandonRomano/oracle"
  flags = ["hcp-ready"]
  component {
    type = "builder"
    name = "Oracle OCI"
    slug = "oci"
  }
  component {
    type = "builder"
    name = "Oracle Cloud Infrastructure Classic"
    slug = "classic"
  }
}
