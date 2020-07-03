# プロジェクトレベルのロール
locals {
  fishapp_node_roles = [
    "roles/logging.logWriter",
    "roles/monitoring.metricWriter",
    "roles/monitoring.viewer",
    "roles/storage.objectViewer",
  ]

  fishapp_kubectl_roles = [
    "roles/container.admin",
  ]

  fishapp_github_actions_roles = [
    "roles/storage.admin"
  ]

  fishapp_config_connector_roles = [
    "roles/editor",
  ]
}

resource "google_service_account" "fishapp_node_account" {
  account_id   = "fishapp-node-account"
  display_name = "[Terraform] fishapp Node Account"
}

resource "google_project_iam_member" "fishapp_node_iam" {
  for_each = toset(local.fishapp_node_roles)
  role     = each.value
  member   = "serviceAccount:${google_service_account.fishapp_node_account.email}"
}

resource "google_service_account" "fishapp_kubectl_account" {
  account_id   = "fishapp-kubectl-account"
  display_name = "[Terraform] fishapp Kubectl Account"
}

resource "google_project_iam_member" "fishapp_kubectl_iam" {
  for_each = toset(local.fishapp_kubectl_roles)
  role     = each.value
  member   = "serviceAccount:${google_service_account.fishapp_kubectl_account.email}"
}

resource "google_service_account_key" "fishapp_kubectl_account_key" {
  service_account_id = google_service_account.fishapp_kubectl_account.name
}

resource "google_service_account" "fishapp_github_actions_account" {
  account_id   = "fishapp-github-actions-account"
  display_name = "[Terraform] fishapp Github Actions Account"
}

resource "google_project_iam_member" "fishapp_github_actions_iam" {
  for_each = toset(local.fishapp_github_actions_roles)
  role     = each.value
  member   = "serviceAccount:${google_service_account.fishapp_github_actions_account.email}"
}

resource "google_service_account_key" "fishapp_github_actions_account_key" {
  service_account_id = google_service_account.fishapp_github_actions_account.name
}

resource "google_service_account" "fishapp_config_connector_account" {
  account_id   = "fishapp-config-conn-account"
  display_name = "[Terraform] fishapp Config Connector Account."
}

resource "google_project_iam_member" "fishapp_config_connector_iam" {
  for_each = toset(local.fishapp_config_connector_roles)
  role     = each.value
  member   = "serviceAccount:${google_service_account.fishapp_config_connector_account.email}"
}

resource "google_service_account_iam_member" "fishapp_GSA_to_KSA_iam" {
  service_account_id = google_service_account.fishapp_config_connector_account.name
  role = "roles/iam.workloadIdentityUser"

  member = "serviceAccount:${var.project}.svc.id.goog[cnrm-system/cnrm-controller-manager-${var.k8s_namespace}]"
}