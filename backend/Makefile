PROJECT_NAME=contacts-list

gen-api:
	cd api && easyjson --all .

download:
	go mod download

tidy:
	go mod tidy

lint:
	golangci-lint run

build-service:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build \
        -installsuffix 'static' \
        -o ./app ./cmd/${PROJECT_NAME}/*


build-docker:
	docker build \
       -t "${CONTAINER_NAME}:${TAG}" \
       --build-arg PROJECT_NAME=${PROJECT_NAME} \
       --network host \
       -f ${DOCKER_FILE_PATH} \
       .

run:
	go run cmd/${PROJECT_NAME}
