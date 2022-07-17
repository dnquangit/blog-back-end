#!/bin/sh

APP_NAME=food-delivery-g04
PLATFORM=linux/amd64
DOCKER_FILE=./Dockerfile.local

echo "Delete old image ..."
docker rmi -f ${APP_NAME}

echo "Docker build image ..."
docker build -t ${APP_NAME} --platform=${PLATFORM} -f ${DOCKER_FILE} .

echo "Docker delete old container ..."
docker rm -f ${APP_NAME}

echo "Docker run container ..."
docker run -d --name ${APP_NAME} \
  --network=fd-network \
  -e VIRTUAL_HOST="demo2.bangthong.one" \
  -e LETSENCRYPT_HOST="demo2.bangthong.one" \
  -e LETSENCRYPT_EMAIL="dnquangit@gmail.com" \
  -e ENV=DOCKER \
  -p 10000:10000 \
  ${APP_NAME}