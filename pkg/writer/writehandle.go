package writer

import (
	"encoding/json"
	"fmt"
	"os"
)



func WriteHandle[T any](data... T){
	arr := []T{}
	arr = append(arr, data...)
	dByte,_ := json.Marshal(arr)
	err := os.WriteFile("db.json",dByte,0644)
	if err != nil{
		fmt.Println(err)
	}
}

