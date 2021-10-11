package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/coast-nav-api/usecases"
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
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(points)
}

func (mc *PointController) Add(w http.ResponseWriter, r *http.Request) {
	mc.PointInteractor.Add()
	setHeaderContentType(w)

}
