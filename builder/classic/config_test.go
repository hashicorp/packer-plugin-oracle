// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package classic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"identity_domain":   "abc12345",
		"username":          "test@hashicorp.com",
		"password":          "testpassword123",
		"api_endpoint":      "https://api-test.compute.test.oraclecloud.com/",
		"dest_image_list":   "/Config-thing/myuser/myimage",
		"source_image_list": "/oracle/public/whatever",
		"shape":             "oc3",
		"image_name":        "TestImageName",
		"ssh_username":      "opc",
	}
}

func TestConfigAutoFillsSourceList(t *testing.T) {
	tc := testConfig()
	var conf Config
	err := conf.Prepare(tc)
	if err != nil {
		t.Fatalf("Should not have error: %s", err.Error())
	}
	if conf.SSHSourceList != "seciplist:/oracle/public/public-internet" {
		t.Fatalf("conf.SSHSourceList should have been "+
			"\"seciplist:/oracle/public/public-internet\" but is \"%s\"",
			conf.SSHSourceList)
	}
}

func TestConfigValidationCatchesMissing(t *testing.T) {
	required := []string{
		"username",
		"password",
		"api_endpoint",
		"identity_domain",
		"dest_image_list",
		"source_image_list",
		"shape",
		"ssh_username",
	}
	for _, key := range required {
		tc := testConfig()
		delete(tc, key)
		var c Config
		err := c.Prepare(tc)
		if err == nil {
			t.Fatalf("Test should have failed when config lacked %s!", key)
		}
	}
}

func TestConfigValidatesObjects(t *testing.T) {
	var objectTests = []struct {
		object string
		valid  bool
	}{
		{"foo-BAR.0_9", true},
		{"%", false},
		{"Matt...?", false},
		{"/Config-thing/myuser/myimage", true},
	}
	for _, s := range []string{"dest_image_list", "image_name"} {
		for _, tt := range objectTests {
			tc := testConfig()
			tc[s] = tt.object
			var c Config
			err := c.Prepare(tc)
			if tt.valid {
				assert.NoError(t, err, tt.object)
			} else {
				assert.Error(t, err, tt.object)
			}
		}
	}
}
