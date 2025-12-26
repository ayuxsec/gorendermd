package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	rawMD           string
	showVersion     = flag.Bool("version", false, "print current version and exit")
	port            = flag.Int("port", 0, "Port to listen on (0 = random)")
	skipOpenBrowser = flag.Bool("skip-open-browser", false, "Skip Auto Open the URL in a Browser")
	browser         = flag.String("browser", "firefox", "Browser to open (runs command <browser> <url> in the cli)")
	browserTimeout  = flag.Int("browser-timeout", 30, "Timeout in seconds to wait for the brownser to respond")
)

func parseArgs() error {
	flag.Usage = printHelp
	flag.Parse()

	// early exit if showVersion is true
	if *showVersion {
		log.Printf("Current gorendermd version is %s", VersionString())
		os.Exit(0)
	}

	// Case 1: stdin is piped
	if isPiped() {
		scanner := bufio.NewScanner(os.Stdin)
		var builder strings.Builder

		for scanner.Scan() {
			builder.WriteString(scanner.Text())
			builder.WriteString("\n")
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading piped input: %w", err)
		}

		rawMD = builder.String()
		return nil
	}

	// Case 2: file input
	if flag.NArg() != 1 {
		printHelp()
		return fmt.Errorf("expected exactly one markdown file")
	}

	mdFilePath := flag.Arg(0)

	data, err := os.ReadFile(mdFilePath)
	if err != nil {
		return err
	}

	rawMD = string(data)
	return nil
}

func printHelp() {
	fmt.Fprint(os.Stderr, "Gorendermd - Render Markdown in a browser via a local HTTP server\n\n")
	fmt.Fprint(os.Stderr, "Usage:\n")
	fmt.Fprint(os.Stderr, "  gorendermd [options] <md_file_path>\n")
	fmt.Fprint(os.Stderr, "  cat <md_file_path> | gorendermd [options]\n\n")
	fmt.Fprint(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprint(os.Stderr, "Examples:\n")
	fmt.Fprint(os.Stderr, "  gorendermd example.md\n")
	fmt.Fprint(os.Stderr, "  cat example.md | gorendermd -skip-open-browser\n")
	fmt.Fprint(os.Stderr, `  gorendermd -browser "brave-browser" -port "8080" -timeout 60 example.md`+"\n\n")
}

func isPiped() bool {
	fi := must(os.Stdin.Stat())
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
