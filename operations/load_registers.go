package operations

import (
	"fmt"
	"chip8/system"
)

type LoadRegistersParser struct {}
func(p LoadRegistersParser) Matches(opcode OpCode) bool {
	return opcode.String()[0] == 'f' && byte(opcode) == 0x65
}

func(p LoadRegistersParser) CreateOp(opcode OpCode) Operation {
	return LoadRegistersOp{
		topRegister: byte(opcode & 0x0F00 >> 8),
	}
}

type LoadRegistersOp struct {
	topRegister byte
}
func(o LoadRegistersOp) String() string {
	return fmt.Sprintf("load_registers(V%X, &I)", o.topRegister)
}

func(o LoadRegistersOp) Execute(machine *system.VirtualMachine) {
	for i := byte(0); i <= o.topRegister; i++ {
		machine.Registers[i] = machine.Memory[machine.IndexRegister + uint16(i)]
	}
}