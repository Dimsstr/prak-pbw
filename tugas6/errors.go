package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "Dimas" {
		return NotFoundError
	}

	// success
	return nil
}

func main() {
	err := GetById("Dimas")
	if err != nil {
		if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else {
			fmt.Println("Unknown error")
		}
	}
}
