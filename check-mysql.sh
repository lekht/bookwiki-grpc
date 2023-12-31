#!/bin/sh
# check-mysql.sh

set -e

until mysql -h $MYSQL_HOST -p=$MYSQL_PORT -u $MYSQL_USER --password=$MYSQL_PASSWORD -e '\q'; do
  >&2 echo "MYSQL is unavailable - sleeping"
  sleep 60
done

>&2 echo "MYSQL is up - executing command"