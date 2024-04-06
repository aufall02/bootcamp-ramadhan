package utils

import "fmt"

type Admin struct {
	IdAdmin uint
	Name    string
	Status  bool
}

func Hello(name string) string {

	return fmt.Sprintf("selamat datang %s", name)
}
