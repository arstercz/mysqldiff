# mysqldiff

Wrap the sys-mysql-diff script to monitor MySQL changes with batch methods.

## Dependency

[sys-mysql-diff](https://github.com/chenzhe07/sys-toolkit#sys-mysql-diff)

## How to build

```
go get github.com/chenzhe07/mysqldiff.git
go build -o mysqldiff *.go
```

## How to use

### privileges
all instances changes can be store in `[backend]` section, `[test3301]` and `[test3306]` can be monitor by `sys-mysql-diff`, any instance that monitored should be with the testxxxx format.

user_mysqlmon user should have the following privileges:
```
grant select,insert,update on mysqldiff.* to user_mysqlmon@`10.0.21.%`;
```
user_mysqldiff user with the following privileges:
```
grant select on *.* to user_mysqldiff@`10.0.21.%` with grant option;
```

#### conf file
sys-mysql-diff monitor all instance when db option is `information_schema`.
```
[backend]
dsn = user_mysqlmon:xxxxxxxx@tcp(10.0.21.17:3306)/mysqldiff?charset=utf8

[test3301]
host = 10.0.21.5
port = 3301
db   = test
user = user_mysqldiff
pass = xxxxxxxx
tag  = host_location

[test3306]
host = 10.0.21.7
port = 3306
db   = percona
user = user_mysqldiff
pass = xxxxxxxx
tag  = host_location
```

### monitor with mysqldiff
```
# ./mysqldiff -conf conf.cnf -verbose
2017/03/15 16:31:27 ---------------------------
changes from 10.0.21.5:3301 
changes from 10.0.21.7:3306 
DROP TABLE `emp`;
SET GLOBAL wait_timeout = 1000;
2017/03/15 16:31:27 insert 10.0.21.17:3306/percona ok
2017/03/15 16:31:27 ---------------------------
```

## License

MIT / BSD

