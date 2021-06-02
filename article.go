// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
    ID      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article 
//func homePage(w http.ResponseWriter, r *http.Request) {
   // fmt.Fprintf(w, "Welcome to the HomePage!")
    //fmt.Println("Endpoint Hit: homePage")
//}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.ID == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range Articles{
        if item.ID == params["id"]{
           Articles = append(Articles[:index], Articles[index+1:]...)
           break  
            }      
	    }
        json.NewEncoder(w).Encode(Articles)
}      

func main() {
    myRouter := mux.NewRouter().StrictSlash(true)
    Articles = append(Articles,Article{
    	ID : "1",
        Title:   "SACHIN",
    	Desc:    "YES DESCRIPTION",
    	Content: "YES CONTENT",
    })
    Articles = append(Articles,Article{
        ID : "2",
    	Title:   "RUBY",
    	Desc:    "NODESCRIPTION",
    	Content: "NO CONTENT",
    })
     Articles = append(Articles,Article{
     	ID : "3",
        Title:   "VIRAT",
     	Desc:    "YNOTES DESCRIPTION",
     	Content: "NOT CONTENT",
     })
    
    // myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe("9878", myRouter))
}