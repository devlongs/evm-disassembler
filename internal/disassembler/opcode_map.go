package disassembler

var OpcodeMap = map[byte]string{
	0x00: "STOP",
	0x01: "ADD",
	0x02: "MUL",
	0x03: "SUB",
	0x04: "DIV",
	0x05: "SDIV",
	// ...
	0x14: "EQ",
	0x15: "ISZERO",
	// ...
	0x50: "POP",
	0x51: "MLOAD",
	0x52: "MSTORE",
	0x53: "MSTORE8",
	// ...
	0x55: "SSTORE",
	// ...
	0x56: "JUMP",
	0x57: "JUMPI",
	0x58: "PC",
	0x59: "MSIZE",
	0x5A: "GAS",
	// ...
	// PUSH range
	// 0x60 => PUSH1, 0x61 => PUSH2, ... up to 0x7F => PUSH32
	0x60: "PUSH1",
	0x61: "PUSH2",
	// ...
	0x7F: "PUSH32",
	// ...
	0x80: "DUP1",
	0x81: "DUP2",
	// ...
	0x8F: "DUP16",
	0x90: "SWAP1",
	0x91: "SWAP2",
	// ...
	0x9F: "SWAP16",
	// ...
}
