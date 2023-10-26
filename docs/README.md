This plugin allows Packer to communicate with the Oracle cloud platform.
It is able to create custom images on both Oracle Cloud Infrastructure and
Oracle Cloud Infrastructure Classic Compute. This plugin comes with builders
designed to support both platforms.

## Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    oracle = {
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

- [oracle-classic](/packer/integrations/hashicorp/oracle/latest/components/builder/oci) - Create custom images in Oracle Cloud Infrastructure
    Classic Compute by launching a source instance and creating an image list from a snapshot of it after provisioning.

- [oracle-oci](/packer/integrations/hashicorp/oracle/latest/components/builder/classic) - Create custom images in Oracle Cloud Infrastructure (OCI) by
    launching a base instance and creating an image from it after provisioning.

## Oracle Classic Authentication

This builder authenticates API calls to Oracle Cloud Infrastructure Classic
Compute using basic authentication (user name and password). To read more, see
the [authentication
documentation](https://docs.oracle.com/en/cloud/iaas/compute-iaas-cloud/stcsa/Authentication.html)



## Oracle OCI Authentication

There are three authentication methods available for the OCI builder. The API
Signing Key of the `DEFAULT` profile in the OCI
 [SDK and CLI Configuration File](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm#SDK_and_CLI_Configuration_File)
 will be used by default.

### Policy Reference

When running the `oracle-oci` builder as a non-administrative user, you will need to have OCI policies defined which will enable Packer to do its job.  

The following represent minimal required policies for the plugin to operate:

```
Allow group PackerGroup to manage instance-family in compartment ${COMPARTMENT_NAME}
Allow group PackerGroup to manage instance-images in compartment ${COMPARTMENT_NAME}
Allow group PackerGroup to use virtual-network-family in compartment ${COMPARTMENT_NAME}
Allow group PackerGroup to use compute-image-capability-schema in tenancy
```

This example assumes the user running Packer is in a group named 'PackerGroup'.  You will need to update ${COMPARTMENT_NAME} to the name of the
compartment in which the plugin is managing resources.

For more details on working with OCI Policies, please refer to the [OCI Policy Documentation](https://docs.oracle.com/en-us/iaas/Content/Identity/policiesgs/get-started-with-policies.htm)

### API Signing Key

The OCI REST API requires that requests be signed with the RSA public key
associated with your
[IAM](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm)
user account. For a comprehensive example of how to configure the required
authentication see the documentation on [Required Keys and
OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm)
([Oracle Cloud
IDs](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm)).

### Instance Principal

If you run Packer on an OCI compute instance, you can configure Packer to use the
[Instance Principal](https://docs.cloud.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm)
associated with that instance to authenticate instead of an API Signing Key that
is associated with a specific user. This method requires the creation of appropriately configured
[Dynamic Groups](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingdynamicgroups.htm)
and [Policies](https://docs.oracle.com/en-us/iaas/Content/Identity/Concepts/policygetstarted.htm).

The [`use_instance_principals`](https://www.packer.io/docs/builders/oracle/oci#use_instance_principals)
parameter is used to enable this method.

The [`oci_instance_principals.pkr.hcl`](/example/oci_instance_principals.pkr.hcl)
and [`oci_instance_principals.json`](/example/oci_instance_principals.json) examples can be used as a reference for
configuring this authentication method.

### Security Token

[Token-based authentication](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/clitoken.htm)
requires you to interactively authenticate a session using the OCI CLI and a web
browser. The token that is created has an extremely limited TTL and must be
refreshed every hour for up to 24 hours. However, this method allows for
authentication using non-SCIM supported federated identity providers.

To use token-based authentication, start a new session using the CLI's
[`oci session authenticate`](https://docs.oracle.com/en-us/iaas/Content/API/SDKDocs/clitoken.htm#Starting)
command and provide the path to the configuration file as the `access_cfg_file`
configuration parameter and the name of the profile as the `access_cfg_file_account`
configuration parameter.

The [`oci_security_token.pkr.hcl`](/example/oci_security_token.pkr.hcl) and
[`oci_security_token.json`](/example/oci_security_token.json) examples can be
used as a reference for this method.
