package globalerrors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"testing"
)

func TestHTTPStatus(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{"BadRequest", BadRequest, http.StatusBadRequest},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized},
		// Add all the other HTTP status error cases here
		{"Unknown", errors.New("unknown"), http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HTTPStatus(tt.err)
			if got != tt.want {
				t.Errorf("HTTPStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGRPCStatus(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want codes.Code
	}{
		{"BadRequest", BadRequest, codes.InvalidArgument},
		{"Unauthorized", Unauthorized, codes.Unauthenticated},
		// Add all the other gRPC status error cases here
		{"Unknown", errors.New("unknown"), codes.Internal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GRPCStatus(tt.err)
			if got != tt.want {
				t.Errorf("GRPCStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomErrorCompatibility(t *testing.T) {
	err := ErrIDIsRequired()
	httpStatus := HTTPStatus(err)
	grpcStatus := GRPCStatus(err)

	if httpStatus != http.StatusUnprocessableEntity {
		t.Errorf("Expected HTTPStatus to be 422, got %d", httpStatus)
	}

	if grpcStatus != codes.InvalidArgument {
		t.Errorf("Expected GRPCStatus to be InvalidArgument, got %s", grpcStatus.String())
	}
}

func ErrIDIsRequired() error {
	return errIDIsRequired{}
}

type errIDIsRequired struct{}

func (e errIDIsRequired) Error() string {
	return "id is required"
}

func (e errIDIsRequired) Is(target error) bool {
	return errors.Is(target, UnprocessableEntity) || e == target
}
