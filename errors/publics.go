package errors

import (
	"github.com/hedzr/errors"
	"strings"
)

// Walk will walk all inner and nested error objects inside err
func Walk(err error, fn func(err error) (stop bool)) {
	if !fn(err) {
		if ee, ok := err.(errors.Walkable); ok {
			ee.Walk(fn)
		}
	}
}

// Range can walk the inner/attached errors inside err
func Range(err error, fn func(err error) (stop bool)) {
	if !fn(err) {
		if ee, ok := err.(errors.Ranged); ok {
			ee.Range(fn)
		}
	}
}

// Equal tests if code number presented recursively
func Equal(err error, code Code) bool {
	if x, ok := err.(interface{ EqualRecursive(code Code) bool }); ok {
		return x.EqualRecursive(code)
	}
	return false
}

// IsAny tests if any codes presented
func IsAny(err error, code ...Code) bool {
	if x, ok := err.(interface{ IsAny(codes ...Code) bool }); ok {
		return x.IsAny(code...)
	}
	return false
}

// IsBoth tests if all codes presented
func IsBoth(err error, code ...Code) bool {
	if x, ok := err.(interface{ IsBoth(codes ...Code) bool }); ok {
		return x.IsBoth(code...)
	}
	return false
}

// TextContains test if a text fragment is included by err
func TextContains(err error, text string) bool {
	return strings.Index(err.Error(), text) >= 0
}

// Attach attaches the errors into `err`
func Attach(err error, errs ...error) {
	if x, ok := err.(interface{ AttachIts(errors ...error) }); ok {
		x.AttachIts(errs...)
	}
}

// Nest wraps/nests the errors into `err`
func Nest(err error, errs ...error) {
	if x, ok := err.(interface{ NestIts(errors ...error) }); ok {
		x.NestIts(errs...)
	}
}

// DumpStacksAsString returns stack tracing information like debug.PrintStack()
func DumpStacksAsString(allRoutines bool) string {
	return errors.DumpStacksAsString(allRoutines)
}

// HasInnerErrors detects if nested or attached errors present
func HasInnerErrors(err error) (yes bool) {
	return errors.HasInnerErrors(err)
}

// HasAttachedErrors detects if attached errors present
func HasAttachedErrors(err error) (yes bool) {
	return errors.HasAttachedErrors(err)
}

// HasWrappedError detects if nested or wrapped errors present
//
// nested error: ExtErr.inner
// wrapped error: fmt.Errorf("... %w ...", err)
func HasWrappedError(err error) (yes bool) {
	return errors.HasWrappedError(err)
}
