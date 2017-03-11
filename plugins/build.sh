#!/bin/bash

set -e

for dir in $(find . -type d)
do
    [ "$dir" = "." ] && continue

    echo "building $dir"
    go build -buildmode=plugin "./$dir" &
done

wait
