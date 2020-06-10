package interfaces

var status string

type HTTPResponse struct {
	Status string `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	ErrorMessage interface{}    `json:"error_message"`
}

func CreateHTTPResponsePayload(data interface{}, status string, errorMessage interface{}) *HTTPResponse{
	return &HTTPResponse{
		Status: status,
		Data: data,
		ErrorMessage: errorMessage,
	}
}