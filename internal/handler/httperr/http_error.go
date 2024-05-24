package httperr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error,omitempty"`
	Code    int      `json:"code"`
	Fields  []Fields `json:"fields,omitempty"`
}

type Fields struct {
	Field   string `json:"field"`
	Value   any    `json:"value,omitempty"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(m, e string, c int, f []Fields) *RestErr {
	return &RestErr{
		Message: m,
		Err:     e,
		Code:    c,
		Fields:  f,
	}
}

func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func UnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func BadRequestValidationError(m string, c []Fields) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Fields:  c,
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

func ForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
