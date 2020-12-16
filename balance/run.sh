#!/bin/sh
# darwest@ebay.com <darryl.west>
# 2018-07-25 07:20:17
#

image=ebayhub/apache-webserver:latest

for i in 0 1 2 3
do
    port="950$i"
    name="apache-server-$i"

    docker run -d --name $name --network=service-net -p $port:80 $image

    sleep 2
    docker exec -it $name /bin/sh -c /update-id.sh

    # echo "open http://localhost:$port"
done


