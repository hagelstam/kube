build:
	@docker build -t maxim/todo-backend:latest .

run:
	@docker run -p 3001:3001 maxim/todo-backend:latest

deploy:
	@k3d image import maxim/todo-backend:latest
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
