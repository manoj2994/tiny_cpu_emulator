# CPU Emulator in Go

This is a simple CPU emulator written in Go.  
It supports basic operations like `LOAD`, `ADD`, `SUB`, `STORE`, `JUMP`, and `HALT`.


## 🧠 How a real CPU works:
At a high level, a CPU performs operations through the fetch-decode-execute cycle. This is the core mechanism behind most CPUs:
- Fetch: The CPU fetches an instruction from memory.
- Decode: It decodes the instruction to understand what operation it needs to perform.
- Execute: It performs the operation, such as an arithmetic operation, moving data, or jumping to another part of the program.

Registers and memory play a key role in this process:
  - Registers: Small, fast storage locations inside the CPU that hold data and addresses.
  - Memory: The larger, slower storage for the program’s data (RAM).

In a real CPU, operations like `ADD`, `SUB`, `LOAD`, `JUMP`, and `HALT` are often controlled by the Control Unit (CU) and Arithmetic Logic Unit (ALU).



## 📂 Project Structure

- `main.go` — Entry point of the application
- `cpu_emulator/` — Contains CPU logic (structs, instructions, memory management)

## 🛠 Features

- Simple CPU with 4 registers and 256 bytes of memory
- Basic instruction set:
  - `LOAD`
  - `ADD`
  - `SUB`
  - `STORE`
  - `JUMP`
  - `HALT`
- Program Counter (`pc`) and Running State (`running`) management
- Easy to extend with more instructions!

## 🚀 How to Run

First, clone the repo:

```bash
git clone https://github.com/manoj2994/tiny_cpu_emulator
cd <repo>
go run main.go
```


## LICENSE
This project is licensed under the MIT License — see the LICENSE file for details.
