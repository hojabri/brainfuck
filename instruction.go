package brainfuck

type InstructionType byte

const (
	Plus      InstructionType = '+'
	Minus     InstructionType = '-'
	Right     InstructionType = '>'
	Left      InstructionType = '<'
	PutChar   InstructionType = '.'
	ReadChar  InstructionType = ','
	OpenLoop  InstructionType = '['
	CloseLoop InstructionType = ']'
	Custom    InstructionType = '{'
)

type Instruction struct {
	Type   InstructionType
	Repeat int
	IndexOfJump int
	CustomCommandName string
	CustomCommandArg int
}
