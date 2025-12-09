// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package classic

import (
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/packer/registry/image"
)

// Artifact is an artifact implementation that contains Image List
// and Machine Image info.
type Artifact struct {
	APIEndpoint      string
	SourceImageList  string
	MachineImageName string
	MachineImageFile string
	ImageListVersion int

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

func (a *Artifact) Id() string {
	return a.MachineImageName
}

func (a *Artifact) String() string {
	return fmt.Sprintf("An image list entry was created: \n"+
		"Name: %s\n"+
		"File: %s\n"+
		"Version: %d",
		a.MachineImageName, a.MachineImageFile, a.ImageListVersion)
}

func (a *Artifact) State(name string) interface{} {
	if name == image.ArtifactStateURI {
		return a.buildHCPackerRegistryMetadata()
	}

	return a.StateData[name]
}

// Destroy deletes the custom image associated with the artifact.
func (a *Artifact) Destroy() error {
	return nil
}

func (a *Artifact) buildHCPackerRegistryMetadata() interface{} {

	img, err := image.FromArtifact(a, image.WithRegion(a.APIEndpoint), image.WithSourceID(a.SourceImageList))

	if err != nil {
		log.Printf("[TRACE] error encountered when creating HCP Packer registry image for artifact: %s", err)
		return nil
	}

	return img
}
