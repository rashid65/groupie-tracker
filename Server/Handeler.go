package groupie

import (
	"log"
	"net/http"
	"html/template"
)
func PageHandle(w http.ResponseWriter, r *http.Request) (int) {
	//Validating the request path
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		t, err := template.ParseFiles("templates/404.html")
		t.Execute(w,r)
		if err != nil{
			log.Println("Error: no error file")
		}
		return 1
	}
	return 0
}
func ServerError(w http.ResponseWriter, r *http.Request, er error){
	w.WriteHeader(http.StatusInternalServerError)
	t, err := template.ParseFiles("templates/500.html")
	t.Execute(w,r)
	log.Println(er)
	if err != nil{
		log.Println("Error: no error file")
	}
}

func ClientError(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusBadRequest)
	t, err := template.ParseFiles("templates/400.html")
	t.Execute(w,r)
	if err != nil{
		log.Println("Error: no error file")
	}
}

func MethodError (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/405.html")
	t.Execute(w,r)
	if err != nil {
		log.Println("Error: no error file")
	}
}