# aws-ip
[![pipeline](https://github.com/pete911/aws-ip/actions/workflows/pipeline.yml/badge.svg)](https://github.com/pete911/aws-ip/actions/workflows/pipeline.yml)

list public IPs of aws services

## build

`go build` or `go install`

## download

- [binary](https://github.com/pete911/aws-ip/releases)

## build/install

### brew

- add tap `brew tap pete911/tap`
- install `brew install pete911/tap/aws-ip`

## release

Releases are published when the new tag is created e.g.
`git tag -m "<message>" v1.0.0 && git push --follow-tags`

## run

### rds

List RDS IPs

#### optional flags

```
--status      print status
--subnet      print subnet names
--subnet-id   print subnet ids
--vpc         print vpc name
--vpc-id      print vpc id
```

```
aws-ip rds
    
NAME      ENDPOINT                                           PORT  IP
db_test   db-test.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com   3306  10.0.0.0
db_test2  db-test2.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.1
db_test3  db-test3.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.2
```

### elasticache

List elastic cache IPs

#### optional flags

```
--status      print status
--subnet      print subnet names
--subnet-id   print subnet ids
--vpc         print vpc name
--vpc-id      print vpc id
```

```
aws-ip elasticache

ID        ENGINE     VERSION  NODE  NODE STATUS  NODE ENDPOINT                                  NODE PORT  IP
test      memcached  1.4.33   0001  available    test.xxxxxx.0001.usw2.cache.amazonaws.com      11211      10.0.0.0
test-001  redis      7.0.7    0001  available    test-001.xxxxxx.0001.usw2.cache.amazonaws.com  6379       10.0.0.1
test-002  redis      7.0.7    0001  available    test-002.xxxxxx.0001.usw2.cache.amazonaws.com  6379       10.0.0.2
```
