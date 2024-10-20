package touchid

import "errors"

// ErrTouchIDNotSupported is the error returned when Touch ID is not supported on the current platform
var ErrTouchIDNotSupported = errors.New("Touch ID authentication is not supported on this platform")

// Authenticate attempts to authenticate using Touch ID
// On Windows, this always returns false with an error indicating lack of support
func Authenticate(dType int, reason string) (bool, error) {
	return false, ErrTouchIDNotSupported
}
