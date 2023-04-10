package users

import (
	"belajar-go-echo/businesses/users"
	"belajar-go-echo/controllers"
	"belajar-go-echo/controllers/users/request"
	"belajar-go-echo/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authUseCase users.Usecase
}

func NewAuthController(authUC users.Usecase) *AuthController {
	return &AuthController{
		authUseCase: authUC,
	}
}

func (ctrl *AuthController) CreateUser(c echo.Context) error {
	userInput := request.User{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	user, err := ctrl.authUseCase.CreateUser(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "error when inserting data", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user created", response.FromDomain(user))
}

func (ctrl *AuthController) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()

	usersData, err := ctrl.authUseCase.GetAllUsers(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	users := []response.User{}

	for _, user := range usersData {
		users = append(users, response.FromDomain(user))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get all users", users)
}

func (ctrl *AuthController) Login(c echo.Context) error {
	userInput := request.User{}
	ctx := c.Request().Context()

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	err := userInput.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	token, err := ctrl.authUseCase.Login(ctx, userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to login", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "login", token)
}