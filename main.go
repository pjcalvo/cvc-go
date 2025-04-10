package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// Define conventional commit types with emojis
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
		Message: "Scope (e.g. auth, api, dashboard, PROJ-123):",
	}, &scope)

	// Prompt for commit message
	var message string
	survey.AskOne(&survey.Input{
		Message: "Commit message:",
	}, &message)

	// Parse the type and emoji
	emoji := strings.Split(selectedType, " ")[0]
	typeKey := strings.Split(selectedType, " ")[1]

	// Build commit message with optional scope
	scopePart := ""
	if scope != "" {
		scopePart = fmt.Sprintf("(%s)", scope)
	}
	commit := fmt.Sprintf("%s %s%s: %s", emoji, typeKey, scopePart, message)

	// Show commit preview
	fmt.Println("\n📝 Commit preview:")
	fmt.Println("---------------------------------")
	fmt.Println(commit)
	fmt.Println("---------------------------------\n")

	// Confirm commit
	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to commit this?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Commit cancelled.")
		return
	}

	// Optional: stage all changes
	exec.Command("git", "add", ".").Run()

	// Run git commit
	cmd := exec.Command("git", "commit", "-m", commit)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Commit failed:")
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println("✅ Commit successful!")
}
