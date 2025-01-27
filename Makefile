##
# Default Makefile Settings
# These will run by default, unless the Makefile re-defines them.
#
# Helpful links:
#
# * https://makefiletutorial.com
# * https://stackoverflow.com/questions/11958626/make-file-warning-overriding-commands-for-target/49804748
# * https://www.gnu.org/software/make/manual/html_node/Make-Control-Functions.html
##

# Shell
SHELL:=bash

# Color stuff
HASH := \#
FOREGROUND:=$(HASH)FFFC67
YELLOW:=gum style --foreground='$(FOREGROUND)' --bold
WHITE:=gum style --bold
ERROR:=gum style --foreground='$(HASH)FFFFFF' --background='$(HASH)CC0000' --bold --padding='0 1'
BORDER:=gum style --foreground='$(HASH)FFFFFF' --border='rounded' --border-foreground='240' --padding='0 1' --margin='1 0'
HEADER:=gum style --foreground='$(FOREGROUND)' --border='rounded' --border-foreground='12' --bold --width='78' --padding='0 1' --margin='1 0 0 0'

# Tooling
GO=$(shell which go)

#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all: help

.PHONY: help
## help: [help]* Prints this help message.
help:
	@ $(WHITE) "Usage:"
	@ echo ""
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | \
		while IFS= read -r line; do \
			if [[ "$$line" == *"]*"* ]]; then \
				$(YELLOW) "$$line"; \
			else \
				echo "$$line"; \
			fi; \
		done | LC_ALL=C sort --ignore-nonprinting --stable --ignore-leading-blanks --field-separator=' ' --key=4

#-------------------------------------------------------------------------------
# Mac-specific functions

.PHONY: clean-ds
## clean-ds: [clean] Clean .DS_Store files.
clean-ds:
	@ $(HEADER) "=====> Cleaning .DS_Store files..."
	find . -type f -name ".DS_Store" | xargs -I% rm -fv "%"

.PHONY: clean-resume
## clean-resume: [clean] Clean resumes directory.
clean-resume:
	@ $(HEADER) "=====> Cleaning resumes directory..."
	rm -Rfv ./resumes

.PHONY: clean-html
## clean-html: [clean] Clean resumes directory.
clean-html:
	@ $(HEADER) "=====> Cleaning HTML render files..."
	rm -fv ./render/*.html

.PHONY: clean
## clean: [clean]* Run cleaning tasks.
clean: clean-ds clean-resume clean-html

#-------------------------------------------------------------------------------

.PHONY: build
## build: [build]* Run building tasks.
build:
	@ $(GO) run main.go generate
	@ cp -fv ./resumes/ryanparman-general-cv.md ./README.md
