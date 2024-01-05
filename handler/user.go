package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/todo-got/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
