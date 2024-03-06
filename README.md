[![ci](https://github.com/sinlov/filebrowser-client/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov/filebrowser-client/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov/filebrowser-client?label=go.mod)](https://github.com/sinlov/filebrowser-client)
[![GoDoc](https://godoc.org/github.com/sinlov/filebrowser-client?status.png)](https://godoc.org/github.com/sinlov/filebrowser-client)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov/filebrowser-client)](https://goreportcard.com/report/github.com/sinlov/filebrowser-client)

[![GitHub license](https://img.shields.io/github/license/sinlov/filebrowser-client)](https://github.com/sinlov/filebrowser-client)
[![codecov](https://codecov.io/gh/sinlov/filebrowser-client/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/filebrowser-client)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov/filebrowser-client)](https://github.com/sinlov/filebrowser-client/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov/filebrowser-client)](https://github.com/sinlov/filebrowser-client/releases)

## for what

- this project used to support go web client for [filebrowser](https://github.com/filebrowser/filebrowser)

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov/filebrowser-client)](https://github.com/sinlov/filebrowser-client/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

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
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov/filebrowser-client | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## env

- minimum go version: go 1.17
- change `go 1.17`, `^1.17`, `1.17.13` to new go version

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
