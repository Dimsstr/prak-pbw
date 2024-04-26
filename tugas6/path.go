package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	fileInfo, err := os.Stat("DimasSatrioParikesit")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("DimasSatrioParikesit adalah sebuah direktori")
	} else {
		fmt.Println("DimasSatrioParikesit adalah sebuah file")
	}
}
