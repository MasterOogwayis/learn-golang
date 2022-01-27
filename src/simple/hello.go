package main

import "fmt"
import "sys"

func main() {
	fmt.Println("Hello World!", "你好，世界！", "καλημ ́ρα κóσμ", "こんにちはせかい")
	n,e := sys.Printf("This is a test, %v \n", "Hello World")
	fmt.Println(n)
	fmt.Println(e)
}
