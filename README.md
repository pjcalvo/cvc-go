# cvc 📝✨

**cvc** is an interactive CLI tool for generating beautiful, emoji-powered [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) with scope, structure, and style.

> ✅ Works on MacOS/Linux  
> 💡 Powered by Go  
> 📦 Easy to install via `go install` or `curl | bash`

---

## ✨ Features

- 🧠 Interactive prompts for type, scope, and message
- 🎨 Emojis for commit type (configurable later)
- 💬 Commit preview with `git status`
- ✅ Safe commit confirmation
- 🔧 Minimal dependencies
- 📄 Fully compliant with Conventional Commits spec

---

## 📦 Installation

### 🛠️ With Go

```sh
go install github.com/pjcalvo/cvc-go/cmd/cvc@latest
```

_Make sure your $GOPATH/bin is in your $PATH._

### 🌀 With curl (coming soon) 

```sh
curl -sSf https://raw.githubusercontent.com/pjcalvo/cvc-go/main/install.sh | sh
```