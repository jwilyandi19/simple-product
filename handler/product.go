package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
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

type UpdateProductRequest struct {
	ProductName      string `json:"product_name"`
	ProductPrice     int    `json:"product_price"`
	ProductExpiredAt string `json:"product_expired_at"`
}

type ProductResponse struct {
	ProductID        int    `json:"product_id"`
	ProductName      string `json:"product_name"`
	ProductPrice     int    `json:"product_price"`
	ProductExpiredAt string `json:"product_expired_at"`
}

type ProductDetailResponse struct {
	ProductID        int    `json:"product_id"`
	ProductName      string `json:"product_name"`
	ProductPrice     int    `json:"product_price"`
	ProductExpiredAt string `json:"product_expired_at"`
	ProductCreatedAt string `json:"product_created_at"`
	ProductUpdatedAt string `json:"product_updated_at"`
}

func NewProductHandler(e *echo.Group, product product.ProductUsecase) {
	handler := &productHandler{
		productUsecase: product,
	}
	e.GET("/", handler.GetProducts)
	e.POST("/create", handler.CreateProduct)
	e.GET("/:id", handler.GetProductDetail)
	e.PUT("/update/:id", handler.UpdateProduct)
	e.DELETE("/delete/:id", handler.DeleteProduct)
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
			ProductID:        product.ID,
			ProductName:      product.Name,
			ProductPrice:     product.Price,
			ProductExpiredAt: product.ExpiredAt.Format("2006-01-02 15:04:05"),
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

func (h *productHandler) GetProductDetail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	newCtx := ctx.Request().Context()

	product, err := h.productUsecase.GetDetailProduct(newCtx, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	productResponse := ProductDetailResponse{
		ProductID:        product.ID,
		ProductName:      product.Name,
		ProductPrice:     product.Price,
		ProductExpiredAt: product.ExpiredAt.Format("2006-01-02 15:04:05"),
		ProductCreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
		ProductUpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return ctx.JSON(http.StatusOK, productResponse)
}

func (h *productHandler) UpdateProduct(ctx echo.Context) error {
	newCtx := ctx.Request().Context()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	var req UpdateProductRequest
	err = json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ResponseError{
			Error: err.Error(),
		})
	}

	expiredAt, _ := time.Parse("2006-01-02 15:04:05", req.ProductExpiredAt)

	arg := domain.UpdateProductRequest{
		ProductID:        id,
		ProductName:      req.ProductName,
		ProductPrice:     req.ProductPrice,
		ProductExpiredAt: expiredAt,
	}

	updated, err := h.productUsecase.UpdateProduct(newCtx, arg)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updated)

}

func (h *productHandler) DeleteProduct(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	newCtx := ctx.Request().Context()

	deleted, err := h.productUsecase.DeleteProduct(newCtx, id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ResponseError{
			Error: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, deleted)
}
