package handler

import (
	"github.com/MikelSot/reto-financiera-compartamos/infrastructure/handler/customerrelation"
	"github.com/MikelSot/reto-financiera-compartamos/model"
)

func InitRoutes(spec model.RouterSpecification) {
	// C
	customerrelation.NewRouter(spec)
}
