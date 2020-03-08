# GitUpdate [![Thanks](http://bit.ly/saythankss)](https://github.com/users/nikitavoloboev/sponsorship)

> Commit and push updated files with file names as commit message

## Install

```Bash
go get -u github.com/nikitavoloboev/gitupdate
```

## Usage

You can either use it by passing it a file path (with git repo) that you want to commit. i.e.

`gitupdate /Users/nikivi/src/cli/gitupdate`

Or if you are already in the git directory you want to commit, run:

`gitupdate .`

This will add all files that have changed since last commit and will include all the file names (without extension) as the commit message. [Example use](https://github.com/nikitavoloboev/knowledge/commits/master).

If you want to only consider top level folders for the commit message, use the `--top` (or `-t` for short) flag.

I personally find it very useful for personal repos where commits are not that important (i.e. notes/dotfiles/docs). I treat these repos as write only so there is no point in wasting time writing a commit message.

## Contributing

See [contribution guidelines](CONTRIBUTING.md#readme).

## Thank you

You can support me on [GitHub](https://github.com/users/nikitavoloboev/sponsorship) or look into [other projects](https://nikitavoloboev.xyz/projects) I shared.

[![MIT](https://img.shields.io/badge/license-MIT-0a0a0a.svg?style=flat&colorA=0a0a0a)](LICENSE) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)
