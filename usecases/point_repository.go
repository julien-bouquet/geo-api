package usecases

import "github.com/geo-api/domain"

type PointRepository interface {
	Get(string) (*domain.Point, error)
	GetAll() (domain.Points, error)
	Add(domain.Point) error
	Delete(domain.Point) error
	Update(domain.Point) error
}
