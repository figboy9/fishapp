start:
	minikube start --mount-string $$(dirname ${PWD})/microservices:/microservices --mount

stop:
	minikube stop

addrepo:
	helm repo add stable https://kubernetes-charts.storage.googleapis.com/
	helm repo add bitnami https://charts.bitnami.com/bitnami
	
dev:
	skaffold dev --port-forward
	kubectl delete pvc \
	redis-data-blacklist-kvs-redis-master-0 \
	redis-data-chat-kvs-redis-master-0

migrdb:
	kubectl apply -f jobs/migrate/
# 正常終了したジョブのみを消す
deljobs:
	kubectl delete job $$(kubectl get job -o=jsonpath='{.items[?(@.status.succeeded==1)].metadata.name}')

.PHONY: rmdata start stop migrdb deljobs addrepo dev