package handlers

import (
	"fmt"
	"net/http"

	"github.com/fabiosebastiano/go-microservices/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: ValidationError
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(*prod)
}
