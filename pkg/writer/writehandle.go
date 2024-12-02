package writer

import (
	"encoding/json"
	"fmt"
	"os"
)


const filename = "db.json"
func WriteHandle[T any](data T){
	payloads := []T{}
	jsonData := Reader(filename)
	err := json.Unmarshal(jsonData,&payloads)
	if err != nil{
		fmt.Println(err)
	}
	payloads = append(payloads, data)
	arrbyte ,err:= json.Marshal(&payloads)
	if err != nil{
		fmt.Println(err)
	}
	_ = os.WriteFile(filename,arrbyte,0644)
}
func Reader(filename string)[]byte{
	data,err := os.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}
	return data
}




