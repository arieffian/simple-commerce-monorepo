package handlers

import (
	"errors"
	"strconv"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	generated "github.com/arieffian/simple-commerces-monorepo/internal/pkg/generated/orders"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/validator"
	"github.com/arieffian/simple-commerces-monorepo/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type orderHandler struct {
	orderRepo repositories.OrderInterface
	validator validator.ValidatorService
	cfg       config.Config
}

var _ OrderService = (*orderHandler)(nil)

type NewOrderHandlerParams struct {
	OrderRepo repositories.OrderInterface
	Validator validator.ValidatorService
	Cfg       config.Config
}

func (h *orderHandler) GetOrdersByUserId(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	pPage := c.Params("page")
	page, err := strconv.Atoi(pPage)

	if err != nil {
		status := fiber.StatusBadRequest
		response := generated.GetOrdersByUserIdResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	if page <= 0 {
		page = 1
	}

	ordersResult, err := h.orderRepo.GetOrdersByUserId(c.Context(), repositories.GetOrdersByUserIdParams{
		UserId: userId,
		Limit:  PER_PAGE,
		Offset: (page - 1) * PER_PAGE,
	})

	if err != nil {
		status := fiber.StatusInternalServerError
		response := generated.GetOrdersByUserIdResponse{
			Code:    int32(status),
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var orders []generated.Order
	for _, order := range ordersResult.Orders {
		var orderDetails []generated.OrderDetail
		for _, orderDetail := range order.OrderDetails {
			orderDetails = append(orderDetails, generated.OrderDetail{
				Qty:       int32(orderDetail.Qty),
				ProductId: orderDetail.ProductID,
				Price:     int64(orderDetail.Price),
				SubTotal:  int64(orderDetail.SubTotal),
			})
		}

		orders = append(orders, generated.Order{
			Id:           order.ID,
			UserId:       order.UserID,
			GrandTotal:   int64(order.GrandTotal),
			CreatedAt:    order.CreatedAt.String(),
			Status:       generated.OrderStatus(order.Status),
			OrderDetails: orderDetails,
		})
	}

	response := generated.GetOrdersByUserIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data:    &orders,
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *orderHandler) CreateNewOrder(c *fiber.Ctx) error {
	params := new(CreateNewOrderParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.CreateNewOrderResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.CreateNewOrderResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	order, err := h.orderRepo.CreateNewOrder(c.Context(), repositories.CreateNewOrderParams{
		UserId: params.UserId,
	})

	if err != nil {
		response := &generated.CreateNewOrderResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	details := []generated.OrderDetail{}
	for _, orderDetail := range order.Order.OrderDetails {
		details = append(details, generated.OrderDetail{
			Qty:       int32(orderDetail.Qty),
			ProductId: orderDetail.ProductID,
			Price:     int64(orderDetail.Price),
			SubTotal:  int64(orderDetail.SubTotal),
			CreatedAt: orderDetail.CreatedAt.String(),
		})
	}

	response := &generated.CreateNewOrderResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Order{
			CreatedAt:    order.Order.CreatedAt.String(),
			GrandTotal:   int64(order.Order.GrandTotal),
			Id:           order.Order.ID,
			UserId:       order.Order.UserID,
			Status:       generated.OrderStatus(order.Order.Status),
			OrderDetails: details,
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *orderHandler) UpdateOrderById(c *fiber.Ctx) error {
	params := new(UpdateOrderByIdParams)
	err := c.BodyParser(params)
	if err != nil {
		response := &generated.UpdateOrderByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	params.ID = c.Params("id")

	err = h.validator.Validate(c.Context(), params)
	if err != nil {
		response := &generated.UpdateOrderByIdResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	order, err := h.orderRepo.UpdateOrderById(c.Context(), repositories.UpdateOrderByIdParams{
		ID:     params.ID,
		Status: params.Status,
	})
	if err != nil {
		response := &generated.UpdateOrderByIdResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var orderDetails []generated.OrderDetail
	for _, orderDetail := range order.Order.OrderDetails {
		orderDetails = append(orderDetails, generated.OrderDetail{
			Qty:       int32(orderDetail.Qty),
			ProductId: orderDetail.ProductID,
			Price:     int64(orderDetail.Price),
			SubTotal:  int64(orderDetail.SubTotal),
		})
	}

	response := &generated.UpdateOrderByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Order{
			Id:           order.Order.ID,
			UserId:       order.Order.UserID,
			GrandTotal:   int64(order.Order.GrandTotal),
			CreatedAt:    order.Order.CreatedAt.String(),
			Status:       generated.OrderStatus(order.Order.Status),
			OrderDetails: orderDetails,
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func (h *orderHandler) GetOrderById(c *fiber.Ctx) error {
	order, err := h.orderRepo.GetOrderById(c.Context(), repositories.GetOrderByIdParams{
		OrderId: c.Params("id"),
	})

	if err != nil {
		status := fiber.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
		}
		response := generated.GetOrderByIdResponse{
			Code:    int32(status),
			Message: "OK",
			Data:    nil,
		}
		return c.Status(int(response.Code)).JSON(response)
	}

	var orderDetails []generated.OrderDetail
	for _, orderDetail := range order.Order.OrderDetails {
		orderDetails = append(orderDetails, generated.OrderDetail{
			Qty:       int32(orderDetail.Qty),
			ProductId: orderDetail.ProductID,
			Price:     int64(orderDetail.Price),
			SubTotal:  int64(orderDetail.SubTotal),
		})
	}

	response := generated.GetOrderByIdResponse{
		Code:    fiber.StatusOK,
		Message: "OK",
		Data: &generated.Order{
			Id:           order.Order.ID,
			UserId:       order.Order.UserID,
			GrandTotal:   int64(order.Order.GrandTotal),
			CreatedAt:    order.Order.CreatedAt.String(),
			Status:       generated.OrderStatus(order.Order.Status),
			OrderDetails: orderDetails,
		},
	}

	return c.Status(int(response.Code)).JSON(response)
}

func NewOrderHandler(p NewOrderHandlerParams) *orderHandler {
	return &orderHandler{
		orderRepo: p.OrderRepo,
		validator: p.Validator,
		cfg:       p.Cfg,
	}
}
