#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2017-09-29 09:49:18
#

image="nats:latest"
docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 $image

