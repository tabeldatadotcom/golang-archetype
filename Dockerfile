FROM ubuntu:jammy
# build args
ARG ENVFILE=".env.example"
ARG GO_BUILD_FILE="build/backend-api"
# default environment variables
ENV DEFAULT_FOLDER=/usr/share/go-docker
ENV DEFAULT_LOG_FILE=/var/log/go-docker

RUN addgroup godocker && \
    adduser norootuser --system --no-create-home

RUN mkdir -p $DEFAULT_FOLDER && \
    chmod -R 777 $DEFAULT_FOLDER && \
    mkdir -p $DEFAULT_LOG_FILE && \
    chmod -R 777 $DEFAULT_LOG_FILE

USER norootuser

WORKDIR $DEFAULT_FOLDER
ENV APP_PORT=8080

# copy binary extract
COPY --chown=norootuser $GO_BUILD_FILE application
COPY --chown=norootuser $ENVFILE .env

CMD ["/usr/share/go-docker/application"]
EXPOSE $APP_PORT/tcp
