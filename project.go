package main
import (
	"log"
	"net/http"
	"encoding/json"
"github.com/gorilla/mux"

)
type Person struct{
	ID string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address *Address `json:"address"`
}
type Address struct{
	city string `json:"city"`
}
var people []Person
func GetPersons(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(people)

}
func GetPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _,item := range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
            return		
		}
	}
	json.NewEncoder(w).Encode(people)
}
func CreatePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
    var Person Person
    _=json.NewDecoder(r.Body).Decode(&people)
	people =append(people, Person)
	json.NewEncoder(w).Encode(people)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
    params := mux.Vars(r)
	for index, item :=range people{
		if item.ID == params["id"]{
			people = append(people[:index],people[index+1:]...)
			var Person Person
			_ = json.NewDecoder(r.Body).Decode(&people)
			item.ID =params["id"]
			people =append(people, Person)
			json.NewEncoder(w).Encode(people)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item :=range people{
		if item.ID == params["id"]{
			people = append(people[:index],people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
func main(){
	myRouter:= mux.NewRouter()
	people = append(people, Person {ID:        "1",FirstName: "SACHIN",LastName:  "BAGHEL",Address:   &Address{city:"DELHI"},})
	people = append(people, Person {ID:        "2",FirstName: "rahul",LastName:  "JHHFGHG",Address:   &Address{city:"GDHGFHGSA"},})
	people = append(people, Person {ID:        "3",FirstName: "hhgkjdgh",LastName:  "GHJASFHJL",Address:   &Address{city:"GVGDGI"},})
	
	myRouter.HandleFunc("/Person",GetPersons ).Methods("GET")
	myRouter.HandleFunc("/people/{ID}",GetPerson).Methods("GET")
	myRouter.HandleFunc("/people",CreatePerson).Methods("POST")
	myRouter.HandleFunc("/people{ID}",UpdatePerson).Methods("PUT")
	myRouter.HandleFunc("/people/{ID}",DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10200", myRouter))

	
	
}