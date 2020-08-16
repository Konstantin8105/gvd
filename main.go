package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	var (
		// packages names pattern
		pkgs = flag.String("pkgs", "std", "packages names pattern")

		// help
		help = flag.Bool("h", false, "print help information")
	)
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}

	words, err := Get(*pkgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	sort.Strings(words)
	for _, word := range words {
		fmt.Fprintf(os.Stdout, "%s\n", word)
	}
}

func Get(pkg string) (words []string, err error) {
	data, err := exec.Command("go", "list", pkg).Output()
	if err != nil {
		return
	}
	pkgs := strings.Fields(string(data))

	for _, pkg := range pkgs {
		if pkg == "" {
			continue
		}
		var ws []string
		ws, err = Words(pkg)
		if err != nil {
			return
		}
		words = append(words, ws...)
	}
	return
}

func Words(pkg string) (words []string, err error) {
	data, err := exec.Command("go", "doc", "-short", pkg).Output()
	if err != nil {
		return
	}
	pkg = CleanPkg(pkg)
	wordsIn := strings.Split(string(data), "\n")
	for i := range wordsIn {
		if wordsIn[i] == "" {
			continue
		}
		cw := CleanWord(wordsIn[i])
		if cw == "" {
			continue
		}
		words = append(words, pkg+"."+cw)
	}
	return
}

func CleanPkg(pkg string) string {
	// change from :
	// unicode/utf16
	// to :
	// utf16
	index := strings.LastIndex(pkg, "/")
	if index >= 0 {
		pkg = pkg[index+1:]
	}
	return pkg
}

func CleanWord(word string) string {
	// change from :
	// func MustHaveGoRun(t testing.TB)
	// to
	// MustHaveGoRun(

	if word == "" {
		return ""
	}
	if word[0] == ' ' {
		return ""
	}

	prefixes := []string{"func ", "type ", "var ", "const "}
	found := false
	for _, pre := range prefixes {
		found = found || strings.HasPrefix(word, pre)
	}
	if !found {
		return ""
	}
	for _, pre := range prefixes {
		word = strings.TrimPrefix(word, pre)
	}

	if word[0] == '(' {
		return ""
	}

	if index := strings.Index(word, " "); index >= 0 {
		s := word[:index]
		if strings.Contains(word, "struct{") {
			s += "{"
		}
		word = s 
	}
	if index := strings.Index(word, "("); index >= 0 {
		word = word[:index+1]
	}

	word = strings.TrimSpace(word)
	return word
}
