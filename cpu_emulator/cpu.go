package cpu_emulator

import (
	"fmt"
	"strconv"
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

func (cpu *CPU) find_instruction_place(pc int) map[int]int {
	check := true
	instructions_addr := make(map[int]int)
	pos := 0
	for check {
		instructions_addr[pos] = pc
		if cpu.Memory[pc] == 5 {
			break
		}
		opcode := cpu.Memory[pc]
		switch opcode {
		case LOAD:
			pc += 3
		case ADD:
			pc += 4
		case STORE:
			pc += 3
		case JUMP:
			pc += 2
		case HALT:
			pc += 1

		}
		pos += 1
	}
	//fmt.Println(instructions_addr)
	//fmt.Println("check the code for halt", cpu.Memory[15])
	return instructions_addr
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
		cpu.register[dest] = cpu.register[src1] + cpu.register[src2]
		fmt.Println("ADD Registers", cpu.register[src1], " & ", cpu.register[src2], " in ", dest)
		cpu.pc += 4
	case STORE:
		dest := cpu.Memory[cpu.pc+2]
		val := cpu.register[cpu.Memory[cpu.pc+1]]
		cpu.Memory[dest] = val
		fmt.Println("STORE Register- R"+strconv.Itoa(cpu.Memory[cpu.pc+1]), "in", dest, " Memory ")
		cpu.pc += 3
	case JUMP:
		addr := cpu.find_instruction_place(0)
		//fmt.Println("Jump pc is ", cpu.pc)
		fmt.Println("Jump to the instruction ", cpu.Memory[cpu.pc+1])
		cpu.pc = addr[cpu.Memory[cpu.pc+1]]
	case HALT:
		fmt.Println("System Received - HALT")
		cpu.Running = false

	}
}
