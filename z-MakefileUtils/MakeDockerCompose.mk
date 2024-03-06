## include z-MakefileUtils/MakeDockerCompose.mk
## MakeDockerCompose.mk settings start
INFO_DOCKER_COMPOSE_DEFAULT_FILE ?=docker-compose.yml
## MakeDockerCompose.mk settings end

ENV_INFO_DOCKER_COMPOSE_FILE =${INFO_DOCKER_COMPOSE_DEFAULT_FILE}

dockerComposeCheck:
	@docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} config -q
	$(info -> docker compose config check pass as: ${ENV_INFO_DOCKER_COMPOSE_FILE})

dockerComposeUp: dockerComposeCheck
	@echo "-> docker compose up as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} up -d --build --remove-orphans

dockerComposePs:
	@echo "-> docker compose ps as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} ps

dockerComposeLogs:
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} logs

dockerComposeFollowLogs:
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} logs -f

dockerComposeRestart: dockerComposeCheck
	@echo "-> docker compose restart as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} restart

dockerComposeDown: dockerComposeCheck
	@echo "-> docker compose down as: ${ENV_INFO_DOCKER_COMPOSE_FILE}"
	docker-compose -f ${ENV_INFO_DOCKER_COMPOSE_FILE} down --remove-orphans --rmi local

helpDockerCompose:
	@echo "=== this make file can include MakeDockerCompose.mk then use"
	@echo "- must has file: ${ENV_INFO_DOCKER_COMPOSE_FILE} by setting: [ INFO_DOCKER_COMPOSE_DEFAULT_FILE ]"
	@echo ""
	@echo "# - then can run as docker-compose build image and up"
	@echo "$$ make dockerComposeUp"
	@echo "# - then see log as docker-compose"
	@echo "$$ make dockerComposeLogs"
	@echo "# - then follow log as docker-compose"
	@echo "$$ make dockerComposeFollowLogs"
	@echo "# - down as docker-compose will auto remove local image"
	@echo "$$ make dockerComposeDown"
	@echo ""