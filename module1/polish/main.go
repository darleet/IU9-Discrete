package main

import (
	"bufio"
	"os"
	"strconv"
)

// multiply обрабатывает умножение в польской записи
func multiply(reader *bufio.Reader, symbol *byte) int {
	*symbol, _ = reader.ReadByte()
	var result = 1
	for *symbol != ')' {
		if *symbol == '(' {
			result *= parsePolish(reader)
		}
		if *symbol >= '0' && *symbol <= '9' {
			result *= int(*symbol - '0')
		}
		*symbol, _ = reader.ReadByte()
	}
	return result
}

// subtract обрабатывает вычитание в польской записи
func subtract(reader *bufio.Reader, symbol *byte) int {
	*symbol, _ = reader.ReadByte()
	var result int
	var isSubtracted bool
	for *symbol != ')' {
		var stepNum int
		if *symbol == '(' {
			stepNum = parsePolish(reader)
		} else if *symbol >= '0' && *symbol <= '9' {
			stepNum = int(*symbol - '0')
		}
		if *symbol >= '0' && *symbol <= '9' || *symbol == '(' {
			if isSubtracted {
				stepNum = -stepNum
			}
			isSubtracted = true
		}
		result += stepNum
		*symbol, _ = reader.ReadByte()
	}
	return result
}

// add обрабатывает сложение в польской записи
func add(reader *bufio.Reader, symbol *byte) int {
	*symbol, _ = reader.ReadByte()
	var result int
	for *symbol != ')' {
		if *symbol == '(' {
			result += parsePolish(reader)
		}
		if *symbol >= '0' && *symbol <= '9' {
			result += int(*symbol - '0')
		}
		*symbol, _ = reader.ReadByte()
	}
	return result
}

// parsePolish выбирает обработчик для действия в польской записи
func parsePolish(reader *bufio.Reader) int {
	var symbol byte
	var result int
	for symbol != '\n' && symbol != ')' {
		symbol, _ = reader.ReadByte()
		switch symbol {
		case '+':
			result += add(reader, &symbol)
		case '-':
			result += subtract(reader, &symbol)
		case '*':
			result += multiply(reader, &symbol)
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)
	_, _ = writer.WriteString(strconv.Itoa(parsePolish(reader)) + "\n")
}
