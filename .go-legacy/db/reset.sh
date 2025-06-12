#!/bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$DIR"

MYSQL_PASS=$(cat "$DIR/password")

sudo docker rm -f jlum_db
sudo rm -rf mount/*

./run.sh

# read temp
sleep 15

# Create the user if it does not exist
sudo docker exec -i jlum_db mysql -uroot -p"$MYSQL_PASS" -e "CREATE USER IF NOT EXISTS 'jlum_user'@'172.17.0.1' IDENTIFIED BY '$MYSQL_PASS';"

# Grant privileges to allow connections from 172.17.0.1
sudo docker exec -i jlum_db mysql -uroot -p"$MYSQL_PASS" -e "GRANT ALL PRIVILEGES ON jlum_db.* TO 'jlum_user'@'172.17.0.1';"
sudo docker exec -i jlum_db mysql -uroot -p"$MYSQL_PASS" -e "FLUSH PRIVILEGES;"

# Initialize the database
sudo docker exec -i jlum_db mysql -ujlum_user -p"$MYSQL_PASS" jlum_db < "$DIR/init.sql"

# Show tables to verify
sudo docker exec -i jlum_db mysql -ujlum_user -p"$MYSQL_PASS" -e 'SHOW TABLES;' jlum_db

