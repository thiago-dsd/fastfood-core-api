package order_entity

import (
	"database/sql/driver"
	"errors"
	"strings"
)

// Items is a custom type for the Items field in the Order struct.
// It represents an array of strings.
type Items []string

// Value implements the driver.Valuer interface to serialize the Items field.
// It converts the array of strings into the PostgreSQL array format: {item1,item2}.
func (i Items) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	// Convert the array of strings into the format {item1,item2}
	str := "{" + strings.Join(i, ",") + "}"
	return str, nil
}

// Scan implements the sql.Scanner interface to deserialize the Items field.
// It converts the PostgreSQL array format (stored as a string or []byte) into a slice of strings.
func (i *Items) Scan(value interface{}) error {
	if value == nil {
		*i = nil
		return nil
	}

	// Convert the value to []byte
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("failed to convert Items to []byte")
	}

	// Remove curly braces and extra spaces
	str := string(bytes)
	str = strings.Trim(str, "{}")

	// Split the string into a slice of strings
	items := strings.Split(str, ",")

	// Remove extra spaces from each item
	for idx, item := range items {
		items[idx] = strings.TrimSpace(item)
	}

	// Assign the slice of strings to the Items field
	*i = items
	return nil
}