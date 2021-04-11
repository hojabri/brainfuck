package brainfuck

import (
	"bytes"
	"io"
	"testing"
)

var fakeScreenContent []byte

type FakeIoReader struct {
	Content bytes.Buffer
}

func NewFakeIoReader(content []byte) *FakeIoReader {
	var newBuffer bytes.Buffer
	newBuffer.Write(content)
	return &FakeIoReader{Content: newBuffer}
}
type FakeIoWriter struct {
	Content bytes.Buffer
}

func NewFakeIoWriter(content []byte) *FakeIoWriter {
	var newBuffer bytes.Buffer
	newBuffer.Read(content)
	return &FakeIoWriter{Content: newBuffer}
}



func (reader *FakeIoReader) Read(p []byte) (n int, err error) {
	return reader.Content.Read(p)
}
func (writer *FakeIoWriter) Write(p []byte) (n int, err error) {
	n, err = writer.Content.Write(p)
	fakeScreenContent = writer.Content.Bytes()
	return
}



func TestMachine_Execute(t *testing.T) {
	type fields struct {
		code               []*Instruction
		instructionPointer int
		memory             [30000]int
		dataPointer        int
		input              io.Reader
		output             io.Writer
		readBuf            []byte
	}
	
	tests := []struct {
		name   string
		fields fields
		want []byte
	}{
		{
			name:   "test1",
			fields: fields{
				code: []*Instruction{
					{
						Type:        Plus,
						Repeat:      3,
						IndexOfJump: 0,
					},
					{
						Type:        Right,
						Repeat:      2,
						IndexOfJump: 0,
					},
					{
						Type:        Plus,
						Repeat:      10,
						IndexOfJump: 0,
					},
					{
						Type:        Left,
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
						Type:        ReadChar,
						Repeat:      1,
						IndexOfJump: 0,
					},
					{
						Type:        PutChar,
						Repeat:      1,
						IndexOfJump: 0,
					},
				},
				instructionPointer: 0,
				memory:             [30000]int{},
				dataPointer:        0,
				input:              NewFakeIoReader([]byte("123")),
				output:             NewFakeIoWriter([]byte("")),
				readBuf:            make([]byte, 1),
			},
			want: []byte("12"),
		},
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machine := &Machine{
				codes:              tt.fields.code,
				InstructionPointer: tt.fields.instructionPointer,
				Memory:             tt.fields.memory,
				DataPointer:        tt.fields.dataPointer,
				input:              tt.fields.input,
				output:             tt.fields.output,
				readBuf:            tt.fields.readBuf,
			}
			
			machine.Execute()

			
			t.Logf("machine Memory: %v" , machine.Memory)
			t.Logf("machine readBuf: %v" , machine.readBuf )
			
			if string(fakeScreenContent) != string(tt.want) {
				t.Errorf("Machine Execution output %v , want: %v",string(fakeScreenContent) ,string(tt.want) )
			}
		})
	}
}
