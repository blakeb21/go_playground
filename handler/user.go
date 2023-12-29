package handler

import (
	"github.com/blakeb21/go_playground/model"
	"github.com/blakeb21/go_playground/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "blake.barnhill@gmail.com",
	}
	return render(c, user.Show(u))
}
