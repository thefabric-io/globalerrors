package globalerrors

import (
	"errors"
	"google.golang.org/grpc/codes"
)

func HTTPStatus(customError error) int {
	return status[int](customError, httpStatusCodesMap)
}

func GRPCStatus(customError error) codes.Code {
	return status[codes.Code](customError, grpcStatusCodesMap)
}

func status[T any](customError error, m map[error]T) T {
	for _, e := range globalErrors {
		if errors.Is(customError, e) {
			s, exists := m[e]
			if exists {
				return s
			}
		}
	}

	return m[InternalServerError]
}

var globalErrors = []error{
	BadRequest,
	Unauthorized,
	PaymentRequired,
	Forbidden,
	NotFound,
	MethodNotAllowed,
	NotAcceptable,
	ProxyAuthRequired,
	RequestTimeout,
	Conflict,
	Gone,
	LengthRequired,
	PreconditionFailed,
	RequestEntityTooLarge,
	RequestURITooLong,
	UnsupportedMediaType,
	RequestedRangeNotSatisfiable,
	ExpectationFailed,
	Teapot,
	MisdirectedRequest,
	UnprocessableEntity,
	Locked,
	FailedDependency,
	TooEarly,
	UpgradeRequired,
	PreconditionRequired,
	TooManyRequests,
	RequestHeaderFieldsTooLarge,
	UnavailableForLegalReasons,
	InternalServerError,
	NotImplemented,
	BadGateway,
	ServiceUnavailable,
	GatewayTimeout,
	HTTPVersionNotSupported,
	VariantAlsoNegotiates,
	InsufficientStorage,
	LoopDetected,
	NotExtended,
	NetworkAuthenticationRequired,
}

var (
	BadRequest                    = errors.New("bad request")
	Unauthorized                  = errors.New("unauthorized")
	PaymentRequired               = errors.New("payment required")
	Forbidden                     = errors.New("forbidden")
	NotFound                      = errors.New("not found")
	MethodNotAllowed              = errors.New("method not allowed")
	NotAcceptable                 = errors.New("not acceptable")
	ProxyAuthRequired             = errors.New("proxy authentication required")
	RequestTimeout                = errors.New("request timeout")
	Conflict                      = errors.New("conflict")
	Gone                          = errors.New("gone")
	LengthRequired                = errors.New("length required")
	PreconditionFailed            = errors.New("precondition failed")
	RequestEntityTooLarge         = errors.New("request entity too large")
	RequestURITooLong             = errors.New("request URI too long")
	UnsupportedMediaType          = errors.New("unsupported media type")
	RequestedRangeNotSatisfiable  = errors.New("requested range not satisfiable")
	ExpectationFailed             = errors.New("expectation failed")
	Teapot                        = errors.New("I'm a teapot")
	MisdirectedRequest            = errors.New("misdirected request")
	UnprocessableEntity           = errors.New("unprocessable entity")
	Locked                        = errors.New("locked")
	FailedDependency              = errors.New("failed dependency")
	TooEarly                      = errors.New("too early")
	UpgradeRequired               = errors.New("upgrade required")
	PreconditionRequired          = errors.New("precondition required")
	TooManyRequests               = errors.New("too many requests")
	RequestHeaderFieldsTooLarge   = errors.New("request header fields too large")
	UnavailableForLegalReasons    = errors.New("unavailable for legal reasons")
	InternalServerError           = errors.New("internal server error")
	NotImplemented                = errors.New("not implemented")
	BadGateway                    = errors.New("bad gateway")
	ServiceUnavailable            = errors.New("service unavailable")
	GatewayTimeout                = errors.New("gateway timeout")
	HTTPVersionNotSupported       = errors.New("HTTP version not supported")
	VariantAlsoNegotiates         = errors.New("variant also negotiates")
	InsufficientStorage           = errors.New("insufficient storage")
	LoopDetected                  = errors.New("loop detected")
	NotExtended                   = errors.New("not extended")
	NetworkAuthenticationRequired = errors.New("network authentication required")
)
