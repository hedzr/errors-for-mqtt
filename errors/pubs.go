package errors

import "github.com/hedzr/errors"

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

// CanWalk tests if err is walkable
func CanWalk(err error) (ok bool) {
	_, ok = err.(errors.Walkable)
	return
}

// CanRange tests if err is range-able
func CanRange(err error) (ok bool) {
	_, ok = err.(errors.Ranged)
	return
}

// CanUnwrap tests if err is unwrap-able
func CanUnwrap(err error) (ok bool) {
	_, ok = err.(interface{ Unwrap() error })
	return
}

// CanIs tests if err is is-able
func CanIs(err error) (ok bool) {
	_, ok = err.(interface{ Is(error) bool })
	return
}

// CanAs tests if err is as-able
func CanAs(err error) (ok bool) {
	_, ok = err.(interface{ As(interface{}) bool })
	return
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
