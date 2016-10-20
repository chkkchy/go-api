#!/bin/sh

if [ -z $1 ]; then
    echo "Usage: cmd [api|admin]" 1>&2
    exit 1
fi

go run src/$1/server.go -alsologtostderr -log_dir=./logs

