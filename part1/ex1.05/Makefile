docker-build:
	@docker build -t gohtml .
	@docker tag gohtml maxim/gohtml:latest

docker-run:
	@docker run gohtml

deploy:
	@kubectl apply -f manifests/deployment.yaml

port-forward:
	@kubectl port-forward gohtml-6c47d9d647-xwgxq 8080:8080
