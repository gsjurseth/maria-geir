FROM alpine:latest

RUN apk add curl net-tools

COPY ./apigee-remote-service-envoy /apigee-remote-service-envoy

ENTRYPOINT /apigee-remote-service-envoy -c /config.yaml -l $DEBUG
