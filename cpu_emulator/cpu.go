package cpu_emulator

import (
	"fmt"
)

type CPU struct {
	register [4]int
	Memory   [256]int
	pc       int
	Running  bool
}

const (
	LOAD = iota
	ADD
	SUB
	STORE
	JUMP
	HALT
)

func (cpu *CPU) Run() {

	for cpu.Running {
		fmt.Println("Current Program counter - ", cpu.pc)
		opcode := cpu.Memory[cpu.pc]
		cpu.execute(opcode)
	}
}

func (cpu *CPU) execute(opcode int) {

	switch opcode {
	case LOAD:
		reg := cpu.Memory[cpu.pc+1]
		val := cpu.Memory[cpu.pc+2]
		cpu.register[reg] = val
		cpu.pc += 3
		fmt.Println("LOAD Register ", reg, " with ", cpu.register[reg])
	case ADD:
		dest := cpu.Memory[cpu.pc+1]
		src1 := cpu.Memory[cpu.pc+2]
		src2 := cpu.Memory[cpu.pc+3]
		cpu.register[dest] = src1 + src2
		fmt.Println("ADD Registers", src1, " & ", src2, " in ", cpu.register[dest])
		cpu.pc += 4
	case HALT:
		fmt.Println("System Received - HALT")
		cpu.Running = false

	}
}
