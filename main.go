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
		cpu_emulator.ADD, 2, 0, 1,
		cpu_emulator.HALT,
	}
	program.Running = true
	program.Run()
}
