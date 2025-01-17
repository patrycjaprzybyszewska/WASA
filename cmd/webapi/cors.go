package main

import (
	"github.com/gorilla/handlers"
	"net/http"
)

// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{
			"content-type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With", "Authorization",
		}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "searcher_id",
			"identifier", "requester_ID", "user_id", "user_name", "commenter_name"})
	)(h)
}
	// Some of the allowed headers are legacy.
