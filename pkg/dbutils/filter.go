package dbutils

import (
	"fmt"
	"strings"
)

func AddFilter(str, connector, key string, paramCount int) string {
	if !strings.Contains(strings.ToLower(str), "where") {
		str = fmt.Sprintf("%s WHERE %s = $%d", str, key, paramCount)
		return str
	}

	str = fmt.Sprintf("%s %s %s = '%d'", str, connector, key, paramCount)
	return str
}
