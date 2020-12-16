#!/bin/sh
# dpw@oakland.localdomain
# 2018.04.30
#

net=service-net

if docker network inspect $net > /dev/null
then
    echo "network $net is up..."
else
    docker network create service-net
fi

docker run --name nats-1 -d -p 0.0.0.0:7222:7222 -p 0.0.0.0:7244:7244 --network=$net -v $(PWD)/gnatsd-A.conf:/tmp/cluster.conf nats -c /tmp/cluster.conf -p 7222 -D -V
docker run --name nats-2 -d -p 0.0.0.0:7223:7222 -p 0.0.0.0:7245:7244 --network=$net -v $(PWD)/gnatsd-B.conf:/tmp/cluster.conf nats -c /tmp/cluster.conf -p 7222 -D -V
docker run --name nats-3 -d -p 0.0.0.0:7224:7222 -p 0.0.0.0:7246:7244 --network=$net -v $(PWD)/gnatsd-C.conf:/tmp/cluster.conf nats -c /tmp/cluster.conf -p 7222 -D -V
