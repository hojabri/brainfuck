package main

import (
	"flag"
	"fmt"
	"github.com/hojabri/brainfuck"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	

	flag.Parse()
	args := flag.Args()
	
	if len(args) > 0 {
		
		filename := args[0]
		program, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("File %s could not be used\n", filename)
			return
		}
		compiler,err := brainfuck.NewCompiler(string(program))
		if err != nil {
			log.Fatalf("error when compiling code: %s", err.Error())
			return
		}
		
		instructions := compiler.Compile()
		
		m := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout, nil)
		m.Execute()
		
		fmt.Println("\nExecution finished")
	} else {
		fmt.Println("no input file specified")
		fmt.Println("Usage: ./main {inputfile.bf}")
	}
	
	
}
