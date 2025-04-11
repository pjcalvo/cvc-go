package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	git "github.com/pjcalvo/cvc-go/internal/commit"
)

func main() {
	// Check if we're inside a Git repository
	if !git.IsGitRepo() {
		fmt.Println("❌ You must be inside a Git repository to use this tool.")
		os.Exit(1)
	}

	var selectedType string
	survey.AskOne(&survey.Select{
		Message: "Select the type of change:",
		Options: git.TypeOptions,
	}, &selectedType)

	// Prompt for scope (optional)
	var scope string
	survey.AskOne(&survey.Input{
		Message: "Scope (e.g. JIRA-001, api, dashboard):",
	}, &scope)

	// Prompt for commit message
	var message string
	survey.AskOne(&survey.Input{
		Message: "Commit message:",
	}, &message)

	// Parse selected type
	parts := strings.Split(selectedType, " ")
	emoji := parts[0]
	typ := parts[1]

	// Format commit message
	commitMsg := git.FormatCommit(emoji, typ, scope, message)

	git.PreviewCommit(commitMsg)

	// Confirm
	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to commit this?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Commit cancelled.")
		return
	}

	fmt.Println("Executing command:")
	fmt.Printf("   git commit -m %q\n", commitMsg)
	fmt.Println("────────────────────────────────────────────")

	// Commit
	cmd := exec.Command("git", "commit", "-m", commitMsg)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Commit failed:")
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println("✅ Commit successful!")
}
