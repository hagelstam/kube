build:
	@docker build -t maxim/gopong:latest .

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
