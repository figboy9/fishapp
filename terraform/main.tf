provider "google" {
  project = var.project
  region  = var.region
}

provider "google-beta" {
  project = var.project
  region  = var.region
}

terraform {
  backend "gcs" {
    bucket = "fishapp-282106-tf-state-prod"
    prefix = "terraform/state"
  }
}

locals {
  fishapp_enable_services = [
    "cloudresourcemanager.googleapis.com",
    "iam.googleapis.com",
    "compute.googleapis.com",
    "container.googleapis.com",
    "dns.googleapis.com",
    "cloudkms.googleapis.com",
  ]
}

resource "google_project_service" "fishapp_service" {
  for_each = toset(local.fishapp_enable_services)
  service  = each.value

  disable_dependent_services = true
}

resource "google_storage_bucket" "tf_state" {
  name               = "${var.project}-tf-state-prod"
  location           = "us-west1"
  storage_class      = "REGIONAL"
  bucket_policy_only = true

  versioning {
    enabled = true
  }

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      num_newer_versions = 5
    }
  }
}

resource "google_container_cluster" "fishapp_cluster" {
  provider = google-beta
  name     = "fishapp-cluster"

  location                 = var.location
  remove_default_node_pool = true
  initial_node_count       = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }

  release_channel {
    channel = "RAPID"
  }

  # config connectorの設定
  addons_config {
    config_connector_config {
      enabled = true
    }
  }

  workload_identity_config {
    identity_namespace = "${var.project}.svc.id.goog"
  }
}
# 公式通りnode_configはnodeに書く
resource "google_container_node_pool" "fishapp_nodes" {
  name       = "fishapp-node-pool"
  location   = var.location
  cluster    = google_container_cluster.fishapp_cluster.name
  node_count = 3

  node_config {
    machine_type = "n1-standard-1"

    metadata = {
      disable-legacy-endpoints = "true"
    }

    service_account = google_service_account.fishapp_node_account.email

    oauth_scopes = [
      # Google APIへのアクセス制御はアプリケーション毎のサービスアカウントで行うため、oauth_scopesはすべて許可
      # https://cloud.google.com/compute/docs/access/create-enable-service-accounts-for-instances#best_practices
      "https://www.googleapis.com/auth/cloud-platform",
    ]
  }
}
