package usecases

import "github.com/geo-api/domain"

type PointInteractor struct {
	PointRepository PointRepository
}

func (mi *PointInteractor) GetAll() (points domain.Points, err error) {
	return mi.PointRepository.GetAll()
}

func (mi *PointInteractor) Add(point domain.Point) (err error) {
	return mi.PointRepository.Add(point)
}
