package serverManager

import (
	"bytes"
	"fmt"
	"httpinterface"
	"log"
	"net/http"
)

var (
	b bytes.Buffer
)

// MakeServer - Create our MUX server
func MakeServer() {
	// Define our mux server

	mux := http.NewServeMux()
	mux.HandleFunc("/rgb/", func(w http.ResponseWriter, r *http.Request) {
		// Call our manager here...
		rgbHTTP := colorHTTPInterface.RgbHandler{
			R: r,
			W: w,
		}

		data, e := rgbHTTP.HandleReq()

		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(data)
		w.Write(data)
	})

	mux.HandleFunc("/hsv", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/cymk", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/ycbcr", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/hsl", func(w http.ResponseWriter, r *http.Request) {

	})
	// Listen our server
	fmt.Println("run server")
	log.Fatal(http.ListenAndServe(":1698", mux))
}
