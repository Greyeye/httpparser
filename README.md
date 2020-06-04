
# httpwrapper (work in progress)
remote http get wrapper

[![Build Status](https://travis-ci.org/Greyeye/httpwrapper.svg?branch=master)](https://travis-ci.org/Greyeye/httpwrapper) - Master  
[![Build Status](https://travis-ci.org/Greyeye/httpwrapper.svg?branch=development)](https://travis-ci.org/Greyeye/httpwrapper) - Development  

[![Go Report Card](https://goreportcard.com/badge/github.com/Greyeye/httpwrapper)](https://goreportcard.com/report/github.com/Greyeye/httpwrapper)  

## TODO
two timeout
context timeout (http call should be canned if parent context is cancelled)
http timeout (session call)

- [ ] implement ctxHttp.Do (context aware http call, will return ctx.Err() when timeout occurs)
- [ ] implement transport timeout (http.Client.transport)

