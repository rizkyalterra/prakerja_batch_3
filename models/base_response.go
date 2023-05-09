package models

type BaseResponse struct {
	Status  bool        `json:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
