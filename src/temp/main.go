package main

import (
	"fmt"
	"params"
	"reflect"
)

func main() {
	var num int8 = 1
	var str = "Hello"

	_, name := "丢弃", "ZSW"

	_, e := fmt.Println(str, num, name)

	println(e)
	println(params.Pi)
	println(params.MaxThread)

	testNumber()
	testString()
}

func testBool() {
	var enable, disable = true, false

	println(enable, disable)
}

func testNumber() {
	var num rune = 12
	println(num)

	var a int32
	var b int32

	c := a + b

	println(c)

	var n int = 1

	println(reflect.TypeOf(n).Name())

}

func testString()  {
	var str string = "字符串"
	name := `name 这也是字符串`
	fmt.Println(str, name)
	c := []byte(name)

	c[0] = 'c'

	s2 := string(c)

	fmt.Printf("s2: %v", s2)

	h := "Hello "
	w := "World!"
	hw := h + w
	println(hw)


	println("c" + h[1:])

	ss := `这是一段文本
                 原样输出， 带换行
           带空格`

	println(ss)

}
