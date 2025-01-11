build:
	@docker build -t maxim/todo-backend:latest .

run:
	@docker run -p 3001:3001 maxim/todo-backend:latest

deploy:
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
