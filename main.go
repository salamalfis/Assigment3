package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Mulai goroutine untuk mengupdate file JSON setiap 15 detik
	go updateJSON()

	// Mulai server web
	http.HandleFunc("/", statusHandler)
	http.Handle("/data.json", http.FileServer(http.Dir("."))) // Serve file data.json
	http.ListenAndServe(":8080", nil)
}

func updateJSON() {
	for {
		// Generate angka acak untuk water dan wind antara 1-100
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// Membuat struktur Data dengan nilai yang di-generate
		data := Data{
			Water: water,
			Wind:  wind,
		}

		// Mengubah struct Data menjadi JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			continue
		}

		// Menuliskan JSON ke dalam file
		err = ioutil.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}

		fmt.Printf("Data updated: Water = %d, Wind = %d\n", water, wind)

		// Tunggu 15 detik untuk update berikutnya
		time.Sleep(15 * time.Second)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
