# git-prompt

git-prompt is a small tool for generating beautifully short yet informative git
status strings mainly for your shell prompt to make it better than ever before.

# Installation

You must have git and Go toolchain properly set up. Simply run `go get
github.com/rce/git-prompt` and you can start using the `git-prompt` command.

# Usage

Call `git-prompt` and it returns the status string for the repository you are
currently in. At the moment it only contains the current branch name. If you
are in a directory that is not part of a git repository (or git status
otherwise returns non-zero exit code), the program will not output anything and
returns a non-zero exit code.

## Example 

This is how I use it in my zsh config. As far as I know only the precmd hook
function which is called right before each prompt is zsh exclusive feature but
I might be wrong.

```shell
PROMPT_TEMPLATE="[%n@%m] %~ GIT_STATUS
$ "

# Generates the prompt
function makeprompt() {
	GIT_STATUS=$(git-prompt)
	if [ $? -ne 0 ]; then
		GIT_STATUS=""
	fi

	echo $PROMPT_TEMPLATE | sed -e "s/GIT_STATUS/$GIT_STATUS/"
}

function precmd() {
	PROMPT=$(makeprompt)
}

PROMPT=$(makeprompt)
```
