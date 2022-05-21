EXE_LINUX = "run_server_linux"
EXE_WIN = "run_server_win.exe"
DOCKER_IMAGE = "wishlist_frontend"
CONTAINER_NAME ="wishlist_frontend"

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## Builds the executable for linux
build:
	@echo "### Building Linux Executable ###"
	@GOOS="linux" go build -o data/${EXE_LINUX} ./src/

## Builds the executable for windows
build-win:
	@echo "### Building Windows Executable ###"
	@GOOS="windows" go build -o data/${EXE_WIN} ./src/

## Builds the docker image
image: build
	@echo "### Building Docker Image ###"
	@docker build -t ${DOCKER_IMAGE} .

## Starts the docker-compose cluster
up: down image
	@echo "### Starting Container ###"
	@docker run -d --name ${CONTAINER_NAME} -p 5000:5000 ${DOCKER_IMAGE} -v "/etc/letsencrypt:/certs:ro"

## Stops the docker-compose cluster
down:
	@echo "### Stopping Container ###"
	@-docker stop ${CONTAINER_NAME}
	@-docker rm ${CONTAINER_NAME}

## Connects to api container
bash-api:
	@docker exec -it wishlistapi_wishlist_api_1 sh

## Connects to db container
bash-db:
	@docker exec -it wishlistapi_wishlist_db_1 sh