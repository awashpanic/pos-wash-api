package constant

import "net/http"

// default error Code
const (
	DefaultUnhandledError = 1000 + iota
	DefaultNotFoundError
	DefaultBadRequestError
	DefaultUnauthorizedError
	DefaultDuplicateDataError
)

var ErrorMessageMap = map[int]string{
	DefaultUnhandledError:     "something went wrong with our side, please try again",
	DefaultNotFoundError:      "data not found",
	DefaultUnauthorizedError:  "you are not authorized to access this api",
	DefaultDuplicateDataError: "your request body was unprocessable, please modify it",
}

var InteralResponseCodeMap = map[int]int{
	http.StatusInternalServerError: DefaultUnhandledError,
	http.StatusNotFound:            DefaultNotFoundError,
	http.StatusBadRequest:          DefaultBadRequestError,
	http.StatusUnauthorized:        DefaultUnauthorizedError,
	http.StatusUnprocessableEntity: DefaultDuplicateDataError,
}
