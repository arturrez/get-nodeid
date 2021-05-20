# get-nodeid

## Build
```shell
go build -o bin/get-nodeid
```
## Test
* generate test staking cert
```shell
 openssl req -x509 -nodes -newkey rsa:4096 -keyout test/staking.key -out test/staking.crt -days 3650
```
* Run 
```shell
./bin/get-nodeid -key=./test/staking.key -cert=./test/staking.crt
```
