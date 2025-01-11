build-all:
	@cd gopong && make build
	@cd golog && make build

deploy-all:
	@cd golog && make deploy
	@cd gopong && make deploy
	@cd postgres && make deploy
	@kubectl apply -f manifests/ingress.yaml
