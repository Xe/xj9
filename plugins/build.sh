#!/bin/bash

set -e
set -x

for dir in $(find . -type d)
do
    [ "$dir" = "." ] && continue
    go build -buildmode=plugin "./$dir" &
done

wait
