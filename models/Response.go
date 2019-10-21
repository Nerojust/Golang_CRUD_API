package models

type Response struct {
	//StatusCode StatusCode         `json:"status_code"`
	Message Message     `json:"'message'"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
type StatusCode int
type Message string

const (
	Success      StatusCode = 200
	BadRequest   StatusCode = 400
	UnAuthorized StatusCode = 401
)
const (
	SUCCESS Message = "Request performed successfully"
	FAILURE Message = "Request failed"
)
