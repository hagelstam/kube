docker-build:
	@docker build -t maxim/golog:latest .

docker-run:
	@docker run maxim/golog:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
	@kubectl apply -f manifests/ingress.yaml

logs:
	@kubectl logs -f golog-847f9ffccc-f29qm
