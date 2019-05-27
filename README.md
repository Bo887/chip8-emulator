# CHIP-8 Emulator
[![Build Status](https://travis-ci.com/Bo887/chip8-emulator.svg)](https://travis-ci.com/Bo887/chip8-emulator)

A CHIP-8 emulator written in golang. See this [page](https://en.wikipedia.org/wiki/CHIP-8) for information about the CHIP-8 architecture and opcodes.

## Installation
Make sure golang is [installed](https://golang.org/doc/install).

To get the package, run:
```
go get github.com/Bo887/chip8-emulator
```

The package should now be located in:
```
$GOPATH/src/github.com/Bo887
```

To install the package, run:
```
make install
```

The binary should be locaed in:
```
$GOPATH/bin
```

To remove the package, just simply delete the directory and the binary.

## Usage
The binary name is `chip8-emulator`. It takes a single argument of the path to the ROM. 

For more information, run:
```
chip8-emulator --help
```

A few ROMS are located in the `roms/` directory from the project root. A lot more are located [here](https://github.com/dmatlack/chip8/tree/master/roms).

For example, running:
```
chip8-emulator roms/chip8-picture
```
from the project root directory will display the chip8 logo.
