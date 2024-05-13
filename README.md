# aws-ip
list public IP of aws service

## rds

List IPs of RDS databases

```
aws-ip rds
    
NAME      ENDPOINT                                           PORT  IP        VPC           SUBNETS
db_test   db-test.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com   3306  10.0.0.0  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
db_test2  db-test2.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.1  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
db_test3  db-test3.xxxxxxxxxxxx.us-west-2.rds.amazonaws.com  3306  10.0.0.2  vpc-xxxxxxxx  subnet-xxxxxxxx, subnet-xxxxxxxx, subnet-xxxxxxxx
```
