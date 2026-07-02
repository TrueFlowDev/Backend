package value_object_test

import (
	"errors"
	"testing"

	"github.com/TrueFlowDev/Backend/internal/module/user/domain/value_object"
)

func TestNewPhone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		want      string
		wantError error
	}{
		{"valid mobile", "09121234567", "+989121234567", nil},
		{"valid mobile with country code", "+989121234567", "+989121234567", nil},
		{"valid mobile with spaces", " 09121234567 ", "+989121234567", nil},
		{"empty phone", "", "", value_object.ErrPhoneRequired},
		{"invalid format", "abcdef", "", value_object.ErrPhoneInvalidFormat},
		{"short phone", "0912", "", value_object.ErrPhoneInvalidFormat},
		{"fixed line", "02112345678", "", value_object.ErrPhoneInvalidFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			phone, err := value_object.NewPhone(tt.input)

			if !errors.Is(err, tt.wantError) {
				t.Fatalf("expected error %v, got %v", tt.wantError, err)
			}

			if err == nil && phone.Value() != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, phone.Value())
			}
		})
	}
}
