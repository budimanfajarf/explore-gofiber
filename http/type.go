package http

type HttpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HttpResponse struct {
	Data  interface{} `json:"data"`
	Meta  interface{} `json:"meta"`
	Error *HttpError  `json:"error"`
}
