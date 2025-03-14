package order_entity

import (
	"github.com/google/uuid"
	common_model "github.com/thiago-dsd/fastfood-core-api/src/common/model"
	user_entity "github.com/thiago-dsd/fastfood-core-api/src/user/entity"
)

type Order struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Description string    `json:"description" gorm:"not null"`
	Items       []string  `json:"items" gorm:"type:text[];not null"`
	User        *user_entity.User `json:"user" gorm:"foreignKey:UserID;references:ID"`

	common_model.Audit
}
