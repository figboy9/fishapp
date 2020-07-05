variable "project" {
  description = "The project ID where all resources will be launched."
  type        = string
  default     = "fishapp-282106"
}

variable "location" {
  description = "The location (region or zone) of the GKE cluster."
  type        = string
  default     = "asia-northeast1-a"
}

variable "region" {
  description = "The region for the network. If the cluster is regional, this must be the same region. Otherwise, it should be the region of the zone."
  type        = string
  default     = "asia-northeast1"
}

variable "k8s_namespace" {
  description = "The Namespace of k8s watch by config connector"
  type        = string
  default     = "default"
}

variable "dns_name" {
  description = "The Domain Name of fishapp"
  type        = string
  default     = "fishapp.work."
}