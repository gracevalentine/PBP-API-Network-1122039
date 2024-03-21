package Model

type Products struct {
	ID           int    `json:"IDProduct"`
	Product_name string `json:"Product Name"`
	Category     string `json:"Category"`
	Price        int    `json:"Price"`
	Quantity     int    `json:"Quantity"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ProductResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Products `json:"data"`
}
type ProductsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Products `json:"data"`
}
