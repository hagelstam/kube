build-writer:
	@cd log-writer && docker build -t maxim/log-writer:latest .

build-reader:
	@cd log-reader && docker build -t maxim/log-reader:latest .

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml
