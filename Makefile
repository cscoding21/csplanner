
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## goinstalls: install golang tools used by this project
.PHONY: goinstalls
goinstalls:
	cd ./csserver && \
	../scripts/install.sh

## up: start the minikube server if not running and configure if necessary
.PHONY: up
up:
	./scripts/run.sh

## down: shut down all app services
.PHONY: down
down:
	tilt down

## reset: delete the minikube cluster and configuration
.PHONY: reset
reset:
	minikube delete

## deps: output a list of all project depencencies and indicate whether or not they are installed
.PHONY: deps
deps:
	./scripts/devdeps.sh

## appserv: run the csserver web app outside of the cluster in a local proccess
.PHONY: appserv
appserv:
	cd ./csserver && \
	go run cmd/appserver/main.go;

## gql: run csserver GraphQL generation
.PHONY: gql
gql:
	cd ./csserver/internal/appserv && \
		go run github.com/99designs/gqlgen generate && \
		cd ../.. && \
		go generate ./...;

## csmap: run csmap to create all configured object mappings
.PHONY: csmap
csmap:
	cd ./csserver && \
	csmap;

## codegen: run pnpm codegen in the csweb project.  This should be run after 'gql'
.PHONY: codegen
codegen:
	cd ./csweb && \
	pnpm codegen;

## servgen: run cstools to generate services from the configuration file
.PHONY: servgen
servgen:
	cd ./csserver/cmd/cstools && \
	go install cstools.go && \
	cstools gen;

## genall: run gql and csmap in succession
.PHONY: genall
genall:
	make gql && \
	make csmap;

## k8s-secrets: delete and recreate k8s secrets in the minikube cluster
.PHONY: k8s-secrets
k8s-secrets:
	./scripts/k8s-secrets.sh;


