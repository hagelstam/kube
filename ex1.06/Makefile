docker-build:
	@docker build -t gohtml .
	@docker tag gohtml maxim/gohtml:latest

docker-run:
	@docker run gohtml

deploy:
	@kubectl apply -f manifests/service.yaml
