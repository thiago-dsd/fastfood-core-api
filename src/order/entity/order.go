package order_entity

import (
	"github.com/google/uuid"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
)

type Order struct {
	UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Description string    `json:"description" gorm:"not null"`
	Items       *Items     `gorm:"type:text[];not null"`

	common_model.Audit
}
