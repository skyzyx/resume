# Reminder of the steps to build the résumé

## Prerequisites

* Modern Python 3.x
* Pandoc

## Notes

The source Markdown (`README.md`) is _always_ the most up-to-date. It contains some HTML blocks with the following CSS classes: `.cv`, `.short`, and `.web`.

* Blocks with the `.web` class display on GitHub as a README, but is hidden in the PDF versions.
* Blocks with the `.cv` class display in the long-form "CV" style résumé. `.short` and `.web` are hidden.
* Blocks with the `.short` class display in the shorter-form "American" style résumé. `.cv` and `.web` are hidden.
* There are two "frame" templates which load the same generated-HTML, but each load different stylesheets to show/hide content as appropriate.

## Steps

1. Build the HTML from the Markdown.

    ```bash
    make build
    ```

1. One version (`resume.html`) will display the long-form "CV" style résumé. Specific styles are defined in `cv.css`.

    * all `<details>` blocks are expanded by default.

1. Second version (`resume-short.html`) will display the shorter-form "American" style résumé. Specific styles are defined in `short.css`. This version omits:

    * any jobs started more than 10 years ago
    * highlighted LinkedIn recommendations
    * generally speaking, everything else should be the same

1. Load each HTML file in a Chromium-based browser. (There's something weird going on with Safari's print rendering.)

1. Print using the system's print dialog (not the Chromium one).

    * Scale-down to 80%
    * Save each as PDF, overwriting the old version
