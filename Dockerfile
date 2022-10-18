FROM alpine:3.16
# build args
ARG ENVFILE=".env.example"
ARG GO_BUILD_FILE="build/backend-api"
# default environment variables
ENV DEFAULT_FOLDER=/usr/share/go-docker
ENV DEFAULT_LOG_FILE=/var/log/go-docker

WORKDIR $DEFAULT_FOLDER
ENV APP_PORT=8080

# copy binary extract
COPY $GO_BUILD_FILE application
COPY $ENVFILE .env

EXPOSE $APP_PORT/tcp

ENTRYPOINT ["/usr/share/go-docker/application"]
