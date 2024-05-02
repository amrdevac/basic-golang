package utils

type BaseTypeResponse struct {
	Status  bool        `json:"status"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
