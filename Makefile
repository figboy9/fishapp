GCP_TERRA_KEY = ~/.ssh/gcp/fishapp/terraform.json
GCP_KUBECTL_KEY = ~/.ssh/gcp/fishapp/kubectl.json
GCP_PROJECT = fishapp-282106
GCP_ZONE = asia-northeast1-a
GCP_CLUSTER = fishapp-cluster

terra:
	docker run -it --rm --name terra --entrypoint sh -w /terraform \
	-v $(PWD)/terraform:/terraform \
	-v $(GCP_TERRA_KEY):/credentials.json \
	-e GOOGLE_APPLICATION_CREDENTIALS=/credentials.json \
	hashicorp/terraform:light

# gcloud, kubectl, helm入りのdocker image
kubectl:
	docker run -it --rm --name kubectl -w /k8s \
	-v $(PWD):/k8s \
	-v $(GCP_KUBECTL_KEY):/credentials.json \
	-e CLOUDSDK_CORE_PROJECT=$(GCP_PROJECT) \
	-e CLOUDSDK_COMPUTE_ZONE=$(GCP_ZONE) \
	ezio1119/cloud-sdk-kubectl-helm sh -c " \
	gcloud auth activate-service-account --key-file=/credentials.json && \
	gcloud container clusters get-credentials $(GCP_CLUSTER) && sh"

kubesec:
	docker run -it --workdir /work --rm --name kubesec \
	-v $(GCP_KUBECTL_KEY):/credentials.json \
	-v $(PWD)/k8s:/work \
	-e GOOGLE_APPLICATION_CREDENTIALS=/credentials.json \
	ezio1119/kubesec

.PHONY: kubesec kubectl terra