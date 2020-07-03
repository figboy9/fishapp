resource "google_compute_global_address" "fishapp_external_address" {
  name = "fishapp-external-address"
  description = "fishapp External Address"
}

# resource "google_dns_managed_zone" "fishapp_zone" {
#   name        = "fishapp-zone"
#   dns_name    = var.dns_name
#   description = "fishapp DNS zone"
# }

# resource "google_dns_record_set" "fishapp_dns_record" {
#   name         = google_dns_managed_zone.fishapp_zone.dns_name
#   managed_zone = google_dns_managed_zone.fishapp_zone.name
#   type         = "A"
#   ttl          = 300

#   rrdatas = [google_compute_global_address.fishapp_external_address.address]
# }