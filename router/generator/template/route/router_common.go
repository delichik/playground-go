package route

import "github.com/labstack/echo/v4"

func InitRouter(g *echo.Group) {
	for _, r := range __routeIncubators {
		__loadCommonRouter(g, r)
		if t, ok := r.(__customRouterHandler); ok {
			t.customRouters(g)
		}
	}
}

type __customRouterHandler interface {
	customRouters(g *echo.Group)
}

var __routeIncubators []any

func registerRouteIncubator(r any) {
	__routeIncubators = append(__routeIncubators, r)
}

func __loadCommonRouter(g *echo.Group, router any) {
	if t, ok := router.(__getter); ok {
		t.getterRouter(g)
	}
}
