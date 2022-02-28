#!/bin/bash


# Build docker images with build kit to use in development environment
COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose build --pull --force-rm --no-cache --build-arg SSH_KEY="$(cat ~/.ssh/id_rsa)"

# Run docker containers for project
docker-compose up

echo 'Dev environment is ready to use' 