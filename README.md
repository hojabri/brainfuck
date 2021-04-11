# brainfuck
A powerful, optimized and extendible Brainfuck library in Golang

## ⚡️ Quickstart

```go
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
```

## ⚙️ Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). 
Install brainfuck library with the [`go get`](https://github.com/hojabri/brainfuck) command:

```bash
go get -u github.com/hojabri/brainfuck
```

## 🎯 Features

-   Supports all basic brainfuck commands
-   Supports custom commands (runtime) and their argument
-   Input from external brainfuck code files, and also user input in stdin
-   Fast and optimized. All repeated command will be grouped by brainfuck compiler before run; all loop opening and closing positions will be stored, so no time will be taken each time to re-calculate their positions.
-   Validation before compiling the code. It checks for proper closing/opening of loop braces. 

## 📖 Examples

####  Custom Commands 

You can include your own custom commands inside brainfuck code
 For example:
 ```
 {memory}10++++{memory}10{increment}150{memory}10
 ```
 Output:
```		Memory: [0 0 0 0 0 0 0 0 0 0]

		Memory: [4 0 0 0 0 0 0 0 0 0]

		Memory: [154 0 0 0 0 0 0 0 0 0]
```
 Example 2:
 ```
 {increment}72.>{increment}101.>{increment}108.>{increment}108.>{increment}111.>{increment}32.>{increment}119.>{increment}111.>{increment}114.>{increment}108.>{increment}100.>{increment}33.{memory}12
 ```
 Output:
```
		Hello world!
		Memory: [72 101 108 108 111 32 119 111 114 108 100 33]
```
Sample code:
```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hojabri/brainfuck"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	
	flag.Parse()
	args := flag.Args()
	
	if len(args) > 0 {
		
		filename := args[0]
		program, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("File %s could not be used\n", filename)
			return
		}
		compiler,err := brainfuck.NewCompiler(string(program))
		if err != nil {
			log.Fatalf("error when compiling code: %s", err.Error())
			return
		}
		instructions := compiler.Compile()
		
		m := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout, customFunction)
		m.Execute()
	
	
	} else {
		var code string
		fmt.Print("Please enter the code here (hit Enter key at the end):")
		reader := bufio.NewReader(os.Stdin)
		
		code, _ = reader.ReadString('\n')
		code = strings.TrimSuffix(code, "\n")
		if strings.Contains(code,",") {
			fmt.Print("Please enter input string to be used in code (hit Enter key at the end):")
		}
		compiler,err := brainfuck.NewCompiler(code)
		if err != nil {
			log.Fatalf("error when compiling code: %s", err.Error())
			return
		}
		instructions := compiler.Compile()
		
		
		
		m := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout, customFunction)
		
		m.Execute()
	
	}
	fmt.Println("\nExecution finished")

}



func customFunction(commandName string, commandArg int, machine *brainfuck.Machine) {
	switch commandName {
	case "{power}":
		machine.Memory[machine.DataPointer] =int(math.Pow(float64(machine.Memory[machine.DataPointer]),float64(commandArg)))
	case "{memory}":
		if commandArg >0 {
			fmt.Printf("\nMemory: %v\n" , machine.Memory[:commandArg])
		} else {
			fmt.Printf("\nMemory: %v\n" , machine.Memory)
		}
	case "{increment}":
		machine.Memory[machine.DataPointer] += commandArg
	}
}
```