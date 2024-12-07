build-all:
	@cd todo-app && make build
	@cd todo-backend && build

deploy:
	@kubectl apply -f manifests/ingress.yaml
	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml

tmp-dir:
	@docker exec k3d-k3s-default-server-0 mkdir -p /tmp/kube
	@docker exec k3d-k3s-default-server-0 chmod 777 /tmp/kube
