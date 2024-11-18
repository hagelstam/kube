docker-build:
	@docker build . -t simple-server:latest

docker-run:
	@docker run simple-server:latest

deploy:
	@kubectl create deployment simple-server --image=simple-server:latest

logs:
	@kubectl logs -f simple-server-f558cfcd5-w6hr2
