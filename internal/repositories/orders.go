package repositories

import (
	"context"
	"fmt"

	"github.com/arieffian/simple-commerces-monorepo/internal/config"
	"github.com/arieffian/simple-commerces-monorepo/internal/database"
	"github.com/arieffian/simple-commerces-monorepo/internal/models"
	"github.com/arieffian/simple-commerces-monorepo/internal/pkg/redis"
	redis_pkg "github.com/redis/go-redis/v9"
)

type orderRepository struct {
	db      *database.DbInstance
	redisDb redis.RedisService
	cfg     config.Config
}

var _ OrderInterface = (*orderRepository)(nil)

type NewOrderRepositoryParams struct {
	Db    *database.DbInstance
	Redis redis.RedisService
	Cfg   config.Config
}

// todo: check cache is enabled or not
func (r *orderRepository) GetOrderById(ctx context.Context, p GetOrderByIdParams) (*GetOrderByIdResponse, error) {
	var order models.Order

	cacheKey := fmt.Sprintf("|%s|id|%s|", r.cfg.Service, p.OrderId)

	err := r.redisDb.GetCache(ctx, cacheKey, &order)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		if err := r.db.Db.Preload("OrderDetails").First(&order, "id = ? AND status <> deleted", p.OrderId).Error; err != nil {
			return nil, err
		}
		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, order, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetOrderByIdResponse{
		Order: order,
	}, nil
}

// todo: check cache is enabled or not
func (r *orderRepository) GetOrdersByUserId(ctx context.Context, p GetOrdersByUserIdParams) (*GetOrdersByUserIdResponse, error) {
	var orders []models.Order

	cacheKey := fmt.Sprintf("|%s|user_id|%s|offset|%d|limit|%d|", r.cfg.Service, p.UserId, p.Offset, p.Limit)

	err := r.redisDb.GetCache(ctx, cacheKey, &orders)

	if err != nil {
		if err != redis_pkg.Nil {
			return nil, err
		}

		err := r.db.Db.Model(&models.Order{}).Preload("OrderDetails").Limit(p.Limit).Offset(p.Offset).Find(&orders).Error

		if err != nil {
			return nil, err
		}

		if err := r.redisDb.SetCacheWithExpiration(context.Background(), cacheKey, orders, r.cfg.CacheTTL); err != nil {
			return nil, err
		}
	}

	return &GetOrdersByUserIdResponse{
		Orders: orders,
	}, nil
}

func (r *orderRepository) CreateNewOrder(ctx context.Context, p CreateNewOrderParams) (*CreateNewOrderResponse, error) {
	order := models.Order{
		UserID:     p.UserId,
		Status:     p.Status,
		GrandTotal: p.GrandTotal,
	}

	tx := r.db.Db.Begin()

	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, detail := range p.OrderDetails {
		orderDetail := models.OrderDetail{
			OrderID:   order.ID,
			ProductID: detail.ProductId,
			SubTotal:  detail.SubTotal,
			Qty:       detail.Qty,
			Price:     detail.Price,
		}

		err := tx.Create(&orderDetail).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	err = r.db.Db.Model(&models.Order{}).Preload("OrderDetails").Find(&order).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &CreateNewOrderResponse{
		Order: order,
	}, nil
}

func (r *orderRepository) UpdateOrderById(ctx context.Context, p UpdateOrderByIdParams) (*UpdateOrderByIdResponse, error) {
	order := models.Order{
		ID:     p.ID,
		Status: p.Status,
	}

	err := r.db.Db.Model(&models.Order{}).Updates(&order).Error
	if err != nil {
		return nil, err
	}

	return &UpdateOrderByIdResponse{
		Order: order,
	}, nil
}

func NewOrderRepository(p NewOrderRepositoryParams) *orderRepository {

	return &orderRepository{
		db:      p.Db,
		redisDb: p.Redis,
		cfg:     p.Cfg,
	}
}
