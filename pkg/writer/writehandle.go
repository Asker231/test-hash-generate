package writer

import (
	"encoding/json"
	"fmt"
	"os"
)



func WriteHandle[T any](data T)error{
	arr := []T{}
	arr = append(arr, data)	
	users,err:=	json.Marshal(&arr)
	if err != nil{
		return  err
	}
	file,err := os.Create("db.json")
    if err != nil{
		return err
	}
	file,err = os.OpenFile(file.Name(),os.O_APPEND|os.O_CREATE,0600)
	if err != nil{
		fmt.Println(err)
	}
	err = os.WriteFile(file.Name(),users,0644)

	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	return nil
}