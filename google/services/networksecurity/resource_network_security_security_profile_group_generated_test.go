// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package networksecurity_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(context),
			},
			{
				ResourceName:            "google_network_security_security_profile_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "parent", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccNetworkSecuritySecurityProfileGroup_networkSecuritySecurityProfileGroupBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_network_security_security_profile_group" "default" {
  name                      = "tf-test-sec-profile-group%{random_suffix}"
  parent                    = "organizations/%{org_id}"
  description               = "my description"
  threat_prevention_profile = google_network_security_security_profile.security_profile.id

  labels = {
    foo = "bar"
  }
}

resource "google_network_security_security_profile" "security_profile" {
    name        = "tf-test-sec-profile%{random_suffix}"
    type        = "THREAT_PREVENTION"
    parent      = "organizations/%{org_id}"
    location    = "global"
}
`, context)
}

func testAccCheckNetworkSecuritySecurityProfileGroupDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_security_security_profile_group" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{NetworkSecurityBasePath}}{{parent}}/locations/{{location}}/securityProfileGroups/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("NetworkSecuritySecurityProfileGroup still exists at %s", url)
			}
		}

		return nil
	}
}
