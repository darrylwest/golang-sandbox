#!/bin/sh
# dpw@alameda.local
# 2017.12.23
#

(PORT=2017 ~/.gopath/bin/wiki  -db $PWD/wiki.db > $PWD/wiki.log 2>&1 ) &
PID=$!

echo $PID > ./wiki.pid

