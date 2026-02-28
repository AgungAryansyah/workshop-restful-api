package usecase

import (
	"context"
	"time"
	"workshop-restful-api-backend/internal/entity"
	"workshop-restful-api-backend/internal/model"
	"workshop-restful-api-backend/internal/repository"

	"github.com/google/uuid"
)

type IRestaurantUsecase interface {
	CreateRestaurant(ctx context.Context, createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error)
	GetRestaurants(ctx context.Context) ([]model.RestaurantResponse, error)
	DeleteRestaurants(ctx context.Context, id uuid.UUID) error
}

type RestaurantUsecase struct {
	restaurantRepository repository.IRestaurantRepository
}

func NewRestaurantUsecase(restaurantRepository repository.IRestaurantRepository) *RestaurantUsecase {
	return &RestaurantUsecase{restaurantRepository}
}

func (r *RestaurantUsecase) CreateRestaurant(ctx context.Context, createRestaurant model.CreateRestaurant) (*model.RestaurantResponse, error) {
	restaurant := entity.Restaurant{
		Id:        uuid.New(),
		Name:      createRestaurant.Name,
		Location:  createRestaurant.Location,
		CreatedAt: time.Now(),
	}

	err := r.restaurantRepository.CreateRestaurant(ctx, restaurant)
	if err != nil {
		return nil, err
	}

	response := model.ToRestourantResponse(restaurant)
	return &response, nil
}

func (r *RestaurantUsecase) GetRestaurants(ctx context.Context) ([]model.RestaurantResponse, error) {
	restaurants, err := r.restaurantRepository.GetRestaurants(ctx)
	if err != nil {
		return nil, err
	}

	responses := model.ToRestaurantResponses(restaurants)
	return responses, nil
}

func (r *RestaurantUsecase) DeleteRestaurants(ctx context.Context, id uuid.UUID) error {
	return r.restaurantRepository.DeleteRestaurants(ctx, id)
}
