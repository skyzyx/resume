#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all:
	@cat Makefile | grep "^[a-z]" | sed 's/://' | awk '{print $$1}'

#-------------------------------------------------------------------------------

.PHONY: build
build:
	pandoc -r gfm -w docx --output=ryanparman-cv.docx README.md
	python3 ./pre-process.py
	pandoc -r gfm -w html5+smart --columns=20000 --eol=lf --output=body.html README.processed.md
	python3 ./post-process.py

.PHONY: serve
serve:
	php -S 0.0.0.0:4000 -t .
