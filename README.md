[![go-ubuntu](https://github.com/sinlov/filebrowser-client/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/filebrowser-client/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/filebrowser-client?status.png)](https://godoc.org/github.com/sinlov/filebrowser-client/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/filebrowser-client)](https://goreportcard.com/report/github.com/sinlov/filebrowser-client)
[![codecov](https://codecov.io/gh/sinlov/filebrowser-client/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/filebrowser-client)

## for what

- this project used to support go web client for [filebrowser](https://github.com/filebrowser/filebrowser)

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlov/filebrowser-client.git

# test depends see full version
$ go list -mod=readonly -v -m -versions github.com/sinlov/filebrowser-client
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlov/filebrowser-client | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.17+

# dev

```bash
make init dep
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# see how to use
$ filebrowser-client -h
```
