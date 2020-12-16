#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-11-15 09:53:40
#

port=4000
echo "listen on port $port"

gobalance tcp -bind ":$port" localhost:9500 localhost:9501 localhost:9502 localhost:9503

