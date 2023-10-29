package storage

import "check-domain-api/internal/models"

//go:generate go run github.com/vektra/mockery/v2@v2.34.2  --all

type DomainRepository interface {
	Save(models.Domain) error
	GetByName(name string) (models.Domain, error)
	GetMin() (models.Domain, error)
	GetMax() (models.Domain, error)
}
