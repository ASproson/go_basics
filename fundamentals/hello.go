package main

import "fmt"

const hello = "Hello, "

func Hello(name string) string {
	if name == "" {
		return "Hello, World"
	}
	return hello + name
}

func main() {
	fmt.Println(Hello("world"))
}

// Go programs will have a main package defined with a main func inside it
// Packages are ways of grouping related Go code together
// fmt is a package that we import
