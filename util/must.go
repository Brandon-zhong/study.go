package util

import (
	"io"
)

// Must panics if err is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// String returns the string or panic.
func MustString(s string, err error) string {
	Must(err)
	return s
}

// Int returns the integer or panic.
func MustInt(a int, err error) int {
	Must(err)
	return a
}

// Bool returns the bool or panic.
func MustBool(a bool, err error) bool {
	Must(err)
	return a
}

// Float64 returns the float64 or panic.
func MustFloat64(a float64, err error) float64 {
	Must(err)
	return a
}

// Byte returns the byte array or panic.
func MustByte(a []byte, err error) []byte {
	Must(err)
	return a
}

// NotEmpty checks string not empty.
func MustNotEmpty(s string) {
	if s == "" {
		panic("given string is empty")
	}
}

// True checks b is true.
func MustTrue(b bool) {
	if !b {
		panic("assertion not true")
	}
}

// Write checks for a io.Write result.
func MustWrite(n int, err error) {
	if err != nil {
		panic(err)
	}
}

// Close closes the file and panic on error.
// Useful in defer statement.
func MustClose(c io.Closer) {
	Must(c.Close())
}
