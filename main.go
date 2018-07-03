package main

import "fmt"

func main() {
	// Integer Palendrome
	fmt.Printf("%v\n", integerPalindrome(12921))

	// Chuck Noris API
	fmt.Printf("%v\n", jokeCategory("dev"))
	fmt.Printf("%v\n", jokeSearch("kill"))

	// Graph Traversal
	fmt.Println("\nGraph A")
	for _, path := range findPath("1 2 3\n2 1 3\n3 1 2 4\n4 3") {
		fmt.Printf("%+v\n", path)
	}

	fmt.Println("\nGraph B")
	for _, path := range findPath("a b\nb a\nc") {
		fmt.Printf("%+v\n", path)
	}

	fmt.Println("\nGraph C")
	for _, path := range findPath("1 4\n2 5\n3 6\n4 1 5 6\n5 2 4 6\n6 3 4 5") {
		fmt.Printf("%+v\n", path)
	}
	fmt.Println("")
}
