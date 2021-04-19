#!/usr/bin/env bash

set -e

ENVIRONMENT_NAME="${SPRING_PROFILES_ACTIVE:-"local"}"
APPLICATION_PORT="${PORT:-"8080"}"
COMMAND=${1:-"app"}
echo $COMMAND
if [ $COMMAND == "migrate" ]; then
		ENVIRONMENT_NAME="migration"
fi
echo "ENVIRONMENT_NAME: ${ENVIRONMENT_NAME}"

case "$COMMAND" in
  migrate|app)
    exec java -Dspring.profiles.active=${ENVIRONMENT_NAME} \
      -Dserver.port=${APPLICATION_PORT} \
      -jar /app/app.jar
    ;;
  *)
    exec sh -c "$*"
    ;;
esac
