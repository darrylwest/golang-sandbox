#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-11-15 10:09:40
#

port=4443
echo "listen on port $port"

# gobalance https -bind ":$port" -cert ../certs/example.crt -key ../certs/example.key localhost:9500 localhost:9501 localhost:9502 localhost:9503
gobalance https -bind ":$port" -cert ../certs/example.crt -key ../certs/example.key localhost:80

