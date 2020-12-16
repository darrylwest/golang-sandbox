#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-11-15 10:05:13
#

name=elixirdev.tk
cert=$name.crt
key=$name.key

openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out $cert -keyout $key

chmod 400 $key

