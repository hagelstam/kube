build:
	@docker build -t maxim/golog:latest .

deploy:
	@kubectl apply -f manifests/config-map.yaml
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
