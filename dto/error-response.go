package dto

type ErrorResponse struct {
	ResponseMeta
	Data interface{} `json:"data"`
}

func DefaultErrorResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("")
}

func DefaultErrorResponseWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      msg,
			ResponseTime: "",
		},
		Data: nil,
	}
}

func DefaultErrorInvalidDataWithMessage(msg string) ErrorResponse {
	return ErrorResponse{
		ResponseMeta: ResponseMeta{
			Success:      false,
			MessageTitle: "Oops, something went wrong.",
			Message:      "Form Invalid data.",
			ResponseTime: "",
		},
		Data: msg,
	}
}

func DefaultBadRequestResponse() ErrorResponse {
	return DefaultErrorResponseWithMessage("Bad request")
}
