#!/bin/sh
# dpw@alameda.local
# 2017.09.27
#

# docker exec -it rabbitmq rabbitmqctl list_exchanges
docker exec -it rabbitmq rabbitmqctl list_queues

