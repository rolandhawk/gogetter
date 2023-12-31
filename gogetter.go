package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/rolandhawk/gogetter/ast"
)

type values struct {
	Package string
	Source  string
	Imports []ast.Import
	Structs []ast.Struct

	Map map[string]int64
}

// Expected usage in gogenerate:
//   go:generate gogetter StructName1 StructName2
func main() {
	sourceFile := os.Getenv("GOFILE")
	packageName := os.Getenv("GOPACKAGE")
	outputFile := flag.String("out", getDefaultOutputFileName(sourceFile), "Custom output file name. Default: [source]_getter.go. Format: *.go")
	dry := flag.Bool("dry", false, "Dry run. If true, it will print the output to stdout instead of write to file.")
	flag.Parse()

	structNames := flag.Args()
	fmt.Println("generate getter for struct", structNames)

	file, err := ast.ParseFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	structs, err := file.GetStructs(structNames...)
	if err != nil {
		fmt.Println(err)
		return
	}

	val := &values{
		Package: packageName,
		Source:  sourceFile,
		Imports: file.GetImports(),
		Structs: structs,
	}

	var buff bytes.Buffer
	t := template.Must(template.New("tmpl").Parse(tmpl))
	err = t.Execute(&buff, val)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Prettify output
	res, err := imports.Process(*outputFile, buff.Bytes(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *dry {
		fmt.Println(string(res))
	} else {
		// Write into file
		err = ioutil.WriteFile(*outputFile, res, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func getDefaultOutputFileName(source string) string {
	filename := strings.TrimSuffix(source, ".go")
	return fmt.Sprintf("%s_getter.go", filename)
}

func extractMap(str string) map[string]string {
	res := make(map[string]string)
	maps := strings.Split(str, ",")
	for _, m := range maps {
		kv := strings.Split(m, "=")
		if len(kv) == 2 {
			res[kv[0]] = kv[1]
		}
	}
	return res
}
