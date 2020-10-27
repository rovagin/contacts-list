PROJECT_NAME=contacts-list-frontend

build-docker:
	docker build \
       -t "${CONTAINER_NAME}:${TAG}" \
       --network host \
       -f ${DOCKER_FILE_PATH} \
       .

