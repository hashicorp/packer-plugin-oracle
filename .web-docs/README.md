This plugin allows Packer to communicate with the Oracle cloud platform.
It is able to create custom images on both Oracle Cloud Infrastructure and
Oracle Cloud Infrastructure Classic Compute. This plugin comes with builders
designed to support both platforms.

## Builders

- [oracle-classic](./builders/classic.mdx) - Create custom images in Oracle Cloud Infrastructure
    Classic Compute by launching a source instance and creating an image list
    from a snapshot of it after provisioning.

- [oracle-oci](./builders/oci.mdx) - Create custom images in Oracle Cloud Infrastructure (OCI) by
    launching a base instance and creating an image from it after provisioning.
