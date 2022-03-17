package dbutils

import (
	"fmt"
	"strings"
)

func AddFilter(str, connector, key, comparator string, paramCount int) string {
	return AddCustomFilter(str, connector, key, comparator, fmt.Sprintf("$%d", paramCount))
}

func AddCustomFilter(str, connector, key, comparator, param string) string {
	if !strings.Contains(strings.ToLower(str), "where") {
		str = fmt.Sprintf("%s WHERE %s %s %s", str, key, comparator, param)
		return str
	}

	str = fmt.Sprintf("%s %s %s %s %s", str, connector, key, comparator, param)
	return str
}
