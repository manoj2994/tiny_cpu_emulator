# CPU Emulator in Go

This is a simple CPU emulator written in Go.  
It supports basic operations like `LOAD`, `ADD`, `SUB`, `STORE`, `JUMP`, and `HALT`.

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
