package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

"github.com/gorilla/mux"
)
type Article struct{

	Title string `json:"title"`
	Desc string `json:"description"`
	Content string `json:"content"`


}
var articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Endpoint Hit: All Article Hit Endpoints")
	json.NewEncoder(w).Encode(articles)
}
func testpostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"test post end point worked") 
}
func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Homepage Endpoint Hit")
}

	
func main(){
	articles = append(articles,Article{Title: "test title", Desc: "test description", Content: "content",})
		
	myRouter:= mux.NewRouter()
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testpostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8086", myRouter))


}