# Cobra Demo

## Purpose

Testing out how cobra CLI's are parsed.

## Overview

To use this repo, you can run something like this (from the directory where you cloned this repo to):

```bash
go build . && ./cobrademo commaTest --slice 'foo=bar,baz' --array 'foo=bar,baz' --slice 'a=b,c' --array 'a=b,c'
```

The results will give you something like this:

```bash
commaTest called

slice value:
foo=bar
baz
a=b
c

array value:
foo=bar,baz
a=b,c
```

## Creation

I created this demo as follows (my GOPATH is $HOME/go):

```bash
GITHUB_USER=brendonthiede
cd $GOPATH/src/github.com/${GITHUB_USER}
go get -u github.com/spf13/cobra/cobra
cobra init --pkg-name github.com/${GITHUB_USER}/cobrademo cobrademo
cd cobrademo
go mod init github.com/${GITHUB_USER}/cobrademo
go mod tidy
cobra add commaTest
```
