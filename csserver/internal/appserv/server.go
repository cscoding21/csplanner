package appserv

import (
	"fmt"
	"net/http"

	"csserver/internal/appserv/handlers"
	"csserver/internal/appserv/middleware"
	"csserver/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/99designs/gqlgen/graphql/playground"
)

func Serve() error {
	config.InitConfig()
	config.InitLogger()

	//---get a new router
	router := getRouter()

	//---add handlers to router
	//	auth handlers
	router.Handle("/login",
		middleware.OrgMiddleware(
			middleware.AuthMiddleware(handlers.GetLoginHandler()),
		))
	router.Handle("/refresh",
		middleware.OrgMiddleware(
			middleware.AuthMiddleware(
				middleware.ValidationMiddleware(handlers.GetRefreshHandler()),
			)),
	)
	router.Handle("/signout",
		middleware.OrgMiddleware(
			middleware.AuthMiddleware(
				middleware.ValidationMiddleware(handlers.GetSignoutHandler()),
			)),
	)

	//	graphql handlers
	router.Handle("/query",
		middleware.OrgMiddleware(
			middleware.ValidationMiddleware(handlers.GetGqlHandler())),
	)
	if config.Config.Server.EnablePlayground {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	}

	//	file handler
	router.Handle("/files/*", http.StripPrefix("/files/", handlers.GetTusdHandler()))

	//---start web server
	return http.ListenAndServe(fmt.Sprintf(":%v", config.Config.Server.ServerPort), router)
}

func getRouter() *chi.Mux {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Tus-Resumable", "Upload-Metadata", "Upload-Length", "Upload-Offset", "Sec-Websocket-Extensions", "Sec-Websocket-Key", "Sec-Websocket-Protocol", "Sec-Websocket-Version"},
		AllowCredentials: true,
		Debug:            false,
		MaxAge:           300, // Maximum value not ignored by any major browsers
	}))

	return router
}
