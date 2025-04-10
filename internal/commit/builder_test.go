package commit

import (
	"testing"
)

func TestFormatCommit(t *testing.T) {
	tests := []struct {
		emoji   string
		typ     string
		scope   string
		message string
		want    string
	}{
		{"✨", "feat", "auth", "add login support", "feat ✨ (auth): add login support"},
		{"🐛", "fix", "", "null pointer error", "fix 🐛 : null pointer error"},
		{"📝", "docs", "readme", "update instructions", "docs 📝 (readme): update instructions"},
	}

	for _, tt := range tests {
		got := FormatCommit(tt.typ, tt.emoji, tt.scope, tt.message)
		if got != tt.want {
			t.Errorf("FormatCommit() = %q; want %q", got, tt.want)
		}
	}
}
