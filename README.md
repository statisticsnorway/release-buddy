# Release buddy
A simple tool to create new (semver) releases.   

## Installation
Requires Go version >= 1.13
```bash
GO111MODULE=on go get github.com/statisticsnorway/release-tool@master
```
You can specify a release tag or commit SHA instead of `master`

The binary will be installed into `$GOPATH/bin` if `$GOPATH` is set, `$HOME/go/bin` otherwise.

## Usage
To release the next patch version, hit enter when prompted, e.g:
```bash
$ release
What is the release version? 1.4.2:
Creating new release: 1.4.2
Done 
```

To release a version different from the next patch, enter the new version when prompted, e.g:
```bash
$ release
What is the release version? 1.4.2: 1.5.0
Creating new release: 1.5.0
Done 
```
