package order

import (
	"site/internal/adapters/api/order"
	"site/internal/domain/models"
)

type service struct {
	storage Storage
}

func NewService(stor Storage) order.Service {
	s := &service{storage: stor}
	return s
}

func (s *service) Create(order *models.Order) error {
	return s.storage.Create(order)
}

func (s *service) Update(UUID string, data map[string]string) error {
	return s.storage.Update(UUID, data)
}

func (s *service) Delete(UUID string) error {
	return s.storage.Delete(UUID)
}

func (s *service) Complite(UUID string) error {
	return s.storage.Complite(UUID)
}

func (s *service) Get(UUID string) (*models.Order, error) {
	return s.storage.Get(UUID)
}

func (s *service) List(params map[string]string) ([]*models.Order, error) {
	return s.storage.List(params)
}
