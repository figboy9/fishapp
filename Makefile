CWD = $(shell pwd)
GCP_TERRA_KEY = ~/.ssh/gcp/fishapp/terraform.json
GCP_KUBECTL_KEY = ~/.ssh/gcp/fishapp/kubectl.json

terra:
	docker run -it --rm --name terra --entrypoint sh -w /terraform \
	-v $(CWD)/terraform:/terraform \
	-v $(GCP_TERRA_KEY):/CREDENTIALS_FILE.json \
	-e GOOGLE_APPLICATION_CREDENTIALS=/CREDENTIALS_FILE.json \
	hashicorp/terraform:light

# gcloud, kubectl, helm入りのdocker image
kubectl:
	docker run -it --rm --name kubectl -w /k8s \
	-v $(CWD)/k8s:/k8s \
	-v $(GCP_KUBECTL_KEY):/CREDENTIALS_FILE.json \
	-e GOOGLE_APPLICATION_CREDENTIALS=/CREDENTIALS_FILE.json \
	ezio1119/cloud-sdk-kubectl-helm sh -c " \
	gcloud auth activate-service-account --key-file=/CREDENTIALS_FILE.json && \
	gcloud config set project $(PROJECT_ID) && \
	gcloud config set compute/zone $(ZONE) && \
	gcloud container clusters get-credentials $(CLUSTER) && \
	sh"

kubesec:
	docker run --workdir /work -it --rm --name kubesec \
	-v $(GCP_KUBECTL_KEY):/CREDENTIALS_FILE.json \
	-v $(CWD)/k8s:/work \
	-e GOOGLE_APPLICATION_CREDENTIALS=/CREDENTIALS_FILE.json \
	ezio1119/kubesec ${CMD}

.PHONY: kubesec kubectl terra