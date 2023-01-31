#!/bin/bash
set -e


PROTO_FILES=("admin/protos/account.proto")
PROTO_FOLDERS=("admin/protos")
COUNT=${#PROTO_FILES[@]}

echo "PWD: ${PWD}, COUNT: ${COUNT}"

for (( n=1; n<=COUNT; n++ ))
do
  FILE=${PROTO_FILES[i]}
  FOLDER=${PROTO_FOLDERS[i]}

  protoc --proto_path="$FOLDER" --go_out="$FOLDER" --go_opt=paths=source_relative --go-grpc_out="$FOLDER" --go-grpc_opt=paths=source_relative "$FILE"
  echo "--> $FILE"
done
