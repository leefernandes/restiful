package restiful

import (
	"net/http"
)

type (
	Handler func (w http.ResponseWriter, r *http.Request) (error)
)

func Handle(handlers ...Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range handlers {
			err := handler(w, r)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
		}
	})
}