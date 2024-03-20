package apperrors

type ErrorCode int

const (
	BadRequest ErrorCode = iota
	Unauthorized
	Forbidden
	NotFound
	InternalServerError
	BadGateway
	ServiceUnavailable
	GatewayTimeout
	DatabaseError
	NullReferenceError
	SQLConnectionError
	SQLQueryError
	SQLNoRowsError
)
