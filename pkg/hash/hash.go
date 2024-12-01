package hash

import "math/rand"


type Closure = func(int)string

func HashGenerate()Closure{
    set := []rune("qwertyuio1pl2kj3hg4fd5sa6zx7cv8bnm9#%&!@") 
	return func(length int)string{
		hash := make([]rune,length)
		for i := range hash{
			hash[i] = set[rand.Intn(len(set))]
		}
		return string(hash)
	}
}