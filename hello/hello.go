package main

import "fmt"

func main() {
	fmt.Println(Hello("Nadia"))
}

func Hello(name string) string {
	return "Hello " + name + "!"
}
