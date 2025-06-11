#!/bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$DIR"

MYSQL_PASS=$(cat "$DIR/password")

sudo docker build -t jlum_db .

sudo docker rm -f jlum_db 2>/dev/null || true

sudo docker run -d \
  -p 3306:3306 \
  -v "$DIR/mount:/mount" \
  --name jlum_db \
  -e MYSQL_ROOT_PASSWORD="$MYSQL_PASS" \
  -e MYSQL_DATABASE=jlum_db \
  -e MYSQL_USER=jlum_user \
  -e MYSQL_PASSWORD="$MYSQL_PASS" \
  jlum_db

