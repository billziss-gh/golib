#!/bin/bash

# requires https://github.com/davecheney/godoc2md

golib="github.com/billziss-gh/golib"
progdir=$(dirname "$0")

(
    sed -n '1,/(GODOC)/p' README.md
    go list -f "{{.ImportPath}} {{.Doc}}" "$golib/..." | sed 's@.*/vendor/@@' |
    while read p d; do
        echo "- [$(basename $p)](#$p) - $d [:book:](https://godoc.org/$p)"
    done
    go list "$golib/..." | sed 's@.*/vendor/@@' |
    while read p; do
        $GOPATH/bin/godoc2md -template "$progdir/godoc2md.templ" $p |
        sed \
            -e "s@/src/$golib/@@g"\
            -e "s@/src/target@$(basename $p)@g"\
            -e "s@?s=[0-9][0-9]*:[0-9][0-9]*#@#@g"
    done
) > $progdir/../README.md.new
mv $progdir/../README.md.new $progdir/../README.md
