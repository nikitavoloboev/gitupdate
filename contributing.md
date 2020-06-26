# Contributing

Thank you for taking the time to contribute! ♥️ You can:

- Submit [bug reports or feature requests](../../issues/new/choose). Contribute to discussions. Fix [open issues](../../issues).
- Improve docs, the code and more! Any idea is welcome.

## Run project

1. Clone repo
2. If you use [VSCode](https://code.visualstudio.com) with [Go](https://github.com/microsoft/vscode-go) plugin, it will install all Go dependencies for you in the background when you open the project.
3. Edit the code & run it with `go run .`.

I use [watchexec](https://github.com/watchexec/watchexec) to develop.

Running `watchexec --exts go "echo -- && go run ."` will automatically rerun `go run .` for you on every Go file changed.
