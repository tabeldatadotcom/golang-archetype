FROM gcr.io/distroless/base-debian10:latest
ARG ENVFILE="env.yaml"

WORKDIR /var/www/apps
COPY build/application backend
COPY $ENVFILE env.yaml

USER nonroot:nonroot

ENTRYPOINT ["/var/www/apps/backend"]

EXPOSE 8080/tcp
