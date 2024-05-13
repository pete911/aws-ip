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

List IPs of RDS databases

```
aws-ip rds
    
NAME      ENDPOINT                                           PORT  IP        VPC           SUBNETS
db_test   db-test.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com   3306  10.0.0.0  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
db_test2  db-test2.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.1  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
db_test3  db-test3.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.2  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
```
