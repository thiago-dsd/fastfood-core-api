package order_model

type UpdateOrder struct {
	Id          string   `json:"id"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
}
