build-all:
	@cd todo-app && make build
	@cd todo-backend && make build
	@cd cron && make build

deploy:
	@kubectl apply -f manifests/ingress.yaml
	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml

deploy-all:
	@cd postgres && make deploy
	@cd todo-backend && make deploy
	@cd todo-app && make deploy
	@make deploy
	@cd cron && make deploy

tmp-dir:
	@docker exec k3d-k3s-default-server-0 mkdir -p /tmp/kube
	@docker exec k3d-k3s-default-server-0 chmod 777 /tmp/kube
