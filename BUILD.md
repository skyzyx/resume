# How to run/build

Generates multiple flavors of résumé into multiple output formats.

## Prerequisites

* [Xcode CLI Tools](https://github.com/northwood-labs/macos-for-development/wiki/Installing-the-Xcode-CLI-Tools) — Includes `make` and other core CLI tools.
* [Homebrew](https://github.com/northwood-labs/macos-for-development/wiki/Installing-Homebrew) — Package manager for macOS.
* [Go](https://go.dev) — The Go programming language and toolchain.
* [Google Chrome](https://www.google.com/chrome/) — Web browser with a programmatic headless mode and reliable PDF generation.
* [Pandoc](https://pandoc.org) — Converts between different textual formats.
* [exiftool](https://exiftool.org) — Edits metadata for images and PDFs.
* [gum](https://github.com/charmbracelet/gum) — Used by the `Makefile` for glamming-up the output.

For Mac users with Hombrew already installed:

```bash
brew update
brew install exiftool go gum pandoc
brew install --cask google-chrome
```

## Usage

After making a change to any of the `.html` or `.gohtml` files, you'll need to run the generation step.

```bash
make build
```

## Flow

The text content of the résumé is written using Go template files.

### Templates

These live in `cmd/templates/`. There are a few boolean variables that can be used in these templates:

* `.IsAll` — This is for the general-purpose (not targeted) roles.
* `.IsCloud` — This is for the Cloud Engineering, DevOps, and Site Reliability Engineering roles.
* `.IsSDE` — This is for the Software Engineering and DevTools roles.
* `.IsTPM` — This is for the PM/TPM roles.

### Steps

1. Looping over each role, the appropriate templates are executed to produce the Markdown-formatted résumé for that role.

1. Once we have the Markdown, we pass it to Pandoc to generate versions in Microsoft Word, OpenDocument, and HTML formats.

1. Once we have the HTML versions, we launch a local web server, load the local web URL in a headless Chrome process, generate a PDF, then terminate the local web server.

1. Once we have the PDF versions, we use `exiftool` to cleanup the metadata. This way, if someone downloads resumes to their computer, and searches for them, the results should match.
