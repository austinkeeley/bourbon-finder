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
		results, err := bourbonfinder.Search(config)
		if err != nil {
			return
		}

		m := bourbonfinder.GroupByStore(results)
		w.Write([]byte("<html><body><pre>"))
		bourbonfinder.PrintGroup(m, w, false)
		w.Write([]byte("</pre></body></html>"))
	})
	log.Println("Starting web server on " + address)
	log.Fatal(http.ListenAndServe(address, nil))
}
