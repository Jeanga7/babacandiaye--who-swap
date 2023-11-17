#!/bin/bash

extensions=(
    "mintlify.document"
    "ritwickdey.LiveServer"
    "fabiospampinato.vscode-todo-plus"
    "shakram02.bash-beautify"
    "rvest.vs-code-prettier-eslint"
    "MS-vsliveshare.vsliveshare"
    "dracula-theme.theme-dracula"
    "Blackboxapp.blackbox"
    "usernamehw.errorlens"
    "esbenp.prettier-vscode"
    "golang.go"
    "streetsidesoftware.code-spell-checker"
    "christian-kohler.path-intellisense"
    "esbenp.prettier-vscode"


)

for extension in "${extensions[@]}"
do
    code --install-extension "$extension"
done
