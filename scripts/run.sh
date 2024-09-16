#!/bin/bash

if [[ $(minikube status | grep Running) ]]; then
    echo "Minikube already running"
else
    echo "Minikube not running: starting..."

    HAS_STATE=$(minikube status | grep Stopped)
    echo "HAS_STATE = ${HAS_STATE}"
    echo "\n\n"
    

    if [[ -z "${HAS_STATE}" ]]; then
        #---configure minikube
        echo "Setting minikube resources"
        minikube config set cpus 4  
        minikube config set memory 16384
    fi

    #---start minikube
    minikube start --driver=docker --container-runtime docker --gpus all  --mount-string="${HOME}/projects/data:/mnt/data" --mount

    if [[ -z ${HAS_STATE} ]]; then
        echo "Setting ingress and linkerd"

        #---enable minikube addons
        minikube addons enable ingress
        minikube addons enable nvidia-device-plugin
        minikube addons enable nvidia-gpu-device-plugin
        minikube addons enable nvidia-driver-installer

        #---check for linkerd dependencies...for initial setup
        linkerd check --pre

        #---install linkerd crds
        linkerd install --crds | kubectl apply -f -

        #---install linkerd
        linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -

        #---install linkerd dashboard
        linkerd viz install | kubectl apply -f -
    fi
fi

echo "Running tilt up"
tilt up
