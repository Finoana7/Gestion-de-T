package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ajouter les en-têtes CORS
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Autoriser les requêtes depuis localhost:5173
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Authorization", "application/json")
		// Si c'est une requête OPTIONS (préflight), renvoyer directement un 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passer au prochain handler
		next.ServeHTTP(w, r)
	})
}
