docker-build:
	@docker build -t golog .
	@docker tag golog maxim/golog:latest

docker-run:
	@docker run golog

deploy:
	@kubectl apply -f manifests/deployment.yaml

logs:
	@kubectl logs -f golog-847f9ffccc-f29qm
