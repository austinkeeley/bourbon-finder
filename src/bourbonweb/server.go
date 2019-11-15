package bourbonweb

import (
	"bourboncommon"
	"bourbonfinder"
	"log"
	"net/http"
)

func StartWebServer(address string, config *bourboncommon.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling incoming request")
		bourbonfinder.Search(config)
	})
	log.Println("Starting web server on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
