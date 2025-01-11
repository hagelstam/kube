build-all:
	@cd gopong && make build
	@cd golog && make build

push-all:
	@cd gopong && make push
	@cd golog && make push

deploy-all:
	@cd golog && make deploy
	@cd gopong && make deploy
	@cd postgres && make deploy
	@kubectl apply -f manifests/ingress.yaml

