package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type Data struct{
  Name string 
}
func home(w http.ResponseWriter,r *http.Request){
  tmpl,_:=template.ParseFiles("index.html")
  data := Data{Name:"Vasanth"}
  tmpl.Execute(w,data)
}
func main(){
  http.HandleFunc("/",home)
  http.ListenAndServe(":8080",nil)
}
