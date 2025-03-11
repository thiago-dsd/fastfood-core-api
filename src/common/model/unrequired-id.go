package common_model

import "github.com/google/uuid"

// Represents an optional UUID.
type UnrequiredId struct {
	Id uuid.UUID `json:"id,omitempty"` // The unique identifier
}
