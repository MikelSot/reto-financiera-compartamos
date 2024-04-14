package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractIDFromURLParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param("id"))
}

func ExtractIDFromURLParamByName(name string, c *gin.Context) (int, error) {
	return strconv.Atoi(c.Param(name))
}
