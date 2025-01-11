build:
	@docker build -t maximilianhagelstam/golog .

push:
	@docker push maximilianhagelstam/golog

deploy:
	@kubectl apply -f manifests/config-map.yaml
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
