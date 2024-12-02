package main

import (
	"github.com/Asker231/test-hash-generate.git/internal/user"
	"github.com/Asker231/test-hash-generate.git/pkg/writer"
)


func main() {
	person := user.NewUser("Asker","Googlse.com");
	person2 := user.NewUser("Jhon","Googles.com");
	writer.WriteHandle(*person)
	writer.WriteHandle(*person2)
}