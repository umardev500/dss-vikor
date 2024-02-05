package constants

import "errors"

var (
	ErrorNotAffected = errors.New("not affected")
	ErrorDuplicate   = errors.New("duplicate entry")
)
