#!/bin/sh
# dpw@seattle.local
# 2015.10.10
#

nohup ./http-server :5001 &
pid1=$!
echo $pid1 > server01.pid

exit 0

nohup ./http-server :5002 &
pid2=$!
echo $pid2 > server02.pid

echo "$pid1 $pid2"


