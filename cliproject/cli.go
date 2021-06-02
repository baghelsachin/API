ackage main

import (
	"encoding/json"
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
)
type Customer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Class string `json:"class"`
	Section string `json:"section"`

}
var students []Customer

func getstudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(students)
}
func getstudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _,item := range students{
		if item.ID == params["id"]{
	json.NewEncoder(w).Encode(item)
	return	
        }
	}
json.NewEncoder(w).Encode(students)
}
func createstudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var student Student
	_=json.NewDecoder(r.Body).Decode(&students)
	students = append(students,student)
	json.NewEncoder(w).Encode(students)
}
func updatestudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params :=mux.Vars(r)
	for index, item := range students{
	    if item.ID == params["id"]{
		  students = append(students[:index],students[index+1:]...)
		   var student Student
		   _ =json.NewDecoder(r.Body).Decode(&students)
		   item.ID = params["id"]
		   students = append(students, student)
		    json.NewEncoder(w).Encode(students)
	          return
	     }
    }
}
func deletestudent(w http.ResponseWriter, r *http.Request){
	 := w.Header().Set("content-type", "application/json")
	
	params := mux.Vars(r)
	for index, student := range students{
	    if student.ID == params["id"]{
         students= append(students[:index],students[index+1:]...)
         break
		}
	} 
	json.NewEncoder(w).Encode(students)
}
func main(){
r := mux.NewRouter()
students = append(students,Student{ ID: "1",Name:"SACHIN", Class : "12",Section: "A"})
students = append(students,Student{ ID: "2",Name:"RUBY", Class : "10",Section: "B"})
students = append(students,Student{ ID: "3",Name:"VIRAT", Class : "9",Section: "R"})
r.HandleFunc("/student", getstudent).Methods("GET")
r.HandleFunc("/student/{id}", getstudent).Methods("GET")
r.HandleFunc("/student", createstudent).Methods("POST")
r.HandleFunc("/student/{id}", updatestudent).Methods("PUT")
r.HandleFunc("/student/{id}", deletestudent).Methods("DELETE")
log.Fatal(http.ListenAndServe(":10078",r))
}