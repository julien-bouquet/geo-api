package usecases

import "github.com/geo-api/domain"

type PointRepository interface {
	GetAll() (domain.Points, error)
	Add(domain.Point) error
}
