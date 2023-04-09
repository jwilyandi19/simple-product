package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	domain "github.com/jwilyandi19/simple-product/domain/order"
	"github.com/jwilyandi19/simple-product/usecase/order"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type orderHandler struct {
	orderUsecase order.OrderUsecase
}

type CreateOrderRequest struct {
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	Descriptions string `json:"descriptions"`
}

type UpdateOrderRequest struct {
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	Descriptions string `json:"descriptions"`
}

type OrderResponse struct {
	OrderID      int    `json:"order_id"`
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	Descriptions string `json:"descriptions"`
}

type OrderDetailResponse struct {
	OrderID        int    `json:"order_id"`
	UserID         int    `json:"user_id"`
	ProductID      int    `json:"product_id"`
	Descriptions   string `json:"descriptions"`
	OrderCreatedAt string `json:"order_created_at"`
	OrderUpdatedAt string `json:"order_updated_at"`
}

func NewOrderHandler(e *echo.Group, order order.OrderUsecase) {
	handler := &orderHandler{
		orderUsecase: order,
	}
	e.GET("/", handler.GetOrders)
	e.POST("/create", handler.CreateOrder)
	e.GET("/:id", handler.GetOrderDetail)
	e.PUT("/update/:id", handler.UpdateOrder)
}

func (h *orderHandler) GetOrders(ctx echo.Context) error {
	newCtx := ctx.Request().Context()

	arg := domain.GetOrderRequest{}
	orders, err := h.orderUsecase.GetOrders(newCtx, arg)

	if err != nil {
		log.Errorf("[GetOrders-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	datas := make([]OrderResponse, 0)
	for _, order := range orders {
		datas = append(datas, OrderResponse{
			OrderID:      order.ID,
			UserID:       order.UserId,
			ProductID:    order.OrderItemId,
			Descriptions: order.Descriptions,
		})
	}

	return ctx.JSON(http.StatusOK, datas)
}

func (h *orderHandler) CreateOrder(ctx echo.Context) error {
	var req CreateOrderRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		log.Errorf("[CreateOrder-Handler] failed to decode: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}
	newCtx := ctx.Request().Context()

	arg := domain.CreateOrderRequest{
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		Descriptions: req.Descriptions,
	}

	created, err := h.orderUsecase.CreateOrder(newCtx, arg)

	if err != nil {
		log.Errorf("[CreateOrder-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, created)
}

func (h *orderHandler) GetOrderDetail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("[GetOrderDetail-Handler] can't get ID: %s", err.Error())
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	newCtx := ctx.Request().Context()

	order, err := h.orderUsecase.GetDetailOrder(newCtx, id)

	if err != nil {
		log.Errorf("[GetOrderDetail-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	orderResponse := OrderDetailResponse{
		OrderID:        order.ID,
		UserID:         order.UserId,
		ProductID:      order.OrderItemId,
		Descriptions:   order.Descriptions,
		OrderCreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		OrderUpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.JSON(http.StatusOK, orderResponse)
}

func (h *orderHandler) UpdateOrder(ctx echo.Context) error {
	newCtx := ctx.Request().Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Errorf("[UpdateOrder-Handler] can't get ID: %s", err.Error())
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	var req UpdateOrderRequest
	err = json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		log.Errorf("[UpdateOrder-Handler] failed to decode: %s", err.Error())
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}

	arg := domain.UpdateOrderRequest{
		OrderID:      id,
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		Descriptions: req.Descriptions,
	}

	updated, err := h.orderUsecase.UpdateOrder(newCtx, arg)

	if err != nil {
		log.Errorf("[UpdateOrder-Handler] %s", err.Error())
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updated)

}
