package order

import "site/internal/domain/models"

type Storage interface {
	Create(*models.Order) error
	Delete(UUID string) error
	Complite(UUID string) error
	Update(UUID string, data map[string]string) error
	Get(UUID string) (*models.Order, error)
	List(params map[string]string) ([]*models.Order, error)
}
