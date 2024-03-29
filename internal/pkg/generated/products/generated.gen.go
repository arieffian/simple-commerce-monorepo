// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package generated

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for CreateNewProductRequestStatus.
const (
	CreateNewProductRequestStatusActive   CreateNewProductRequestStatus = "active"
	CreateNewProductRequestStatusDisabled CreateNewProductRequestStatus = "disabled"
)

// Defines values for ProductStatus.
const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusDisabled ProductStatus = "disabled"
)

// Defines values for UpdateProductByIdRequestStatus.
const (
	Active   UpdateProductByIdRequestStatus = "active"
	Disabled UpdateProductByIdRequestStatus = "disabled"
)

// CreateNewProductRequest defines model for CreateNewProductRequest.
type CreateNewProductRequest struct {
	Creator     string                        `json:"creator"`
	Description string                        `json:"description"`
	Name        string                        `json:"name"`
	Price       int32                         `json:"price"`
	Qty         int32                         `json:"qty"`
	Sku         string                        `json:"sku"`
	Status      CreateNewProductRequestStatus `json:"status"`
}

// CreateNewProductRequestStatus defines model for CreateNewProductRequest.Status.
type CreateNewProductRequestStatus string

// CreateNewProductResponse defines model for CreateNewProductResponse.
type CreateNewProductResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// DeleteProductByIdResponse defines model for DeleteProductByIdResponse.
type DeleteProductByIdResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// ErrorBadRequest defines model for ErrorBadRequest.
type ErrorBadRequest struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorInternalServer defines model for ErrorInternalServer.
type ErrorInternalServer struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorUnauthorized defines model for ErrorUnauthorized.
type ErrorUnauthorized struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorUnexpected defines model for ErrorUnexpected.
type ErrorUnexpected struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// GetProductByIdResponse defines model for GetProductByIdResponse.
type GetProductByIdResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// GetProductBySKUResponse defines model for GetProductBySKUResponse.
type GetProductBySKUResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// GetProductBySlugResponse defines model for GetProductBySlugResponse.
type GetProductBySlugResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// GetProductsResponse defines model for GetProductsResponse.
type GetProductsResponse struct {
	Code    int32      `json:"code"`
	Data    *[]Product `json:"data,omitempty"`
	Message string     `json:"message"`
}

// Product defines model for Product.
type Product struct {
	Creator     string        `json:"creator"`
	Description string        `json:"description"`
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Price       int32         `json:"price"`
	Qty         int32         `json:"qty"`
	Sku         string        `json:"sku"`
	Slug        string        `json:"slug"`
	Status      ProductStatus `json:"status"`
	Updater     string        `json:"updater"`
}

// ProductStatus defines model for Product.Status.
type ProductStatus string

// UpdateProductByIdRequest defines model for UpdateProductByIdRequest.
type UpdateProductByIdRequest struct {
	Creator     string                         `json:"creator"`
	Description string                         `json:"description"`
	Name        string                         `json:"name"`
	Price       int32                          `json:"price"`
	Qty         int32                          `json:"qty"`
	Sku         string                         `json:"sku"`
	Status      UpdateProductByIdRequestStatus `json:"status"`
}

// UpdateProductByIdRequestStatus defines model for UpdateProductByIdRequest.Status.
type UpdateProductByIdRequestStatus string

// UpdateProductByIdResponse defines model for UpdateProductByIdResponse.
type UpdateProductByIdResponse struct {
	Code    int32    `json:"code"`
	Data    *Product `json:"data,omitempty"`
	Message string   `json:"message"`
}

// BadRequestError defines model for BadRequestError.
type BadRequestError = ErrorBadRequest

// InternalServerError defines model for InternalServerError.
type InternalServerError = ErrorInternalServer

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError = ErrorUnauthorized

// UnexpectedError defines model for UnexpectedError.
type UnexpectedError = ErrorUnexpected

// CreateNewProductJSONBody defines parameters for CreateNewProduct.
type CreateNewProductJSONBody = CreateNewProductRequest

// UpdateProductByIdJSONBody defines parameters for UpdateProductById.
type UpdateProductByIdJSONBody = UpdateProductByIdRequest

// CreateNewProductJSONRequestBody defines body for CreateNewProduct for application/json ContentType.
type CreateNewProductJSONRequestBody = CreateNewProductJSONBody

// UpdateProductByIdJSONRequestBody defines body for UpdateProductById for application/json ContentType.
type UpdateProductByIdJSONRequestBody = UpdateProductByIdJSONBody
