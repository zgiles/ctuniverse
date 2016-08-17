
# ctuniverse
Space object tracker for CT moon lander  

[![Go Report Card](https://goreportcard.com/badge/github.com/zgiles/ctuniverse)](https://goreportcard.com/report/github.com/zgiles/ctuniverse)
[![Build Status](https://drone.io/github.com/zgiles/ctuniverse/status.png)](https://drone.io/github.com/zgiles/ctuniverse/latest)

## Requirements
* Go  

## What does it do
Provides a way to send and receive object positions (and attributes) for everything in the universe between game clients via websockets
See PROTOCOL.md for protocol description.  
Link for the game soon..  

## Building
Building needs a few pre-req's..  
1) The bin-data needs to be generated so a static directory is not needed.
```
go get github.com/elazarl/go-bindata/...
go get github.com/zgiles/ctuniverse/...
cd $GOPATH/src/github.com/zgiles/ctuniverse/cmd/ctuniverse
go generate

```
2) Build with the hash and date in the main file..  
This app provides an ability to print the git hash it was pulled from and build time, if it is included at build-time.  
The app can also be installed using standard `go get` methods.  
Example for full `go get`:  
```
go get -ldflags "-X main.buildtime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD`" github.com/zgiles/ctuniverse/...
```

Example of just building:  
```
go build -ldflags "-X main.buildtime=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD`" github.com/zgiles/ctuniverse/...
```

## Release History
* 0.1.0 - Initial release  

## License
Copyright 2016 Zachary Giles  
MIT License (Expat)  

Please see the LICENSE file  
