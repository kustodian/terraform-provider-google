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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexTaskIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexTaskIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_task_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-task%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccDataplexTaskIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_task_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-task%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexTaskIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccDataplexTaskIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_dataplex_task_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-task%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataplexTaskIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project_name":  envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexTaskIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_dataplex_task_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_dataplex_task_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-task%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDataplexTaskIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_dataplex_task_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/lakes/%s/tasks/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-lake%s", context["random_suffix"]), fmt.Sprintf("tf-test-task%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDataplexTaskIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}

resource "google_dataplex_task_iam_member" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccDataplexTaskIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_dataplex_task_iam_policy" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_dataplex_task_iam_policy" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  depends_on = [
    google_dataplex_task_iam_policy.foo
  ]
}
`, context)
}

func testAccDataplexTaskIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}

data "google_iam_policy" "foo" {
}

resource "google_dataplex_task_iam_policy" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccDataplexTaskIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}

resource "google_dataplex_task_iam_binding" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccDataplexTaskIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "%{project_name}"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "%{project_name}"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "%{project_name}"
    
}

resource "google_dataplex_task_iam_binding" "foo" {
  project = google_dataplex_task.example.project
  location = google_dataplex_task.example.location
  lake = google_dataplex_task.example.lake
  task_id = google_dataplex_task.example.task_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
