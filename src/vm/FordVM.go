package main

import "fmt"

const (
	STACK_LIMIT uint16 = 1024
)

var (
	code      []byte
	ip        byte
	sp        uint16
	stack     []FordValue
	constants []FordValue
)

func READ_BYTE() byte {
	opCode := code[ip]
	ip++
	return opCode
}

func GET_CONST() *FordValue {
	constIndex := READ_BYTE()
	constant := constants[constIndex]
	return &constant
}

func binaryOp(operator byte) {
	op2 := pop().AsNumber()
	op1 := pop().AsNumber()
	var result FordValue
	switch operator {
	case OP_ADD:
		result = createFordValue("number", op1+op2)
		break
	case OP_SUB:
		result = createFordValue("number", op1-op2)
		break
	case OP_MUL:
		result = createFordValue("number", op1*op2)
		break
	case OP_DIV:
		result = createFordValue("number", op1/op2)
		break
	}
	push(&result)
}

func push(value *FordValue) {
	if sp == STACK_LIMIT {
		panic("push: stack overflow.\n")
	}
	stack[sp] = *value
	sp++
}

func pop() FordValue {
	if sp == 0 {
		panic("pop: empty stack.\n")
	}

	sp--
	return stack[sp]
}

func exec(program string) FordValue {

	// 1. parse the program into an AST
	// ast = parser.parse(program)

	// 2. compile ast to bytecode
	// code = compiler.compile(ast)

	// constant pool
	// constants = append(constants, createFordValue("string", "Hello, "))
	// constants = append(constants, createFordValue("string", "world!"))
	constants = append(constants, createFordValue("number", 10))
	constants = append(constants, createFordValue("number", 3))

	// instruction pointer
	code = []byte{OP_CONST, 0, OP_CONST, 1, OP_ADD, OP_HALT}
	ip = 0

	// stack pointer
	stack = make([]FordValue, STACK_LIMIT)
	sp = 0

	return eval()
}

func eval() FordValue {

	for {
		opCode := READ_BYTE()
		switch opCode {
		case OP_HALT:
			return pop()
		case OP_CONST:
			push(GET_CONST())
			break
		case OP_ADD:

			op2 := pop()
			op1 := pop()

			if op1.IsNumber() && op2.IsNumber() {
				v1 := op1.AsNumber()
				v2 := op2.AsNumber()
				result := createFordValue("number", v1+v2)
				push(&result)
			} else if op1.IsString() && op2.IsString() {
				v1 := op1.AsString()
				v2 := op2.AsString()
				result := createFordValue("string", v1+v2)
				push(&result)
			}

			break
		case OP_SUB:
			binaryOp(OP_SUB)
			break
		case OP_MUL:
			binaryOp(OP_MUL)
			break
		case OP_DIV:
			binaryOp(OP_DIV)
			break
		default:
			panic("unknown opcode.")
		}

	}
}

func main() {
	result := exec("")

	fmt.Println("result", result, result.AsNumber())
}
