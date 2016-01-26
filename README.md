[![Build Status](https://travis-ci.org/brownjohnf/slit.svg?branch=master)](https://travis-ci.org/brownjohnf/slit)

# slit

Like cut + bits of awk, but easier to use 90% of the time.

## Install

```
go install github.com/brownjohnf/goutils/slit
```

## Usage

```
Usage of slit:
  -d string
        delimiter on which to split columns (defaults to whitespace)
  -f string
        field(s) to select (default "1")
  -s int
        lines to skip (headers, etc)
  -v    verbose errors
```

From stdout:

```
df | slit -f 2
```

Local file:

```
slit -f 2 < some_file.txt
```

## Building

```
$ go get github.com/brownjohnf/goutils/slit
$ cd $GOPATH
$ go build # or go install
```

