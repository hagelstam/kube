docker-build:
	@docker build -t maxim/todo-backend:latest .

docker-run:
	@docker run -p 8080:8080 maxim/todo-backend:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml
	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml

tmp-kube:
	@docker exec k3d-k3s-default-server-0 mkdir -p /tmp/kube
	@docker exec k3d-k3s-default-server-0 chmod 777 /tmp/kube