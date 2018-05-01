package api

import  (
	"net/http"
	"github.com/srleyva/CertAPI/pkg/certs"
	"bytes"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
)

// Set up new Server
func NewPkiServer() http.Handler {
	err := certs.Initialize()
	if err != nil {
		panic(fmt.Errorf("error initializing: %s", err))
	}
	r := mux.NewRouter()
	r.HandleFunc("/pki/NewCSR", NewCSRHandler)
	r.HandleFunc("/pki/config", ConfigHandler)
	return r
}


// Routing for /pki/NewCSR
func NewCSRHandler(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r.Body)
	if csr, err := certs.NewCSR(buffer.String()); err == nil {
		w.Write(csr)
	} else {
		w.Write([]byte("Error! No CSR Generated"))
	}
}

// Routing for /certs/config
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := certs.GetConfig()
	json.NewEncoder(w).Encode(config)

}