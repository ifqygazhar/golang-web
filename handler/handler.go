package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error, hubungi admin", http.StatusInternalServerError)
		return
	}
	// data := map[string]interface{}{
	// 	"title":   "Golang web",
	// 	"content": "saya belajar golang web",
	// }

	// data := entity.Product{Id: 1, Name: "Laptop", Price: 250000, Stock: 3}

	data := []entity.Product{
		{Id: 1, Name: "Laptop", Price: 250000, Stock: 8},
		{Id: 3, Name: "Kalung", Price: 250000, Stock: 11},
		{Id: 2, Name: "Rog", Price: 250000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error, hubungi admin", http.StatusInternalServerError)
		return
	}

}
func HeloHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func AzharHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello saya azhar"))
}
func ProductHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}
	//fmt.Fprint(w, "Product page : ", idNumber)
	data := map[string]interface{}{
		"content": idNumber,
	}
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "error hubungi admin", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error hubungi admin", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah get"))
	case "POST":
		w.Write([]byte("ini post"))
	default:
		http.Error(w, "error", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error hubungi admin", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "error hubungi admin", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "error hubungi admin", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "error hubungi admin", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")

		data := map[string]interface{}{
			"name":  name,
			"pesan": pesan,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "error hubungi admin", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "error hubungi admin", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "error hubungi admin", http.StatusBadRequest)
}
