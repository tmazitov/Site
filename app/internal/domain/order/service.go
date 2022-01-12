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

func (s *service) Update(orderId int) error {
	return s.storage.Update(orderId)
}

func (s *service) Delete(orderId int) error {
	return s.storage.Delete(orderId)
}

func (s *service) Complite(orderId int) error {
	return s.storage.Complite(orderId)
}

func (s *service) Get(orderId int) (*models.Order, error) {
	return s.storage.Get(orderId)
}
