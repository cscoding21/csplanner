up:
	./scripts/run.sh

down:
	tilt down

deps:
	./scripts/devdeps.sh

gql:
	cd ./csserver/internal/appserv && \
		go run github.com/99designs/gqlgen generate && \
		cd ../.. && \
		go generate ./...;