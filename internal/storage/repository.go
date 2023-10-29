package storage

import (
	"check-domain-api/internal/models"
	"database/sql"
	"errors"
)

var (
	ErrDomainNotFound = errors.New("Domain not found")
)

type Domain struct {
	db *sql.DB
}

func NewDomain(db *sql.DB) *Domain {
	return &Domain{db: db}
}

func (d *Domain) Save(dom models.Domain) error {
	_, err := d.db.Exec("INSERT INTO domain(domain,latency,last_update, available) VALUES ($1, $2, $3, $4)", dom.Domain, dom.Latency, dom.Last_update, dom.Available)
	return err

}

func (d *Domain) GetByName(name string) (models.Domain, error) {
	var domain models.Domain
	err := d.db.QueryRow("SELECT * FROM domain WHERE domain=$1", name).Scan(&domain.ID, &domain.Domain, &domain.Latency, &domain.Available, &domain.Last_update)
	if err == sql.ErrNoRows {
		return domain, ErrDomainNotFound
	}
	return domain, err
}

func (d *Domain) GetMin() (models.Domain, error) {

	var domain models.Domain

	err := d.db.QueryRow("SELECT * FROM domain ORDER BY latency ASC").Scan(&domain.ID, &domain.Domain, &domain.Latency, &domain.Available, &domain.Last_update)
	if err == sql.ErrNoRows {
		return domain, ErrDomainNotFound
	}
	return domain, err
}
func (d *Domain) GetMax() (models.Domain, error) {

	var domain models.Domain

	err := d.db.QueryRow("SELECT * FROM domain ORDER BY latency DESC").Scan(&domain.ID, &domain.Domain, &domain.Latency, &domain.Available, &domain.Last_update)
	if err == sql.ErrNoRows {
		return domain, ErrDomainNotFound
	}
	return domain, err
}
