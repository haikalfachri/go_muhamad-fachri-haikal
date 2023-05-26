package response

type Response[T any] struct {
	Status  string 	`json:"status"`
	Message string 	`json:"message"`
	Error	string 	`json:"error,omitempty"`
	Data    T      	`json:"data,omitempty"`
}