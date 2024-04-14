package model

import (
	"fmt"
)

type ErrorDetail struct {
	Field       string    `json:"field,omitempty"`
	Issue       IssueType `json:"issue,omitempty"`
	Code        string    `json:"code,omitempty"`
	Description string    `json:"description"`
}

type ErrorDetails []ErrorDetail

func (e *ErrorDetails) Add(field ErrorDetail) {
	*e = append(*e, field)
}

func (e *ErrorDetails) AddMissingRequireParam(nameField string) {
	*e = append(*e, ErrorDetail{
		Field:       nameField,
		Issue:       IssueMissingRequiredParam,
		Description: "can not be empty",
	})
}

func (e *ErrorDetails) AddInvalidSyntaxParam(nameField string, syntaxType string) {
	*e = append(*e, ErrorDetail{
		Field:       nameField,
		Issue:       IssueInvalidSyntaxParam,
		Description: fmt.Sprintf("must be of type %s", syntaxType),
	})
}

func (e *ErrorDetails) AddMissingRequireFieldBody(nameField string) {
	*e = append(*e, ErrorDetail{
		Field:       nameField,
		Issue:       IssueMissingRequiredFieldBody,
		Description: "can not be empty",
	})
}
