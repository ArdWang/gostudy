package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s :="Yes我爱慕课网!" // UTF-8 编码
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%s\n", []byte(s))
	for _, b := range [] byte(s){
		fmt.Printf("%x ",b)
	}
	fmt.Println()

	for i, ch := range s{ // ch is a rune 4字节整数
		fmt.Printf("(%d %x)",i,ch)
	}

	fmt.Println()

	fmt.Println("Rune count:",utf8.RuneCountInString(s))

	//fmt.Printf("%X\n", []byte(s))

	bytes :=[] byte(s)
	for len(bytes) >0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}
	fmt.Println()

	for i, ch :=range [] rune(s){
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

}
