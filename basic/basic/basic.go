package main

// 2020/7/7 20：49 学到2-3章节

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
 	bb = true
)


// bb: = true  //这个是不可以

func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction(){
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter()  {
	a,b,c,s :=3,4,true,"def"
	b = 5
	c = false
	fmt.Println(a,b,c,s)

	//euler()
}



func euler()  {
	fmt.Printf("%.3f",
		cmplx.Exp(1i*math.Pi)+1,
		)
}

//func triangle()  {
//	var a, b int = 3,4
//	var c int
//	c = int(math.Sqrt(float64(a*a+b*b))) // sqrt 求平方
//	fmt.Println(c)
//}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a,b))
}

func calcTriangle(a, b int) int{
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	return c
}

//常量
const (
	filename = "abc.txt"
	a,b = 3,4
)

func consts(){
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c)
}


//枚举类型
func enums(){

	const (
		cpp = iota //自增值
		_
		python
		golang
		javascript
	)

	// b ,kb , mb , gb ,tb, pb
	const (
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {
	//fmt.Println("Hello world!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//
	//fmt.Println(aa,ss,bb)
	//
	//euler()
	//
	//triangle()
	//consts()

	enums()
}

/// 使用 var 关键字
/// var a,b,c bool
/// :只能在函数内部使用 函数外部是不能使用得
/// bool, string
/// u ,int ,int8 ,int16 ,int32 ,int64, uintptr(指针)
/// byte, rune(字符型)
/// float32 ,float64, complex64 ,complex128 (复数)

