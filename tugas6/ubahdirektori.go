package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new directory
	err := os.Chmod("DimasSatrioParikesit", 0002)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Directory 'DimasSatrioParikesit' permissions successfully changed")
}
