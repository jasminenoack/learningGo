package main // handler echoes the HTTP request.
import (
	"lissajous"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	http.HandleFunc("/two", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous2(w)
	})
	http.HandleFunc("/three", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous3(w)
	})
	http.HandleFunc("/four", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous4(w)
	})

	http.HandleFunc("/five", func(w http.ResponseWriter, r *http.Request) {
		lissajous.ColorGifs(w)
	})
	http.HandleFunc("/six", func(w http.ResponseWriter, r *http.Request) {
		lissajous.CheckerboardMoving(w)
	})

	http.HandleFunc("/seven", func(w http.ResponseWriter, r *http.Request) {
		lissajous.CheckerboardChanging(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
