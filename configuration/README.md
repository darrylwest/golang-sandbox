# Configuration : Consul & Vault

## Consul Install

It's easiest to run consul from a docker container like this:

`docker run -d -p 8500:8500 --name=dev-consul consul`

### Local Install

Sometimes it's easier to just install consul for local development and test.

download from the site
* unzip
* mv to /usr/local/bin/

## Vault Install

`docker run --cap-add=IPC_LOCK -d -p 8200:8200 --name=dev-vault vault`

`docker run --cap-add=IPC_LOCK -e 'VAULT_DEV_ROOT_TOKEN_ID=ulid' -p 8200:8200 vault`

### Local Install

* download from the site
* unzip
* mv to /usr/local/bin/


## Go API

`go get github.com/hashicorp/consul/api`





## TOML

* [TOML Starter](https://npf.io/2014/08/intro-to-toml/)
* [TOML github reference](https://npf.io/2014/08/intro-to-toml/)
* [The TOML Spec](https://github.com/toml-lang/toml)

###### darryl.west | Version 2017.04.29