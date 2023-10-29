package service

import (
	"check-domain-api/internal/models"
)

type DomainService interface {
	Save(domain models.Domain) error
	GetByName(name string) (models.Domain, error)
	GetMin() (models.Domain, error)
	GetMax() (models.Domain, error)
}

type Domain struct {
	repo DomainService
}

func NewDomain(repo DomainService) *Domain {
	return &Domain{
		repo: repo,
	}
}

func (d *Domain) Save(dom models.Domain) error {
	return d.repo.Save(dom)
}

func (d *Domain) GetByName(name string) (models.Domain, error) {
	return d.repo.GetByName(name)
}

func (d *Domain) GetMin() (models.Domain, error) {
	return d.repo.GetMin()

}
func (d *Domain) GetMax() (models.Domain, error) {
	return d.repo.GetMax()

}
