#! /usr/bin/env python

with open('README.md', 'r') as fin:
    with open('README.processed.md', 'w') as fout:
        content = fin.read().replace('<!-- <div class="short">', '<div class="short">').replace('</div> -->', '</div>')
        fout.write(content)
