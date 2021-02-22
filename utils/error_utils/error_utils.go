package error_utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type messageErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *messageErr) Message() string {
	return e.ErrMessage
}

func (e *messageErr) Status() int {
	return e.ErrStatus
}

func (e *messageErr) Error() string {
	return e.ErrError
}

// NewNotFoundError func for new not found http error
func NewNotFoundError(message string) MessageErr {

	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

// NewBadRequestError func for new bad request http error
func NewBadRequestError(message string) MessageErr {

	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

// NewUnprocessableEntityError func for new unprocessable entity http error
func NewUnprocessableEntityError(message string) MessageErr {

	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}

// NewAPIErrFromBytes func for new API error from body in byte array
func NewAPIErrFromBytes(body []byte) (MessageErr, error) {

	var result messageErr
	if err := json.Unmarshal(body, &result); err != nil {
		log.Print(err)
		return nil, err
	}
	fmt.Printf("result: %+v\n", result)
	return &result, nil
}

// NewInternalServerError func for new internar server error
func NewInternalServerError(message string) MessageErr {

	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}
