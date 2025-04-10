package commit

import "fmt"

// FormatCommit builds the conventional commit string
func FormatCommit(emoji, typ, scope, message string) string {
	scopePart := ""
	if scope != "" {
		scopePart = fmt.Sprintf("(%s)", scope)
	}
	return fmt.Sprintf("%s %s%s: %s", emoji, typ, scopePart, message)
}
