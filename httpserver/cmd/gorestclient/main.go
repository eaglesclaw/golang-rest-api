package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Product struct{
	ProductID int
	ProductName string
}

func main(){

	url := "http://127.0.0.1:8080/product"


	resp, err := http.Get(url)
	checkError(err)

	jsonDataFromHttp , err := ioutil.ReadAll(resp.Body)
	checkError(err)

	var jsonData []Product
	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	checkError(err)
	fmt.Println(jsonData)

}


func checkError(e error){
	if e != nil {panic(e)}
}