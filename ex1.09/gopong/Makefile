build:
	@docker build -t maxim/gopong:latest .

run:
	@docker run -p 8081:8081 maxim/gopong:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
