# Reminder of the steps to build the résumé

## Prerequisites

* `brew install php`
* Modern Python 3.x
* Pandoc
* Google Chrome

## Notes

The source Markdown (`README.md`) is _always_ the most up-to-date. It contains some HTML blocks with the following CSS classes: `.cv`, `.short`, and `.web`.

* Blocks with the `.web` class display on GitHub as a README, but is hidden in the PDF versions.
* Blocks with the `.cv` class display in the long-form "CV" style résumé. `.short` and `.web` are hidden.
* Blocks with the `.short` class display in the shorter-form "American" style résumé. `.cv` and `.web` are hidden.
* There are two "frame" templates which load the same generated-HTML, but each load different stylesheets to show/hide content as appropriate.

## Steps

1. Run the local web server.

    ```bash
    make serve
    ```

1. Generate the Microsoft Word, HTML, and PDF versions from the Markdown.

    ```bash
    make build pdf
    ```
