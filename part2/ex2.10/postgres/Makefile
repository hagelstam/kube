deploy:
	@sops -d secret.enc.yaml | kubectl apply -f -
	@kubectl apply -f service.yaml
	@kubectl apply -f statefulset.yaml
