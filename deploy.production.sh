#!/bin/sh

APP_NAME=blog-server
PLATFORM=linux/arm64
DOCKER_FILE=./Dockerfile.production
DOCKER_SERVER_PORT=10000
LOCAL_SERVER_PORT=8080

echo "Delete old image ..."
docker rmi -f ${APP_NAME}

echo "Docker build image ..."
docker build -t ${APP_NAME} --platform=${PLATFORM} -f ${DOCKER_FILE} .

echo "Docker delete old container ..."
docker rm -f ${APP_NAME}

echo "Docker run container ..."
docker run -d --name ${APP_NAME} \
  --network=fd-network \
  -e ENV=DOCKER \
  -p ${LOCAL_SERVER_PORT}:${DOCKER_SERVER_PORT} \
  ${APP_NAME}