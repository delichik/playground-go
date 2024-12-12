package route

import (
	"github.com/guregu/null/v5"
	"github.com/labstack/echo/v4"
	"github.com/pion/webrtc/v3/pkg/null"

	"router/model"
)

func init() {
	registerRouteIncubator(algAbility{})
}

type getAlgAbilityReq struct {
	Id null.Int `json:"id"`
}

type algAbility struct {
	getter[model.AlgAbility, getAlgAbilityReq]
}

func (algAbility) routers(g *echo.Group) {}
