package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	fset = token.NewFileSet()
)

var (
	filename = flag.String("file-name", "", "support one file")
	filedir  = flag.String("file-dir", "", "support dir, each file")
	outdir   = flag.String("out-dir", "", "export dir,default empty is origin file")
)

func init() {
	flag.Parse()
}

func main() {
	log.SetPrefix("[go-remove-all-comment]")

	if len(*filename) > 0 {
		if err := parseFile(getPath(*filename)); err != nil {
			log.Panic(err)
		}
	}

	if len(*filedir) > 0 {
		_filedir := getPath(*filedir)
		filepath.Walk(_filedir, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				if err := parseFile(path); err != nil {
					log.Panic(err)
				}
			}
			return err
		})
	}
}

func parseFile(filename string) error {
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil
	}
	f.Comments = make([]*ast.CommentGroup, 0)

	transformContent, err := printFile(f)

	if err != nil {
		panic(err)
	}

	if len(*outdir) > 0 {
		_outdir := getPath(*outdir)

		filename = path.Join(_outdir, filename)
		_filedir := path.Dir(filename)
		if _, err := os.Stat(_filedir); err != nil {
			fmt.Println(os.MkdirAll(_filedir, fs.ModePerm))
		}
	}

	return ioutil.WriteFile(filename, transformContent, fs.ModePerm)
}

func getPath(cpath string) (filepath string) {
	if path.IsAbs(cpath) {
		filepath = cpath
	} else {
		pwd, _ := os.Getwd()
		filepath = path.Join(pwd, cpath)
	}
	return
}
