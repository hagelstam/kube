docker-build:
	@docker build -t goserver .
	@docker tag goserver maxim/goserver:latest

docker-run:
	@docker run goserver

deploy:
	@kubectl apply -f manifests/deployment.yaml

logs:
	@kubectl logs -f goserver-5bb8f74fc5-bql72
