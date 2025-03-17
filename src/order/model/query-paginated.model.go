package order_model

import (
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	database_model "github.com/thiago-dsd/fastfood-core-api/src/database/model"
)

type QueryPaginated struct {
	Description string   `json:"description,omitempty"`
	Items       []string `json:"items,omitempty"`

	common_model.UnrequiredId
	database_model.Paginate
	database_model.DateOrder
	database_model.DateWhere
}