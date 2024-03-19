
.PHONY: deploy-local-tf
deploy-local-tf:
	@echo "Deploying to local environment using Terraform"
	@cd infra/local && terraform init && terraform apply -auto-approve

.PHONY: up rebuild-up
up:
	@echo "Starting the services using docker compose"
	docker compose up -d

rebuild-up:
	@echo "Rebuilding and starting the services using docker compose"
	docker compose up --build -d

.PHONY: down
down:
	@echo "Stopping the services using docker compose"
	docker compose down

.PHONY: watch-log-all
watch-log-all:
	@echo "Watching logs for all services"
	docker compose logs -f

.PHONY: watch-log-api
watch-log-api:
	@echo "Watching logs for api-dev service"
	docker compose logs -f api-dev
