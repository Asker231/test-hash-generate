package main

import (
	"fmt"

	"github.com/Asker231/test-hash-generate.git/pkg/hash"
)
func main() {
	hash := hash.HashGenerate()
	fmt.Println(hash(10))
}