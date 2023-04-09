package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	domain "github.com/jwilyandi19/simple-product/domain/user"
	"github.com/jwilyandi19/simple-product/usecase/user"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type userHandler struct {
	userUsecase user.UserUsecase
}

type CreateUserRequest struct {
	UserFullName string `json:"full_name"`
}

type UpdateUserRequest struct {
	UserID       int    `json:"user_id"`
	UserFullName string `json:"full_name"`
}

type UserResponse struct {
	UserID       int    `json:"user_id"`
	UserFullName string `json:"full_name"`
}

type UserDetailResponse struct {
	UserID        int    `json:"user_id"`
	UserFullName  string `json:"full_name"`
	UserCreatedAt string `json:"user_created_at"`
	UserUpdatedAt string `json:"user_updated_at"`
}

func NewUserHandler(e *echo.Group, user user.UserUsecase) {
	handler := &userHandler{
		userUsecase: user,
	}
	e.GET("/", handler.GetUsers)
	e.POST("/create", handler.CreateUser)
	e.GET("/:id", handler.GetUserDetail)
	e.PUT("/update/:id", handler.UpdateUser)
	e.DELETE("/delete/:id", handler.DeleteUser)
}

func (h *userHandler) GetUsers(ctx echo.Context) error {
	newCtx := ctx.Request().Context()

	arg := domain.GetUserRequest{}
	users, err := h.userUsecase.GetUsers(newCtx, arg)

	if err != nil {
		log.Errorf("[GetUsers-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	datas := make([]UserResponse, 0)
	for _, user := range users {
		datas = append(datas, UserResponse{
			UserID:       user.ID,
			UserFullName: user.FullName,
		})
	}

	return ctx.JSON(http.StatusOK, datas)
}

func (h *userHandler) CreateUser(ctx echo.Context) error {
	var req CreateUserRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		log.Errorf("[CreateUser-Handler] failed to decode: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}
	newCtx := ctx.Request().Context()

	arg := domain.CreateUserRequest{
		FullName: req.UserFullName,
	}

	created, err := h.userUsecase.CreateUser(newCtx, arg)

	if err != nil {
		log.Errorf("[CreateUser-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, created)
}

func (h *userHandler) GetUserDetail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("[GetUserDetail-Handler] can't get id: %s", err.Error())
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	newCtx := ctx.Request().Context()

	user, err := h.userUsecase.GetDetailUser(newCtx, id)

	if err != nil {
		log.Errorf("[GetUserDetail-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	userResponse := UserDetailResponse{
		UserID:        user.ID,
		UserFullName:  user.FullName,
		UserCreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UserUpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.JSON(http.StatusOK, userResponse)
}

func (h *userHandler) UpdateUser(ctx echo.Context) error {
	newCtx := ctx.Request().Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("[UpdateUser-Handler] can't get id: %s", err.Error())
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	var req UpdateUserRequest
	err = json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		log.Errorf("[UpdateUser-Handler] failed to decode: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}

	arg := domain.UpdateUserRequest{
		UserID:   id,
		FullName: req.UserFullName,
	}

	updated, err := h.userUsecase.UpdateUser(newCtx, arg)

	if err != nil {
		log.Errorf("[UpdateUser-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updated)

}

func (h *userHandler) DeleteUser(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("[DeleteUser-Handler] can't get id: %s", err.Error())
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	newCtx := ctx.Request().Context()

	deleted, err := h.userUsecase.DeleteUser(newCtx, id)

	if err != nil {
		log.Errorf("[DeleteUser-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, deleted)
}
