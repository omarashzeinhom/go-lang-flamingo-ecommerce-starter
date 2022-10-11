package main

import (
	"helloworld/helloworld"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
)

func main() {
	flamingo.App([]dingo.Module{
		new(helloworld.Module),
	})
}
