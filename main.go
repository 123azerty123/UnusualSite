package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

var t *template.Template
var err error

type indexData struct {
	SPName string
	SPLink string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // text/plain
	spNum := rand.Intn(len(suggestedPages))
	err := t.ExecuteTemplate(w, "index.html", indexData{suggestedPages[spNum][0], suggestedPages[spNum][1]})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func tests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // text/plain
	err := t.ExecuteTemplate(w, "menu.html", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func formTest(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "form-test.html", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.FormValue("fname"), r.FormValue("lname"))
}

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // text/plain

	switch r.URL.Path {
	case "/":
		home(w, r)
	case "/tests":
		tests(w, r)
	case "/form-test":
		formTest(w, r)
	case "/furtis_music":
		http.Redirect(w, r, furtisMusics[rand.Intn(len(furtisMusics))], http.StatusSeeOther)
	case "/easter_egg":
		http.Redirect(w, r, "https://youtu.be/MTW4sIL9Dpw", http.StatusSeeOther)
		fmt.Fprint(w, "No fighter has escaped it.")
	default:
		fmt.Fprint(w, "<h1 style='color:red;font-size: 20vh;text-align: center;margin-top: 40vh;'>404</h1>")
	}
}

func main() {
	t, err = template.ParseGlob("*.html") //template.ParseFiles("index.html",)
	if err != nil {
		fmt.Println(err)
		return
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	fs = http.FileServer(http.Dir("styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", fs))

	http.HandleFunc("/", handlerFunction)
	http.ListenAndServe("", nil)

	//https:
	//http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
	//go run "C:\Program Files\Go\src\crypto\tls\generate_cert.go" --host=localhost
}

var suggestedPages [][2]string = [][2]string{
	{"The Unusual Discord", "https://discord.gg/k3fGJwJqzu"},
	{"An Easter Egg", "./easter_egg"},
	{"A Musical Work Of Art", "./furtis_music"},
}
var furtisMusics []string = []string{
	"https://youtu.be/Eu1V8ECWcUE",
	"https://youtu.be/OmXMaGgF2gc",
	"https://youtu.be/5kvy1qguzps",
	"https://youtu.be/Yke56tRxbSg",
	"https://youtu.be/xzPGYpA8TXI",
	"https://youtu.be/b_Smll3A8Zg",
	"https://youtu.be/bD2iLtJzfag",
	"https://youtu.be/MmkAuMnYPGk",
}
var footer string = "<div id=footer><p>Not associated with or approved by Mojang Studios</p></div>"
