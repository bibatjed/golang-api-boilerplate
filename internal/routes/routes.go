package routes

import (
	"github.com/gorilla/mux"
	"golang-api-boilerplate/internal/handlers"
	"log"
	"net/http"
)

func InitializeRoutes(log *log.Logger) *mux.Router {
	productHandler := handlers.NewProduct(log)
	servMux := mux.NewRouter()
	// registers product handler methods to serve request on api end points with specific http methods.
	getHandler := servMux.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/products", productHandler.GetProducts)

	postHandler := servMux.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/products", productHandler.AddProduct)

	putHandler := servMux.Methods(http.MethodPut).Subrouter()
	putHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)

	patchHandler := servMux.Methods(http.MethodPatch).Subrouter()
	patchHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProductAttribute)

	deleteHandler := servMux.Methods(http.MethodDelete).Subrouter()
	deleteHandler.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct)

	return servMux
}
