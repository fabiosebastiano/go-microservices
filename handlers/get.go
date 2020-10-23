package handlers

import (
	"net/http"

	"github.com/fabiosebastiano/go-microservices/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts ritorna una lista di tutti i prodotti
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
