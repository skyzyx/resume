#-------------------------------------------------------------------------------
# Running `make` will show the list of subcommands that will run.

all:
	@cat Makefile | grep "^[a-z]" | sed 's/://' | awk '{print $$1}'

#-------------------------------------------------------------------------------

.PHONY: serve
serve:
	php -S 0.0.0.0:4000 -t .

.PHONY: build
build:
	pandoc -r gfm -w docx --output=ryanparman-cv.docx README.md
	python3 ./pre-process.py
	pandoc -r gfm -w html5+smart --columns=20000 --eol=lf --output=body.html README.processed.md
	python3 ./post-process.py

.PHONY: pdf
pdf:
	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
		--headless \
		--virtual-time-budget=5000 \
		--no-pdf-header-footer \
		--print-to-pdf="./ryanparman-cv.pdf" \
		http://0.0.0.0:4000/resume.html

	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
		--headless \
		--virtual-time-budget=5000 \
		--no-pdf-header-footer \
		--print-to-pdf="./ryanparman-resume-short.pdf" \
		http://0.0.0.0:4000/resume-short.html

	exiftool -all:all= ryanparman-*.pdf -overwrite_original

	exiftool \
		-Title="Ryan Parman résumé" \
		-Author="Ryan Parman" \
		-Subject="Cloud-native engineering leader with a focus on reliability, scalability, and security for the modern web." \
		-Keywords="cloud, engineer, devops, devsecops, director, principal, linux, go, golang, python, bash, javascript, aws, amazon web services" \
		ryanparman-*.pdf
