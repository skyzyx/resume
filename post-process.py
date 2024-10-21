#! /usr/bin/env python

with open('page-cv.tmpl', 'r') as tin:
    with open('body.html', 'r') as fin:
        with open('resume.html', 'w') as fout:
            content = tin.read().replace('@@BODY@@', fin.read())
            fout.write(content)

with open('page-short.tmpl', 'r') as tin:
    with open('body.html', 'r') as fin:
        with open('resume-short.html', 'w') as fout:
            content = tin.read().replace('@@BODY@@', fin.read())
            fout.write(content)
