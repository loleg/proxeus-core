BUILD_WITH_DOCKER=true make init server-docker
docker build .
docker-compose build
docker-compose restart
