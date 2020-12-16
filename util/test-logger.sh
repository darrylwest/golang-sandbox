#!/bin/sh
# dpw@MacBook-Pro.localdomain
# 2018.08.30
#

COUNTER=0
while [  $COUNTER -lt 10 ]; do
    echo The counter is $COUNTER
    sleep 1
    let COUNTER=COUNTER+1 
done
