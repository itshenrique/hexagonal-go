#!/bin/bash
# Config
docker exec nito-db-1 mysqldump -h"localhost" -P"3306" -u"$1" -p"$2" $3 > ./database/DB.sql