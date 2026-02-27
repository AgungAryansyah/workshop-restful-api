package repository

import (
	"workshop-restful-api-backend/internal/entity"

	"gorm.io/gorm"
)

type IRestaurantRepository interface {
	CreateRestaurant(restaurant entity.Restaurant) error
	GetRestaurants() ([]entity.Restaurant, error)
}

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db}
}

func (r *RestaurantRepository) CreateRestaurant(restaurant entity.Restaurant) error {
	result := r.db.Create(&restaurant)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *RestaurantRepository) GetRestaurants() ([]entity.Restaurant, error) {
	restaurants := []entity.Restaurant{}

	result := r.db.Find(&restaurants)
	if result.Error != nil {
		return nil, result.Error
	}

	return restaurants, nil
}
