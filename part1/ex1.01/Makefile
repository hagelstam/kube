docker-build:
	@docker build . -t log-output:latest

docker-run:
	@docker run log-output:latest

deploy:
	@kubectl create deployment log-output --image=log-output:latest

logs:
	@kubectl logs -f log-output-6c8db5d59f-7bdqw
