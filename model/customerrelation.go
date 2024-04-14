package model

type CustomerRelation struct {
	Customer Customer `json:"customer"`
	City     City     `json:"city"`
}

type CustomerRelations []CustomerRelation
