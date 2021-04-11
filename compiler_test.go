package brainfuck

import (
	"reflect"
	"testing"
)

func TestCompiler_Compile(t *testing.T) {
	type fields struct {
		code         string
		codeLength   int
		position     int
		instructions []*Instruction
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Instruction
	}{
		{
			name: "test1",
			fields: fields{
				code:         "+++>---",
				codeLength:   len("+++>---"),
				position:     0,
				instructions: []*Instruction{},
			},
			want: []*Instruction{
				{
					Type:        Plus,
					Repeat:      3,
					IndexOfJump: 0,
				},
				{
					Type:        Right,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        Minus,
					Repeat:      3,
					IndexOfJump: 0,
				},
			},
		},
		{
			name: "test2",
			fields: fields{
				code:         "+++[>,.<-]",
				codeLength:   len("+++[>,.<-]"),
				position:     0,
				instructions: []*Instruction{},
			},
			want: []*Instruction{
				{
					Type:        Plus,
					Repeat:      3,
					IndexOfJump: 0,
				},
				{
					Type:        OpenLoop,
					Repeat:      1,
					IndexOfJump: 7,
				},
				{
					Type:        Right,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        ReadChar,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        PutChar,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        Left,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        Minus,
					Repeat:      1,
					IndexOfJump: 0,
				},
				{
					Type:        CloseLoop,
					Repeat:      1,
					IndexOfJump: 1,
				},
			},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compiler := &Compiler{
				code:         tt.fields.code,
				codeLength:   tt.fields.codeLength,
				position:     tt.fields.position,
				instructions: tt.fields.instructions,
			}
			if got := compiler.Compile(); !reflect.DeepEqual(&got, &tt.want) {
				t.Errorf("Compile() = %v, \nfunctionOutputWant %v", &got, &tt.want)
			}
		})
	}
}

func TestCompiler_appendToInstructions(t *testing.T) {
	type fields struct {
		code         string
		codeLength   int
		position     int
		instructions []*Instruction
	}
	type args struct {
		instructionType InstructionType
		repeat          int
		indexOfJump     int
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		functionOutputWant int
		instructionsWant   []*Instruction
	}{
		{
			name:   "test1",
			fields: fields{},
			args: args{
				instructionType: Plus,
				repeat:          3,
				indexOfJump:     0,
			},
			functionOutputWant: 0,
			instructionsWant: []*Instruction{
				{
					Type:        Plus,
					Repeat:      3,
					IndexOfJump: 0,
				},
			},
		},
		{
			name:   "test2",
			fields: fields{},
			args: args{
				instructionType: Right,
				repeat:          2,
				indexOfJump:     0,
			},
			functionOutputWant: 0,
			instructionsWant: []*Instruction{
				{
					Type:        Right,
					Repeat:      2,
					IndexOfJump: 0,
				},
			},
			// TODO: Add more test cases.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compiler := &Compiler{
				code:         tt.fields.code,
				codeLength:   tt.fields.codeLength,
				position:     tt.fields.position,
				instructions: tt.fields.instructions,
			}
			if got := compiler.appendToInstructions(tt.args.instructionType, tt.args.repeat, tt.args.indexOfJump); got != tt.functionOutputWant {
				t.Errorf("appendToInstructions() = %v, functionOutputWant %v", got, tt.functionOutputWant)
			}
			
			if !reflect.DeepEqual(compiler.instructions, tt.instructionsWant) {
				t.Errorf("instruction after running appendToInstructions() = %v, functionOutputWant %v", compiler.instructions, tt.functionOutputWant)
			}
			
		})
	}
}

func TestCompiler_groupRepeatedInstructions(t *testing.T) {
	type fields struct {
		code         string
		codeLength   int
		position     int
		instructions []*Instruction
	}
	type args struct {
		instructionType InstructionType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*Instruction
	}{
		{
			name: "test1",
			fields: fields{
				code:         "++++",
				codeLength:   4,
				position:     0,
				instructions: nil,
			},
			args: args{
				instructionType: Plus,
			},
			want: []*Instruction{
				{
					Type:        Plus,
					Repeat:      4,
					IndexOfJump: 0,
				},
			},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compiler := &Compiler{
				code:         tt.fields.code,
				codeLength:   tt.fields.codeLength,
				position:     tt.fields.position,
				instructions: tt.fields.instructions,
			}
			
			compiler.groupRepeatedInstructions(tt.args.instructionType)
			
			if !reflect.DeepEqual(compiler.instructions, tt.want) {
				t.Errorf("instruction after running groupRepeatedInstructions() = %v, want %v", compiler.instructions, tt.want)
			}
			
			
		})
	}
}

func TestNewCompiler(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want *Compiler
	}{
		{
			name: "test1",
			args: args{code: "+++>---"},
			want: &Compiler{
				code:       "+++>---",
				codeLength: len("+++>---"),
				position:   0,
				instructions: []*Instruction{},
			},
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCompiler(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCompiler() = %v, functionOutputWant %v", got, tt.want)
			}
		})
	}
}
