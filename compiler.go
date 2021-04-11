package brainfuck

import (
	"VaultningMiddleware/pkg/log"
	"regexp"
	"strconv"
)

// SpecialInstructionRegex is to enable adding new custom commands
// They should be in the format of {commandName}argument
// For example: {multiply}3 for multiply current cell by 3
// or {square}2 , ...
var SpecialInstructionRegex = regexp.MustCompile("^({[a-z]+})([0-9]+)?")
var err error

type Compiler struct {
	code       string
	codeLength int
	position   int
	
	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

func (compiler *Compiler) Compile() []*Instruction {
	var loopStack []int
	for compiler.position < compiler.codeLength {
		current := compiler.code[compiler.position]
		
		switch current {
		case '+':
			compiler.groupRepeatedInstructions(Plus)
		case '-':
			compiler.groupRepeatedInstructions(Minus)
		case '<':
			compiler.groupRepeatedInstructions(Left)
		case '>':
			compiler.groupRepeatedInstructions(Right)
		case '.':
			compiler.groupRepeatedInstructions(PutChar)
		case ',':
			compiler.groupRepeatedInstructions(ReadChar)
		case '[':
			insPos := compiler.appendToInstructions(OpenLoop, 1, 0,"",0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Pop position of last OpenLoop ("[") instruction off stack
			indexOfOpenLoop := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			
			// Emit the new CloseLoop ("]") instruction,
			// with correct position as argument
			indexOfCloseLoop := compiler.appendToInstructions(CloseLoop, 1, indexOfOpenLoop,"",0)
			
			// Patch the old OpenLoop ("[") instruction with new position
			compiler.instructions[indexOfOpenLoop].IndexOfJump = indexOfCloseLoop
		case '{':
			match := SpecialInstructionRegex.FindStringSubmatch(compiler.code[compiler.position:])
			if len(match) > 0 {
		
				var customCommandName string
				var customCommandArg int
				customCommandName = match[1]
				if match[2] !="" {
					customCommandArg,err = strconv.Atoi(match[2])
					if err!=nil {
						log.Error("custom command argument should be an integer number")
					}
				}
				compiler.appendToInstructions(Custom, 1, 0,customCommandName,customCommandArg)
				compiler.position += len(match[0])-1
			}
			
		}
		
		compiler.position++
	}
	
	return compiler.instructions
}

func (compiler *Compiler) groupRepeatedInstructions(instructionType InstructionType) {
	repeat := 1
	
	for compiler.position < compiler.codeLength-1 && compiler.code[compiler.position+1] == byte(instructionType) {
		repeat++
		compiler.position++
	}
	
	compiler.appendToInstructions(instructionType, repeat, 0,"",0)
}

func (compiler *Compiler) appendToInstructions(instructionType InstructionType, repeat int, indexOfJump int, customCommandName string, customCommandArg int) int {
	instruction := &Instruction{
		Type:              instructionType,
		Repeat:            repeat,
		IndexOfJump:       indexOfJump,
		CustomCommandName: customCommandName,
		CustomCommandArg:  customCommandArg,
	}
	compiler.instructions = append(compiler.instructions, instruction)
	return len(compiler.instructions) - 1
}
