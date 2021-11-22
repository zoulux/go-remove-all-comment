package main

import (
	"bytes"
	"go/ast"
	"go/printer"
)

func printFile(file1 *ast.File) ([]byte, error) {
	printConfig := printer.Config{Mode: printer.RawFormat}

	var buf1 bytes.Buffer
	if err := printConfig.Fprint(&buf1, fset, file1); err != nil {
		return nil, err
	}
	src := buf1.Bytes()

	return src, nil
}
