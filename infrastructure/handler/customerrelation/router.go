package customerrelation

import (
	"github.com/gin-gonic/gin"

	"github.com/MikelSot/reto-financiera-compartamos/domain/city"
	"github.com/MikelSot/reto-financiera-compartamos/domain/citycustomer"
	"github.com/MikelSot/reto-financiera-compartamos/domain/customer"
	"github.com/MikelSot/reto-financiera-compartamos/domain/customerrelation"
	"github.com/MikelSot/reto-financiera-compartamos/infrastructure/handler/response"
	cityStorage "github.com/MikelSot/reto-financiera-compartamos/infrastructure/postgres/city"
	cityCustomerStorage "github.com/MikelSot/reto-financiera-compartamos/infrastructure/postgres/citycustomer"
	customerStorage "github.com/MikelSot/reto-financiera-compartamos/infrastructure/postgres/customer"
	"github.com/MikelSot/reto-financiera-compartamos/model"
)

const (
	_routesPrefix = "/api/v1/compartamos"
)

func NewRouter(spec model.RouterSpecification) {
	handler := buildHandler(spec)

	routes(spec.Api, handler)
}

func buildHandler(spec model.RouterSpecification) handler {
	response := response.New(spec.Logger)

	customerUseCase := customer.New(customerStorage.New(spec.DB))
	cityUseCase := city.New(cityStorage.New(spec.DB))
	cityCustomer := citycustomer.New(cityCustomerStorage.New(spec.DB))

	customerRelationUseCase := customerrelation.New(customerUseCase, cityUseCase, cityCustomer)

	return newHandler(customerRelationUseCase, response)
}

func routes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group(_routesPrefix, middlewares...)

	routes.POST("", h.Create)
	routes.PUT("/:id", h.Update)
	routes.DELETE("/customers/:id/cities/:city-id", h.Delete)

	routes.GET("/customers", h.GetAllCustomers)
}
