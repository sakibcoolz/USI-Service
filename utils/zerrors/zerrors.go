package zerrors

import (
	"net/http"

	"USI-Service/utils/zerrors/apperrors"

	"github.com/gin-gonic/gin"
)

var (
	errorMessages = map[string]string{
		"BadRequest":          "Bad request",
		"Unauthorized":        "Unauthorized",
		"Forbidden":           "Forbidden",
		"NotFound":            "Not found",
		"InternalServerError": "Internal server error",
		"BadGateway":          "Bad gateway",
		"ServiceUnavailable":  "Service unavailable",
		"GatewayTimeout":      "Gateway timeout",
		"DatabaseError":       "Database error",
		"NullReferenceError":  "Null reference error",
		"SQLConnectionError":  "SQL connection error",
		"SQLQueryError":       "SQL query error",
		"SQLNoRowsError":      "No rows in result set",
	}

	errorStatusCodes = map[string]int{
		"BadRequest":          http.StatusBadRequest,
		"Unauthorized":        http.StatusUnauthorized,
		"Forbidden":           http.StatusForbidden,
		"NotFound":            http.StatusNotFound,
		"InternalServerError": http.StatusInternalServerError,
		"BadGateway":          http.StatusBadGateway,
		"ServiceUnavailable":  http.StatusServiceUnavailable,
		"GatewayTimeout":      http.StatusGatewayTimeout,
		"DatabaseError":       http.StatusInternalServerError,
		"NullReferenceError":  http.StatusInternalServerError,
	}

	// nolint
	errorDetails = map[string]string{
		"DatabaseError":      "Failed to connect to the database",
		"NullReferenceError": "Attempted to access a null object",
	}
)

type Error interface {
	error
	Code() string
	Message() string
	Details() string
}

type AppError struct {
	code    apperrors.ErrorCode
	message string
	err     error
	errMsg  string
}

type Respose struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Code() string {
	return e.code.String()
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) Details() string {
	return e.err.Error()
}

func (e *AppError) Response(c *gin.Context) {
	var response Respose
	response.Code = errorStatusCodes[e.code.String()]
	response.Message = e.message
	response.Details = e.errMsg
	if response.Code == 0 {
		response.Code = http.StatusInternalServerError
	}
	c.JSON(response.Code, response)
}

func NewAppError(code apperrors.ErrorCode, message string, err error, errMsg string) Error {
	return &AppError{code: code, message: message, err: err, errMsg: errMsg}
}

func Errors(key apperrors.ErrorCode, err error) error {
	return NewAppError(key, errorMessages[key.String()], err, err.Error())
}
