#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m'
CHECKMARK="\xE2\x9C\x94"

dependencies=("go" "npm" "python" "node" "docker" "kubectl" "minikube" "helm" "tilt" "surreal" "pnpm" "poetry" "make" "ollama" "terraform" "linkerd" "pyenv")


for i in ${!dependencies[@]};
do
    dep=${dependencies[$i]}

    if [[ $(which ${dep}) ]]; then
        echo -e " ${CHECKMARK} ${dep} ${GREEN}already installed at $(which ${dep}) ${NC}"
    else
        echo -e " * ${dep} ${YELLOW}not installed ${NC}"   
    fi
done


