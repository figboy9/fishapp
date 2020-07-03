output "fishapp_kubectl_account_key" {
  description = "fishapp Kubectl Account Key"
  value       = base64decode(google_service_account_key.fishapp_kubectl_account_key.private_key)
}

output "fishapp_github_actions_account_key" {
  description = "fishapp Github Actions Account Key"
  value       = base64decode(google_service_account_key.fishapp_github_actions_account_key.private_key)
}

