package helloworld

import (
	"context"
	"net/http"
	"strings"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

}

type routes struct{}

func (*routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/", "home")
	registry.HandleAny("home", indexHandler)
}

func indexHandler(ctx context.Context, req *web.Request) web.Result {
	return &web.Response{
		Status: http.StatusOK,
		Body:   strings.NewReader("Hello World"),
	}

}
