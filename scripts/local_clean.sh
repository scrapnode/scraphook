#!/bin/bash
set -e

docker compose -f docker-compose.service.yaml down
docker compose down
sleep 3
docker volume prune -f