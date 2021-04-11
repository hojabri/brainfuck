package brainfuck

import (
	"io"
)

type customCommandExecutor func(commandName string, commandArg int, machine *Machine)

type Machine struct {
	codes              []*Instruction
	InstructionPointer int
	
	Memory      [30000]int
	DataPointer int
	
	input  io.Reader
	output io.Writer
	
	readBuf []byte
	
	customCommandExecutor
}

func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer, customFunc customCommandExecutor) *Machine {
	return &Machine{
		codes:                 instructions,
		input:                 in,
		output:                out,
		readBuf:               make([]byte, 1),
		customCommandExecutor: customFunc,
	}
}

func (machine *Machine) Execute() {
	for machine.InstructionPointer < len(machine.codes) {
		instruction := machine.codes[machine.InstructionPointer]
		
		switch instruction.Type {
		case Plus:
			machine.Memory[machine.DataPointer] += instruction.Repeat
		case Minus:
			machine.Memory[machine.DataPointer] -= instruction.Repeat
		case Right:
			machine.DataPointer += instruction.Repeat
		case Left:
			machine.DataPointer -= instruction.Repeat
		case PutChar:
			for i := 0; i < instruction.Repeat; i++ {
				machine.printChar()
			}
		case ReadChar:
			for i := 0; i < instruction.Repeat; i++ {
				machine.readChar()
			}
		case OpenLoop:
			if machine.Memory[machine.DataPointer] == 0 {
				machine.InstructionPointer = instruction.IndexOfJump
				continue
			}
		case CloseLoop:
			if machine.Memory[machine.DataPointer] != 0 {
				machine.InstructionPointer = instruction.IndexOfJump
				continue
			}
		case Custom:
			machine.customCommandExecutor(instruction.CustomCommandName,instruction.CustomCommandArg, machine)
		}
		
		machine.InstructionPointer++
		if machine.DataPointer >= len(machine.Memory) || machine.DataPointer <0 {
			break
		}
	}

}

func (machine *Machine) readChar() {
	_, err := machine.input.Read(machine.readBuf)
	if err != nil {
		panic(err)
	}
	machine.Memory[machine.DataPointer] = int(machine.readBuf[0])
}

func (machine *Machine) printChar() {
	machine.readBuf[0] = byte(machine.Memory[machine.DataPointer])
	
	_, err := machine.output.Write(machine.readBuf)
	if err != nil {
		panic(err)
	}
	//fmt.Print(string(byte(machine.Memory[machine.DataPointer])))
}
