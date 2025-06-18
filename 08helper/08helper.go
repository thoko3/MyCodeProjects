package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func greet(name string) string{
	return "Hello" + name + "!"
}
func main() {
reader := bufio.NewReader(os.Stdin)
fmt.Print("What is your name?")
nameinput,_:=reader.ReadString('\n')
nameinput=strings.TrimSpace(nameinput)
message :=greet("nameinput ")
fmt.Println(message)
}