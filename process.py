#! /usr/bin/env python

with open('page.tmpl', 'r') as tin:
    with open('body.html', 'r') as fin:
        with open('resume.html', 'w') as fout:
            content = tin.read().replace('@@BODY@@', fin.read())
            fout.write(content)
