package usecases

import "github.com/coast-nav-api/domain"

type PointRepository interface {
	GetAll() (domain.Points, error)
	Add(domain.Point) error
}
