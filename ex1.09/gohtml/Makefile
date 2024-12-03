build:
	@docker build -t maxim/gohtml:latest .

run:
	@docker run -p 8080:8080 maxim/gohtml:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml
