#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all:
	@cat Makefile | grep "^[a-z]" | sed 's/://' | awk '{print $$1}'

#-------------------------------------------------------------------------------

.PHONY: build
build:
	pandoc -r gfm -w docx --output=ryanparman-resume.docx README.md
	pandoc -r gfm -w html5+smart --output=body.html README.md
	python ./process.py

.PHONY: serve
serve:
	php -S 0.0.0.0:4000 -t .
