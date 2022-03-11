#!/bin/bash

help()
{
cat << EOF
usage: $0 options
Create backup of your database.
OPTIONS:
  -n    Name of the environment you are running at
EOF
}

while getopts ":n:" OPTION
do
    case $OPTION in
    n)
      NAME=$OPTARG
      ;;
    *)
      help
      exit 1
      ;;
    esac
done

if [ -z "$NAME" ]; then
  printf "\n"
  printf '\e[1;91m%-6s\e[m' "NAME MUST BE PROVIDED WITH -n FLAG (production, staging, local), name should be the same as the sourced .env file"
  printf "\n\n"
  help
  exit 1
fi

DIR="backup"
if [ ! -d $DIR ]
then
    mkdir -p $DIR
fi

HOST=$DB_HOST
PORT=$DB_PORT
PASS=$DB_PASS
USER=$DB_USER

DB=$DB_NAME

export MYSQL_PWD=$PASS

echo "##### Backup of $DATABASE #####"
mysqldump --column-statistics=0 --set-gtid-purged=OFF --single-transaction --events --routines --port=$PORT -u$USER -h$HOST $DB > $DIR/$NAME-$DB.sql