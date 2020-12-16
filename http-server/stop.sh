#!/bin/sh
# dpw@seattle.local
# 2015.10.10
#

echo "killing servers..."

for f in *.pid
do
    kill -9 `cat $f`
done

echo "servers killed..."
