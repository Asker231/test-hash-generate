package main

import (
	"fmt"

	"github.com/Asker231/test-hash-generate.git/internal/user"
	"github.com/Asker231/test-hash-generate.git/pkg/writer"
)


func main() {
	person := user.NewUser("Asker","Google.com");
	person2 := user.NewUser("Jhon","Google.com");
	fmt.Printf("Id: %s\n Name: %s\n Url:%s\n",person.Id,person.Name,person.Url)
	err := writer.WriteHandle(*person)
	if err !=nil{
		fmt.Print(err)
	}
	err = writer.WriteHandle(*person2)
	if err != nil{
		fmt.Println(err)
	}
}