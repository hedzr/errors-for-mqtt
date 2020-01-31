package errors

import (
	"fmt"
	"github.com/hedzr/errors"
	"io"
	"log"
	"strconv"
	"strings"
)

// NoCannedError detects mqttError object is not an error or not an canned-error (inners is empty)
func (e *MqttError) NoCannedError() bool {
	return e.Number() == OK || e.HasAttachedErrors()
}

// HasAttachedErrors tests if any errors attached (nor nested) to `e` or not
func (e *MqttError) HasAttachedErrors() bool {
	return len(e.GetAttachedErrors()) != 0
}

// Code put another code into CodedErr
func (e *MqttError) Code(code Code) *MqttError {
	_ = e.CodedErr.Code(errors.Code(code))
	return e
}

// Equal compares code
func (e *MqttError) Equal(code Code) bool {
	return e.CodedErr.Equal(errors.Code(code))
}

// EqualRecursive compares with code
func (e *MqttError) EqualRecursive(code Code) bool {
	// return e.CodedErr.EqualRecursive(code)

	// EqualRecursive compares with code
	if e.Equal(code) {
		return true
	}

	b := false
	errors.Walk(e, func(err error) (stop bool) {
		log.Printf("  ___E : %+v", err)
		if c, ok := err.(interface{ Equal(code Code) bool }); ok {
			if c.Equal(Code(code)) {
				b = true
				return true
			}
		} else if c, ok := err.(interface{ Equal(code errors.Code) bool }); ok {
			if c.Equal(errors.Code(code)) {
				b = true
				return true
			}
		}
		return false
	})
	return b
}

// Number returns the code number
func (e *MqttError) Number() Code {
	return Code(e.CodedErr.Number())
}

//
func (e *MqttError) IsBoth(code ...Code) bool {
	for _, c := range code {
		if !e.EqualRecursive(c) {
			return false
		}
	}
	return true
}

//
func (e *MqttError) IsAny(code ...Code) bool {
	for _, c := range code {
		if e.EqualRecursive(c) {
			return true
		}
	}
	return false
}

func (e *MqttError) Error() string {
	var b strings.Builder
	strings.Repeat(" ", 32)
	b.WriteString(strconv.Itoa(int(e.CloseReason)))
	b.WriteRune('|')
	b.WriteString(e.CloseReason.String())
	b.WriteRune('|')
	b.WriteString(e.CodedErr.Error())
	return b.String()
}

// Template setup a string format template.
// Coder could compile the error object with formatting args later.
//
// Note that `ExtErr.Template()` had been overrided here
func (e *MqttError) Template(tmpl string) errors.Templater {
	_ = e.CodedErr.Template(tmpl)
	return e
}

// Formatf compiles the final msg with string template and args
//
// Note that `ExtErr.Template()` had been overridden here
func (e *MqttError) Formatf(args ...interface{}) errors.Templater {
	_ = e.CodedErr.Format(args...)
	return e
}

// Msg encodes a formattable msg with args into ExtErr
//
// Note that `ExtErr.Template()` had been overridden here
func (e *MqttError) Msg(msg string, args ...interface{}) *MqttError {
	_ = e.CodedErr.Msg(msg, args...)
	return e
}

// Attach attaches the nested errors into CodedErr
//
// Note that `ExtErr.Template()` had been overridden here
func (e *MqttError) Attach(errors ...error) *MqttError {
	e.CodedErr.AttachIts(errors...)
	return e
}

// Nest attaches the nested errors into CodedErr
//
// Note that `ExtErr.Template()` had been overridden here
func (e *MqttError) Nest(errors ...error) *MqttError {
	e.CodedErr.NestIts(errors...)
	return e
}

// AttachIts attaches the nested errors into CodedErr
func (e *MqttError) AttachIts(errors ...error) {
	e.CodedErr.AttachIts(errors...)
}

// NestIts attaches the nested errors into CodedErr
func (e *MqttError) NestIts(errors ...error) {
	e.CodedErr.NestIts(errors...)
}

func (e *MqttError) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, e.Error())
			e.Stack().Format(st, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(st, e.Error())
	case 'q':
		fmt.Fprintf(st, "%q", e.Error())
	}
}
