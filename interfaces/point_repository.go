package interfaces

import (
	"context"

	"github.com/geo-api/domain"
)

type PointRepository struct {
	NoSQLHandler NoSQLHandler
}

var ctx = context.TODO()

func (mr *PointRepository) Get(name string) (point *domain.Point, err error) {
	doc, err := mr.NoSQLHandler.Get(map[string]interface{}{"name": name})
	if err != nil {
		return nil, err
	}

	defer doc.Close()
	if err = doc.Err(); err != nil {
		return nil, err
	}

	err = doc.Read(point)
	if err != nil {
		return nil, err
	}

	return point, err
}

func (mr *PointRepository) GetAll() (points domain.Points, err error) {
	docs, err := mr.NoSQLHandler.Get(nil)

	if err != nil {
		return nil, err
	}
	defer docs.Close()

	for docs.Next(ctx) {
		var point domain.Point
		err := docs.Read(&point)
		if err != nil {
			return points, err
		}

		points = append(points, &point)
	}

	if err = docs.Err(); err != nil {
		return nil, err
	}

	return points, err
}

func (mr *PointRepository) Add(point domain.Point) (err error) {
	err = mr.NoSQLHandler.Add(point)
	if err != nil {
		return err
	}
	return nil
}

func (mr *PointRepository) Delete(point domain.Point) (err error) {
	err = mr.NoSQLHandler.Delete(point)
	if err != nil {
		return err
	}
	return nil
}
