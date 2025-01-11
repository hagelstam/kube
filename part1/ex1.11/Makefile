build-log:
	@cd golog && docker build -t maxim/golog:latest .

build-pong:
	@cd gopong && docker build -t maxim/gopong:latest .

deploy:
	@kubectl apply -f manifests/ingress.yaml

	@kubectl apply -f manifests/golog-deployment.yaml
	@kubectl apply -f manifests/golog-service.yaml
	
	@kubectl apply -f manifests/gopong-deployment.yaml
	@kubectl apply -f manifests/gopong-service.yaml

	@kubectl apply -f manifests/persistentvolume.yaml
	@kubectl apply -f manifests/persistentvolumeclaim.yaml
