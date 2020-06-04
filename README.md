# httpwrapper (work in progress)
remote http get wrapper

[![Build Status](https://travis-ci.org/Greyeye/httpwrapper.svg?branch=master)](https://travis-ci.org/Greyeye/httpwrapper) - Master  
[![Build Status](https://travis-ci.org/Greyeye/httpwrapper.svg?branch=development)](https://travis-ci.org/Greyeye/httpwrapper) - Development  

[![Go Report Card](https://goreportcard.com/badge/github.com/Greyeye/httpwrapper)](https://goreportcard.com/report/github.com/Greyeye/httpwrapper)  

## TODO
two timeouts
context timeout (http call should be canned if parent context is cancelled)
http timeout (session call)  
POST/PUT/OPTION


- [ ] implement ctxHttp.Do (context aware http call, will return ctx.Err() when timeout occurs)
- [ ] implement transport timeout (http.Client.transport)
- [ ] implenent POST, PUT, and multi-part upload
- [ ] implement OPTION
- [ ] implement better JSON validation
- [ ] custom header

## Install  

Supports Go 1.12.17 - 1.14.1

### Using go modules (aka. `go mod`)

In your go files, simply use:
``` go
import "github.com/Greyeye/httpwrapper"
```

Then next `go mod tidy` or `go test` invocation will automatically
populate your `go.mod` with the last httpwrapper release.  

### Using `$GOPATH`

```shell script
go get github.com/Greyeye/httpwrapper
```

automatically downloads to `$GOPATH/src`. Then in your
go files use:
```go
import "github.com/Greyeye/httpwrapper"
```

