FROM alpine:latest
# build args
ARG ENVFILE=".env.example"
ARG GO_BUILD_FILE="build/backend-api"
# default environment variables
ENV APP_PORT=8080

# install dependency
RUN apk --no-cache add ca-certificates

# create linux user account
RUN addgroup -S ec2-group && \
    adduser -S ec2-user -G ec2-group
USER ec2-user

# copy binary extract
WORKDIR /var/www/apps
COPY $GO_BUILD_FILE backend
COPY $ENVFILE .env

ENTRYPOINT ["./backend"]
EXPOSE $APP_PORT/tcp
