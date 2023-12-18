package controller

import (
	"github.com/labstack/echo/v4"

	"go-web-starter/model"
	"go-web-starter/view/user"
)

type UserController struct {
}

func (cr *UserController) ShowDashboard(c echo.Context) error {
    u := model.User{
        Email: "test@test.ch",
    }
    return render(c, user.ShowDashboard(u))
}
