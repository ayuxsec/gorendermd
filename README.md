# gorendermd

Render Markdown in a browser via a local HTTP server

![preview](.github/images/out.gif)

## Install

If you have Go installed and configured:

```bash
go install github.com/ayuxsec/gorendermd@latest
```

otherwise [download a release for your platform](https://github.com/ayuxsec/gorendermd/releases)

## Usage

```console
$ ./gorendermd -h
Gorendermd - Render Markdown in a browser via a local HTTP server

Usage:
  gorendermd [options] <md_file_path>
  cat <md_file_path> | gorendermd [options]

Options:
  -browser string
        Browser to open (runs command <browser> <url> in the cli) (default "firefox")
  -browser-timeout int
        Timeout in seconds to wait for the brownser to respond (default 30)
  -port int
        Port to listen on (0 = random)
  -skip-open-browser
        Skip Auto Open the URL in a Browser
  -version
        print current version and exit

Examples:
  gorendermd example.md
  cat example.md | gorendermd -skip-open-browser
  gorendermd -browser "brave-browser" -port "8080" -timeout 60 example.md
```

## Contributing

Sure, PRs are welcome!
