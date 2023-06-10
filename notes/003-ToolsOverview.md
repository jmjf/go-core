# Tools overview

## Installing

Go to <https://go.dev>, Getting Started (currently), and install the appropriate version for your platform.

Or use the Docker container in a dev container, like I'm doing.

## go command

The Go CLI is `go`.

`go help` shows the commands and other details.

`go help <command>` for more details on the command.

The documentation on commands is decently detailed (and also available online a <https://go.dev>).

Some key commands:

* `build` -- used this already to compile logfilter01 and web01
* `doc` -- pulls documentation comments from files and prints documentation
  * `go doc json`, `go doc json.Encoder`, `go doc json.Encoder.Encode` to see examples of different levels of documentation
  * For standard packages (like `json`) same docs are available on the website
  * VS Code and similar LSP editors make the CLI option less critical, but the docs need to be there and knowing this is available can be helpful
* `get` -- pulls and installs packages and dependencies -- similar to `npm install` for Node folks
* `fmt` -- formats package sources to golang standards (opinionated, but reasonable)
  * Default formatting uses tabs, so 2 spaces vs 4 spaces wars are easily solved (but not tabs vs. spaces folks)
* `test` -- basic testing tool and test runner
* `run` -- compiles and runs an application in a temporary directory

I definitely want to know more about `doc` and how to write comments that it can use. Also `test` because testing is critical for anything real world.

## Editors and VS Code setup

Can use VS Code, vim, IntelliJ and others. Pick what works for you. I'm using VS Code, so notes below are for VS Code. Find what works for your editor and preferences.

* The key extension for VS Code is the Go extension by Google.
  * After installing, go to the command palette, find and select "Go: Install/Update Tools"
  * Select all tools and install them
  * These are tools the extension will want.
  * If you don't do this, the extension will remind you when it needs them.
  * In the extension settings, Go > Tools Management : Auto Update seems to allow automatic update to keep them fresh.
* I'm using dev containers, so have the Dev Containers plugin installed.
* I also have (personal preferences)
  * Theme/color scheme -- pick your preference
  * A spell checker (Code Spell Checker by Street Side Software) to reduce basic spelling errors in code (far too common and can be misleading or confusing)
  * markdownlint (David Anson), because I write markdown and want it to be uniform (pick what works for you)
  * GitLens to improve the basic version control experience (other options exist)

## Project setup and organization

* Create a directory that will hold the code.
* `go mod init <modulename>` to setup the module
  * `go mod init github.com/jmjf/example` (assumes you're using a repo named example in my GitHub account)
  * Creates the basic `go.mod` file
  *
* `main.go` is the start of the application

Based on this discussion, I removed `go.work`. I can run within a specific directory with `go run` or build with `go build`, so the workspace thing was apparently misleading guidance from tooling.

**COMMIT: DOCS: add notes on tools; clean up workspace/module structure based on what I learned**
