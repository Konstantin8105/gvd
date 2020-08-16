# gvd
Golang vim dictionary

```vim
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Go dictionary
"
" result of command : :set iskeyword?
" iskeyword=@,48-57,_,192-255
" add point in list
"
"	Each file should contain a list of words.  This can be one word per line, 
"	or several words per line, separated by non-keyword characters
"	(white space is preferred).  Maximum line length is 510 bytes.
"
:set iskeyword+=. " add point for dictionary work
:set dictionary+=$HOME/dotfiles/gonames.txt
```

Run `go list ... | less` and output:
```
archive/tar
archive/zip
bufio
bytes
cmd/addr2line
cmd/api
cmd/asm
cmd/asm/internal/arch
cmd/asm/internal/asm
cmd/asm/internal/flags
cmd/asm/internal/lex
cmd/buildid
cmd/cgo
cmd/compile
...
```

Run `go doc -short  math` and output:
```
const E = 2.71828182845904523536028747135266249775724709369995957496696763 ...
const MaxFloat32 = 3.40282346638528859811704183484516925440e+38 ...
const MaxInt8 = 1<<7 - 1 ...
func Abs(x float64) float64
func Acos(x float64) float64
func Acosh(x float64) float64
func Asin(x float64) float64
func Asinh(x float64) float64
func Atan(x float64) float64
func Atan2(y, x float64) float64
func Atanh(x float64) float64
func Cbrt(x float64) float64
...
```

Example of Vim dictionary:
```
strings.Index
strings.Replace
...
```

Get from Vim dictionary: "i_CTRL+X_CTRL+K"

https://stackoverflow.com/questions/28166249/how-to-list-installed-go-packages

```Go

import (
    "fmt"

    "golang.org/x/tools/go/packages"
)

func main() {
    pkgs, err := packages.Load(nil, "std")
    if err != nil {
        panic(err)
    }

    fmt.Println(pkgs)
    // Output: [archive/tar archive/zip bufio bytes compress/bzip2 ... ]
}
```

```Go
package main

import (
    "fmt"

    "golang.org/x/tools/go/packages"
)

var standardPackages = make(map[string]struct{})

func init() {
    pkgs, err := packages.Load(nil, "std")
    if err != nil {
        panic(err)
    }

    for _, p := range pkgs {
        standardPackages[p.PkgPath] = struct{}{}
    }
}

func isStandardPackage(pkg string) bool {
    _, ok := standardPackages[pkg]
    return ok
}

func main() {
    fmt.Println(isStandardPackage("fmt"))  // true
    fmt.Println(isStandardPackage("nope")) // false
}
```

```
cmd := exec.Command("go", "list", "std")
p, err := cmd.Output()
if err != nil {
    // handle error
}
stdPkgs = strings.Fields(string(p))
```
```
$ go list std

archive/tar
archive/zip
bufio
bytes
compress/bzip2
compress/flate
compress/gzip
compress/lzw
compress/zlib
container/heap
container/list
container/ring
```
