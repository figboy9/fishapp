resource "google_kms_key_ring" "fishapp_kubesec_key_ring" {
  name     = "fishapp-kubesec-key-ring"
  location = var.region
}

resource "google_kms_crypto_key" "fishapp_kubesec_crypto_key" {
  name     = "fishapp-kubesec-crypto-key"
  key_ring = google_kms_key_ring.fishapp_kubesec_key_ring.id
}

resource "google_kms_crypto_key_iam_member" "fishapp_kubesec_iam" {
  crypto_key_id = google_kms_crypto_key.fishapp_kubesec_crypto_key.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${google_service_account.fishapp_kubectl_account.email}"
}