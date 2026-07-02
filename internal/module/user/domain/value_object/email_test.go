package value_object_test

import (
	"errors"
	"testing"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

func TestNewEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{"valid email", "Ali@test.com", nil},
		{"empty email", "", value_object.ErrEmailRequired},
		{"invalid email", "invalid", value_object.ErrEmailInvalidFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := value_object.NewEmail(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v, got %v", tt.wantErr, err)
			}
		})
	}
}
