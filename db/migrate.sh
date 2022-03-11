#!/bin/bash
TIMEFORMAT="%E"

help()
{
cat << EOF
usage: $0 options
Applying migration scripts to database.
Migrations scripts must be named timestamp_title.up.sql and timestamp.down.sql
Scripts must be placed in migrations/.....sql directory.
OPTIONS:
  -m <UP|DOWN>    UP for scripts to be be migrated or down for roll back
                  (default: UP)
EOF
}

source ../.env

HOST=$DB_HOST
PORT=$DB_PORT
PASS=$DB_PASS
USER=$DB_USER

DB=$DB_NAME

while getopts “:m:” OPTION
do
    case $OPTION in
    m)
      MIGRATE=$OPTARG
      if [ ! "$MIGRATE" == "UP" ] && [ ! "$MIGRATE" == "DOWN" ]; then
        printf "\n"
        printf '\e[1;91m%-6s\e[m' "MIGTATION DIRECTION CAN BE UP OR DOWN SPECIFIED WITH -m FLAG"
        printf "\n\n"
        help
        exit 1
      fi
      ;;
    *)
      help
      exit 1
      ;;
    esac
done

export MYSQL_PWD=$PASS

get_script_name () {
  echo `echo $1 | awk -F'/' '{print $NF}' | awk -F'.' '{print $NR}'`;
}

check_if_script_in_db () {
  local script=$1
  local err=0
  local name=""
  for name in `mysql -h$HOST -u$USER -D$DB -BN -e "select name from migrations"`; do
    if [[ $name = $script ]]; then
      err=1;
      break;
    fi
  done 
  echo $err
}

rollback_script () {
  local script=$1
  local err=0;
  local id=0
  local output=""
  local time=""
  local script_name=""
  local script_run=0
  local exit_code=0
  
  script_name=$(get_script_name $script)
  script_run=$(check_if_script_in_db $script_name)
  
  if [[ $script_run -eq "1" ]]; then
    output=$( { time mysql -h$HOST -u$USER -D$DB < $script; } 2>&1 )
    exit_code=$?
    time=`echo $output | sed 's/,/./'`
    if [[ $exit_code = 0 ]]; then 
      printf "$script $time"
      printf "\n"
    else
      printf '\e[1;91m%-6s\e[m' "SCRIPT $script FAILED TO RUN"
      printf "\n"
      printf "$output"
      printf "\n"
      exit 1
    fi
    mysql -h$HOST -u$USER -D$DB -e "DELETE FROM migrations WHERE name=\"$script_name\"";

  fi
}

migrate_script () {
  local script=$1
  local err=0;
  local id=0
  local output=""
  local time=""
  local script_name=""
  local script_run=0
  local exit_code=0
  
  script_name=$(get_script_name $script)
  script_run=$(check_if_script_in_db $script_name)
  
  if [[ $script_run -eq "0" ]]; then
    mysql -h$HOST -u$USER -D$DB -e "INSERT INTO migrations (name,date_of_running,hostname) VALUE (\"$script_name\",NOW(),\"$HOSTNAME\")";
    id=`mysql -h$HOST -u$USER -D$DB -BN -e "SELECT id FROM migrations WHERE name=\"$script_name\"";`;
    output=$( { time mysql -h$HOST -u$USER -D$DB < $script; } 2>&1 )
    exit_code=$?
    time=`echo $output | sed 's/,/./'`
    mysql -h$HOST -u$USER -D$DB -e "UPDATE migrations SET exit_code=$exit_code, time_of_running=\"$time\" WHERE id=$id";
    if [[ $exit_code = 0 ]]; then 
      printf "$script $time"
      printf "\n"
    else
      printf '\e[1;91m%-6s\e[m' "SCRIPT $script FAILED TO RUN"
      printf "\n"
      printf "$output"
      printf "\n"
      exit 1
    fi
  fi 
}

SCRIPT_DIRECTORY="migrations"
if [[ ! -d "$SCRIPT_DIRECTORY" ]]; then
  printf "\n"
  printf '\e[1;91m%-6s\e[m' "$SCRIPT_DIRECTORY DOES NOT EXIST."
  printf "\n\n"
  exit 1
fi

if ! `mysqlshow -u$USER -h$HOST "$DB" > /dev/null 2>&1`; then 
  printf "\n"
  printf '\e[1;91m%-6s\e[m' "DATABASE $DB DOES NOT EXIST OR WE HAD DIFFICULTIES TO LOG IN."
  printf "\n\n"
  exit 1
fi

mysql -h$HOST -u$USER -D$DB -e "CREATE TABLE IF NOT EXISTS migrations (id int(11) NOT NULL AUTO_INCREMENT, name varchar(150) NOT NULL, hostname varchar(150) DEFAULT NULL, exit_code tinyint(4) DEFAULT NULL, date_of_running datetime NOT NULL, time_of_running float(10,4) DEFAULT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8";

if [[ "$MIGRATE" = "UP" ]]; then
  for script in `ls -1 $SCRIPT_DIRECTORY/*.up.sql | sort`; do
    if [[ -f $script ]]; then
      migrate_script $script
    fi
  done
elif [[ "$MIGRATE" = "DOWN" ]]; then
  for script in `ls -1 $SCRIPT_DIRECTORY/*.down.sql | sort -r`; do
    if [[ -f $script ]]; then
      rollback_script $script
    fi
  done
fi

exit 0