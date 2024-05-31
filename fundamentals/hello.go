package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

// Go programs will have a main package defined with a main func inside it
// Packages are ways of grouping related Go code together
// fmt is a package that we import
