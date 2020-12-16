#!/bin/sh
# dpw@seattle.local
# 2016.06.12
#

wd=$HOME/golang-sandbox
image=darrylwest/alpine-go

docker run --rm -it \
    -v $wd:/src \
    --name go-tcp-client \
    --link go-service \
    -p 6061:8080 \
    $image /bin/sh
