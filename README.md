# EVM Disassembler

A robust REST API built in Go that disassembles Ethereum Virtual Machine (EVM) bytecode into human-readable opcodes. Designed for developers, researchers, and enthusiasts who want to easily decode and explore smart contract bytecode.

## Overview
EVM Disassembler is a REST API backend providing a single POST /api/disassemble endpoint. Given a string of hex-encoded EVM bytecode, it will decode the bytecode and convert it to a sequence of EVM instructions (OP CODES). This project is perfect for anyone building blockchain explorers, automated analysis tools, or educational resources around Ethereum’s EVM.

## Features
- Disassembly Made Easy – Converts tricky hex bytecode to a readable list of opcodes.
- Configurable – Leverages Viper for flexible YAML, JSON, or environment-based configuration.
- Structured Logging – Employs Zap for performance-friendly, JSON-based logs.
- RESTful – Uses Gin to expose a lightweight yet powerful HTTP endpoint.
- Dockerized – Includes a multi-stage Dockerfile for convenient container-based deployment.
- Modular Code Structure – Organized under cmd/, internal/, and configs/ for easy maintenance and future growth.
- Extensible – Updating the opcode map allows support for custom or newly introduced opcodes.

## Architecture
```python
evm-disassembler/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── configs/
│   └── config.yaml           # Default YAML configuration
├── internal/
│   ├── api/
│   │   ├── handlers.go       # HTTP handlers
│   │   ├── router.go         # Router and middleware setup
│   │   └── middlewares.go    # Custom middlewares 
│   ├── config/
│   │   └── config.go         # Loads configuration via Viper
│   ├── disassembler/
│   │   ├── disassembler.go   # Disassembler core logic
│   │   └── opcode_map.go     # Maps opcodes to human-readable names
│   ├── logger/
│   │   └── logger.go         # Initializes the Zap logger
├── pkg/                      # Optional: shared libraries or packages
├── Dockerfile                # Multi-stage Docker build
├── go.mod                 
├── go.sum                   
└── README.md                 # Project documentation
```