# go application for redis caching

## Prerequisites

* install go

* install redis  q
  make sure redis server is running on default port 6379  
  run command to check: ``redis-cli`` .  
  If you wish to run it on separate port change the configuration at [cache.go](cache/cache.go#L31)

## Steps to build

**go build** will build and create an executable redis-caching in the root folder  
It will also download the dependencies for the first time

## Steps to run

needs one of the three commands (get, set, delete)

* get command
  * needs flag --key
  
  ``./redis-caching get --key-=key``

* set command
  * needs flag --key
  * optional flag --value **default value=""**  
  * optional flag --ex **default ex=0**

  ``./redis-caching  --key-=key --value=value --ex=ex``

* delete command
  * needs flag --key
  
  ``./redis-caching delete --key-=key``
