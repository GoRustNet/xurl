package errs

import (
	"encoding/json"

	"github.com/GoRustNet/xurl/str"
)

type Error struct {
	Cause   error  `json:"cause"`
	Message string `json:"message"`
	Type    Type   `json:"type"`
}

func (e *Error) Error() string {
	if str.IsNotEmpty(e.Message) {
		return e.Message
	}
	return getDefinedMsg(e.Type)
}

func (e *Error) Debug() string {
	buf, _ := json.Marshal(e)
	return string(buf)
}

func NewError(types Type, cause error, message string) *Error {
	return &Error{
		Cause:   cause,
		Message: message,
		Type:    types,
	}
}

func FromStrWithType(types Type, message string) *Error {
	return NewError(types, nil, message)
}
func FromStr(message string) *Error {
	return FromStrWithType(TypeCommon, message)
}

func FromCause(types Type, cause error) *Error {
	return NewError(types, cause, str.Empty)
}

func FromCauseWithoutType(cause error) *Error {
	return FromCause(TypeCommon, cause)
}
