#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH="amd64"

# Define the URL for the binary based on OS
URL="https://github.com/pjcalvo/cvc-go/releases/latest/download/cvc-${OS}"

# Create a bin directory in the user's home if it doesn't exist
mkdir -p $HOME/bin

echo "🔽 Downloading cvc CLI for ${OS}/${ARCH}..."
curl -L "$URL" -o $HOME/bin/cvc

# Make the binary executable
chmod +x $HOME/bin/cvc

# Add the bin directory to the PATH if it's not already there
if ! grep -q "$HOME/bin" "$HOME/.bashrc"; then
  echo 'export PATH=$HOME/bin:$PATH' >> $HOME/.bashrc
  echo "Added $HOME/bin to PATH in .bashrc"
  source $HOME/.bashrc
  echo "🔄 Applied changes to .bashrc"
fi 

if ! grep -q "$HOME/bin" "$HOME/.zshrc"; then
  echo 'export PATH=$HOME/bin:$PATH' >> $HOME/.zshrc
  echo "Added $HOME/bin to PATH in .zshrc"
  echo "⚠️ Run 'source $HOME/.zshrc' to apply changes to your terminal session."
fi

echo "✅ Installed cvc to $HOME/bin/cvc"
echo "You can now run 'cvc' from anywhere!"
