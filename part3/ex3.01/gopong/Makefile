build:
	@docker build -t maximilianhagelstam/gopong .

push:
	@docker push maximilianhagelstam/gopong

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
