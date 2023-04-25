# Oracle Plugin

This plugin allows Packer to communicate with the Oracle cloud platform.
It is able to create custom images on both Oracle Cloud Infrastructure and
Oracle Cloud Infrastructure Classic Compute. This plugin comes with builders
designed to support both platforms.

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    oracle = {
      version = ">= 1.0.1"
      source  = "github.com/hashicorp/oracle"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/hashicorp/packer-plugin-oracle/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `go build` from the root
directory. Upon successful compilation, a `packer-plugin-oracle` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


## Plugin Contents

### Builders

- [oracle-classic](./builders/classic.mdx) - Create custom images in Oracle Cloud Infrastructure
    Classic Compute by launching a source instance and creating an image list
    from a snapshot of it after provisioning.

- [oracle-oci](./builders/oci.mdx) - Create custom images in Oracle Cloud Infrastructure (OCI) by
    launching a base instance and creating an image from it after provisioning.
