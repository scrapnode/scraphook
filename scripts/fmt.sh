#!/bin/bash

set -e

## this will retrieve all of the .go files that have been
## changed since the last commit
STAGED_GO_FILES=$(git diff --cached --name-only -- '*.go')

## we can check to see if this is empty
if [[ $STAGED_GO_FILES != "" ]]; then
    for file in $STAGED_GO_FILES; do
        if test -f "file"; then
          go fmt "$file"
          git add "$file"
        fi
    done
fi
