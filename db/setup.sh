#!/bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$DIR"

MYSQL_PASS=$(cat "$DIR/password")

sudo docker exec -i jlum_db mysql -ujlum_user -p"$MYSQL_PASS" jlum_db < "$DIR/init.sql"
sudo docker exec -i jlum_db mysql -ujlum_user -p"$MYSQL_PASS" -e 'SHOW TABLES;' jlum_db

