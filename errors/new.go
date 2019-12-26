package errors

//
//
//

// New MqttError error object with message and allows attach more nested errors
func New(msg string, args ...interface{}) *MqttError {
	e := &MqttError{}
	return e.Msg(msg, args...)
}

// NewTemplate ExtErr error object with string template and allows attach more nested errors
func NewTemplate(tmpl string) *MqttError {
	e := &MqttError{}
	return e.Template(tmpl)
}

// NewWithError MqttError error object with nested errors
func NewWithError(errs ...error) *MqttError {
	return New("unknown error").Attach(errs...)
}

// NewCodedError error object with nested errors
func NewCodedError(code Code, errs ...error) *MqttError {
	e := &MqttError{}
	return e.Code(code).Attach(errs...)
}
