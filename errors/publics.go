package errors

import (
	"github.com/hedzr/errors"
	"strings"
)

//
//
//

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap returns the result of calling the Unwrap method on err, if err's
// type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

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

// Attach attaches the nested errors into CodedErr
func Attach(err error, errs ...error) error {
	if x, ok := err.(interface{ AttachIts(errors ...error) }); ok {
		x.AttachIts(errs...)
	}
	return err
}

// Nest attaches the nested errors into CodedErr
func Nest(err error, errs ...error) error {
	if x, ok := err.(interface{ NestIts(errors ...error) }); ok {
		x.NestIts(errs...)
	}
	return err
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
