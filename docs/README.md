This plugin allows Packer to communicate with the Oracle cloud platform.
It is able to create custom images on both Oracle Cloud Infrastructure and
Oracle Cloud Infrastructure Classic Compute. This plugin comes with builders
designed to support both platforms.

## Installation

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    amazon = {
      source  = "github.com/hashicorp/oracle"
      version = "~> 1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/oracle
```

## Components

### Builders

- [oracle-classic](./builders/classic.mdx) - Create custom images in Oracle Cloud Infrastructure
    Classic Compute by launching a source instance and creating an image list
    from a snapshot of it after provisioning.

- [oracle-oci](./builders/oci.mdx) - Create custom images in Oracle Cloud Infrastructure (OCI) by
    launching a base instance and creating an image from it after provisioning.
