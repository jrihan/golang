package main

import (
	"errors"
	"fmt"
)

func main() {

	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// int, uint
	// uintptr
	// byte
	// rune
	// float32, float64
	// complex64, complex128
	// string
	// bool

	var numero int16 = 10000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numero5 float32 = 123.45
	fmt.Println(numero5)

	var numero6 float64 = 123.4567890123456789
	fmt.Println(numero6)

	var string1 string = "String 1"
	fmt.Println(string1)

	char := 'B'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error = errors.New("Erro genérico")
	fmt.Println(erro)
}