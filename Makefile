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

.PHONY: clean
## clean: [clean]* Run cleaning tasks.
clean: clean-ds

#-------------------------------------------------------------------------------

.PHONY: build
## build: [build]* Run building tasks.
build:
	@ $(GO) run main.go generate
	@ cp -fv ./resumes/ryanparman-general-cv.md ./README.md

# .PHONY: build
# build:
# 	pandoc -r gfm -w docx --output=ryanparman-cv.docx README.md
# 	python3 ./pre-process.py
# 	pandoc -r gfm -w html5+smart --columns=20000 --eol=lf --output=body.html README.processed.md
# 	python3 ./post-process.py

# .PHONY: pdf
# pdf:
# 	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
# 		--headless \
# 		--virtual-time-budget=5000 \
# 		--no-pdf-header-footer \
# 		--print-to-pdf="./ryanparman-cv.pdf" \
# 		http://0.0.0.0:4000/resume.html

# 	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
# 		--headless \
# 		--virtual-time-budget=5000 \
# 		--no-pdf-header-footer \
# 		--print-to-pdf="./ryanparman-resume-short.pdf" \
# 		http://0.0.0.0:4000/resume-short.html

# 	exiftool -all:all= ryanparman-*.pdf -overwrite_original

# 	exiftool \
# 		-Title="Ryan Parman résumé" \
# 		-Author="Ryan Parman" \
# 		-Subject="Cloud-native engineering leader with a focus on reliability, scalability, and security for the modern web." \
# 		-Keywords="cloud, engineer, devops, devsecops, director, principal, linux, go, golang, python, bash, javascript, aws, amazon web services" \
# 		ryanparman-*.pdf
