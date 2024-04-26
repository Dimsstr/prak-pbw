package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a new directory
	err := os.Mkdir("DimasSatrioParikesit", 0777)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Directory 'DimasSatrioParikesit' successfully created")
}
