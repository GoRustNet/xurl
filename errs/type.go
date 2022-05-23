package errs

import "database/sql"

type Type int

const (
	TypeNone Type = iota
	TypeCommon
	TypeDb
	TypeExists
	TypeBcrypt
	TypeNotExists
	TypeShortUrl
	TypeFormBind
	TypeInvalidParam
	TypeParse
)

func DbError(cause error) *Error {
	return FromCause(TypeDb, cause)
}

func FromOptStrWithType(types Type, msgs ...string) *Error {
	msg := getDefinedMsg(types)
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return FromStrWithType(TypeExists, msg)
}

func ExistsError(msgs ...string) *Error {
	return FromOptStrWithType(TypeExists, msgs...)
}

func BcryptError(cause error) *Error {
	return FromCause(TypeBcrypt, cause)
}

func NotExistsError(msgs ...string) *Error {
	return FromOptStrWithType(TypeNotExists, msgs...)
}

func NotExistsOrDbError(cause error, msgs ...string) *Error {
	if cause == sql.ErrNoRows {
		return NotExistsError(msgs...)
	}
	return DbError(cause)
}
func ShortUrlError(cause error) *Error {
	return FromCause(TypeShortUrl, cause)
}
func FormBindError(cause error) *Error {
	return FromCause(TypeFormBind, cause)
}
func InvalidParam(msgs ...string) *Error {
	return FromOptStrWithType(TypeInvalidParam, msgs...)
}
func ParseError(cause error) *Error {
	return FromCause(TypeParse, cause)
}
