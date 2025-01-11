build-all:
	@cd todo-app && make docker-build
	@cd todo-backend && make docker-build

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml
	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml

tmp-dir:
	@docker exec k3d-k3s-default-server-0 mkdir -p /tmp/kube
	@docker exec k3d-k3s-default-server-0 chmod 777 /tmp/kube
