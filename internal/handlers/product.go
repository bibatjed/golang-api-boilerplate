package handlers

import (
	"github.com/gorilla/mux"
	"golang-api-boilerplate/internal/models"
	"log"
	"net/http"
	"strconv"
)

//Product struct wth a logger attribute
type Product struct {
	l *log.Logger
}

//NewProduct - creates a product handler with a given logger.
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

//GetProducts : get list of all products
func (p Product) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("handling GET")

	productList := models.GetProducts()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error while marshalling product list", http.StatusInternalServerError)
	}
}

//AddProduct : this will add a new product coming from POST request into existing list of products.
func (p Product) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("handling POST")

	newProduct := &models.Product{} // this prod will have address of Product struct

	err := newProduct.FromJSONtoProduct(r.Body)

	if err != nil {
		http.Error(rw, "error while adding new product", http.StatusInternalServerError)
	}

	models.AddProduct(newProduct)

	rw.WriteHeader(http.StatusCreated)
}

//UpdateProduct : updating a product
func (p Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("handling UPDATE")

	uriParams := mux.Vars(r)
	id, err := strconv.Atoi(uriParams["id"])

	if err != nil {
		http.Error(rw, "unable to parse uri parma product id", http.StatusBadRequest)
		return
	}

	prod := &models.Product{}

	err = prod.FromJSONtoProduct(r.Body)
	if err != nil {
		p.l.Println("unable to deserialize input json", err)
		http.Error(rw, "invalid request", http.StatusBadRequest)
	}

	err = models.UpdateProduct(id, prod)
	if err == models.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Internal server Error while updating product", http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusNoContent)

}

//UpdateProductAttribute : partial update of product attributes
func (p Product) UpdateProductAttribute(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("handling PATCH")
	uriParam := mux.Vars(r)
	pID, err := strconv.Atoi(uriParam["id"])
	p.l.Println("p id-", pID)
	if err != nil {
		http.Error(rw, "unable to parse ui param product id", http.StatusBadRequest)
		return
	}

	prod := &models.Product{}

	err = prod.FromJSONtoProduct(r.Body)
	if err != nil {
		http.Error(rw, "unable to parse request payload", http.StatusBadRequest)
		return
	}

	err = models.UpdateProductAttribute(pID, prod)
	if err != nil {
		http.Error(rw, "failed to perform partial update", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}

//DeleteProduct :
func (p Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	uriParams := mux.Vars(r)
	prodID, err := strconv.Atoi(uriParams["id"])

	if err != nil {
		http.Error(rw, "unable to convert prod id", http.StatusBadRequest)
		return
	}

	err = models.DeleteProductByID(prodID)

	if err == models.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "error deleting product", http.StatusInternalServerError)
		return
	}

}