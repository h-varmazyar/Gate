#!/bin/sh
#service_name="${1}"

echo "running"

is_up=false
is_down=false

while true; do
  case "$1" in
    -s|--service)
      service_name="$2"
      shift 2;;
    up)
      echo "up"
      is_up=true
      shift ;;
    down)
      echo "down"
      is_down=true
      shift ;;
    --|*)
      break;;
  esac
done

echo $is_up
echo $is_down

# shellcheck disable=SC2004
# shellcheck disable=SC2039
if (($is_up && $is_down)) || ((!$is_up && !$is_down)) ; then
  echo "invalid deploy type"
  exit 1
fi

if $is_up ; then
  if [ "$service_name" = "" ]; then
    echo "running all services"
    docker-compose -f ./deploy/docker-compose.yml up -d
  else
    echo "running service $service_name"
    docker-compose -f ./deploy/docker-compose.yml up -d "$service_name"
  fi
fi

if $is_down ; then
  if [ "$service_name" = "" ]; then
      echo "running all services"
      docker-compose -f ./deploy/docker-compose.yml down
    else
      echo "running service $service_name"
      docker-compose -f ./deploy/docker-compose.yml down "$service_name"
  fi
fi

echo "done"