<h1 align="center">gorendermd</h1>
<p align="center">Render Markdown in a browser via a local HTTP server</p>
<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/ayuxsec/gorendermd?style=flat-square">
  <img src="https://img.shields.io/github/license/ayuxsec/gorendermd?style=flat-square">
</p>

## Install

```sh
go install github.com/ayuxsec/gorendermd@latest
```

otherwise [download a release for your platform](https://github.com/ayuxsec/gorendermd/releases)

## Usage Example

```sh
echo "# Hello World" | gorendermd
```

![preview](.github/images/out.gif)

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
