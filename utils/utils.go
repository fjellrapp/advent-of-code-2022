package utils

import (
	"fmt"
	"os"
)

func ReadAndReturnFile(file string) []byte {
	dat, err := os.ReadFile(file)
	Check(err)
	return dat
}
func OpenFile(file string) *os.File {
	opened, err := os.Open(file)
	Check(err)
	return opened
}
func Check(e error) {
	if e != nil {
		panicMessage := fmt.Sprintf("Fails: %f", e)
		panic(panicMessage)
	}
}
