# GitUpdate [![Thanks](http://bit.ly/saythankss)](https://github.com/users/nikitavoloboev/sponsorship)

> Commit and push updated files with file names as commit message

## Install

```Bash
go install github.com/nikitavoloboev/gitupdate@latest
```

## Usage

You can either use it by passing it a file path (with git repo) that you want to commit. i.e.

`gitupdate /Users/nikivi/src/cli/gitupdate`

Or if you are already in the git directory you want to commit, run:

`gitupdate .`

This will add all files that have changed since last commit and will include all the file names (without extension) as the commit message. [Example use](https://github.com/nikitavoloboev/knowledge/commits/main).

If you want to only consider top level folders for the commit message, use the `--top` (or `-t` for short) flag.

I personally find it very useful for personal repos where commits are not that important (i.e. notes/dotfiles/docs). I treat these repos as write only so there is no point in wasting time writing a commit message.

## Run

1. Clone repo
2. If you use [VSCode](https://code.visualstudio.com) with [Go](https://github.com/microsoft/vscode-go) plugin, it will install all Go dependencies for you in the background when you open the project.
3. Edit the code & run it with `go run .`.

I use [watchexec](https://github.com/watchexec/watchexec) to develop.

Running `watchexec --exts go "echo -- && go run ."` will automatically rerun `go run .` for you on every Go file changed.

## Contribute

Have idea how to improve or something is broken?

Search for [existing issues](../../issues) or open [new one](../../issues/new/choose).

PRs with code/docs changes are welcome.

## Thank you

You can support me on [GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects) I shared.

I also have [personal Discord](https://discord.com/invite/TVafwaD23d) you can join for more indepth discussions.

[![MIT](http://bit.ly/mitbadge)](https://choosealicense.com/licenses/mit/) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)
