EXE_LINUX = "run_server_linux"
EXE_WIN = "run_server_win.exe"
DOCKER_IMAGE = "wishlist_frontend"
CONTAINER_NAME ="wishlist_frontend"

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Builds the executable for linux
	@echo "### Building Linux Executable ###"
	@GOOS="linux" go build -o data/${EXE_LINUX} ./src/

build-win: ## Builds the executable for windows
	@echo "### Building Windows Executable ###"
	@GOOS="windows" go build -o data/${EXE_WIN} ./src/

image: build ## Builds the docker image
	@echo "### Building Docker Image ###"
	@docker build -t ${DOCKER_IMAGE} .

up: down image ## Starts the docker-compose cluster
	@echo "### Starting Container ###"
	@docker run -d --name ${CONTAINER_NAME} -v "/etc/letsencrypt:/certs:ro" -p 5000:5000 ${DOCKER_IMAGE}

down: ## Stops the docker-compose cluster
	@echo "### Stopping Container ###"
	@-docker stop ${CONTAINER_NAME}
	@-docker rm ${CONTAINER_NAME}

bash-api: ## Connects to api container
	@docker exec -it wishlistapi_wishlist_api_1 sh

bash-db: ## Connects to db container
	@docker exec -it wishlistapi_wishlist_db_1 sh