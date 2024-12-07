build:
	@docker build -t maxim/todo-app:latest .

run:
	@docker run -p 3000:3000 maxim/todo-app:latest

deploy:
	@k3d image import maxim/todo-app:latest
	@kubectl apply -f manifests/deployment.yaml
	@kubectl apply -f manifests/service.yaml
