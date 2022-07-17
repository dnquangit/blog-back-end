#!/bin/sh

APP_NAME=food-delivery-g04
PLATFORM=linux/arm64
DOCKER_FILE=./Dockerfile.production

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
  -p 8080:10000 \
  ${APP_NAME}