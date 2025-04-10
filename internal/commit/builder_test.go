package commit

import (
	"testing"
)

func TestCommitFormat(t *testing.T) {
	emoji := "✨"
	typ := "feat"
	scope := "auth"
	message := "add login"

	expected := "✨ feat(auth): add login"
	result := FormatCommit(emoji, typ, scope, message)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
