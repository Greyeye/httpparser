# httpwrapper (work in progress)
remote http get wrapper


## TODO
two timeout
context timeout (http call should be canned if parent context is cancelled)
http timeout (session call)

- [ ] implement ctxHttp.Do (context aware http call, will return ctx.Err() when timeout occurs)
- [ ] implement transport timeout (http.Client.transport)

