package constants

import "errors"

type ErrType error

var (
	ErrorNotAffected = errors.New("not affected")
)
