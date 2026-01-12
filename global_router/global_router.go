package global_router

import "net/http"

func GlobalRouter(mux *http.ServeMux) http.Handler {
	handleReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-control-allow-origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handleReq)
}