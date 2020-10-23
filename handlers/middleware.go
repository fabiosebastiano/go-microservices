package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fabiosebastiano/go-microservices/product-api/data"
)

// MiddlewareProductValidation validates the product in the request and calls next if ok
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := data.FromJSON(prod, r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializing product ", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] VALIDATING product ", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %v", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
