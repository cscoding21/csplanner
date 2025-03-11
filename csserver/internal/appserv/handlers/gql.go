package handlers

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/gorilla/websocket"

	"csserver/internal/appserv/graph"
	"csserver/internal/appserv/graph/resolver"
	"csserver/internal/appserv/middleware"

	log "github.com/sirupsen/logrus"
)

// return a configured graphql service handler
func GetGqlHandler() *handler.Server {
	gqlService := handler.New(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &resolver.Resolver{}},
		))

	gqlService.AddTransport(transport.Options{})
	gqlService.AddTransport(transport.GET{})
	gqlService.AddTransport(transport.POST{})
	gqlService.AddTransport(transport.MultipartForm{})

	//gqlService.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	gqlService.Use(extension.Introspection{})
	// gqlService.Use(extension.AutomaticPersistedQuery{
	// 	Cache: lru.New[string](100),
	// })

	gqlService.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				log.Debugf("CHECK ORIGIN HOST: %s", r.Host)
				return true //r.Host == "localhost:4000"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			return middleware.WebSocketInit(ctx, initPayload)
		},
	})

	return gqlService
}
