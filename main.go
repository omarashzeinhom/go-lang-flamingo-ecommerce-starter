package main

import (
	"helloworld/helloworld"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/core/zap"
)

func main() {
	flamingo.App([]dingo.Module{
		new(zap.Module),           // the log formatter
		new(requestlogger.Module), // requestlogger shows the request logs
		new(gotemplate.Module),    // enables template engine module for gotemplate
		new(helloworld.Module),
	})
}
