// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oci

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestStepCreateInstance(t *testing.T) {
	state := testState()
	state.Put("publicKey", "key")

	step := new(stepCreateInstance)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	instanceIDRaw, ok := state.GetOk("instance_id")
	if !ok {
		t.Fatalf("should have machine")
	}

	step.Cleanup(state)

	if driver.TerminateInstanceID != instanceIDRaw.(string) {
		t.Fatalf(
			"should've deleted instance (%s != %s)",
			driver.TerminateInstanceID, instanceIDRaw.(string))
	}
}

func TestStepCreateInstance_InstanceOptions(t *testing.T) {
	runTest := func(t *testing.T, value *bool, expected *bool) {
		state := testState() // testState already calls Prepare on a base config
		state.Put("publicKey", "key")

		config := state.Get("config").(*Config)
		config.InstanceOptionsAreLegacyImdsEndpointsDisabled = value
		step := new(stepCreateInstance)
		defer func() {
			runErr := state.Get("error")
			step.Cleanup(state)
			if runErr == nil {
				if _, ok := state.GetOk("error"); ok {
					t.Logf("Warning: Cleanup reported an error: %v", state.Get("error"))
				}
			}
		}()

		driver := state.Get("driver").(*driverMock)
		// driver.cfg is already pointing to the config from the state, which we've modified and re-prepared.

		if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
			if err, ok := state.GetOk("error"); ok {
				t.Fatalf("bad action: %#v with error: %v. Expected ActionContinue", action, err)
			}
			t.Fatalf("bad action: %#v. Expected ActionContinue", action)
		}

		if expected == nil {
			if driver.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled != nil {
				t.Errorf("Expected CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled to be nil, got %v", *driver.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled)
			}
		} else {
			if driver.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled == nil {
				t.Errorf("Expected CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled to be %v, got nil", *expected)
			} else if *driver.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled != *expected {
				t.Errorf("Expected CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled to be %v, got %v", *expected, *driver.CapturedInstanceOptionsAreLegacyImdsEndpointsDisabled)
			}
		}
	}

	t.Run("True", func(t *testing.T) {
		bTrue := true
		runTest(t, &bTrue, &bTrue)
	})

	t.Run("False", func(t *testing.T) {
		bFalse := false
		runTest(t, &bFalse, &bFalse)
	})

	t.Run("Nil", func(t *testing.T) {
		runTest(t, nil, nil)
	})
}

func TestStepCreateInstance_CreateInstanceErr(t *testing.T) {
	state := testState()
	state.Put("publicKey", "key")

	step := new(stepCreateInstance)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)
	driver.CreateInstanceErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}

	if _, ok := state.GetOk("instance_id"); ok {
		t.Fatalf("should NOT have instance_id")
	}

	step.Cleanup(state)

	if driver.TerminateInstanceID != "" {
		t.Fatalf("Should not have tried to terminate an instance")
	}
}

func TestStepCreateInstance_WaitForInstanceStateErr(t *testing.T) {
	state := testState()
	state.Put("publicKey", "key")

	step := new(stepCreateInstance)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)
	driver.WaitForInstanceStateErr = errors.New("error")

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("bad action: %#v", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}
}

func TestStepCreateInstance_TerminateInstanceErr(t *testing.T) {
	state := testState()
	state.Put("publicKey", "key")

	step := new(stepCreateInstance)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	_, ok := state.GetOk("instance_id")
	if !ok {
		t.Fatalf("should have machine")
	}

	driver.TerminateInstanceErr = errors.New("error")
	step.Cleanup(state)

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}
}

func TestStepCreateInstanceCleanup_WaitForInstanceStateErr(t *testing.T) {
	state := testState()
	state.Put("publicKey", "key")

	step := new(stepCreateInstance)
	defer step.Cleanup(state)

	driver := state.Get("driver").(*driverMock)

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("bad action: %#v", action)
	}

	driver.WaitForInstanceStateErr = errors.New("error")
	step.Cleanup(state)

	if _, ok := state.GetOk("error"); !ok {
		t.Fatalf("should have error")
	}
}
