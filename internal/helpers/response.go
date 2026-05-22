package helpers

type APIResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(code int, message string, data interface{}) APIResponse {
	return APIResponse{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
}
