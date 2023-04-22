package error

import (
	"runtime/debug"
)

type BaseError interface {
	Error() string
	Cause() string
	StackTrace() string
}

type baseError struct {
	msg        string
	cause      error
	stackTrace string
}

func (e *baseError) Error() string {
	return e.msg
}

func (e *baseError) Cause() string {
	if e.cause != nil {
		return e.cause.Error()
	} else {
		return ""
	}
}

func (e *baseError) StackTrace() string {
	return e.stackTrace
}

func NewBaseError(msg string) BaseError {
	return &baseError{msg: msg}
}

func NewBaseErrorWithCause(msg string, cause error) error {
	return &baseError{msg: msg, cause: cause, stackTrace: string(debug.Stack())}
}

type NotFoundError struct {
	*baseError
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{&baseError{msg: msg}}
}

type AuthError struct {
	*baseError
}

func NewAuthError(msg string, cause error) error {
	return &AuthError{&baseError{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}

type BadArgumentError struct {
	*baseError
}

func NewBadArgumentError(msg string) error {
	return &BadArgumentError{&baseError{msg: msg, stackTrace: string(debug.Stack())}}
}

func NewBadArgumentErrorWithCause(msg string, cause error) error {
	return &BadArgumentError{&baseError{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}

type ConflictError struct {
	*baseError
}

func NewConflictError(msg string) error {
	return &ConflictError{&baseError{msg: msg, stackTrace: string(debug.Stack())}}
}

func NewConflictErrorWithCause(msg string, cause error) error {
	return &ConflictError{&baseError{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}

type UnprocessableEntityError struct {
	*baseError
}

func NewUnprocessableEntityError(msg string) error {
	return &UnprocessableEntityError{&baseError{msg: msg, stackTrace: string(debug.Stack())}}
}

func NewUnprocessableEntityErrorWithCause(msg string, cause error) error {
	return &UnprocessableEntityError{&baseError{msg: msg, cause: cause, stackTrace: string(debug.Stack())}}
}
