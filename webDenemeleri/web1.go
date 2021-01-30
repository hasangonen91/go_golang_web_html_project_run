package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Merhaba Dunya %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9191", nil)

	fmt.Println("Web Sunucu")
}
/*
erisime izin ver dedikten sonra tarayıcıya localhost:9191 diye arattiktan sonra 
Merhaba Dunya yazan bir sekme acilacak localhost:9191 kismina "/naber hasan" yazarsaniz sekmede 
Merhaba Dunya naber hasan diye yenilenecektir
*/