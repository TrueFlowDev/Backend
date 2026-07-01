package value_object_test

import (
	"errors"
	"testing"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

func TestNewHashedPassword(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{"valid hash", "$2a$10$abcdefghijklmnopqrstuv1234567890abcdefghijklmnopqrstuv", "$2a$10$abcdefghijklmnopqrstuv1234567890abcdefghijklmnopqrstuv", nil},
		{"trim spaces", "  hash  ", "hash", nil},
		{"empty hash", "", "", value_object.ErrHashedPasswordRequired},
		{"spaces only", "   ", "", value_object.ErrHashedPasswordRequired},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hash, err := value_object.NewHashedPassword(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if err == nil && hash.Value() != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, hash.Value())
			}
		})
	}
}
