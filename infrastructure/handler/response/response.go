package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	rkginctx "github.com/rookie-ninja/rk-gin/v2/middleware/context"

	"github.com/MikelSot/reto-financiera-compartamos/model"
)

const (
	BadRequest model.StatusCode = "INVALID_REQUEST"
)

const (
	BadRequestDescription string = "Request is not well-formed, syntactically incorrect, or violates schema."
)

var (
	_badRequestDefault = model.TypeCode{
		StatusHttp:  http.StatusBadRequest,
		Code:        BadRequest,
		Description: BadRequestDescription,
	}
)

type API struct {
	logger model.Logger
}

func New(logger model.Logger) API {
	return API{logger}
}

func (a API) OK(data any) (int, any) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡listo!"}},
	}
}

func (a API) Created(data any) (int, any) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordCreated, Message: "¡listo!"}},
	}
}

func (a API) Updated() (int, any) {
	return http.StatusOK, model.MessageResponse{
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}

func (a API) Deleted() (int, any) {
	return http.StatusOK, model.MessageResponse{
		Messages: model.Responses{{Code: RecordDeleted, Message: "¡listo!"}},
	}
}

func (a API) BindFailed(c *gin.Context, err error) (int, any) {
	// use of c.Bind add err
	c.Errors = nil

	// add error to rk gin
	event := rkginctx.GetEvent(c)
	event.AddErr(err)

	fun, _, line, _ := runtime.Caller(1)

	e := model.NewError()
	e.SetError(err)
	e.SetCode(_badRequestDefault.Code)
	e.SetAPIMessage(_badRequestDefault.Description)
	e.SetStatusHTTP(_badRequestDefault.StatusHttp)
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho("c.Bind()")

	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var timeErr *time.ParseError

	switch {
	case errors.As(err, &timeErr):
		e.Fields.Add(model.ErrorDetail{
			Field:       "body",
			Description: fmt.Sprintf("Invalid json syntax time (%s).", timeErr.Value),
			Issue:       model.IssueBodyError,
		})
	case errors.As(err, &syntaxError):
		e.Fields.Add(model.ErrorDetail{
			Field:       "body",
			Description: fmt.Sprintf("Invalid json syntax"),
			Issue:       model.IssueBodyError,
		})
	case errors.As(err, &unmarshalTypeError):
		e.Fields.Add(model.ErrorDetail{
			Field:       "body",
			Description: fmt.Sprintf("The required data type must be %s.", unmarshalTypeError.Type),
			Issue:       model.IssueBodyError,
		})
	}

	_ = c.Error(e)

	a.logger.Warnf("%s", e.Error())

	return _badRequestDefault.StatusHttp, model.MessageResponse{
		Errors: model.Responses{{Code: _badRequestDefault.Code, Message: _badRequestDefault.Description}},
	}
}

// UnexpectedError returns an unexpected error
func (a API) UnexpectedError(c *gin.Context, who string, err error) (int, any) {
	fun, _, line, _ := runtime.Caller(1)

	e := model.NewError()
	e.SetError(err)
	e.SetAPIMessage("¡Uy! metimos la pata, disculpanos lo solucionaremos")
	e.SetCode(UnexpectedError)
	e.SetStatusHTTP(http.StatusInternalServerError)
	e.SetEndpoint(c.FullPath())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho(who)

	a.logger.Errorf("%s", e.Error())

	return http.StatusInternalServerError, model.MessageResponse{
		Errors: model.Responses{{Code: UnexpectedError, Message: "¡Uy! metimos la pata, disculpanos lo solucionaremos"}},
	}
}

func (a API) ErrorHandled(c *gin.Context, who string, e *model.Error) (int, any) {
	fun, _, line, _ := runtime.Caller(1)

	e.SetCode(Failure)
	e.SetEndpoint(c.FullPath())
	e.SetWhere(fmt.Sprintf("%s:%d", runtime.FuncForPC(fun).Name(), line))
	e.SetWho(who)

	if !e.HasStatusHTTP() {
		e.SetStatusHTTP(http.StatusBadRequest)
	}

	if e.StatusHTTP() < http.StatusInternalServerError {
		a.logger.Warnf("%s", e.Error())
		return e.StatusHTTP(), model.MessageResponse{
			Errors: model.Responses{
				{Code: Failure, Message: e.APIMessage()},
			},
		}
	}

	a.logger.Errorf("%s", e.Error())

	return e.StatusHTTP(), model.MessageResponse{
		Errors: model.Responses{{Code: Failure, Message: "¡Uy! metimos la pata, disculpanos lo solucionaremos"}},
	}
}

func (a API) Error(c *gin.Context, who string, err error) (int, any) {
	e := model.NewError()
	if errors.As(err, &e) {
		return a.ErrorHandled(c, who, e)
	}

	return a.UnexpectedError(c, who, err)

}
