package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Product struct{
	ProductID int
	ProductName string
}

var apiProductList []Product

func dataGet(){
	product1 := Product{1,"Bilgisayar"}
	product2 := Product{2, "Fritöz"}
	apiProductList = append(apiProductList, product1, product2)
}


func Api(){
	fmt.Println("Api")

	dataGet()

	r := mux.NewRouter()


	// http.HandleFunc("/",index)
	// http.HandleFunc("/post",post)
	r.HandleFunc("/", index)
	r.HandleFunc("/post", post) //gorila mux http yerine yazılır
	r.HandleFunc("/post/{category}/{id}", post)

	r.HandleFunc("/product", product)
	r.HandleFunc("/user", user)

	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", r) // gorilla da r gelir buraya


}

func index(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("index page!!!!"))
}

func post(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	category := vars["category"]
	id := vars["id"]

	var a string
	if r.Method == "POST"{
		a="post"
	} else if r.Method == "GET"{
		a="get"
	}

	w.Write([]byte("post page!!!!, post id : "+ id + "\ncategory : " + category+"\nmethod :"+a))
}

func product(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	j,err := json.Marshal(apiProductList)
	if err !=nil {fmt.Println(err)}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(j))
}

func user(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL.Path)
	
	log.Println(r.FormValue("username"))
	log.Println(r.FormValue("password"))


	fmt.Fprintf(w, "Success")
}


