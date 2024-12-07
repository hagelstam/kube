build:
	@docker build -t maxim/cron:latest .

deploy:
	@k3d image import maxim/cron:latest
	@kubectl apply -f cron.yaml
