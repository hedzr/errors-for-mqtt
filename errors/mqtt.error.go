package errors

import (
	"github.com/hedzr/errors"
)

// Code is a type of error codes wrapped `errors.Code`
type Code errors.Code

// New create a new *MqttError object
func (c Code) New(msg string, args ...interface{}) *MqttError {
	return NewCodedError(c).Msg(msg, args...)
}

func (c Code) NewE(e error, msg string, args ...interface{}) *MqttError {
	return NewCodedError(c).Msg(msg, args...).Attach(e)
}

// Register register a code and its token string for using later
func (c Code) Register(codeName string) (errno Code) {
	return Code(errors.Code(c).Register(codeName))
}

//

const (
	FIRST_ERROR Code = iota + 9999
	ErrCodeWrongState
	ErrCodeOverflow
	ErrCodePacketCorrupt
	ErrCodePacketIllegal
	ErrCodeWrongParserDefinition
	ErrCodeWrongPrerequisiteState
	ErrCodeWrongPacketType
	ErrCodeTooMuchNestedStates
	ErrCodeConnectTimeout
	ErrCodeUnauthorized
	ErrCodeUnsupportedProtocolLevel
	ErrCodeBadIO
	ErrCodeErrors // generic nested errors
	ErrCodeMqttStateWrong
	ErrCodeMqttBadClientId
	// ErrCodeIllegalClientId
	ErrCodePacketIdentifierExisted
	ErrCodeBadTopicFilter
	ErrCodeBadTopicName
	ErrCodePacketIncomplete
	ErrCodeStateLoops
	LASTERROR

	// ErrorNumberAbsent errors.Code = FIRST_ERROR - 1
	// NoError           errors.Code = iota
	OK      Code = iota
	NoError Code = iota

	MaxNestedLoops = 32
)

var (
	ErrOverflow                 = ErrCodeOverflow.New("data overflow")
	ErrWrongState               = ErrCodeWrongState.New("wrong state")
	ErrStateLoops               = ErrCodeStateLoops.New("unconditional state migrates nested loops")
	ErrConnectTimeout           = ErrCodeConnectTimeout.New("connect timeout")
	ErrUnauthorized             = ErrCodeUnauthorized.New("unauthorized")
	ErrUnsupportedProtocolLevel = ErrCodeUnsupportedProtocolLevel.New("unsupported protocol level")
	// ErrPacketIncomplete       = (ErrCodePacketIncomplete).New("packet incomplete")
)

func init() {

	// register the error codes with code name

	ErrCodeOverflow.Register("DATA_OVERFLOW")
	ErrCodeWrongState.Register("WRONG_STATE")
	ErrCodePacketCorrupt.Register("")
	ErrCodePacketIllegal.Register("")
	ErrCodeWrongParserDefinition.Register("")
	ErrCodeWrongPrerequisiteState.Register("")
	ErrCodeWrongPacketType.Register("")
	ErrCodeTooMuchNestedStates.Register("")
	ErrCodeConnectTimeout.Register("connect timeout")
	ErrCodeUnauthorized.Register("UNAUTHORIZED")
	ErrCodeUnsupportedProtocolLevel.Register("unsupported protocol level")
	ErrCodeBadIO.Register("")
	ErrCodeErrors.Register("")
	ErrCodeMqttStateWrong.Register("")
	ErrCodeMqttBadClientId.Register("")
	// ErrCodeIllegalClientId
	ErrCodePacketIdentifierExisted.Register("")
	ErrCodeBadTopicFilter.Register("")
	ErrCodeBadTopicName.Register("")
	ErrCodePacketIncomplete.Register("PACKET_INCOMP")
	ErrCodeStateLoops.Register("unconditional state migrates nested loops")
}

type MqttError struct {
	CloseReason
	errors.CodedErr
}

func (e *MqttError) Reason(reason CloseReason) *MqttError {
	e.CloseReason = reason
	return e
}

func (e *MqttError) GetReason() CloseReason {
	return e.CloseReason
}
