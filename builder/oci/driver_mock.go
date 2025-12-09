// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package oci

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/core"
)

// driverMock implements the Driver interface and communicates with Oracle
// OCI.
type driverMock struct {
	CreateInstanceID  string
	CreateInstanceErr error

	CreateImageID  string
	CreateImageErr error

	UpdateSchemaID  string
	UpdateSchemaErr error

	DeleteImageID  string
	DeleteImageErr error

	GetInstanceIPErr error

	TerminateInstanceID  string
	TerminateInstanceErr error

	WaitForImageCreationErr error

	WaitForInstanceStateErr error

	cfg                                                   *Config
	CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled *bool
}

// CreateInstance creates a new compute instance.
func (d *driverMock) CreateInstance(ctx context.Context, publicKey string) (string, error) {
	if d.CreateInstanceErr != nil {
		return "", d.CreateInstanceErr
	}

	d.CreateInstanceID = "ocid1..."
	if d.cfg != nil {
		// Capture the value from the Config struct that the step is expected to use.
		// This assumes that if cfg.InstanceOptionsAreLegacyImdsEndpointsDisabled is set,
		// stepCreateInstance would correctly prepare it for the actual launch details.
		d.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled = d.cfg.InstanceOptionsAreLegacyImdsEndpointsDisabled
	}

	return d.CreateInstanceID, nil
}

// CreateImage creates a new custom image.
func (d *driverMock) CreateImage(ctx context.Context, id string) (core.Image, error) {
	if d.CreateImageErr != nil {
		return core.Image{}, d.CreateImageErr
	}
	d.CreateImageID = id
	return core.Image{Id: &id}, nil
}

// CreateImage creates a new custom image.
func (d *driverMock) UpdateImageCapabilitySchema(ctx context.Context, imageId string) (core.UpdateComputeImageCapabilitySchemaResponse, error) {
	if d.UpdateSchemaErr != nil {
		return core.UpdateComputeImageCapabilitySchemaResponse{}, d.UpdateSchemaErr
	}
	d.UpdateSchemaID = imageId
	return core.UpdateComputeImageCapabilitySchemaResponse{}, nil
}

// DeleteImage mocks deleting a custom image.
func (d *driverMock) DeleteImage(ctx context.Context, id string) error {
	if d.DeleteImageErr != nil {
		return d.DeleteImageErr
	}

	d.DeleteImageID = id

	return nil
}

// GetInstanceIP returns the public or private IP corresponding to the given instance id.
func (d *driverMock) GetInstanceIP(ctx context.Context, id string) (string, error) {
	if d.GetInstanceIPErr != nil {
		return "", d.GetInstanceIPErr
	}
	if d.cfg.UsePrivateIP {
		return "private_ip", nil
	}
	return "ip", nil
}

// TerminateInstance terminates a compute instance.
func (d *driverMock) TerminateInstance(ctx context.Context, id string) error {
	if d.TerminateInstanceErr != nil {
		return d.TerminateInstanceErr
	}

	d.TerminateInstanceID = id

	return nil
}

// WaitForImageCreation waits for a provisioning custom image to reach the
// "AVAILABLE" state.
func (d *driverMock) WaitForImageCreation(ctx context.Context, id string) error {
	return d.WaitForImageCreationErr
}

// WaitForInstanceState waits for an instance to reach the a given terminal
// state.
func (d *driverMock) WaitForInstanceState(ctx context.Context, id string, waitStates []string, terminalState string) error {
	return d.WaitForInstanceStateErr
}
