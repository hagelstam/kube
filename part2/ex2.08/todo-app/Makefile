build:
	@docker build -t maxim/todo-app:latest .

run:
	@docker run -p 3000:3000 maxim/todo-app:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
