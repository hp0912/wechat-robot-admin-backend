package utils

import (
	"fmt"
	"strings"
)

func MySQLStringLiteral(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, `'`, `\'`)
	return fmt.Sprintf("'%s'", value)
}
