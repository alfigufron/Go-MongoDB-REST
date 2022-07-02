package helpers

type ResponseMeta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

func ResponseBody(status string, message string, data interface{}, code int) *APIResponse {
	var httpCode int

	if status == "success" {
		httpCode = 200
	} else {
		httpCode = 400
	}

	body := &APIResponse{
		Meta: ResponseMeta{
			Status:  status,
			Code:    httpCode,
			Message: message,
		},
		Data: data,
	}

	return body
}
