package route

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"go-test/cmd/router/generator/template/dao"
	"go-test/cmd/router/model"
)

type __getter interface {
	getterRouter(g *echo.Group)
}

type getter[DM model.DatabaseModel, REQ any]struct{}

func (getter[DM, REQ]) getterRouter(g *echo.Group) {
	var _dm DM
	var _req REQ
	name := _dm.ModelName()
	class := _dm.ModelClass()
	_ = copyContent[REQ, DM](&_req)
	g.POST(fmt.Sprintf("/%s/%s/get", class, name), func(c echo.Context) error {
		req := new(REQ)
		m := copyContent[REQ, DM](req)
		rsp := dao.Get(m)
		return c.JSON(http.StatusOK, rsp)
	})
}
