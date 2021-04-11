package main

import (
	"bufio"
	"fmt"
	"github.com/hojabri/brainfuck"
	"log"
	"os"
	"strings"
)

func main() {
	var code string
	
	fmt.Print("Please enter the code here (hit Enter key at the end):")
	reader := bufio.NewReader(os.Stdin)
	
	// Read code string from stdin (for example: ++++>>--,.)
	code, _ = reader.ReadString('\n')
	
	// Remove new-line character
	code = strings.TrimSuffix(code, "\n")
	
	// Show input prompt, if the code needs an input from the user
	if strings.Contains(code,",") {
		fmt.Print("Please enter input string to be used in code (hit Enter key at the end):")
	}
	
	// Initializing brainfuck compiler for compiling the raw code
	compiler,err := brainfuck.NewCompiler(code)
	if err != nil {
		log.Fatalf("error when compiling code: %s", err.Error())
		return
	}
	instructions := compiler.Compile()
	
	// Initializing brainfuck machine
	m := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout, nil)
	
	// Execute brainfuck instructions
	m.Execute()
	
	fmt.Println("\nExecution finished")
}
