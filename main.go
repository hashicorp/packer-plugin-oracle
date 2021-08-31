package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	classicbuilder "github.com/hashicorp/packer-plugin-oracle/builder/classic"
	ocibuilder "github.com/hashicorp/packer-plugin-oracle/builder/oci"
	"github.com/hashicorp/packer-plugin-oracle/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("classic", new(classicbuilder.Builder))
	pps.RegisterBuilder("oci", new(ocibuilder.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
