build:
	@docker build -t maxim/gopic:latest .

run:
	@docker run -p 3000:3000 maxim/gopic:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml
	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml

mkdir-kube:
	@docker exec k3d-k3s-default-server-0 mkdir -p /tmp/kube
	@docker exec k3d-k3s-default-server-0 chmod 777 /tmp/kube
