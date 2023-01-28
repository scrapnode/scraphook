#!/bin/bash
set -e

REPOS=("git@github.com:scrapnode/scrapcore.git")
PACKAGES=("github.com/scrapnode/scrapcore")
COUNT=${#REPOS[@]}

for (( n=1; n<=COUNT; n++ ))
do
    REPO=${REPOS[i]}
    latest=$(git ls-remote "${REPO}" refs/heads/master | cut -f 1)

    PACKAGE=${PACKAGES[i]}
    echo "${PACKAGE}@${latest}"

    go get "${PACKAGE}@${latest}"
    go mod tidy
    
    git add go.mod go.sum
    git commit -m "chore: bump ${REPO} to ${latest}"
done