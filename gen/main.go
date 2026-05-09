package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("../tailwind.js")

	out := fmt.Sprintf(`package tailwind

const JS = %q
`, string(data))
	fmt.Print(out)
}
