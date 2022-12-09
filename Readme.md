# get-nodeid

## Build
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/get-nodeid main.go
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
or generate 
```shell
./bin/get-nodeid --generate --key 1.key --cert 1.crt
```
