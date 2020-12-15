package main

import (
	"fmt"
	"net/http"
)

func adminCheck(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("user") != "admin" {
				http.Error(w, "Not Authorized", 401)
				return
			}

			fmt.Fprintln(w, "Admin Portal")
			h.ServeHTTP(w, r)
		})
}

func newStatusCode(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		h.ServeHTTP(w, r)
	})
}

func addHeader(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Foo", "Bar")
		h.ServeHTTP(w, r)
	})
}

func writeResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello PerfGo!")
}

func main() {
	handler := http.HandlerFunc(writeResponse)
	http.Handle("/", addHeader(newStatusCode(handler)))
	http.Handle("/onlyHeader", addHeader(handler))
	http.Handle("/admin", adminCheck(handler))
	http.ListenAndServe(":1234", nil)
}
