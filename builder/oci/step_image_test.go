// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oci

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepImage(t *testing.T) {
	state := testState()
	state.Put("instance_id", "ocid1...")

	step := new(stepImage)
	defer step.Cleanup(state)

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("image"); !ok {
		t.Fatalf("should have image")
	}
}

func TestStepImage_CreateImageErr(t *testing.T) {
	state := testState()
	state.Put("instance_id", "ocid1...")

	step := new(stepImage)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)
	driver.CreateImageErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}

	if _, ok := state.GetOk("image"); ok {
		t.Fatalf("should NOT have image")
	}
}

func TestStepImage_WaitForImageCreationErr(t *testing.T) {
	state := testState()
	state.Put("instance_id", "ocid1...")

	step := new(stepImage)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)
	driver.WaitForImageCreationErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}

	if _, ok := state.GetOk("image"); ok {
		t.Fatalf("should not have image")
	}
}
