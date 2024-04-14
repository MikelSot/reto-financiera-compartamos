package customerrelation

import (
	"github.com/gin-gonic/gin"

	"github.com/MikelSot/reto-financiera-compartamos/domain/customerrelation"
	"github.com/MikelSot/reto-financiera-compartamos/infrastructure/handler/request"
	"github.com/MikelSot/reto-financiera-compartamos/infrastructure/handler/response"
	"github.com/MikelSot/reto-financiera-compartamos/model"
)

type handler struct {
	useCase customerrelation.UseCase

	response response.ApiResponse
}

func newHandler(useCase customerrelation.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) Create(c *gin.Context) {
	m := model.CustomerRelation{}

	if err := c.Bind(&m); err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	if err := h.useCase.Create(m); err != nil {
		c.JSON(h.response.Error(c, "useCase.Create()", err))

		return
	}

	c.JSON(h.response.Created(m))
}

func (h handler) Update(c *gin.Context) {
	m := model.CustomerRelation{}

	if err := c.Bind(&m); err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}
	m.Customer.Id = uint(ID)

	if err := h.useCase.Update(m); err != nil {
		c.JSON(h.response.Error(c, "useCase.Update()", err))

		return
	}

	c.JSON(h.response.Updated())
}

func (h handler) Delete(c *gin.Context) {
	ID, err := request.ExtractIDFromURLParam(c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	cityID, err := request.ExtractIDFromURLParamByName("city-id", c)
	if err != nil {
		c.JSON(h.response.BindFailed(c, err))

		return
	}

	if err := h.useCase.Delete(uint(ID), uint(cityID)); err != nil {
		c.JSON(h.response.Error(c, "useCase.Delete()", err))

		return
	}

	c.JSON(h.response.Deleted())
}

func (h handler) GetAllCustomers(c *gin.Context) {
	ms, err := h.useCase.GetAllCustomers()
	if err != nil {
		c.JSON(h.response.Error(c, "useCase.GetAllCustomers()", err))
		return
	}

	c.JSON(h.response.OK(ms))
}
