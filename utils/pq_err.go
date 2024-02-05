package utils

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/umardev500/spk/constants"
)

func ParsePostgresError(err error) error {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	code := pqErr.Code
	if code == constants.DuplicateKeyViolationCode {
		return constants.ErrorDuplicate
	}

	return nil
}

func CombinePqErr(from string, dest *string) {
	if dest != nil && from != "" {
		msg := fmt.Sprintf("%s. %s", from, *dest)
		*dest = msg
	}
}
