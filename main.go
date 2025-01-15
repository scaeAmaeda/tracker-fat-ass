package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	go mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		db := ConnectDB("db.sqlite")
		defer db.Close()
		type ListMeals struct {
			ListMeals []Meal
		}
		data := ListMeals{ListMeals: GetMeals(db)}
		tmplIndex := template.Must(template.ParseFiles("./index.html"))
		tmplIndex.Execute(w, data)
	})

	go mux.HandleFunc("POST /meal", func(w http.ResponseWriter, r *http.Request) {
		is300g, _ := strconv.Atoi(r.FormValue("is300g"))
		isnosugar, _ := strconv.Atoi(r.FormValue("isnosugar"))
		meal := Meal{Date: r.FormValue("dateMeal"), Moment: r.FormValue("moment"), Is300g: is300g, IsNoSugar: isnosugar}
		db := ConnectDB("db.sqlite")
		defer db.Close()
		AddMeal(db, meal)
		type ListMeals struct {
			ListMeals []Meal
		}
		data := ListMeals{ListMeals: GetMeals(db)}
		tmplIndex := template.Must(template.ParseFiles("./index.html"))
		tmplIndex.Execute(w, data)
	})

	log.Println("Application runs on http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", mux))

}
