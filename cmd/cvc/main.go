package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pjcalvo/cvc-go/internal/commit"
)

func main() {
	typeOptions := []string{
		"✨ feat - A new feature",
		"🐛 fix - A bug fix",
		"📝 docs - Documentation only changes",
		"🎨 style - Code style changes (formatting, etc)",
		"♻️ refactor - Code refactor, no behavior change",
		"⚡ perf - Performance improvements",
		"✅ test - Adding or updating tests",
		"🔧 chore - Build or tooling changes",
		"🚀 ci - CI related changes",
	}

	var selectedType string
	survey.AskOne(&survey.Select{
		Message: "Select the type of change:",
		Options: typeOptions,
	}, &selectedType)

	// Prompt for scope (optional)
	var scope string
	survey.AskOne(&survey.Input{
		Message: "Scope (e.g. auth, api, dashboard):",
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
	commitMsg := commit.FormatCommit(emoji, typ, scope, message)

	commit.PreviewCommit(commitMsg)

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

	// Optionally stage all changes
	exec.Command("git", "add", ".").Run()

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
