#!/bin/bash
set -e

echo "entrypoint shell started"

# migrate files
#migrate --srouce file://./migrate/mysql --database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DATABASE"

# Then exec the container's main process (what's set as CMD in the Dockerfile).
exec "$@"
