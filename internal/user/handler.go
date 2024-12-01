package user
import "github.com/Asker231/test-hash-generate.git/pkg/hash"

type User struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func NewUser(name,url string)*User{
	id := hash.HashGenerate()
	return &User{
		Id: id(10),
		Name: name,
		Url: url,
	}
}

