#!/bin/bash
set -e


PROTO_FILES=(
  "admin/protos/account.proto"
  "admin/protos/webhook.proto"
  "admin/protos/endpoint.proto"
)
PROTO_FOLDERS=(
  "admin/protos"
  "admin/protos"
  "admin/protos"
)
COUNT=${#PROTO_FILES[@]}

echo "PWD: ${PWD}, COUNT: ${COUNT}"

for (( n=0; n<=COUNT-1; n++ ))
do
  FILE=${PROTO_FILES[n]}
  FOLDER=${PROTO_FOLDERS[n]}

  protoc --proto_path="$FOLDER" --go_out="$FOLDER" --go_opt=paths=source_relative --go-grpc_out="$FOLDER" --go-grpc_opt=paths=source_relative "$FILE"
  echo "--> $FILE"
done
