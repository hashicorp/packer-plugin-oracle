package oci

import (
	"reflect"
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/packer/registry/image"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func TestArtifactImpl(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packersdk.Artifact); !ok {
		t.Fatalf("Artifact should be artifact")
	}
}

func TestArtifactState_StateData(t *testing.T) {
	expectedData := "this is the data"
	artifact := &Artifact{
		StateData: map[string]interface{}{"state_data": expectedData},
	}

	// Valid state
	result := artifact.State("state_data")
	if result != expectedData {
		t.Fatalf("Bad: State data was %s instead of %s", result, expectedData)
	}

	// Invalid state
	result = artifact.State("invalid_key")
	if result != nil {
		t.Fatalf("Bad: State should be nil for invalid state data name")
	}

	// Nil StateData should not fail and should return nil
	artifact = &Artifact{}
	result = artifact.State("key")
	if result != nil {
		t.Fatalf("Bad: State should be nil for nil StateData")
	}
}

func TestArtifactState_hcpPackerRegistryMetadata(t *testing.T) {
	artifact := &Artifact{
		Image: core.Image{
			CompartmentId:          stringPtr("ocid1.compartment.oc1..aaa"),
			Id:                     stringPtr("ocid1.image.oc1.phx.aaa"),
			OperatingSystem:        stringPtr("Oracle Linux"),
			OperatingSystemVersion: stringPtr("7.2"),
			BaseImageId:            stringPtr("ocid1.image.oc1.phx.aaabase"),
			LaunchMode:             core.ImageLaunchModeParavirtualized,
			BillableSizeInGBs:      int64Ptr(10),
		},
		Region: "us-phoenix-1",
	}

	result := artifact.State(image.ArtifactStateURI)

	expected := &image.Image{
		ImageID:        "ocid1.image.oc1.phx.aaa",
		ProviderName:   BuilderId,
		ProviderRegion: "us-phoenix-1",
		Labels: map[string]string{
			"billable_size_in_gbs":     "10",
			"compartment_id":           "ocid1.compartment.oc1..aaa",
			"launch_mode":              string(core.ImageLaunchModeParavirtualized),
			"operating_system":         "Oracle Linux",
			"operating_system_version": "7.2",
		},
		SourceImageID: "ocid1.image.oc1.phx.aaabase",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Bad: HCP registry metadata was %#v instead of %#v", result, expected)
	}
}

func stringPtr(s string) *string {
	return &s
}

func int64Ptr(int64 int64) *int64 {
	return &int64
}
