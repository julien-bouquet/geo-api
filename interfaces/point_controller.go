package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/geo-api/domain"
	"github.com/geo-api/usecases"
)

type PointController struct {
	PointInteractor usecases.PointInteractor
}

func NewPointController(noSQLHandler NoSQLHandler) *PointController {
	return &PointController{
		PointInteractor: usecases.PointInteractor{
			PointRepository: &PointRepository{
				NoSQLHandler: noSQLHandler,
			},
		},
	}
}

func setHeaderContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (mc *PointController) List(w http.ResponseWriter, r *http.Request) {
	setHeaderContentType(w)

	points, err := mc.PointInteractor.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

func (mc *PointController) Add(w http.ResponseWriter, r *http.Request) {
	setHeaderContentType(w)

	point := domain.Point{}
	err := json.NewDecoder(r.Body).Decode(&point)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	mc.PointInteractor.Add(point)

}
