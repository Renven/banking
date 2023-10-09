package app

import (
	"log"
	"net/http"

	"github.com/Renven/banking/domain"
	"github.com/Renven/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}
