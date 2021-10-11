package usecases

import "github.com/coast-nav-api/domain"

type PointInteractor struct {
	PointRepository PointRepository
}

func (mi *PointInteractor) GetAll() (points domain.Points, err error) {
	return mi.PointRepository.GetAll()
}

func (mi *PointInteractor) Add() (err error) {
	return
}
