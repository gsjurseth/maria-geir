FROM openjdk:12-alpine3.9 AS openjdk

ENV APP_DIR /app
WORKDIR $APP_DIR

RUN apk add --no-cache tini

COPY target/*.jar /app/app.jar
COPY init.sh /app/

EXPOSE 8080

ENTRYPOINT ["tini", "-s", "--", "sh", "init.sh"]

