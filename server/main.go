package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//todos:
// 1.) implement db connection storing votes, users
//
// 2.) implement google photos api
// 	2.1) implement oauth flows
// 	2.2) https://pkg.go.dev/golang.org/x/oauth2/google
// 3.) implement user authentication
// 	3.1) implement jwts
// 	3.2) implement user registration
// 	3.3) implement user login
//

func addHandlers(mux *http.ServeMux, apiConf apiConfig) {
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	mux.HandleFunc("GET /picker/qr", apiConf.showQRCode)
	//todo
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	// dbUrl := os.Getenv("DB_URL")

	if port == "" {
		log.Println("Could not load port number from ENV file, defaulting to port 8080")
		port = "8080"
	}
	// if dbUrl == "" {
	// 	log.Fatal("Could not retreive connection string from ENV file")
	// }
	apiConf := apiConfig{}
	mux := http.NewServeMux()
	addHandlers(mux, apiConf)

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
