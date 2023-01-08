#!/bin/sh
set -e

# split args string to args param and pass it to our application command
# /app/scraphook "service start scrap.service.webhook.rest" -> NOT WORK because the app receives string as 1 args
# /app/scraphook service start scrap.service.webhook.rest -> WORK because the app receives args
set -- /app/scraphook $@

# then execute the command with args
exec $@