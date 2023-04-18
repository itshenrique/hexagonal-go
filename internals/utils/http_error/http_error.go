package http_error

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type IHttpError interface {
	Error() string
	Status() int
	Message() string
	Causes() []interface{}
}

type HttpError struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e HttpError) Message() string {
	return e.ErrMessage
}

func (e HttpError) Status() int {
	return e.ErrStatus
}

func (e HttpError) Causes() []interface{} {
	return e.ErrCauses
}

func NewRestError(message string, status int, err string, causes []interface{}) IHttpError {
	return HttpError{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

func NewRestErrorFromBytes(bytes []byte) (IHttpError, error) {
	var apiErr HttpError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) IHttpError {
	return HttpError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) IHttpError {
	return HttpError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewUnauthorizedError(message string) IHttpError {
	return HttpError{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewForbiddenError(message string) IHttpError {
	return HttpError{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "forbidden",
	}
}

func NewInternalServerError(message string, err error) IHttpError {
	result := HttpError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internals_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
