# HTTPParser (work in progress)
remote http get wrapper

[![Build Status](https://travis-ci.org/Greyeye/httpparser.svg?branch=master)](https://travis-ci.org/Greyeye/httpwrapper) - Master  
[![Build Status](https://travis-ci.org/Greyeye/httpparser.svg?branch=development)](https://travis-ci.org/Greyeye/httpwrapper) - Development  

[![Go Report Card](https://goreportcard.com/badge/github.com/Greyeye/httpwrapper)](https://goreportcard.com/report/github.com/Greyeye/httpwrapper)  

## TODO

- [ ] implement transport timeout (http.Client.transport)
- [ ] implenent POST, PUT, and multi-part upload
- [ ] implement OPTION
- [ ] implement better JSON validation
- [ ] custom header

## Install  

Supports Go 1.13.12 - 1.14.7

### Using go modules (aka. `go mod`)

In your go files, simply use:
``` go
import "github.com/Greyeye/httpparser"
```

Then next `go mod tidy` or `go test` invocation will automatically
populate your `go.mod` with the last httpparser release.  

### Using `$GOPATH`

```shell script
go get github.com/Greyeye/httpparser
```

automatically downloads to `$GOPATH/src`. Then in your
go files use:
```go
import "github.com/Greyeye/httpparser"
```

