package repository

import (
	"context"
	"workshop-restful-api-backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IItemRepository interface {
	GetRestaurantItems(ctx context.Context, restaurantID uuid.UUID) ([]entity.Item, error)
	CreateItem(ctx context.Context, item entity.Item) error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

func (r *ItemRepository) GetRestaurantItems(ctx context.Context, restaurantID uuid.UUID) ([]entity.Item, error) {
	Item, err := gorm.G[entity.Item](r.db).Where("restaurant_id = ?", restaurantID).Find(ctx)
	if err != nil {
		return nil, err
	}

	if len(Item) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return Item, nil
}

func (r *ItemRepository) CreateItem(ctx context.Context, item entity.Item) error {
	err := gorm.G[entity.Item](r.db).Create(ctx, &item)
	if err != nil {
		return err
	}

	return nil
}
