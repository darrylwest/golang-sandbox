# Nats Notes

## Overview

## Backups

Trigger backups to all client machines in a round robin fashion with redundant messages to check completion.  Backup multiple times per day, essentially whenever a file changes or is added.  

* Add a second backup hub.
* run slow-backups to S3 or digital ocean (s3cmd)

## File-Sync

Distribute latest versions of watched files to all clients.

## Monitoring

Insure that client applications are running, e.g., docker and it's containers, proxy apps, backups, etc.

### Dashboard UI

Create a dashboard to show status and activity on all clients.

### Machine List

* seattle: docker, nats hub, erlang node
* alameda: docker, nats hub, erlang node
* oakland: docker, nats hub, erlang node

* marin: nats subscriber

* digital ocean: envoy edge proxy, nats hub (native)

###### darryl.west | 2018.04.30
