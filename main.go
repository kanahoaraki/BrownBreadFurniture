package ma

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "shop.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// step 1: validate form
	// step 2: send message in an email
	// step 3: redirect to confirmation page.
}
func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "confirmation.html", nil)
}
