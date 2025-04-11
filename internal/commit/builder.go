package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// FormatCommit builds the conventional commit string
func FormatCommit(emoji, typ, scope, message string) string {
	scopePart := ""
	if scope != "" {
		scopePart = fmt.Sprintf("(%s)", scope)
	}
	return fmt.Sprintf("%s%s: %s %s", typ, scopePart, emoji, message)
}

func PreviewCommit(commitMsg string) {
	commitStatus, err := exec.Command("git", "status", "--short").Output()
	if err != nil {
		fmt.Println("❌ Failed to get git status:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\n📝 Commit Message:")
	fmt.Printf("   %s\n", commitMsg)

	fmt.Println("\n📂 Git Status:")
	lines := strings.Split(strings.TrimSpace(string(commitStatus)), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		fmt.Printf("   %s\n", line)
	}

	fmt.Println("\n────────────────────────────────────────────")
}

// IsGitRepo checks if the current directory is inside a Git repository
func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}
