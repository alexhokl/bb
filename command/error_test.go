package command

import (
	"testing"
)

func TestStatusError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      StatusError
		expected string
	}{
		{
			name: "not found error",
			err: StatusError{
				Status:     "Not Found",
				StatusCode: 404,
			},
			expected: "Status: Not Found, Code: 404",
		},
		{
			name: "unauthorized error",
			err: StatusError{
				Status:     "Unauthorized",
				StatusCode: 401,
			},
			expected: "Status: Unauthorized, Code: 401",
		},
		{
			name: "internal server error",
			err: StatusError{
				Status:     "Internal Server Error",
				StatusCode: 500,
			},
			expected: "Status: Internal Server Error, Code: 500",
		},
		{
			name: "bad request error",
			err: StatusError{
				Status:     "Bad Request",
				StatusCode: 400,
			},
			expected: "Status: Bad Request, Code: 400",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("StatusError.Error() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestStatusError_ImplementsError(t *testing.T) {
	var _ error = StatusError{}
}
