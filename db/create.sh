#!/bin/bash

help()
{
cat << EOF
usage: $0 options
Create properly named migration files.
OPTIONS:
  -n <name>       Create name of the migration script
EOF
}

while getopts “:n:” OPTION
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
  printf '\e[1;91m%-6s\e[m' "NAME MUST BE PROVIDED WITH -n FLAG"
  printf "\n\n"
  help
  exit 1
fi

SCRIPT_DIRECTORY="migrations"
mkdir -p $SCRIPT_DIRECTORY
time=`date +'%Y%m%d%H%M'`
touch $SCRIPT_DIRECTORY/$time-$NAME.up.sql
touch $SCRIPT_DIRECTORY/$time-$NAME.down.sql

echo "$SCRIPT_DIRECTORY/$time-$NAME.up.sql created"
echo "$SCRIPT_DIRECTORY/$time-$NAME.down.sql created"

exit 0