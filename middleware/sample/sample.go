package sample

import (
	"context"
	"net/http"

	"github.com/goadesign/goa"
)

func New() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			rw.Header().Set("X-Sample", "sample")
			return h(ctx, rw, req)
		}
	}
}
