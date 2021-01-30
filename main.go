package main

import (
	"net/http"

	_ "./statik" //Oluşturulmuş statik.go dosyasının konumu
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, _ := fs.New()
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	http.ListenAndServe(":9191", nil) //localhost:9191 adres belirleme
}

//bu kisim public ve static icin
