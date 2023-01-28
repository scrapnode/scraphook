#!/bin/bash
set -e

# ignore docker compose orphans warning because we want to start 2 different compose file sequential
export COMPOSE_IGNORE_ORPHANS=true

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLEANUP=${CLEANUP:-"false"}

if [ "$CLEANUP" == "true" ]; then
  /bin/bash "$__dir/local_clean.sh"
fi

docker compose up -d
docker compose -f docker-compose.service.yaml up -d