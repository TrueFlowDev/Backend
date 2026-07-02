package entity

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrOTPExpired        = errors.New("otp expired")
	ErrOTPLocked         = errors.New("otp locked")
	ErrOTPInvalidCode    = errors.New("otp code is invalid")
	ErrCodeIsRequired    = errors.New("code is required")
	ErrInvalidExpireTime = errors.New("expire time is invalid")
)

const MaxOtpAttempts = 5

type OTP struct {
	code      string
	attempts  uint8
	expiresAt time.Time
}

func NewOTP(
	code string,
	expiresAt time.Time,
) (OTP, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return OTP{}, ErrCodeIsRequired
	}

	if expiresAt.IsZero() || expiresAt.Before(time.Now().UTC()) {
		return OTP{}, ErrInvalidExpireTime
	}

	return OTP{code: code, attempts: 0, expiresAt: expiresAt}, nil
}

// <-- Getters -->

func (o *OTP) Code() string         { return o.code }
func (o *OTP) Attempts() uint8      { return o.attempts }
func (o *OTP) ExpiresAt() time.Time { return o.expiresAt }

// <-- Setters -->

func (o *OTP) Verify(code string) error {
	code = strings.TrimSpace(code)

	if o.isLocked() {
		return ErrOTPLocked
	}
	if o.isExpired() {
		return ErrOTPExpired
	}

	if o.code != code {
		o.attempts++
		return ErrOTPInvalidCode
	}

	return nil
}

// <-- Helpers -->

func (o *OTP) isExpired() bool { return time.Now().UTC().After(o.expiresAt) }
func (o *OTP) isLocked() bool  { return o.attempts >= MaxOtpAttempts }
