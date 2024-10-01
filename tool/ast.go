package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/golu360/lox/utils"
)

func main() {
	cliFlags := utils.GetCLIFlags()
	cliFlags.Parse()

	log.Printf("Output Dir : %s \n", cliFlags.Output)
	fields := []string{"Binary   : Left Expr[Binary], Operator parser.Token, Right Expr[Binary]",
		"Grouping : Expression Expr[Grouping]",
		"Literal  : Value interface{}",
		"Unary    : Operator parser.Token, Right Expr[Unary]"}
	defineAst(cliFlags.Output, "Expr", fields)
}

func defineAst(outPutDir string, baseName string, types []string) {
	outputPath := path.Join(outPutDir, baseName+".go")
	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		log.Println("AST File could not be generated")
		os.Exit(65)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "package %s\n\n", strings.ToLower(baseName))
	fmt.Fprintln(writer, "import (")
	fmt.Fprintln(writer, `"github.com/golu360/lox/parser"`)
	fmt.Fprintln(writer, ")")
	fmt.Fprintln(writer, "")
	fmt.Fprintf(writer, "type %s [R any] interface {\n", baseName)
	fmt.Fprintln(writer, "Accept(visitor Visitor[R]) R")
	fmt.Fprintln(writer, "}")
	defineVisitor(writer, baseName, types)

	for _, fType := range types {
		structName := strings.TrimSpace(strings.Split(fType, ":")[0])
		fieldSet := strings.TrimSpace(strings.Split(fType, ":")[1])
		defineType(writer, structName, baseName, fieldSet)
	}
	writer.Flush()

}

func defineType(writer *bufio.Writer, structName string, baseName, fieldList string) {
	fmt.Fprintf(writer, "type  %s struct {", structName)
	fmt.Fprintln(writer, "")
	fields := strings.Split(fieldList, ",")
	for _, field := range fields {
		fmt.Fprintf(writer, "\t%s\n", field)
	}

	fmt.Fprintf(writer, "\n}\n")
	fmt.Fprintf(writer, "func (t *%s) Accept(v Visitor[%s]){\n", structName, structName)
	fmt.Fprintf(writer, "\tv.Visit%s%s(t) \n", structName, baseName)
	fmt.Fprintln(writer, "}")

	fmt.Fprintln(writer, "")
}

func defineVisitor(writer *bufio.Writer, structName string, types []string) {
	fmt.Fprintf(writer, "type Visitor[%s any] interface{\n", structName)
	for _, fType := range types {
		typeName := strings.TrimSpace(strings.Split(fType, ":")[0])
		fmt.Fprintf(writer, "Visit%s%s(*%s)\n", typeName, structName, typeName)
	}
	fmt.Fprintln(writer, "}")

}
