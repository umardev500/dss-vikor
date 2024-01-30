package utils

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func LogBuilder(uuid uuid.UUID, msg string, data interface{}, err error) string {
	var msgStr = color.CyanString(msg)
	if err != nil {
		msgStr = color.RedString(msg)
	}

	logStr := fmt.Sprintf("uuid: %s | msg: %s | data: %v", color.GreenString(uuid.String()), msgStr, data)
	return logStr
}
