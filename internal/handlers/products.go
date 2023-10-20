package handlers

import (
	"errors"
	"strconv"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	generated "github.com/arieffian/simple-commerces-monorepo/internal/pkg/generated/products"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type productHandler struct {
	productRepo repositories.ProductInterface
	validator   validator.ValidatorService
	cfg         config.Config
}

var _ ProductService = (*productHandler)(nil)

type NewProductHandlerParams struct {
	ProductRepo repositories.ProductInterface
	Validator   validator.ValidatorService
	Cfg         config.Config
}

// // todo: add validation
func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	pPage := c.Params("page")
	page, err := strconv.Atoi(pPage)

	if err != nil {
		status := fiber.StatusBadRequest
		response := generated.GetProductsResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	if page <= 0 {
		page = 1
	}

	productsResult, err := h.productRepo.GetProducts(c.Context(), repositories.GetProductsParams{
		Limit:  PER_PAGE,
		Offset: (page - 1) * PER_PAGE,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		response := generated.GetProductsResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var products []generated.Product
	for _, product := range productsResult.Products {
		products = append(products, generated.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       int32(product.Price),
			Qty:         int32(product.Qty),
			Sku:         product.SKU,
			Slug:        product.Slug,
			Status:      generated.ProductStatus(product.Status),
		})
	}

	response := generated.GetProductsResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    &products,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	params := new(CreateNewProductParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.CreateNewProductResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.CreateNewProductResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	product, err := h.productRepo.CreateNewProduct(c.Context(), repositories.CreateNewProductParams{
		Name:        params.Name,
		SKU:         params.SKU,
		Price:       params.Price,
		Qty:         params.Qty,
		Description: params.Description,
		Status:      params.Status,
		CreatedBy:   params.CreatedBy,
	})

	if err != nil {
		response := &generated.CreateNewProductResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := &generated.CreateNewProductResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) UpdateProductById(c *fiber.Ctx) error {
	params := new(UpdateProductByIdParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.UpdateProductByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	params.ID = c.Params("id")

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.UpdateProductByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	product, err := h.productRepo.UpdateProductById(c.Context(), repositories.UpdateProductByIdParams{
		ID:          params.ID,
		Name:        params.Name,
		SKU:         params.SKU,
		Price:       params.Price,
		Qty:         params.Qty,
		Description: params.Description,
		Status:      params.Status,
		UpdatedBy:   params.UpdatedBy,
	})
	if err != nil {
		response := &generated.UpdateProductByIdResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := &generated.UpdateProductByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) DeleteProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.productRepo.DeleteProductById(c.Context(), repositories.DeleteProductByIdParams{
		ProductId: id,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.DeleteProductByIdResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.DeleteProductByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    nil,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) GetProductById(c *fiber.Ctx) error {
	product, err := h.productRepo.GetProductById(c.Context(), repositories.GetProductByIdParams{
		ProductId: c.Params("id"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetProductByIdResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetProductByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) GetProductBySKU(c *fiber.Ctx) error {
	product, err := h.productRepo.GetProductBySKU(c.Context(), repositories.GetProductBySKUParams{
		ProductSKU: c.Params("sku"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetProductBySKUResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetProductBySKUResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *productHandler) GetProductBySlug(c *fiber.Ctx) error {
	product, err := h.productRepo.GetProductBySlug(c.Context(), repositories.GetProductBySlugParams{
		ProductSlug: c.Params("slug"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetProductBySlugResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	response := generated.GetProductBySlugResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Product{
			Id:          product.Product.ID,
			Name:        product.Product.Name,
			Description: product.Product.Description,
			Price:       int32(product.Product.Price),
			Qty:         int32(product.Product.Qty),
			Sku:         product.Product.SKU,
			Slug:        product.Product.Slug,
			Status:      generated.ProductStatus(product.Product.Status),
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func NewProductHandler(p NewProductHandlerParams) *productHandler {
	return &productHandler{
		productRepo: p.ProductRepo,
		validator:   p.Validator,
		cfg:         p.Cfg,
	}
}
