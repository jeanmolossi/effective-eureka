package httputil

// HttpOkWithMessage returns a 200 response with a message
type HttpOkWithMessage struct {
	Message string `json:"message" example:"feedback message"`
}

// HttpInternalServerErr returns a error message
type HttpInternalServerErr struct {
	Message string `json:"error" example:"internal server error"`
}

// HttpBadRequestErr returns a error message
type HttpBadRequestErr struct {
	Message string `json:"error" example:"bad request"`
}

// HttpForbiddenErr returns a error message
type HttpForbiddenErr struct {
	Message string `json:"error" example:"forbidden"`
}

// HttpUnauthorizedErr returns a error message
type HttpUnauthorizedErr struct {
	Message string `json:"error" example:"unauthorized"`
}
