#!/bin/bash

extensions=(
    "mintlify.document"
    "ritwickdey.LiveServer"
    "fabiospampinato.vscode-todo-plus"
    "shakram02.bash-beautify"
    "rvest.vs-code-prettier-eslint"
    "MS-vsliveshare.vsliveshare"
    # "dracula-theme.theme-dracula"
    # "Blackboxapp.blackbox"
    "usernamehw.errorlens"
)

for extension in "${extensions[@]}"
do
    code --install-extension "$extension"
done
