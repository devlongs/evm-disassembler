package disassembler

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// EVMOp defines a single EVM opcode and optional data (for PUSH instructions).
type EVMOp struct {
	Opcode string `json:"opcode"`
	Data   string `json:"data,omitempty"`
}

// DisassembleEVMBytecode disassembles a hex string of EVM bytecode into EVMOps.
func DisassembleEVMBytecode(hexStr string) ([]EVMOp, error) {
	// Clean input
	hexStr = strings.TrimPrefix(hexStr, "0x")
	hexStr = strings.ReplaceAll(hexStr, " ", "")
	hexStr = strings.ToLower(hexStr)

	// Decode hex
	code, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %v", err)
	}

	var ops []EVMOp
	i := 0
	for i < len(code) {
		opcode := code[i]
		i++

		// PUSH instructions range from 0x60 to 0x7F
		if opcode >= 0x60 && opcode <= 0x7F {
			pushSize := int(opcode - 0x5F)
			// Safely handle case where code might end prematurely
			if i+pushSize > len(code) {
				pushSize = len(code) - i
			}

			data := code[i : i+pushSize]
			i += pushSize

			ops = append(ops, EVMOp{
				Opcode: fmt.Sprintf("PUSH%d", pushSize),
				Data:   fmt.Sprintf("0x%s", hex.EncodeToString(data)),
			})
		} else {
			// Lookup opcode
			name, ok := OpcodeMap[opcode]
			if !ok {
				name = fmt.Sprintf("UNKNOWN (0x%x)", opcode)
			}
			ops = append(ops, EVMOp{
				Opcode: name,
			})
		}
	}

	return ops, nil
}
