// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/packer-plugin-sdk/packer/registry/image"
	"github.com/oracle/oci-go-sdk/v65/core"
)

// Artifact is an artifact implementation that contains a built Custom Image.
type Artifact struct {
	Image  core.Image
	Region string
	driver Driver

	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

// BuilderId uniquely identifies the builder.
func (a *Artifact) BuilderId() string {
	return BuilderId
}

// Files lists the files associated with an artifact. We don't have any files
// as the custom image is stored server side.
func (a *Artifact) Files() []string {
	return nil
}

// Id returns the OCID of the associated Image.
func (a *Artifact) Id() string {
	return *a.Image.Id
}

func (a *Artifact) String() string {
	var displayName string
	if a.Image.DisplayName != nil {
		displayName = *a.Image.DisplayName
	}

	return fmt.Sprintf(
		"An image was created: '%v' (OCID: %v) in region '%v'",
		displayName, *a.Image.Id, a.Region,
	)
}

func (a *Artifact) State(name string) interface{} {
	if name == image.ArtifactStateURI {
		return a.buildHCPackerRegistryMetadata()
	}

	return a.StateData[name]
}

// Destroy deletes the custom image associated with the artifact.
func (a *Artifact) Destroy() error {
	return a.driver.DeleteImage(context.TODO(), *a.Image.Id)
}

func (a *Artifact) buildHCPackerRegistryMetadata() interface{} {

	labels := map[string]interface{}{}

	if a.Image.BillableSizeInGBs != nil {
		labels["billable_size_in_gbs"] = strconv.FormatInt(*a.Image.BillableSizeInGBs, 10)
	}

	if a.Image.CompartmentId != nil {
		labels["compartment_id"] = *a.Image.CompartmentId
	}

	if a.Image.LaunchMode != "" {
		labels["launch_mode"] = string(a.Image.LaunchMode)
	}

	if a.Image.OperatingSystem != nil {
		labels["operating_system"] = *a.Image.OperatingSystem
	}

	if a.Image.OperatingSystemVersion != nil {
		labels["operating_system_version"] = *a.Image.OperatingSystemVersion
	}

	img, err := image.FromArtifact(a, image.WithRegion(a.Region), image.WithSourceID(*a.Image.BaseImageId), image.SetLabels(labels))

	if err != nil {
		log.Printf("[TRACE] error encountered when creating HCP Packer registry image for artifact: %s", err)
		return nil
	}

	return img
}
