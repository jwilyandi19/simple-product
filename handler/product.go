package handler

import (
	"encoding/json"
	"net/http"
	"time"

	domain "github.com/jwilyandi19/simple-product/domain/product"
	"github.com/jwilyandi19/simple-product/usecase/product"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUsecase product.ProductUsecase
}

type ResponseError struct {
	Error string `json:"error"`
}

type CreateProductRequest struct {
	ProductName      string `json:"product_name"`
	ProductPrice     int    `json:"product_price"`
	ProductExpiredAt string `json:"product_expired_at"`
}

type ProductResponse struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
}

func NewProductHandler(e *echo.Group, product product.ProductUsecase) {
	handler := &productHandler{
		productUsecase: product,
	}
	e.GET("/", handler.GetProducts)
	e.POST("/create", handler.CreateProduct)
}

func (h *productHandler) GetProducts(ctx echo.Context) error {
	newCtx := ctx.Request().Context()

	arg := domain.GetProductRequest{}
	products, err := h.productUsecase.GetProducts(newCtx, arg)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	datas := make([]ProductResponse, 0)
	for _, product := range products {
		datas = append(datas, ProductResponse{
			ProductID:    product.ID,
			ProductName:  product.Name,
			ProductPrice: product.Price,
		})
	}

	return ctx.JSON(http.StatusOK, datas)
}

func (h *productHandler) CreateProduct(ctx echo.Context) error {
	var req CreateProductRequest
	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}
	newCtx := ctx.Request().Context()
	expiredAt, _ := time.Parse("2006-01-02 15:04:05", req.ProductExpiredAt)

	arg := domain.CreateProductRequest{
		ProductName:      req.ProductName,
		ProductPrice:     req.ProductPrice,
		ProductExpiredAt: expiredAt,
	}

	created, err := h.productUsecase.CreateProduct(newCtx, arg)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, created)
}
