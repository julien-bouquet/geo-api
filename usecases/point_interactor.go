package usecases

import "github.com/geo-api/domain"

type PointInteractor struct {
	PointRepository PointRepository
}

func (mi *PointInteractor) Get(name string) (point *domain.Point, err error) {
	return mi.PointRepository.Get(name)
}

func (mi *PointInteractor) GetAll() (points domain.Points, err error) {
	return mi.PointRepository.GetAll()
}

func (mi *PointInteractor) Add(point domain.Point) (err error) {
	return mi.PointRepository.Add(point)
}

func (mi *PointInteractor) Delete(point domain.Point) (err error) {
	return mi.PointRepository.Delete(point)
}

func (mi *PointInteractor) Update(point domain.Point) (err error) {
	return mi.PointRepository.Update(point)
}
