#!/bin/bash

set -e

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH="amd64"

URL="https://github.com/pjcalvo/cvc-go/releases/latest/download/cvc-${OS}"

echo "🔽 Downloading cvc CLI for ${OS}/${ARCH}..."
curl -L "$URL" -o /usr/local/bin/cvc
chmod +x /usr/local/bin/cvc

echo "✅ Installed cvc to /usr/local/bin/cvc"
