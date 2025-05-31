package main

import (
	"fmt"

	"github.com/manoj2994/tiny_cpu_emulator/cpu_emulator"
)

func main() {

	fmt.Println("Sai Ram")
	program := &cpu_emulator.CPU{}
	program.Memory = [256]int{
		cpu_emulator.LOAD, 0, 10,
		cpu_emulator.LOAD, 1, 20,
		cpu_emulator.JUMP, 5,
		cpu_emulator.ADD, 2, 0, 1, // ADD R0 + R1 in R2
		cpu_emulator.STORE, 1, 100, // STORE R1 in location 100
		//cpu_emulator.JUMP, 5, // JUMP PC to 1
		cpu_emulator.HALT,
	}
	program.Running = true
	program.Run()
	fmt.Println(program.Memory)
}
