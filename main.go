package main

import (
	"fmt"

	say_hello "github.com/onggolt-dev/go-mod-sample"
	say_hello_v2 "github.com/onggolt-dev/go-mod-sample/v2"
)

func main() {
	fmt.Println(say_hello.SayHello("Tiyan"))
	fmt.Println(say_hello_v2.SayHello("Tiyan", "male"))
}
