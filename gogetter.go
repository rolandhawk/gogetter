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
	customZeroValue := flag.String("custom", "", "Custom zero value map. Format: type1=zeroval1,type2=zeroval2")
	dry := flag.Bool("dry", false, "Dry run. If true, it will print the output to stdout.")
	flag.Parse()

	sourceFile := os.Getenv("GOFILE")
	packageName := os.Getenv("GOPACKAGE")
	structNames := flag.Args()
	outputFile := getFileName(sourceFile)

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

	customZeroValueMap := extractMap(*customZeroValue)
	for i, str := range val.Structs {
		for j, f := range str.Fields {
			if c, ok := customZeroValueMap[f.Type]; ok {
				val.Structs[i].Fields[j].ZeroValue = c
			}
		}
	}

	var buff bytes.Buffer
	t := template.Must(template.New("tmpl").Parse(tmpl))
	err = t.Execute(&buff, val)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Prettify output
	res, err := imports.Process(outputFile, buff.Bytes(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *dry {
		fmt.Println(string(res))
	} else {
		// Write into file
		err = ioutil.WriteFile(outputFile, res, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func getFileName(source string) string {
	filename := strings.Split(source, ".")
	filename = filename[:len(filename)-1] // remove .go suffix
	filename = append(filename, "getter", "go")

	return strings.Join(filename, ".")
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
