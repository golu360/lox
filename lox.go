package main

import (
	"bufio"
	"log"
	"os"

	"github.com/golu360/lox/parser"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		log.Println("Usage : lox [script]")
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runShell()
	}
}

func runShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		report(0, scanner.Err().Error())
	}
}

func runFile(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		report(0, err.Error())
	}
	run(string(b))
}

func report(where int, message string) {
	log.Fatalf("[ERROR] : %s at %d", message, where)
	os.Exit(64)
}

func run(source string) {
	tokenizer := parser.NewTokenizer(source)
	log.Println(tokenizer.GetTokens())
}
