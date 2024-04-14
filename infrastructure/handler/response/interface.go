package response

import (
	"github.com/gin-gonic/gin"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

const (
	BindFailed model.StatusCode = "bind_failed"
	// Failure sends the custom error message and API message from the logic
	Failure        model.StatusCode = "failure"
	Ok             model.StatusCode = "ok"
	RecordCreated  model.StatusCode = "record_created"
	RecordUpdated  model.StatusCode = "record_updated"
	RecordDeleted  model.StatusCode = "record_deleted"
	RecordNotFound model.StatusCode = "record_not_found"
	// UnexpectedError is a server error
	UnexpectedError model.StatusCode = "unexpected_error"
	// AuthError is any of authorization errors
	AuthError model.StatusCode = "authorization_error"
)

// ApiResponse interface that must be implemented for handler http responses of framework gin
type ApiResponse interface {
	OK(data any) (int, any)
	Created(data any) (int, any)
	Updated() (int, any)
	Deleted() (int, any)
	// BindFailed set the data with status code 400
	BindFailed(c *gin.Context, err error) (int, any)
	Error(c *gin.Context, who string, err error) (int, any)
}
