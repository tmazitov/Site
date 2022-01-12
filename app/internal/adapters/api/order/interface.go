package order

import "site/internal/domain/models"

type Service interface {
	Create(*models.Order) error
	Delete(orderId int) error
	Get(orderId int) (*models.Order, error)
	Complite(orderId int) error
	Update(orderId int) error
}
