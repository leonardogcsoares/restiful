//Package restiful is a fork that serves as facilitator for Middleware implementation
// using httprouter
package restiful

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler function is a wrapper for the way HttpRouter deals with httpHandlers
type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error

// Handle calls each of the Handlers for the API endpoint in order.
func Handle(handlers ...Handler) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		for _, handler := range handlers {
			err := handler(w, r, ps)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	})
}
