<h1 align="center">Welcome to gorendermd 👋</h1>
<p>
  <a href="#" target="_blank">
    <img alt="License: The Unlicense" src="https://img.shields.io/badge/License-The Unlicense-yellow.svg" />
  </a>
</p>

> Render Markdown in a browser via a local HTTP server

![preview](.github/images/out.gif)

## Install

```sh
go install github.com/ayuxsec/gorendermd@latest
```

otherwise [download a release for your platform](https://github.com/ayuxsec/gorendermd/releases)

## Usage Example

```sh
echo "# Hello World" | gorendermd
```

## Flags

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

## Author

- Github: [@ayuxsec](https://github.com/ayuxsec)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
