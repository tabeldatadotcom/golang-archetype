FROM gcr.io/distroless/base-debian10:latest
ARG ENVFILE=".env"
ARG GO_BUILD_FILE="build/backend-api"

WORKDIR /var/www/apps
COPY $GO_BUILD_FILE backend
COPY $ENVFILE .env

USER nonroot:nonroot

ENTRYPOINT ["/var/www/apps/backend"]

EXPOSE 8080/tcp
