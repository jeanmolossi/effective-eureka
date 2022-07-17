package httputil

// HttpOkWithMessage returns a 200 response with a message
type HttpOkWithMessage struct {
	Message string `json:"message" example:"feedback message"`
}

// HttpInternalServerErr returns a error message
type HttpInternalServerErr struct {
	Message string `json:"error" example:"internal server error"`
}
