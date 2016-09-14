#! /usr/bin/env bash

# Pandoc must be available on the path.
pandoc \
    -r markdown \
    -w html5 \
    --smart \
    --output=body.html \
    README.md

# Merge the Pandoc HTML with the page template.
python ./process.py

# Generate the PDF (problem with fonts)
# wkhtmltopdf \
#     --no-background \
#     --print-media-type \
#     --enable-local-file-access \
#     --enable-smart-shrinking \
#     --encoding "UTF-8" \
#     --page-size "Letter" \
#     --title "Ryan Parman • Résumé" \
#     resume.html \
#     ryanparman-resume.pdf
