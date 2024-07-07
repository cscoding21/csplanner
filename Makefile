up:
	./scripts/run.sh

down:
	tilt down

reset:
	minikube delete

deps:
	./scripts/devdeps.sh

appserv:
	cd ./csserver && \
	go run cmd/appserver/main.go;

gql:
	cd ./csserver/internal/appserv && \
		go run github.com/99designs/gqlgen generate && \
		cd ../.. && \
		go generate ./...;