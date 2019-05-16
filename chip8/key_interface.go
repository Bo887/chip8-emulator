package chip8

import (
    "io"
)

var Keypad = [...]uint8 {
    '1', '2', '3', '4',
    'Q', 'W', 'E', 'R',
    'A', 'S', 'D', 'F',
    'Z', 'X', 'C', 'V',
}

func ReadInput(reader io.ByteReader) {
    value, err := reader.ReadByte()
    println(value)
    println(err)
}
