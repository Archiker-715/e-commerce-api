package entity

type SearchRequest struct {
	Query filter `json:"query"`
}

type filter struct {
	Term    term     `json:"term"`
	Range   prPrice  `json:"range"`
	Boolean boolean  `json:"bool"`
	Filter  prFilter `json:"filter"`
}

type term struct {
	Product string `json:"product"`
}

type prPrice struct {
	UnitPrice rang `json:"unit_price"`
}

type rang struct {
	Gte uint `json:"gte"`
	Lte uint `json:"lte"`
}

type boolean struct {
	Must    []filter `json:"must"`
	MustNot []filter `json:"must_not"`
	Should  []filter `json:"should"`
}

type prFilter struct {
	Exists exists `json:"exists"`
}

type exists struct {
	Field string `json:"field"` // здесь вместо field в json надо вставлять наименование поля
}
