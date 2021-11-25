package users

import (
	"miniproject/business/users"
	"miniproject/controllers"
	"miniproject/controllers/users/requests"
	"miniproject/controllers/users/responses"
	"miniproject/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController *UserController) Login(c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	request := requests.UserLogin{}
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, token, err = userController.UserUseCase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.UsersLogin(login, token))
}

func (userController *UserController) Register(c echo.Context) error {
	request := requests.UserRegister{}
	var err error
	err = c.Bind(&request)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.UserUseCase.Register(ctx, *request.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromUsersRegister(data))
}

func (userController *UserController) GetAllUsers(c echo.Context) error {
	requests := c.Request().Context()
	user, err := userController.UserUseCase.GetAllUsers(requests)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromUsersAll(user))
}

func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := requests.UserRegister{}
	err = c.Bind(&request)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := userController.UserUseCase.Update(ctx, *request.ToDomain(), convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, responses.FromUsersRegister(data))
}

func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	convId, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	request := c.Request().Context()
	err = userController.UserUseCase.Delete(request, convId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}